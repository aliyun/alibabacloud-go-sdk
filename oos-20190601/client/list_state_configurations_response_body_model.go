// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStateConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListStateConfigurationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListStateConfigurationsResponseBody
	GetRequestId() *string
	SetStateConfigurations(v []*ListStateConfigurationsResponseBodyStateConfigurations) *ListStateConfigurationsResponseBody
	GetStateConfigurations() []*ListStateConfigurationsResponseBodyStateConfigurations
}

type ListStateConfigurationsResponseBody struct {
	// The pagination token that was used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAASO3cL82+b5iji7bfPNpMh8=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1306108F-610C-40FD-AAD5-DA13E8B00BE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the desired-state configurations.
	StateConfigurations []*ListStateConfigurationsResponseBodyStateConfigurations `json:"StateConfigurations,omitempty" xml:"StateConfigurations,omitempty" type:"Repeated"`
}

func (s ListStateConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStateConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListStateConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStateConfigurationsResponseBody) GetStateConfigurations() []*ListStateConfigurationsResponseBodyStateConfigurations {
	return s.StateConfigurations
}

func (s *ListStateConfigurationsResponseBody) SetNextToken(v string) *ListStateConfigurationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListStateConfigurationsResponseBody) SetRequestId(v string) *ListStateConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStateConfigurationsResponseBody) SetStateConfigurations(v []*ListStateConfigurationsResponseBodyStateConfigurations) *ListStateConfigurationsResponseBody {
	s.StateConfigurations = v
	return s
}

func (s *ListStateConfigurationsResponseBody) Validate() error {
	if s.StateConfigurations != nil {
		for _, item := range s.StateConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStateConfigurationsResponseBodyStateConfigurations struct {
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
	// Collect inventory data
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
	// { "ResourceType": "ALIYUN::ECS::Instance", "Filters": [ { "Type": "All", "RegionId": "cn-hangzhou", "Parameters": { "RegionId": "cn-hangzhou", "Status": "Running" } } ] }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The template ID.
	//
	// example:
	//
	// t-ajshjalscfhjk2214
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template.
	//
	// example:
	//
	// v2
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the configuration was updated.
	//
	// example:
	//
	// 2021-04-22T03:13:32Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListStateConfigurationsResponseBodyStateConfigurations) String() string {
	return dara.Prettify(s)
}

func (s ListStateConfigurationsResponseBodyStateConfigurations) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetConfigureMode() *string {
	return s.ConfigureMode
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetDescription() *string {
	return s.Description
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetParameters() *string {
	return s.Parameters
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetStateConfigurationId() *string {
	return s.StateConfigurationId
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetTargets() *string {
	return s.Targets
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetConfigureMode(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ConfigureMode = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetCreateTime(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.CreateTime = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetDescription(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Description = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetParameters(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Parameters = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetResourceGroupId(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetScheduleExpression(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ScheduleExpression = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetScheduleType(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ScheduleType = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetStateConfigurationId(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.StateConfigurationId = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTags(v map[string]interface{}) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Tags = v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTargets(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Targets = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTemplateId(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.TemplateId = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTemplateName(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.TemplateName = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTemplateVersion(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.TemplateVersion = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetUpdateTime(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.UpdateTime = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) Validate() error {
	return dara.Validate(s)
}
