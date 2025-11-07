// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmConfig(v *CreateApplicationRequestAlarmConfig) *CreateApplicationRequest
	GetAlarmConfig() *CreateApplicationRequestAlarmConfig
	SetApplicationSource(v string) *CreateApplicationRequest
	GetApplicationSource() *string
	SetClientToken(v string) *CreateApplicationRequest
	GetClientToken() *string
	SetDescription(v string) *CreateApplicationRequest
	GetDescription() *string
	SetName(v string) *CreateApplicationRequest
	GetName() *string
	SetRegionId(v string) *CreateApplicationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateApplicationRequest
	GetResourceGroupId() *string
	SetServiceId(v string) *CreateApplicationRequest
	GetServiceId() *string
	SetTags(v map[string]interface{}) *CreateApplicationRequest
	GetTags() map[string]interface{}
}

type CreateApplicationRequest struct {
	// The configurations of application alerts.
	AlarmConfig *CreateApplicationRequestAlarmConfig `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty" type:"Struct"`
	// The source of application.
	//
	// example:
	//
	// {"Platform":"github","Owner":"githubUser","RepoName":"githubUser/repoName"}
	ApplicationSource *string `json:"ApplicationSource,omitempty" xml:"ApplicationSource,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// TF-CreateApplication-1647587475-84104b89-eba5-47a8-b2fd-807b8b7d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the Compute Nest service that corresponds to the application template.
	//
	// example:
	//
	// service-79538e30e44541b699d8
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) GetAlarmConfig() *CreateApplicationRequestAlarmConfig {
	return s.AlarmConfig
}

func (s *CreateApplicationRequest) GetApplicationSource() *string {
	return s.ApplicationSource
}

func (s *CreateApplicationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateApplicationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationRequest) GetName() *string {
	return s.Name
}

func (s *CreateApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateApplicationRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateApplicationRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateApplicationRequest) SetAlarmConfig(v *CreateApplicationRequestAlarmConfig) *CreateApplicationRequest {
	s.AlarmConfig = v
	return s
}

func (s *CreateApplicationRequest) SetApplicationSource(v string) *CreateApplicationRequest {
	s.ApplicationSource = &v
	return s
}

func (s *CreateApplicationRequest) SetClientToken(v string) *CreateApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationRequest) SetDescription(v string) *CreateApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequest) SetName(v string) *CreateApplicationRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationRequest) SetRegionId(v string) *CreateApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationRequest) SetResourceGroupId(v string) *CreateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationRequest) SetServiceId(v string) *CreateApplicationRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateApplicationRequest) SetTags(v map[string]interface{}) *CreateApplicationRequest {
	s.Tags = v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
	if s.AlarmConfig != nil {
		if err := s.AlarmConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApplicationRequestAlarmConfig struct {
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

func (s CreateApplicationRequestAlarmConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestAlarmConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestAlarmConfig) GetContactGroups() []*string {
	return s.ContactGroups
}

func (s *CreateApplicationRequestAlarmConfig) GetHealthCheckUrl() *string {
	return s.HealthCheckUrl
}

func (s *CreateApplicationRequestAlarmConfig) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *CreateApplicationRequestAlarmConfig) SetContactGroups(v []*string) *CreateApplicationRequestAlarmConfig {
	s.ContactGroups = v
	return s
}

func (s *CreateApplicationRequestAlarmConfig) SetHealthCheckUrl(v string) *CreateApplicationRequestAlarmConfig {
	s.HealthCheckUrl = &v
	return s
}

func (s *CreateApplicationRequestAlarmConfig) SetTemplateIds(v []*string) *CreateApplicationRequestAlarmConfig {
	s.TemplateIds = v
	return s
}

func (s *CreateApplicationRequestAlarmConfig) Validate() error {
	return dara.Validate(s)
}
