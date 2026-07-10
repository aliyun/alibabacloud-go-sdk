// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody
	GetApplications() []*ListApplicationsResponseBodyApplications
	SetRequestId(v string) *ListApplicationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationsResponseBody
	GetTotalCount() *int64
}

type ListApplicationsResponseBody struct {
	// The list of application information.
	Applications []*ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) GetApplications() []*ListApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationsResponseBody) SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) SetTotalCount(v int64) *ListApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsResponseBody) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsResponseBodyApplications struct {
	// The application creation type.
	//
	// example:
	//
	// user_custom
	ApplicationCreationType *string `json:"ApplicationCreationType,omitempty" xml:"ApplicationCreationType,omitempty"`
	// The application ID.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application identity type.
	//
	// example:
	//
	// application
	ApplicationIdentityType *string `json:"ApplicationIdentityType,omitempty" xml:"ApplicationIdentityType,omitempty"`
	// The application name.
	//
	// example:
	//
	// SAML Application
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application creation source. Valid values:
	//
	// - urn:alibaba:idaas:app:source:template: Application template.
	//
	// - urn:alibaba:idaas:app:source:standard: Standard protocol.
	//
	// example:
	//
	// urn:alibaba:idaas:app:source:standard
	ApplicationSourceType *string `json:"ApplicationSourceType,omitempty" xml:"ApplicationSourceType,omitempty"`
	// The application template ID.
	//
	// example:
	//
	// apt_xxx_xxx
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The client ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The time when the application was created, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The application description.
	//
	// example:
	//
	// A test application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The features supported by the application, returned as a JSON array string. Valid values:
	//
	// - sso: Single sign-on.
	//
	// - slo: Single logout.
	//
	// - provision: Account synchronization.
	//
	// - api_invoke: API access.
	//
	// - m2m_client: M2M Client capability.
	//
	// - resource_server: API service capability.
	//
	// - other: Fallback.
	//
	// example:
	//
	// ["sso", "provision"]
	Features *string `json:"Features,omitempty" xml:"Features,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The URL of the application logo.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01lvYwpv1aGowQXDML9_!!6000000003303-0-tps-580-580.jpg
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The ServiceCode of the Alibaba Cloud service that manages the application template.
	//
	// example:
	//
	// rpa
	ManagedServiceCode *string `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	// The unique identifier of the ResourceServer.
	//
	// example:
	//
	// https://example.com
	ResourceServerIdentifier *string `json:"ResourceServerIdentifier,omitempty" xml:"ResourceServerIdentifier,omitempty"`
	// The resource server source type.
	//
	// example:
	//
	// urn:cloud:idaas:resourceserver:source:custom
	ResourceServerSourceType *string `json:"ResourceServerSourceType,omitempty" xml:"ResourceServerSourceType,omitempty"`
	// The resource server status.
	//
	// example:
	//
	// enabled
	ResourceServerStatus *string `json:"ResourceServerStatus,omitempty" xml:"ResourceServerStatus,omitempty"`
	// Indicates whether the application template is managed by an Alibaba Cloud service.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The single sign-on (SSO) logon protocol. Valid values:
	//
	// - saml2: SAML 2.0 protocol.
	//
	// - oidc: OpenID Connect protocol.
	//
	// - oauth2/m2m: OAuth 2.0 protocol.
	//
	// - oidc+oauth2/m2m: OpenID Connect and OAuth 2.0 protocols.
	//
	// example:
	//
	// saml2
	SsoType *string `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
	// The application status. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// - deleted: Soft deleted.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the application was last updated, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationCreationType() *string {
	return s.ApplicationCreationType
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationIdentityType() *string {
	return s.ApplicationIdentityType
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationSourceType() *string {
	return s.ApplicationSourceType
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *ListApplicationsResponseBodyApplications) GetClientId() *string {
	return s.ClientId
}

func (s *ListApplicationsResponseBodyApplications) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationsResponseBodyApplications) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationsResponseBodyApplications) GetFeatures() *string {
	return s.Features
}

func (s *ListApplicationsResponseBodyApplications) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsResponseBodyApplications) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *ListApplicationsResponseBodyApplications) GetManagedServiceCode() *string {
	return s.ManagedServiceCode
}

func (s *ListApplicationsResponseBodyApplications) GetResourceServerIdentifier() *string {
	return s.ResourceServerIdentifier
}

func (s *ListApplicationsResponseBodyApplications) GetResourceServerSourceType() *string {
	return s.ResourceServerSourceType
}

func (s *ListApplicationsResponseBodyApplications) GetResourceServerStatus() *string {
	return s.ResourceServerStatus
}

func (s *ListApplicationsResponseBodyApplications) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListApplicationsResponseBodyApplications) GetSsoType() *string {
	return s.SsoType
}

func (s *ListApplicationsResponseBodyApplications) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationsResponseBodyApplications) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationCreationType(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationCreationType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationId(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationIdentityType(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationIdentityType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationName(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationSourceType(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationSourceType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationTemplateId(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationTemplateId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetClientId(v string) *ListApplicationsResponseBodyApplications {
	s.ClientId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetCreateTime(v int64) *ListApplicationsResponseBodyApplications {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetDescription(v string) *ListApplicationsResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetFeatures(v string) *ListApplicationsResponseBodyApplications {
	s.Features = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetInstanceId(v string) *ListApplicationsResponseBodyApplications {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetLogoUrl(v string) *ListApplicationsResponseBodyApplications {
	s.LogoUrl = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetManagedServiceCode(v string) *ListApplicationsResponseBodyApplications {
	s.ManagedServiceCode = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetResourceServerIdentifier(v string) *ListApplicationsResponseBodyApplications {
	s.ResourceServerIdentifier = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetResourceServerSourceType(v string) *ListApplicationsResponseBodyApplications {
	s.ResourceServerSourceType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetResourceServerStatus(v string) *ListApplicationsResponseBodyApplications {
	s.ResourceServerStatus = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetServiceManaged(v bool) *ListApplicationsResponseBodyApplications {
	s.ServiceManaged = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetSsoType(v string) *ListApplicationsResponseBodyApplications {
	s.SsoType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetStatus(v string) *ListApplicationsResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetUpdateTime(v int64) *ListApplicationsResponseBodyApplications {
	s.UpdateTime = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
