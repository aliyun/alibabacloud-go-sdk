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
	// The returned application information.
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The request ID.
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
	if s.Application != nil {
		if err := s.Application.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyApplication struct {
	// The status of the Developer API feature for the application. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// disabled
	ApiInvokeStatus *string `json:"ApiInvokeStatus,omitempty" xml:"ApiInvokeStatus,omitempty"`
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
	// The application identity type. Valid values:
	//
	// - application: Application.
	//
	// - agent: Agent.
	//
	// example:
	//
	// application
	ApplicationIdentityType *string `json:"ApplicationIdentityType,omitempty" xml:"ApplicationIdentityType,omitempty"`
	// The application name.
	//
	// example:
	//
	// SAML application
	ApplicationName  *string                                                `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ApplicationOwner *GetApplicationResponseBodyApplicationApplicationOwner `json:"ApplicationOwner,omitempty" xml:"ApplicationOwner,omitempty" type:"Struct"`
	// The source from which the application was created. Valid values:
	//
	// - urn:alibaba:idaas:app:source:template: Application template.
	//
	// - urn:alibaba:idaas:app:source:standard: Standard protocol.
	//
	// example:
	//
	// urn:alibaba:idaas:app:source:template
	ApplicationSourceType *string `json:"ApplicationSourceType,omitempty" xml:"ApplicationSourceType,omitempty"`
	// The ID of the application template associated during creation. This value is returned only when the application was created from an application template.
	//
	// example:
	//
	// apt_rpa_tdsxxx
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The application visibility.
	ApplicationVisibility []*string `json:"ApplicationVisibility,omitempty" xml:"ApplicationVisibility,omitempty" type:"Repeated"`
	// The access authorization type of the application. Valid values:
	//
	// - authorize_required: Explicit authorization is required for access.
	//
	// - default_all: All members have access permissions by default.
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
	// The time when the application was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime   *int64                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomFields []*GetApplicationResponseBodyApplicationCustomFields `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	// Indicates whether the custom Subject field in the token is enabled. After this feature is enabled, the issued Access Token changes from \\<clientId\\> to \\<clientId\\>:\\<client.activeSubjectUrn\\>, where client.activeSubjectUrn is configured in the attribute mapping of the federated identity credential of the application.
	//
	// example:
	//
	// enabled
	CustomSubjectStatus *string `json:"CustomSubjectStatus,omitempty" xml:"CustomSubjectStatus,omitempty"`
	// The application description.
	//
	// example:
	//
	// An application for test environment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The features supported by the application, returned as a JSON array string. Valid values:
	//
	// - sso: Single sign-on.
	//
	// - provision: Account synchronization.
	//
	// - api_invoke: API access.
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
	// The M2MClient status.
	//
	// example:
	//
	// enabled
	M2MClientStatus *string `json:"M2MClientStatus,omitempty" xml:"M2MClientStatus,omitempty"`
	// The ServiceCode of the cloud service that manages the application template.
	//
	// example:
	//
	// rpa
	ManagedServiceCode *string `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	// The unique identifier of the ResourceServer, which corresponds to the ResourceServer audience.
	//
	// example:
	//
	// https://www.example.com
	ResourceServerIdentifier *string `json:"ResourceServerIdentifier,omitempty" xml:"ResourceServerIdentifier,omitempty"`
	// The resource server source type.
	//
	// example:
	//
	// urn:cloud:idaas:resourceserver:source:custom
	ResourceServerSourceType *string `json:"ResourceServerSourceType,omitempty" xml:"ResourceServerSourceType,omitempty"`
	// The ResourceServer status.
	//
	// example:
	//
	// enabled
	ResourceServerStatus *string `json:"ResourceServerStatus,omitempty" xml:"ResourceServerStatus,omitempty"`
	// Indicates whether the application template is managed by a cloud service.
	//
	// example:
	//
	// true
	ServiceManaged          *bool     `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	SmartConfigCapabilities []*string `json:"SmartConfigCapabilities,omitempty" xml:"SmartConfigCapabilities,omitempty" type:"Repeated"`
	// The single sign-on protocol. Valid values:
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
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the application was last updated. The value is a UNIX timestamp in milliseconds.
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

func (s *GetApplicationResponseBodyApplication) GetApplicationCreationType() *string {
	return s.ApplicationCreationType
}

func (s *GetApplicationResponseBodyApplication) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationResponseBodyApplication) GetApplicationIdentityType() *string {
	return s.ApplicationIdentityType
}

func (s *GetApplicationResponseBodyApplication) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *GetApplicationResponseBodyApplication) GetApplicationOwner() *GetApplicationResponseBodyApplicationApplicationOwner {
	return s.ApplicationOwner
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

func (s *GetApplicationResponseBodyApplication) GetCustomFields() []*GetApplicationResponseBodyApplicationCustomFields {
	return s.CustomFields
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

func (s *GetApplicationResponseBodyApplication) GetResourceServerSourceType() *string {
	return s.ResourceServerSourceType
}

func (s *GetApplicationResponseBodyApplication) GetResourceServerStatus() *string {
	return s.ResourceServerStatus
}

func (s *GetApplicationResponseBodyApplication) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *GetApplicationResponseBodyApplication) GetSmartConfigCapabilities() []*string {
	return s.SmartConfigCapabilities
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

func (s *GetApplicationResponseBodyApplication) SetApplicationCreationType(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationCreationType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationId(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationIdentityType(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationIdentityType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationName(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationOwner(v *GetApplicationResponseBodyApplicationApplicationOwner) *GetApplicationResponseBodyApplication {
	s.ApplicationOwner = v
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

func (s *GetApplicationResponseBodyApplication) SetCustomFields(v []*GetApplicationResponseBodyApplicationCustomFields) *GetApplicationResponseBodyApplication {
	s.CustomFields = v
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

func (s *GetApplicationResponseBodyApplication) SetResourceServerSourceType(v string) *GetApplicationResponseBodyApplication {
	s.ResourceServerSourceType = &v
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

func (s *GetApplicationResponseBodyApplication) SetSmartConfigCapabilities(v []*string) *GetApplicationResponseBodyApplication {
	s.SmartConfigCapabilities = v
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
	if s.ApplicationOwner != nil {
		if err := s.ApplicationOwner.Validate(); err != nil {
			return err
		}
	}
	if s.CustomFields != nil {
		for _, item := range s.CustomFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationResponseBodyApplicationApplicationOwner struct {
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	UserIds  []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyApplicationApplicationOwner) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationApplicationOwner) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationApplicationOwner) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *GetApplicationResponseBodyApplicationApplicationOwner) GetUserIds() []*string {
	return s.UserIds
}

func (s *GetApplicationResponseBodyApplicationApplicationOwner) SetGroupIds(v []*string) *GetApplicationResponseBodyApplicationApplicationOwner {
	s.GroupIds = v
	return s
}

func (s *GetApplicationResponseBodyApplicationApplicationOwner) SetUserIds(v []*string) *GetApplicationResponseBodyApplicationApplicationOwner {
	s.UserIds = v
	return s
}

func (s *GetApplicationResponseBodyApplicationApplicationOwner) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyApplicationCustomFields struct {
	FieldName  *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
}

func (s GetApplicationResponseBodyApplicationCustomFields) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationCustomFields) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationCustomFields) GetFieldName() *string {
	return s.FieldName
}

func (s *GetApplicationResponseBodyApplicationCustomFields) GetFieldValue() *string {
	return s.FieldValue
}

func (s *GetApplicationResponseBodyApplicationCustomFields) SetFieldName(v string) *GetApplicationResponseBodyApplicationCustomFields {
	s.FieldName = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationCustomFields) SetFieldValue(v string) *GetApplicationResponseBodyApplicationCustomFields {
	s.FieldValue = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationCustomFields) Validate() error {
	return dara.Validate(s)
}
