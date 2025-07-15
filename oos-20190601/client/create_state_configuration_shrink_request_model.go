// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStateConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateStateConfigurationShrinkRequest
	GetClientToken() *string
	SetConfigureMode(v string) *CreateStateConfigurationShrinkRequest
	GetConfigureMode() *string
	SetDescription(v string) *CreateStateConfigurationShrinkRequest
	GetDescription() *string
	SetParameters(v string) *CreateStateConfigurationShrinkRequest
	GetParameters() *string
	SetRegionId(v string) *CreateStateConfigurationShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateStateConfigurationShrinkRequest
	GetResourceGroupId() *string
	SetScheduleExpression(v string) *CreateStateConfigurationShrinkRequest
	GetScheduleExpression() *string
	SetScheduleType(v string) *CreateStateConfigurationShrinkRequest
	GetScheduleType() *string
	SetTagsShrink(v string) *CreateStateConfigurationShrinkRequest
	GetTagsShrink() *string
	SetTargets(v string) *CreateStateConfigurationShrinkRequest
	GetTargets() *string
	SetTemplateName(v string) *CreateStateConfigurationShrinkRequest
	GetTemplateName() *string
	SetTemplateVersion(v string) *CreateStateConfigurationShrinkRequest
	GetTemplateVersion() *string
}

type CreateStateConfigurationShrinkRequest struct {
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s CreateStateConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStateConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateStateConfigurationShrinkRequest) GetConfigureMode() *string {
	return s.ConfigureMode
}

func (s *CreateStateConfigurationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateStateConfigurationShrinkRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateStateConfigurationShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStateConfigurationShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateStateConfigurationShrinkRequest) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *CreateStateConfigurationShrinkRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateStateConfigurationShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateStateConfigurationShrinkRequest) GetTargets() *string {
	return s.Targets
}

func (s *CreateStateConfigurationShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateStateConfigurationShrinkRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *CreateStateConfigurationShrinkRequest) SetClientToken(v string) *CreateStateConfigurationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetConfigureMode(v string) *CreateStateConfigurationShrinkRequest {
	s.ConfigureMode = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetDescription(v string) *CreateStateConfigurationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetParameters(v string) *CreateStateConfigurationShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetRegionId(v string) *CreateStateConfigurationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetResourceGroupId(v string) *CreateStateConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetScheduleExpression(v string) *CreateStateConfigurationShrinkRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetScheduleType(v string) *CreateStateConfigurationShrinkRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTagsShrink(v string) *CreateStateConfigurationShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTargets(v string) *CreateStateConfigurationShrinkRequest {
	s.Targets = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTemplateName(v string) *CreateStateConfigurationShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTemplateVersion(v string) *CreateStateConfigurationShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
