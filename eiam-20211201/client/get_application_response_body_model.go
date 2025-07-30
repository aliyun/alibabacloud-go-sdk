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
	// The details of the application.
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
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
	return dara.Validate(s)
}

type GetApplicationResponseBodyApplication struct {
	// The status of the Developer API feature. Valid values:
	//
	// 	- Enabled: The Developer API feature is enabled.
	//
	// 	- Disabled: The Developer API feature is disabled.
	//
	// example:
	//
	// disabled
	ApiInvokeStatus *string `json:"ApiInvokeStatus,omitempty" xml:"ApiInvokeStatus,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
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
	// urn:alibaba:idaas:app:source:template
	ApplicationSourceType *string `json:"ApplicationSourceType,omitempty" xml:"ApplicationSourceType,omitempty"`
	// The ID of the template based on which the application is created. This parameter is returned only if the application is created based on a template.
	//
	// example:
	//
	// apt_rpa_tdsxxx
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// Application visibility
	ApplicationVisibility []*string `json:"ApplicationVisibility,omitempty" xml:"ApplicationVisibility,omitempty" type:"Repeated"`
	// The authorization type of the EIAM application. Valid values:
	//
	// 	- authorize_required: Only the user with explicit authorization can access the application.
	//
	// 	- default_all: By default, all users can access the application.
	//
	// example:
	//
	// authorize_required
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
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
	CreateTime          *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomSubjectStatus *string `json:"CustomSubjectStatus,omitempty" xml:"CustomSubjectStatus,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// The application is applicable to the test environment.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The features that are supported by the application. The value is a JSON array. Valid values:
	//
	// 	- sso: The application supports SSO.
	//
	// 	- provision: The application supports account synchronization.
	//
	// 	- api_invoke: The application supports custom APIs.
	//
	// example:
	//
	// ["sso", "provision"]
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
	// M2M client status.
	//
	// example:
	//
	// enabled
	M2MClientStatus *string `json:"M2MClientStatus,omitempty" xml:"M2MClientStatus,omitempty"`
	// The service code of the cloud service that manages the application template.
	//
	// example:
	//
	// rpa
	ManagedServiceCode *string `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	// Unique identifier of the resource server
	//
	// example:
	//
	// https://www.example.com
	ResourceServerIdentifier *string `json:"ResourceServerIdentifier,omitempty" xml:"ResourceServerIdentifier,omitempty"`
	// Resource server status.
	//
	// example:
	//
	// disabled	enabled
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
	// example:
	//
	// saml2
	SsoType *string `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- Enabled: The application is enabled.
	//
	// 	- Disabled: The application is disabled.
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

func (s GetApplicationResponseBodyApplication) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplication) GetApiInvokeStatus() *string {
	return s.ApiInvokeStatus
}

func (s *GetApplicationResponseBodyApplication) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationResponseBodyApplication) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *GetApplicationResponseBodyApplication) GetApplicationSourceType() *string {
	return s.ApplicationSourceType
}

func (s *GetApplicationResponseBodyApplication) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *GetApplicationResponseBodyApplication) GetApplicationVisibility() []*string {
	return s.ApplicationVisibility
}

func (s *GetApplicationResponseBodyApplication) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *GetApplicationResponseBodyApplication) GetClientId() *string {
	return s.ClientId
}

func (s *GetApplicationResponseBodyApplication) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetApplicationResponseBodyApplication) GetCustomSubjectStatus() *string {
	return s.CustomSubjectStatus
}

func (s *GetApplicationResponseBodyApplication) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyApplication) GetFeatures() *string {
	return s.Features
}

func (s *GetApplicationResponseBodyApplication) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationResponseBodyApplication) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *GetApplicationResponseBodyApplication) GetM2MClientStatus() *string {
	return s.M2MClientStatus
}

func (s *GetApplicationResponseBodyApplication) GetManagedServiceCode() *string {
	return s.ManagedServiceCode
}

func (s *GetApplicationResponseBodyApplication) GetResourceServerIdentifier() *string {
	return s.ResourceServerIdentifier
}

func (s *GetApplicationResponseBodyApplication) GetResourceServerStatus() *string {
	return s.ResourceServerStatus
}

func (s *GetApplicationResponseBodyApplication) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *GetApplicationResponseBodyApplication) GetSsoType() *string {
	return s.SsoType
}

func (s *GetApplicationResponseBodyApplication) GetStatus() *string {
	return s.Status
}

func (s *GetApplicationResponseBodyApplication) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetApplicationResponseBodyApplication) SetApiInvokeStatus(v string) *GetApplicationResponseBodyApplication {
	s.ApiInvokeStatus = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationId(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationName(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationSourceType(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationSourceType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationTemplateId(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationTemplateId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationVisibility(v []*string) *GetApplicationResponseBodyApplication {
	s.ApplicationVisibility = v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAuthorizationType(v string) *GetApplicationResponseBodyApplication {
	s.AuthorizationType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetClientId(v string) *GetApplicationResponseBodyApplication {
	s.ClientId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCreateTime(v int64) *GetApplicationResponseBodyApplication {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCustomSubjectStatus(v string) *GetApplicationResponseBodyApplication {
	s.CustomSubjectStatus = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDescription(v string) *GetApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetFeatures(v string) *GetApplicationResponseBodyApplication {
	s.Features = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetInstanceId(v string) *GetApplicationResponseBodyApplication {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetLogoUrl(v string) *GetApplicationResponseBodyApplication {
	s.LogoUrl = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetM2MClientStatus(v string) *GetApplicationResponseBodyApplication {
	s.M2MClientStatus = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetManagedServiceCode(v string) *GetApplicationResponseBodyApplication {
	s.ManagedServiceCode = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetResourceServerIdentifier(v string) *GetApplicationResponseBodyApplication {
	s.ResourceServerIdentifier = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetResourceServerStatus(v string) *GetApplicationResponseBodyApplication {
	s.ResourceServerStatus = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetServiceManaged(v bool) *GetApplicationResponseBodyApplication {
	s.ServiceManaged = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetSsoType(v string) *GetApplicationResponseBodyApplication {
	s.SsoType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetStatus(v string) *GetApplicationResponseBodyApplication {
	s.Status = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetUpdateTime(v int64) *GetApplicationResponseBodyApplication {
	s.UpdateTime = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) Validate() error {
	return dara.Validate(s)
}
