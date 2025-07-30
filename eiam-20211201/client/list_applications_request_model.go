// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *ListApplicationsRequest
	GetApplicationIds() []*string
	SetApplicationName(v string) *ListApplicationsRequest
	GetApplicationName() *string
	SetAuthorizationType(v string) *ListApplicationsRequest
	GetAuthorizationType() *string
	SetInstanceId(v string) *ListApplicationsRequest
	GetInstanceId() *string
	SetM2MClientStatus(v string) *ListApplicationsRequest
	GetM2MClientStatus() *string
	SetPageNumber(v int64) *ListApplicationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListApplicationsRequest
	GetPageSize() *int64
	SetResourceServerStatus(v string) *ListApplicationsRequest
	GetResourceServerStatus() *string
	SetSsoType(v string) *ListApplicationsRequest
	GetSsoType() *string
	SetStatus(v string) *ListApplicationsRequest
	GetStatus() *string
}

type ListApplicationsRequest struct {
	// The IDs of the applications.
	//
	// example:
	//
	// Ram Account SSO
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The name of the application. Only fuzzy match from the leftmost character is supported.
	//
	// example:
	//
	// Ram Account SSO
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The authorization of the application. Valid values:
	//
	// 	- authorize_required: Only the user with explicit authorization can access the application.
	//
	// 	- default_all: By default, all users can access the application.
	//
	// example:
	//
	// authorize_required
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Used to determine whether M2M client identity is enabled.
	//
	// - enabled
	//
	// - disabled
	//
	// example:
	//
	// enabled
	M2MClientStatus *string `json:"M2MClientStatus,omitempty" xml:"M2MClientStatus,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Used to determine whether the ResourceServer capability is enabled.
	//
	// - enabled
	//
	// - disabled
	//
	// example:
	//
	// enabled
	ResourceServerStatus *string `json:"ResourceServerStatus,omitempty" xml:"ResourceServerStatus,omitempty"`
	// SSO type.
	//
	// - oidc
	//
	// - saml2
	//
	// - oauth2/m2m
	//
	// example:
	//
	// oauth2/m2m
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
}

func (s ListApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListApplicationsRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationsRequest) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *ListApplicationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsRequest) GetM2MClientStatus() *string {
	return s.M2MClientStatus
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

func (s *ListApplicationsRequest) GetSsoType() *string {
	return s.SsoType
}

func (s *ListApplicationsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationsRequest) SetApplicationIds(v []*string) *ListApplicationsRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListApplicationsRequest) SetApplicationName(v string) *ListApplicationsRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationsRequest) SetAuthorizationType(v string) *ListApplicationsRequest {
	s.AuthorizationType = &v
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

func (s *ListApplicationsRequest) SetSsoType(v string) *ListApplicationsRequest {
	s.SsoType = &v
	return s
}

func (s *ListApplicationsRequest) SetStatus(v string) *ListApplicationsRequest {
	s.Status = &v
	return s
}

func (s *ListApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
