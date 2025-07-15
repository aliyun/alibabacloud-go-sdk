// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStateConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateStateConfigurationResponseBody
	GetRequestId() *string
	SetStateConfiguration(v *CreateStateConfigurationResponseBodyStateConfiguration) *CreateStateConfigurationResponseBody
	GetStateConfiguration() *CreateStateConfigurationResponseBodyStateConfiguration
}

type CreateStateConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1306108F-610C-40FD-AAD5-DA13E8B00BE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the desired-state configuration.
	StateConfiguration *CreateStateConfigurationResponseBodyStateConfiguration `json:"StateConfiguration,omitempty" xml:"StateConfiguration,omitempty" type:"Struct"`
}

func (s CreateStateConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStateConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStateConfigurationResponseBody) GetStateConfiguration() *CreateStateConfigurationResponseBodyStateConfiguration {
	return s.StateConfiguration
}

func (s *CreateStateConfigurationResponseBody) SetRequestId(v string) *CreateStateConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStateConfigurationResponseBody) SetStateConfiguration(v *CreateStateConfigurationResponseBodyStateConfiguration) *CreateStateConfigurationResponseBody {
	s.StateConfiguration = v
	return s
}

func (s *CreateStateConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateStateConfigurationResponseBodyStateConfiguration struct {
	// The configuration mode. Valid values:
	//
	// example:
	//
	// ApplyAndAutoCorrect
	ConfigureMode *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	// The time when the desired-state configuration was created.
	//
	// example:
	//
	// 2021-03-22T03:13:32Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	// example:
	//
	// sc-a538febe18fabcdef
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
	// {     "ResourceType": "ALIYUN::ECS::Instance",     "Filters": [       {         "Type": "All",         "RegionId": "cn-hangzhou",         "Parameters": {           "RegionId": "cn-hangzhou",           "Status": "Running"         }       }     ]   }
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
}

func (s CreateStateConfigurationResponseBodyStateConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateStateConfigurationResponseBodyStateConfiguration) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetConfigureMode() *string {
	return s.ConfigureMode
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetDescription() *string {
	return s.Description
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetStateConfigurationId() *string {
	return s.StateConfigurationId
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetTargets() *string {
	return s.Targets
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetConfigureMode(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ConfigureMode = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetCreateTime(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.CreateTime = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetDescription(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Description = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetParameters(v map[string]interface{}) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Parameters = v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetResourceGroupId(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetScheduleExpression(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleExpression = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetScheduleType(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleType = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetStateConfigurationId(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.StateConfigurationId = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTags(v map[string]interface{}) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Tags = v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTargets(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Targets = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTemplateId(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateId = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTemplateName(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateName = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTemplateVersion(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) Validate() error {
	return dara.Validate(s)
}
