// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sender

import (
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/pkg/errors"

	api "yunion.io/x/onecloud/pkg/apis/notify"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/notify/models"
)

type SMobileSender struct {
	config map[string]api.SNotifyConfigContent
}

func (smsSender *SMobileSender) GetSenderType() string {
	return api.MOBILE
}

func (smsSender *SMobileSender) Send(args api.SendParams) error {
	smsSendParams := api.SSMSSendParams{
		TemplateParas:       args.Message,
		To:                  args.Receivers.Contact,
		RemoteTemplate:      args.RemoteTemplate,
		RemoteTemplateParam: args.RemoteTemplateParam,
	}
	smsdriver := models.GetSMSDriver(models.ConfigMap[api.MOBILE].Content.SmsDriver)
	return smsdriver.Send(smsSendParams, false, &api.NotifyConfig{
		SNotifyConfigContent: *models.ConfigMap[api.MOBILE].Content,
		Attribution:          models.ConfigMap[api.MOBILE].Attribution,
		DomainId:             models.ConfigMap[api.MOBILE].DomainId,
	})
}

func (smsSender *SMobileSender) ValidateConfig(config api.NotifyConfig) (string, error) {
	driver := models.GetSMSDriver(config.SmsDriver)
	if driver == nil {
		return "", errors.Wrap(errors.ErrNotFound, "driver disabled")
	}
	return "", nil
}

func (smsSender *SMobileSender) UpdateConfig(config api.NotifyConfig) error {
	q := models.ConfigManager.Query()
	q = q.Equals("type", api.MOBILE)
	confs := []models.SConfig{}
	db.FetchModelObjects(models.ConfigManager, q, &confs)
	if len(confs) == 0 {
		return errors.Wrapf(errors.ErrNotFound, "config type:%s", api.MOBILE)
	}
	_, err := db.Update(&confs[0], func() error {
		confs[0].Content = &config.SNotifyConfigContent
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "update config")
	}
	models.ConfigMap[api.MOBILE] = models.SConfig{
		Content: &config.SNotifyConfigContent,
	}
	return nil
}

func (smsSender *SMobileSender) AddConfig(config api.NotifyConfig) error {
	return cloudprovider.ErrNotImplemented
}

func (smsSender *SMobileSender) DeleteConfig(config api.NotifyConfig) error {
	return cloudprovider.ErrNotImplemented
}

func (smsSender *SMobileSender) ContactByMobile(mobile, domainId string) (string, error) {
	return "", nil
}

func (smsSender *SMobileSender) IsPersonal() bool {
	return true
}

func (smsSender *SMobileSender) IsRobot() bool {
	return false
}

func (smsSender *SMobileSender) IsValid() bool {
	return len(smsSender.config) > 0
}

func (smsSender *SMobileSender) IsPullType() bool {
	return false
}

func (smsSender *SMobileSender) IsSystemConfigContactType() bool {
	return true
}

func (smsSender *SMobileSender) GetAccessToken(key string) error {
	return nil
}

func (smsSender *SMobileSender) RegisterConfig(config models.SConfig) {
	models.ConfigMap[config.Type] = config
}

func init() {
	models.Register(&SMobileSender{
		config: map[string]api.SNotifyConfigContent{},
	})
}
