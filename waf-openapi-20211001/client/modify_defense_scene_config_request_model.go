// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseSceneConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *ModifyDefenseSceneConfigRequest
	GetConfigKey() *string
	SetConfigValue(v string) *ModifyDefenseSceneConfigRequest
	GetConfigValue() *string
	SetDefenseScene(v string) *ModifyDefenseSceneConfigRequest
	GetDefenseScene() *string
	SetInstanceId(v string) *ModifyDefenseSceneConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefenseSceneConfigRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefenseSceneConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyDefenseSceneConfigRequest struct {
	// The name of the configuration item to modify.
	//
	// If **DefenseScene*	- is set to **apisec**, the valid value is:
	//
	// - **autoEnabled**: indicates whether core API security detection is automatically enabled for new resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// autoEnableStatus
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value to set for the configuration item.
	//
	// > The value of this parameter depends on the value of **ConfigKey**. For more information, see **Description of mitigation setting parameters**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The protection scenario for which you want to modify the mitigation settings. Valid values:
	//
	// - **apisec**: API security.
	//
	// This parameter is required.
	//
	// example:
	//
	// apisec
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
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

func (s ModifyDefenseSceneConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseSceneConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *ModifyDefenseSceneConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ModifyDefenseSceneConfigRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *ModifyDefenseSceneConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseSceneConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefenseSceneConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefenseSceneConfigRequest) SetConfigKey(v string) *ModifyDefenseSceneConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *ModifyDefenseSceneConfigRequest) SetConfigValue(v string) *ModifyDefenseSceneConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *ModifyDefenseSceneConfigRequest) SetDefenseScene(v string) *ModifyDefenseSceneConfigRequest {
	s.DefenseScene = &v
	return s
}

func (s *ModifyDefenseSceneConfigRequest) SetInstanceId(v string) *ModifyDefenseSceneConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseSceneConfigRequest) SetRegionId(v string) *ModifyDefenseSceneConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseSceneConfigRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseSceneConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseSceneConfigRequest) Validate() error {
	return dara.Validate(s)
}
