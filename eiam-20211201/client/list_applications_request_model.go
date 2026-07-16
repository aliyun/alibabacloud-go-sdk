// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCreationType(v string) *ListApplicationsRequest
	GetApplicationCreationType() *string
	SetApplicationIdentityType(v string) *ListApplicationsRequest
	GetApplicationIdentityType() *string
	SetApplicationIds(v []*string) *ListApplicationsRequest
	GetApplicationIds() []*string
	SetApplicationName(v string) *ListApplicationsRequest
	GetApplicationName() *string
	SetApplicationTemplateId(v string) *ListApplicationsRequest
	GetApplicationTemplateId() *string
	SetAuthorizationType(v string) *ListApplicationsRequest
	GetAuthorizationType() *string
	SetCustomFields(v []*ListApplicationsRequestCustomFields) *ListApplicationsRequest
	GetCustomFields() []*ListApplicationsRequestCustomFields
	SetInstanceId(v string) *ListApplicationsRequest
	GetInstanceId() *string
	SetM2MClientStatus(v string) *ListApplicationsRequest
	GetM2MClientStatus() *string
	SetManagedServiceCode(v string) *ListApplicationsRequest
	GetManagedServiceCode() *string
	SetPageNumber(v int64) *ListApplicationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListApplicationsRequest
	GetPageSize() *int64
	SetResourceServerStatus(v string) *ListApplicationsRequest
	GetResourceServerStatus() *string
	SetServiceManaged(v bool) *ListApplicationsRequest
	GetServiceManaged() *bool
	SetSsoType(v string) *ListApplicationsRequest
	GetSsoType() *string
	SetStatus(v string) *ListApplicationsRequest
	GetStatus() *string
}

type ListApplicationsRequest struct {
	// The application creation type. If this parameter is left empty, applications of the user_custom type are queried by default. To query applications of all types, set this parameter to all.
	//
	// example:
	//
	// system_init
	ApplicationCreationType *string `json:"ApplicationCreationType,omitempty" xml:"ApplicationCreationType,omitempty"`
	// The application identity type. If this parameter is left empty, applications of the application type are queried by default. To query applications of all identity types, set this parameter to all.
	//
	// example:
	//
	// application
	ApplicationIdentityType *string `json:"ApplicationIdentityType,omitempty" xml:"ApplicationIdentityType,omitempty"`
	// The list of application IDs.
	//
	// example:
	//
	// Ram Account SSO
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The application name. Only left fuzzy match is supported.
	//
	// example:
	//
	// Ram Account SSO
	ApplicationName       *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The application access authorization type. Valid values:
	//
	// - authorize_required: Explicit authorization is required for access.
	//
	// - default_all: All members have access permissions by default.
	//
	// example:
	//
	// authorize_required
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The list of custom fields.
	CustomFields []*ListApplicationsRequestCustomFields `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the M2M Client identity is enabled.
	//
	// example:
	//
	// enabled
	M2MClientStatus    *string `json:"M2MClientStatus,omitempty" xml:"M2MClientStatus,omitempty"`
	ManagedServiceCode *string `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the ResourceServer capability is enabled.
	//
	// example:
	//
	// enabled
	ResourceServerStatus *string `json:"ResourceServerStatus,omitempty" xml:"ResourceServerStatus,omitempty"`
	ServiceManaged       *bool   `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The SSO type filter condition. Multiple types can be separated by commas, such as oauth2/m2m,oidc+oauth2/m2m.
	//
	// example:
	//
	// oauth2/m2m
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
}

func (s ListApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) GetApplicationCreationType() *string {
	return s.ApplicationCreationType
}

func (s *ListApplicationsRequest) GetApplicationIdentityType() *string {
	return s.ApplicationIdentityType
}

func (s *ListApplicationsRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListApplicationsRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationsRequest) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *ListApplicationsRequest) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *ListApplicationsRequest) GetCustomFields() []*ListApplicationsRequestCustomFields {
	return s.CustomFields
}

func (s *ListApplicationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsRequest) GetM2MClientStatus() *string {
	return s.M2MClientStatus
}

func (s *ListApplicationsRequest) GetManagedServiceCode() *string {
	return s.ManagedServiceCode
}

func (s *ListApplicationsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListApplicationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationsRequest) GetResourceServerStatus() *string {
	return s.ResourceServerStatus
}

func (s *ListApplicationsRequest) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListApplicationsRequest) GetSsoType() *string {
	return s.SsoType
}

func (s *ListApplicationsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationsRequest) SetApplicationCreationType(v string) *ListApplicationsRequest {
	s.ApplicationCreationType = &v
	return s
}

func (s *ListApplicationsRequest) SetApplicationIdentityType(v string) *ListApplicationsRequest {
	s.ApplicationIdentityType = &v
	return s
}

func (s *ListApplicationsRequest) SetApplicationIds(v []*string) *ListApplicationsRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListApplicationsRequest) SetApplicationName(v string) *ListApplicationsRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationsRequest) SetApplicationTemplateId(v string) *ListApplicationsRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *ListApplicationsRequest) SetAuthorizationType(v string) *ListApplicationsRequest {
	s.AuthorizationType = &v
	return s
}

func (s *ListApplicationsRequest) SetCustomFields(v []*ListApplicationsRequestCustomFields) *ListApplicationsRequest {
	s.CustomFields = v
	return s
}

func (s *ListApplicationsRequest) SetInstanceId(v string) *ListApplicationsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsRequest) SetM2MClientStatus(v string) *ListApplicationsRequest {
	s.M2MClientStatus = &v
	return s
}

func (s *ListApplicationsRequest) SetManagedServiceCode(v string) *ListApplicationsRequest {
	s.ManagedServiceCode = &v
	return s
}

func (s *ListApplicationsRequest) SetPageNumber(v int64) *ListApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsRequest) SetPageSize(v int64) *ListApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsRequest) SetResourceServerStatus(v string) *ListApplicationsRequest {
	s.ResourceServerStatus = &v
	return s
}

func (s *ListApplicationsRequest) SetServiceManaged(v bool) *ListApplicationsRequest {
	s.ServiceManaged = &v
	return s
}

func (s *ListApplicationsRequest) SetSsoType(v string) *ListApplicationsRequest {
	s.SsoType = &v
	return s
}

func (s *ListApplicationsRequest) SetStatus(v string) *ListApplicationsRequest {
	s.Status = &v
	return s
}

func (s *ListApplicationsRequest) Validate() error {
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

type ListApplicationsRequestCustomFields struct {
	// The custom field identifier. Valid values:
	//
	// - agent_type: the agent type.
	//
	// example:
	//
	// agent_type
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The custom field value.
	//
	// example:
	//
	// x-claw
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
}

func (s ListApplicationsRequestCustomFields) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsRequestCustomFields) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequestCustomFields) GetFieldName() *string {
	return s.FieldName
}

func (s *ListApplicationsRequestCustomFields) GetFieldValue() *string {
	return s.FieldValue
}

func (s *ListApplicationsRequestCustomFields) SetFieldName(v string) *ListApplicationsRequestCustomFields {
	s.FieldName = &v
	return s
}

func (s *ListApplicationsRequestCustomFields) SetFieldValue(v string) *ListApplicationsRequestCustomFields {
	s.FieldValue = &v
	return s
}

func (s *ListApplicationsRequestCustomFields) Validate() error {
	return dara.Validate(s)
}
