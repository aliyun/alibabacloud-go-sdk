// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStateConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateStateConfigurationRequest
	GetClientToken() *string
	SetConfigureMode(v string) *CreateStateConfigurationRequest
	GetConfigureMode() *string
	SetDescription(v string) *CreateStateConfigurationRequest
	GetDescription() *string
	SetParameters(v string) *CreateStateConfigurationRequest
	GetParameters() *string
	SetRegionId(v string) *CreateStateConfigurationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateStateConfigurationRequest
	GetResourceGroupId() *string
	SetScheduleExpression(v string) *CreateStateConfigurationRequest
	GetScheduleExpression() *string
	SetScheduleType(v string) *CreateStateConfigurationRequest
	GetScheduleType() *string
	SetTags(v map[string]interface{}) *CreateStateConfigurationRequest
	GetTags() map[string]interface{}
	SetTargets(v string) *CreateStateConfigurationRequest
	GetTargets() *string
	SetTemplateName(v string) *CreateStateConfigurationRequest
	GetTemplateName() *string
	SetTemplateVersion(v string) *CreateStateConfigurationRequest
	GetTemplateVersion() *string
}

type CreateStateConfigurationRequest struct {
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
	// The description of the desired-state configuration.
	//
	// example:
	//
	// The region ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters.
	//
	// example:
	//
	// {     "policy": {       "ACS:Application": {         "Collection": "Enabled"       },       "ACS:Network": {         "Collection": "Enabled"       }     }   }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	// The schedule expression. The interval between two schedules must be a minimum of 30 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// The ID of the resource group.
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The schedule type. Set the value to rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The tags to be added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "inventory"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The resources to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// {     "ResourceType": "ALIYUN::ECS::Instance",     "Filters": [       {         "Type": "All",         "RegionId": "cn-hangzhou",         "Parameters": {           "RegionId": "cn-hangzhou",           "Status": "Running"         }       }     ]   }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The name of the template. The name must be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number of the template. If you do not specify this parameter, the latest version of the template is used.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s CreateStateConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStateConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateStateConfigurationRequest) GetConfigureMode() *string {
	return s.ConfigureMode
}

func (s *CreateStateConfigurationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateStateConfigurationRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateStateConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStateConfigurationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateStateConfigurationRequest) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *CreateStateConfigurationRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateStateConfigurationRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateStateConfigurationRequest) GetTargets() *string {
	return s.Targets
}

func (s *CreateStateConfigurationRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateStateConfigurationRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *CreateStateConfigurationRequest) SetClientToken(v string) *CreateStateConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetConfigureMode(v string) *CreateStateConfigurationRequest {
	s.ConfigureMode = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetDescription(v string) *CreateStateConfigurationRequest {
	s.Description = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetParameters(v string) *CreateStateConfigurationRequest {
	s.Parameters = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetRegionId(v string) *CreateStateConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetResourceGroupId(v string) *CreateStateConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetScheduleExpression(v string) *CreateStateConfigurationRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetScheduleType(v string) *CreateStateConfigurationRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetTags(v map[string]interface{}) *CreateStateConfigurationRequest {
	s.Tags = v
	return s
}

func (s *CreateStateConfigurationRequest) SetTargets(v string) *CreateStateConfigurationRequest {
	s.Targets = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetTemplateName(v string) *CreateStateConfigurationRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetTemplateVersion(v string) *CreateStateConfigurationRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStateConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
