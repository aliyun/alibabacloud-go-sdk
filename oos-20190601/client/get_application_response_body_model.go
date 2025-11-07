// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody
	GetApplication() *GetApplicationResponseBodyApplication
	SetRequestId(v string) *GetApplicationResponseBody
	GetRequestId() *string
}

type GetApplicationResponseBody struct {
	// The information about the application.
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 51004B8A-6D9A-5ACB-9158-6C6794496AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) GetApplication() *GetApplicationResponseBodyApplication {
	return s.Application
}

func (s *GetApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationResponseBody) SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody {
	s.Application = v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) Validate() error {
	if s.Application != nil {
		if err := s.Application.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyApplication struct {
	// The configurations of application alerts.
	AlarmConfig *GetApplicationResponseBodyApplicationAlarmConfig `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty" type:"Struct"`
	// The source of application.
	//
	// example:
	//
	// {"Platform":"gitee","Owner":"aliyun-computenest","RepoName":"aliyun-computenest/java-springboot-demo","Name":"java-springboot-demo"}
	ApplicationSource *string `json:"ApplicationSource,omitempty" xml:"ApplicationSource,omitempty"`
	// The type of the application.
	//
	// Valid values:
	//
	// 	- ComputeNest
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Custom
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DingTalk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DingTalk
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2021-09-07T09:17:46Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2021-09-07T09:17:46Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetApplicationResponseBodyApplication) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplication) GetAlarmConfig() *GetApplicationResponseBodyApplicationAlarmConfig {
	return s.AlarmConfig
}

func (s *GetApplicationResponseBodyApplication) GetApplicationSource() *string {
	return s.ApplicationSource
}

func (s *GetApplicationResponseBodyApplication) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *GetApplicationResponseBodyApplication) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetApplicationResponseBodyApplication) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyApplication) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyApplication) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetApplicationResponseBodyApplication) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetApplicationResponseBodyApplication) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetApplicationResponseBodyApplication) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetApplicationResponseBodyApplication) SetAlarmConfig(v *GetApplicationResponseBodyApplicationAlarmConfig) *GetApplicationResponseBodyApplication {
	s.AlarmConfig = v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationSource(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationSource = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationType(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCreateDate(v string) *GetApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDescription(v string) *GetApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetName(v string) *GetApplicationResponseBodyApplication {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetResourceGroupId(v string) *GetApplicationResponseBodyApplication {
	s.ResourceGroupId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetServiceId(v string) *GetApplicationResponseBodyApplication {
	s.ServiceId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetTags(v map[string]interface{}) *GetApplicationResponseBodyApplication {
	s.Tags = v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetUpdateDate(v string) *GetApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) Validate() error {
	if s.AlarmConfig != nil {
		if err := s.AlarmConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyApplicationAlarmConfig struct {
	// The alert contact list.
	ContactGroups []*string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	// The health check URL of the application.
	//
	// example:
	//
	// /api/health/
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The ID of the alert template.
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyApplicationAlarmConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationAlarmConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationAlarmConfig) GetContactGroups() []*string {
	return s.ContactGroups
}

func (s *GetApplicationResponseBodyApplicationAlarmConfig) GetHealthCheckUrl() *string {
	return s.HealthCheckUrl
}

func (s *GetApplicationResponseBodyApplicationAlarmConfig) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *GetApplicationResponseBodyApplicationAlarmConfig) SetContactGroups(v []*string) *GetApplicationResponseBodyApplicationAlarmConfig {
	s.ContactGroups = v
	return s
}

func (s *GetApplicationResponseBodyApplicationAlarmConfig) SetHealthCheckUrl(v string) *GetApplicationResponseBodyApplicationAlarmConfig {
	s.HealthCheckUrl = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationAlarmConfig) SetTemplateIds(v []*string) *GetApplicationResponseBodyApplicationAlarmConfig {
	s.TemplateIds = v
	return s
}

func (s *GetApplicationResponseBodyApplicationAlarmConfig) Validate() error {
	return dara.Validate(s)
}
