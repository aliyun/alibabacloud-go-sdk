// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *ModifyAccessConfigRequest
	GetConfigKey() *string
	SetConfigValue(v string) *ModifyAccessConfigRequest
	GetConfigValue() *string
	SetInstanceId(v string) *ModifyAccessConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyAccessConfigRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyAccessConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyAccessConfigRequest struct {
	// The key of the configuration item to be modified. Valid values:
	//
	// - **auto_access**: specifies whether to enable the automatic provisioning feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// auto_access
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value of the configuration item. The value is a JSON string. For more information about the fields, see the following section.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"status":1,"products":["alb"]}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s ModifyAccessConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *ModifyAccessConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ModifyAccessConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAccessConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccessConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyAccessConfigRequest) SetConfigKey(v string) *ModifyAccessConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *ModifyAccessConfigRequest) SetConfigValue(v string) *ModifyAccessConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *ModifyAccessConfigRequest) SetInstanceId(v string) *ModifyAccessConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAccessConfigRequest) SetRegionId(v string) *ModifyAccessConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccessConfigRequest) SetResourceManagerResourceGroupId(v string) *ModifyAccessConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyAccessConfigRequest) Validate() error {
	return dara.Validate(s)
}
