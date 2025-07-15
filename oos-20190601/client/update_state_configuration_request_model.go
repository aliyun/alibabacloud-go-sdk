// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStateConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateStateConfigurationRequest
	GetClientToken() *string
	SetConfigureMode(v string) *UpdateStateConfigurationRequest
	GetConfigureMode() *string
	SetDescription(v string) *UpdateStateConfigurationRequest
	GetDescription() *string
	SetParameters(v map[string]interface{}) *UpdateStateConfigurationRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *UpdateStateConfigurationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateStateConfigurationRequest
	GetResourceGroupId() *string
	SetScheduleExpression(v string) *UpdateStateConfigurationRequest
	GetScheduleExpression() *string
	SetScheduleType(v string) *UpdateStateConfigurationRequest
	GetScheduleType() *string
	SetStateConfigurationId(v string) *UpdateStateConfigurationRequest
	GetStateConfigurationId() *string
	SetTags(v map[string]interface{}) *UpdateStateConfigurationRequest
	GetTags() map[string]interface{}
	SetTargets(v string) *UpdateStateConfigurationRequest
	GetTargets() *string
}

type UpdateStateConfigurationRequest struct {
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The resources to be queried.
	//
	// example:
	//
	// { "ResourceType": "ALIYUN::ECS::Instance", "Filters": [ { "Type": "All", "RegionId": "cn-hangzhou", "Parameters": { "RegionId": "cn-hangzhou", "Status": "Running" } } ] }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
}

func (s UpdateStateConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStateConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateStateConfigurationRequest) GetConfigureMode() *string {
	return s.ConfigureMode
}

func (s *UpdateStateConfigurationRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateStateConfigurationRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *UpdateStateConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateStateConfigurationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateStateConfigurationRequest) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *UpdateStateConfigurationRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *UpdateStateConfigurationRequest) GetStateConfigurationId() *string {
	return s.StateConfigurationId
}

func (s *UpdateStateConfigurationRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateStateConfigurationRequest) GetTargets() *string {
	return s.Targets
}

func (s *UpdateStateConfigurationRequest) SetClientToken(v string) *UpdateStateConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetConfigureMode(v string) *UpdateStateConfigurationRequest {
	s.ConfigureMode = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetDescription(v string) *UpdateStateConfigurationRequest {
	s.Description = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetParameters(v map[string]interface{}) *UpdateStateConfigurationRequest {
	s.Parameters = v
	return s
}

func (s *UpdateStateConfigurationRequest) SetRegionId(v string) *UpdateStateConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetResourceGroupId(v string) *UpdateStateConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetScheduleExpression(v string) *UpdateStateConfigurationRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetScheduleType(v string) *UpdateStateConfigurationRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetStateConfigurationId(v string) *UpdateStateConfigurationRequest {
	s.StateConfigurationId = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetTags(v map[string]interface{}) *UpdateStateConfigurationRequest {
	s.Tags = v
	return s
}

func (s *UpdateStateConfigurationRequest) SetTargets(v string) *UpdateStateConfigurationRequest {
	s.Targets = &v
	return s
}

func (s *UpdateStateConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
