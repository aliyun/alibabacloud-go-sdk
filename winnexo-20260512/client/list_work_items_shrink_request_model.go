// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListWorkItemsShrinkRequest
	GetAppToken() *string
	SetBusinessStatusesShrink(v string) *ListWorkItemsShrinkRequest
	GetBusinessStatusesShrink() *string
	SetCollaborationGroupId(v string) *ListWorkItemsShrinkRequest
	GetCollaborationGroupId() *string
	SetItemTypesShrink(v string) *ListWorkItemsShrinkRequest
	GetItemTypesShrink() *string
	SetKeyword(v string) *ListWorkItemsShrinkRequest
	GetKeyword() *string
	SetMaxResults(v int64) *ListWorkItemsShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListWorkItemsShrinkRequest
	GetNextToken() *string
	SetOpsUserCode(v string) *ListWorkItemsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListWorkItemsShrinkRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListWorkItemsShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListWorkItemsShrinkRequest
	GetPageSize() *int64
	SetPrioritiesShrink(v string) *ListWorkItemsShrinkRequest
	GetPrioritiesShrink() *string
	SetScope(v string) *ListWorkItemsShrinkRequest
	GetScope() *string
	SetTenantId(v string) *ListWorkItemsShrinkRequest
	GetTenantId() *string
}

type ListWorkItemsShrinkRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The business status filter, which supports seven states.
	//
	// example:
	//
	// ["PROCESSING"]
	BusinessStatusesShrink *string `json:"businessStatuses,omitempty" xml:"businessStatuses,omitempty"`
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The work item type filter.
	//
	// example:
	//
	// ["WORKFLOW"]
	ItemTypesShrink *string `json:"itemTypes,omitempty" xml:"itemTypes,omitempty"`
	// The keyword for fuzzy search in the title or description.
	//
	// example:
	//
	// example
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of entries to return per page. Valid values: 1 to 100. This parameter takes precedence over pageSize if both are specified.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token, which is obtained from the nextToken value returned in the previous response. This parameter takes precedence over page if both are specified.
	//
	// example:
	//
	// example
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The page number. This parameter is ignored if nextToken is specified.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The priority filter.
	//
	// example:
	//
	// ["HIGH"]
	PrioritiesShrink *string `json:"priorities,omitempty" xml:"priorities,omitempty"`
	// The personal scope filter.
	//
	// example:
	//
	// ASSIGNED_TO_ME
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListWorkItemsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWorkItemsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListWorkItemsShrinkRequest) GetBusinessStatusesShrink() *string {
	return s.BusinessStatusesShrink
}

func (s *ListWorkItemsShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkItemsShrinkRequest) GetItemTypesShrink() *string {
	return s.ItemTypesShrink
}

func (s *ListWorkItemsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListWorkItemsShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListWorkItemsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkItemsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListWorkItemsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListWorkItemsShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkItemsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkItemsShrinkRequest) GetPrioritiesShrink() *string {
	return s.PrioritiesShrink
}

func (s *ListWorkItemsShrinkRequest) GetScope() *string {
	return s.Scope
}

func (s *ListWorkItemsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkItemsShrinkRequest) SetAppToken(v string) *ListWorkItemsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetBusinessStatusesShrink(v string) *ListWorkItemsShrinkRequest {
	s.BusinessStatusesShrink = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetCollaborationGroupId(v string) *ListWorkItemsShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetItemTypesShrink(v string) *ListWorkItemsShrinkRequest {
	s.ItemTypesShrink = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetKeyword(v string) *ListWorkItemsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetMaxResults(v int64) *ListWorkItemsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetNextToken(v string) *ListWorkItemsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetOpsUserCode(v string) *ListWorkItemsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetOpsUserId(v string) *ListWorkItemsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetPage(v int64) *ListWorkItemsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetPageSize(v int64) *ListWorkItemsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetPrioritiesShrink(v string) *ListWorkItemsShrinkRequest {
	s.PrioritiesShrink = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetScope(v string) *ListWorkItemsShrinkRequest {
	s.Scope = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) SetTenantId(v string) *ListWorkItemsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkItemsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
