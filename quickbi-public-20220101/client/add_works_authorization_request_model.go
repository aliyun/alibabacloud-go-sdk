// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorksAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthPoints(v int32) *AddWorksAuthorizationRequest
	GetAuthPoints() *int32
	SetAuthorizeScope(v int32) *AddWorksAuthorizationRequest
	GetAuthorizeScope() *int32
	SetAuthorizedId(v string) *AddWorksAuthorizationRequest
	GetAuthorizedId() *string
	SetExpireDay(v string) *AddWorksAuthorizationRequest
	GetExpireDay() *string
	SetResourceId(v string) *AddWorksAuthorizationRequest
	GetResourceId() *string
	SetResourceType(v string) *AddWorksAuthorizationRequest
	GetResourceType() *string
}

type AddWorksAuthorizationRequest struct {
	// The permissions to grant. Valid values:
	//
	// `1`: View
	//
	// `3`: View and Export
	//
	// `11`: Edit, View, and Export
	//
	// **Note**: If AuthPoints is set to 11, the authorization is permanent and the ExpireDay parameter is ignored.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthPoints *int32 `json:"AuthPoints,omitempty" xml:"AuthPoints,omitempty"`
	// The type of the principal. Valid values:
	//
	// - `0`: User. Set AuthorizedId to the user ID.
	//
	// - `1`: User group. Set AuthorizedId to the user group ID.
	//
	// - `2`: All members of an organization. Set AuthorizedId to the organization ID.
	//
	// - `3`: All members of a workspace. Set AuthorizedId to the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AuthorizeScope *int32 `json:"AuthorizeScope,omitempty" xml:"AuthorizeScope,omitempty"`
	// The ID of the principal to be authorized. The AuthorizeScope parameter specifies the type of principal.
	//
	// This parameter is required.
	//
	// example:
	//
	// ASDAS-WDAS****ASDA
	AuthorizedId *string `json:"AuthorizedId,omitempty" xml:"AuthorizedId,omitempty"`
	// The expiration date for the permissions.
	//
	// Format: `YYYY-MM-DD`.
	//
	// **Note**: This parameter is required if AuthPoints is not 11. The authorization must be valid for at least one day after the authorization date.
	//
	// example:
	//
	// 2099-12-31
	ExpireDay *string `json:"ExpireDay,omitempty" xml:"ExpireDay,omitempty"`
	// The ID of the work.
	//
	// This parameter is required.
	//
	// example:
	//
	// al*************7ufv
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the work. Valid values:
	//
	// - `dashboard`: A dashboard.
	//
	// - `report`: A report.
	//
	// - `dashboardOfflineQuery`: An ad-hoc query.
	//
	// - `cube`: A dataset.
	//
	// - `datasource`: A data source.
	//
	// - `screen`: A data screen.
	//
	// - `ANALYSIS`: An ad-hoc analysis.
	//
	// - `dataForm`: A data form.
	//
	// This parameter is required.
	//
	// example:
	//
	// dashboard
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s AddWorksAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorksAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *AddWorksAuthorizationRequest) GetAuthPoints() *int32 {
	return s.AuthPoints
}

func (s *AddWorksAuthorizationRequest) GetAuthorizeScope() *int32 {
	return s.AuthorizeScope
}

func (s *AddWorksAuthorizationRequest) GetAuthorizedId() *string {
	return s.AuthorizedId
}

func (s *AddWorksAuthorizationRequest) GetExpireDay() *string {
	return s.ExpireDay
}

func (s *AddWorksAuthorizationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddWorksAuthorizationRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddWorksAuthorizationRequest) SetAuthPoints(v int32) *AddWorksAuthorizationRequest {
	s.AuthPoints = &v
	return s
}

func (s *AddWorksAuthorizationRequest) SetAuthorizeScope(v int32) *AddWorksAuthorizationRequest {
	s.AuthorizeScope = &v
	return s
}

func (s *AddWorksAuthorizationRequest) SetAuthorizedId(v string) *AddWorksAuthorizationRequest {
	s.AuthorizedId = &v
	return s
}

func (s *AddWorksAuthorizationRequest) SetExpireDay(v string) *AddWorksAuthorizationRequest {
	s.ExpireDay = &v
	return s
}

func (s *AddWorksAuthorizationRequest) SetResourceId(v string) *AddWorksAuthorizationRequest {
	s.ResourceId = &v
	return s
}

func (s *AddWorksAuthorizationRequest) SetResourceType(v string) *AddWorksAuthorizationRequest {
	s.ResourceType = &v
	return s
}

func (s *AddWorksAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
