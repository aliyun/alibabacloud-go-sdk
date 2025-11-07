// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmConfig(v *UpdateApplicationRequestAlarmConfig) *UpdateApplicationRequest
	GetAlarmConfig() *UpdateApplicationRequestAlarmConfig
	SetDeleteAlarmRulesBeforeUpdate(v bool) *UpdateApplicationRequest
	GetDeleteAlarmRulesBeforeUpdate() *bool
	SetDescription(v string) *UpdateApplicationRequest
	GetDescription() *string
	SetName(v string) *UpdateApplicationRequest
	GetName() *string
	SetRegionId(v string) *UpdateApplicationRequest
	GetRegionId() *string
	SetTags(v map[string]interface{}) *UpdateApplicationRequest
	GetTags() map[string]interface{}
}

type UpdateApplicationRequest struct {
	// The configurations of application alerts.
	AlarmConfig *UpdateApplicationRequestAlarmConfig `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty" type:"Struct"`
	// Specifies whether to delete existing alert rules before applying the alert template. Default value: false.
	//
	// example:
	//
	// false
	DeleteAlarmRulesBeforeUpdate *bool `json:"DeleteAlarmRulesBeforeUpdate,omitempty" xml:"DeleteAlarmRulesBeforeUpdate,omitempty"`
	// The description to be updated for the application.
	//
	// example:
	//
	// test application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// My-Application
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequest) GetAlarmConfig() *UpdateApplicationRequestAlarmConfig {
	return s.AlarmConfig
}

func (s *UpdateApplicationRequest) GetDeleteAlarmRulesBeforeUpdate() *bool {
	return s.DeleteAlarmRulesBeforeUpdate
}

func (s *UpdateApplicationRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationRequest) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApplicationRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateApplicationRequest) SetAlarmConfig(v *UpdateApplicationRequestAlarmConfig) *UpdateApplicationRequest {
	s.AlarmConfig = v
	return s
}

func (s *UpdateApplicationRequest) SetDeleteAlarmRulesBeforeUpdate(v bool) *UpdateApplicationRequest {
	s.DeleteAlarmRulesBeforeUpdate = &v
	return s
}

func (s *UpdateApplicationRequest) SetDescription(v string) *UpdateApplicationRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationRequest) SetName(v string) *UpdateApplicationRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationRequest) SetRegionId(v string) *UpdateApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationRequest) SetTags(v map[string]interface{}) *UpdateApplicationRequest {
	s.Tags = v
	return s
}

func (s *UpdateApplicationRequest) Validate() error {
	if s.AlarmConfig != nil {
		if err := s.AlarmConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationRequestAlarmConfig struct {
	// The alert contact groups.
	ContactGroups []*string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	// The health check URL of the application.
	//
	// example:
	//
	// /healthcheck/tcp50122
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The alert templates.
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s UpdateApplicationRequestAlarmConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationRequestAlarmConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequestAlarmConfig) GetContactGroups() []*string {
	return s.ContactGroups
}

func (s *UpdateApplicationRequestAlarmConfig) GetHealthCheckUrl() *string {
	return s.HealthCheckUrl
}

func (s *UpdateApplicationRequestAlarmConfig) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *UpdateApplicationRequestAlarmConfig) SetContactGroups(v []*string) *UpdateApplicationRequestAlarmConfig {
	s.ContactGroups = v
	return s
}

func (s *UpdateApplicationRequestAlarmConfig) SetHealthCheckUrl(v string) *UpdateApplicationRequestAlarmConfig {
	s.HealthCheckUrl = &v
	return s
}

func (s *UpdateApplicationRequestAlarmConfig) SetTemplateIds(v []*string) *UpdateApplicationRequestAlarmConfig {
	s.TemplateIds = v
	return s
}

func (s *UpdateApplicationRequestAlarmConfig) Validate() error {
	return dara.Validate(s)
}
