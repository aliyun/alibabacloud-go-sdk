// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStateConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateStateConfigurationResponseBody
	GetRequestId() *string
	SetStateConfiguration(v []*UpdateStateConfigurationResponseBodyStateConfiguration) *UpdateStateConfigurationResponseBody
	GetStateConfiguration() []*UpdateStateConfigurationResponseBodyStateConfiguration
}

type UpdateStateConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1306108F-610C-40FD-AAD5-DA13E8B00BE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the configuration.
	StateConfiguration []*UpdateStateConfigurationResponseBodyStateConfiguration `json:"StateConfiguration,omitempty" xml:"StateConfiguration,omitempty" type:"Repeated"`
}

func (s UpdateStateConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStateConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStateConfigurationResponseBody) GetStateConfiguration() []*UpdateStateConfigurationResponseBodyStateConfiguration {
	return s.StateConfiguration
}

func (s *UpdateStateConfigurationResponseBody) SetRequestId(v string) *UpdateStateConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBody) SetStateConfiguration(v []*UpdateStateConfigurationResponseBodyStateConfiguration) *UpdateStateConfigurationResponseBody {
	s.StateConfiguration = v
	return s
}

func (s *UpdateStateConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateStateConfigurationResponseBodyStateConfiguration struct {
	// The configuration mode. Valid values:
	//
	// example:
	//
	// ApplyAndAutoCorrect
	ConfigureMode *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	// The time when the configuration was created.
	//
	// example:
	//
	// 2021-03-22T03:13:32Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the desired-state configuration.
	//
	// example:
	//
	// collect inventory data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters.
	//
	// example:
	//
	// {"policy": {"ACS:Network": {"Collection": "Enabled"}, "ACS:Application": {"Collection": "Enabled"}}}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The CRON expression.
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
	// example:
	//
	// StateConfigurationId
	StateConfigurationId *string `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	// The tags added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "inventory"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The queried resources.
	//
	// example:
	//
	// { "ResourceType": "ALIYUN::ECS::Instance", "Filters": [ { "Type": "All", "RegionId": "cn-hangzhou", "Parameters": { "RegionId": "cn-hangzhou", "Status": "Running" } } ] }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The template ID.
	//
	// example:
	//
	// t-1234asadf
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The name of the template version.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the configuration was updated.
	//
	// example:
	//
	// 2021-03-22T03:13:32Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateStateConfigurationResponseBodyStateConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateStateConfigurationResponseBodyStateConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetConfigureMode() *string {
	return s.ConfigureMode
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetDescription() *string {
	return s.Description
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetParameters() *string {
	return s.Parameters
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetStateConfigurationId() *string {
	return s.StateConfigurationId
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetTargets() *string {
	return s.Targets
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetConfigureMode(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ConfigureMode = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetCreateTime(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.CreateTime = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetDescription(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Description = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetParameters(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Parameters = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetResourceGroupId(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetScheduleExpression(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleExpression = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetScheduleType(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleType = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetStateConfigurationId(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.StateConfigurationId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTags(v map[string]interface{}) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Tags = v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTargets(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Targets = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTemplateId(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTemplateName(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateName = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTemplateVersion(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetUpdateTime(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) Validate() error {
	return dara.Validate(s)
}
