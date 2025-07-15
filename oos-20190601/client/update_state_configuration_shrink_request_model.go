// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStateConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateStateConfigurationShrinkRequest
	GetClientToken() *string
	SetConfigureMode(v string) *UpdateStateConfigurationShrinkRequest
	GetConfigureMode() *string
	SetDescription(v string) *UpdateStateConfigurationShrinkRequest
	GetDescription() *string
	SetParametersShrink(v string) *UpdateStateConfigurationShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *UpdateStateConfigurationShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateStateConfigurationShrinkRequest
	GetResourceGroupId() *string
	SetScheduleExpression(v string) *UpdateStateConfigurationShrinkRequest
	GetScheduleExpression() *string
	SetScheduleType(v string) *UpdateStateConfigurationShrinkRequest
	GetScheduleType() *string
	SetStateConfigurationId(v string) *UpdateStateConfigurationShrinkRequest
	GetStateConfigurationId() *string
	SetTagsShrink(v string) *UpdateStateConfigurationShrinkRequest
	GetTagsShrink() *string
	SetTargets(v string) *UpdateStateConfigurationShrinkRequest
	GetTargets() *string
}

type UpdateStateConfigurationShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// DASKJJLKADS-AHKLJHJSAKL-AJK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration mode. Valid values: ApplyOnce: The configuration is applied only once. After a configuration is updated, the new configuration is applied. ApplyAndMonitor: The configuration is applied only once. After the configuration is applied, the system only checks whether the configuration is migrated in the future. ApplyAndAutoCorrect: The configuration is always applied.
	//
	// example:
	//
	// ApplyOnce
	ConfigureMode *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	// The description.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters.
	//
	// example:
	//
	// { "policy": { "ACS:Application": { "Collection": "Enabled" }, "ACS:Network": { "Collection": "Enabled" } } }
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The schedule expression.
	//
	// example:
	//
	// 1 hour
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The schedule type.
	//
	// example:
	//
	// rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The ID of the desired-state configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// sc-asfgdhj12345
	StateConfigurationId *string `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	// The tags to be added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "sc"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The resources to be queried.
	//
	// example:
	//
	// { "ResourceType": "ALIYUN::ECS::Instance", "Filters": [ { "Type": "All", "RegionId": "cn-hangzhou", "Parameters": { "RegionId": "cn-hangzhou", "Status": "Running" } } ] }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
}

func (s UpdateStateConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStateConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateStateConfigurationShrinkRequest) GetConfigureMode() *string {
	return s.ConfigureMode
}

func (s *UpdateStateConfigurationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateStateConfigurationShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *UpdateStateConfigurationShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateStateConfigurationShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateStateConfigurationShrinkRequest) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *UpdateStateConfigurationShrinkRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *UpdateStateConfigurationShrinkRequest) GetStateConfigurationId() *string {
	return s.StateConfigurationId
}

func (s *UpdateStateConfigurationShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateStateConfigurationShrinkRequest) GetTargets() *string {
	return s.Targets
}

func (s *UpdateStateConfigurationShrinkRequest) SetClientToken(v string) *UpdateStateConfigurationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetConfigureMode(v string) *UpdateStateConfigurationShrinkRequest {
	s.ConfigureMode = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetDescription(v string) *UpdateStateConfigurationShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetParametersShrink(v string) *UpdateStateConfigurationShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetRegionId(v string) *UpdateStateConfigurationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetResourceGroupId(v string) *UpdateStateConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetScheduleExpression(v string) *UpdateStateConfigurationShrinkRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetScheduleType(v string) *UpdateStateConfigurationShrinkRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetStateConfigurationId(v string) *UpdateStateConfigurationShrinkRequest {
	s.StateConfigurationId = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetTagsShrink(v string) *UpdateStateConfigurationShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetTargets(v string) *UpdateStateConfigurationShrinkRequest {
	s.Targets = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
