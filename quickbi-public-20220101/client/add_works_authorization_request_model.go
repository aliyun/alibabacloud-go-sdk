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
	// The type of granted permissions. Valid values:
	//
	// - 1: View.
	//
	// - 3: View and export.
	//
	// - 11: Edit, view, and export.
	//
	// Note: When the value is set to 11, the authorization does not expire.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthPoints *int32 `json:"AuthPoints,omitempty" xml:"AuthPoints,omitempty"`
	// The authorization scope. Valid values:
	//
	// - 0: User. The AuthorizedId parameter corresponds to the user ID in Quick BI.
	//
	// - 1: User group. The AuthorizedId parameter corresponds to the user group ID in Quick BI.
	//
	// - 2: All organization members. The AuthorizedId parameter corresponds to the organization ID in Quick BI.
	//
	// - 3: All workspace members. The AuthorizedId parameter corresponds to the workspace ID in Quick BI.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AuthorizeScope *int32 `json:"AuthorizeScope,omitempty" xml:"AuthorizeScope,omitempty"`
	// The ID of the authorized object. For more information, see the description of AuthorizeScope.
	//
	// This parameter is required.
	//
	// example:
	//
	// ASDAS-WDAS****ASDA
	AuthorizedId *string `json:"AuthorizedId,omitempty" xml:"AuthorizedId,omitempty"`
	// The expiration date of the authorization. Format: yyyy-MM-dd. Example: 2022-02-16. Note: The authorization period must be at least N+1 days. This parameter is required when the AuthPoints parameter is not set to 11.
	//
	// example:
	//
	// 2099-12-31
	ExpireDay *string `json:"ExpireDay,omitempty" xml:"ExpireDay,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// al*************7ufv
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - dashboard: Dashboard.
	//
	// - report: Workbook.
	//
	// - dashboardOfflineQuery: Ad hoc query.
	//
	// - cube: Dataset.
	//
	// - datasource: Data source.
	//
	// - screen: Data screen.
	//
	// - ANALYSIS: Ad hoc analysis.
	//
	// - dataForm: Data entry.
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
