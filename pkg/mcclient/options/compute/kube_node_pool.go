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

package compute

import (
	"yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/mcclient/options"
)

type KubeNodePoolListOptions struct {
	options.BaseListOptions
}

func (opts *KubeNodePoolListOptions) Params() (jsonutils.JSONObject, error) {
	return options.ListStructToParams(opts)
}

type KubeNodePoolIdOption struct {
	ID string `help:"KubeNodePool Id"`
}

func (opts *KubeNodePoolIdOption) GetId() string {
	return opts.ID
}

func (opts *KubeNodePoolIdOption) Params() (jsonutils.JSONObject, error) {
	return nil, nil
}

type KubeNodePoolCreateOptions struct {
	options.BaseCreateOptions
	NetworkIds           []string `metavar:"NETWORK"`
	InstanceTypes        []string `metavar:"INSTANCE_TYPE"`
	MinInstanceCount     int
	MaxInstanceCount     int
	DesiredInstanceCount int
	RootDiskSizeGb       int
	CloudKubeClusterId   string `metavar:"CLUSTER"`
}

func (opts *KubeNodePoolCreateOptions) Params() (jsonutils.JSONObject, error) {
	return jsonutils.Marshal(opts), nil
}
