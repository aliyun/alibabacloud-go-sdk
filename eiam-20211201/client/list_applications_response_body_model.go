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
	// The details of the applications.
	Applications []*ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned entries.
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
	// example:
	//
	// user_created
	ApplicationCreationType *string `json:"ApplicationCreationType,omitempty" xml:"ApplicationCreationType,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId           *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ApplicationIdentityType *string `json:"ApplicationIdentityType,omitempty" xml:"ApplicationIdentityType,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// SAML Application
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The origin of the application. Valid values:
	//
	// 	- urn:alibaba:idaas:app:source:template: The application is created based on a template.
	//
	// 	- urn:alibaba:idaas: The application is created based on the standard protocol.
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
	// The time when the application was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// A single application. The code is pkces.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The features that are supported by the application. The value is a JSON array. Valid values:
	//
	// 	- sso: The application supports SSO.
	//
	// 	- slo: The application supports SLO.
	//
	// 	- provision: The application supports account synchronization.
	//
	// 	- api_invoke: The application supports custom APIs.
	//
	// 	- m2m_client: The application supports M2M Client.
	//
	// 	- resource_server: The application supports Resource Server.
	//
	// 	- other: undertake.
	//
	// example:
	//
	// ["sso","slo", "provision","api_invoke", "m2m_client","resource_server","other"]
	Features *string `json:"Features,omitempty" xml:"Features,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The URL of the application icon.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01lvYwpv1aGowQXDML9_!!6000000003303-0-tps-580-580.jpg
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The service code of the cloud service that manages the application template.
	//
	// example:
	//
	// rpa
	ManagedServiceCode *string `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	// example:
	//
	// test
	ResourceServerIdentifier *string `json:"ResourceServerIdentifier,omitempty" xml:"ResourceServerIdentifier,omitempty"`
	// example:
	//
	// urn:cloud:idaas:resourceserver:source:custom
	ResourceServerSourceType *string `json:"ResourceServerSourceType,omitempty" xml:"ResourceServerSourceType,omitempty"`
	// example:
	//
	// enabled
	ResourceServerStatus *string `json:"ResourceServerStatus,omitempty" xml:"ResourceServerStatus,omitempty"`
	// Indicates whether the application template is managed by a cloud service.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The type of the single sign-on (SSO) protocol. Valid values:
	//
	// 	- saml2: the Security Assertion Markup Language (SAML) 2.0 protocol.
	//
	// 	- oidc: the OpenID Connect (OIDC) protocol.
	//
	// 	- oauth2/m2m: the OAuth2.0  protocol M2M.
	//
	// example:
	//
	// saml2
	SsoType *string `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- enabled: The application is enabled.
	//
	// 	- disabled: The application is disabled.
	//
	// 	- deleted: The application is deleted.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the application was last updated. The value is a UNIX timestamp. Unit: milliseconds.
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
