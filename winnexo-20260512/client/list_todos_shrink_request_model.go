// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTodosShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListTodosShrinkRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ListTodosShrinkRequest
	GetCollaborationGroupId() *string
	SetKeyword(v string) *ListTodosShrinkRequest
	GetKeyword() *string
	SetMaxResults(v int64) *ListTodosShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListTodosShrinkRequest
	GetNextToken() *string
	SetOpsUserCode(v string) *ListTodosShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListTodosShrinkRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListTodosShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListTodosShrinkRequest
	GetPageSize() *int64
	SetPrioritiesShrink(v string) *ListTodosShrinkRequest
	GetPrioritiesShrink() *string
	SetReadStatusShrink(v string) *ListTodosShrinkRequest
	GetReadStatusShrink() *string
	SetSortBy(v string) *ListTodosShrinkRequest
	GetSortBy() *string
	SetSourceTypesShrink(v string) *ListTodosShrinkRequest
	GetSourceTypesShrink() *string
	SetStatusesShrink(v string) *ListTodosShrinkRequest
	GetStatusesShrink() *string
	SetTenantId(v string) *ListTodosShrinkRequest
	GetTenantId() *string
}

type ListTodosShrinkRequest struct {
	// The platform application token issued by the console. If specified, the operation is performed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration workspace ID. This is an optional filter. If not specified, all workspaces are included.
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The keyword for fuzzy search in the title or summary.
	//
	// example:
	//
	// Approval
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of entries to return per page. Valid values: 1 to 100. This parameter takes precedence over pageSize if both are specified.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Use the nextToken value returned in the previous response. This parameter takes precedence over page if both are specified.
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
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The priority filter. Valid values: HIGH, MEDIUM, LOW.
	//
	// example:
	//
	// ["HIGH","MEDIUM"]
	PrioritiesShrink *string `json:"priorities,omitempty" xml:"priorities,omitempty"`
	// The read status filter. Valid values: UNREAD, READ. The status is determined by the readAt time of the assignee.
	//
	// example:
	//
	// ["UNREAD","READ"]
	ReadStatusShrink *string `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
	// The sorting method.
	//
	// example:
	//
	// DEFAULT
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The to-do type filter. Valid values: WORKFLOW, OPEN_API.
	//
	// example:
	//
	// ["WORKFLOW","OPEN_API"]
	SourceTypesShrink *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// The to-do status filter. Valid values: PENDING, COMPLETED, CANCELLED.
	//
	// example:
	//
	// ["PENDING","COMPLETED"]
	StatusesShrink *string `json:"statuses,omitempty" xml:"statuses,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListTodosShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTodosShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTodosShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListTodosShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListTodosShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTodosShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListTodosShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTodosShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListTodosShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListTodosShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListTodosShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTodosShrinkRequest) GetPrioritiesShrink() *string {
	return s.PrioritiesShrink
}

func (s *ListTodosShrinkRequest) GetReadStatusShrink() *string {
	return s.ReadStatusShrink
}

func (s *ListTodosShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTodosShrinkRequest) GetSourceTypesShrink() *string {
	return s.SourceTypesShrink
}

func (s *ListTodosShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListTodosShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListTodosShrinkRequest) SetAppToken(v string) *ListTodosShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ListTodosShrinkRequest) SetCollaborationGroupId(v string) *ListTodosShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListTodosShrinkRequest) SetKeyword(v string) *ListTodosShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListTodosShrinkRequest) SetMaxResults(v int64) *ListTodosShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTodosShrinkRequest) SetNextToken(v string) *ListTodosShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTodosShrinkRequest) SetOpsUserCode(v string) *ListTodosShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListTodosShrinkRequest) SetOpsUserId(v string) *ListTodosShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListTodosShrinkRequest) SetPage(v int64) *ListTodosShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListTodosShrinkRequest) SetPageSize(v int64) *ListTodosShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTodosShrinkRequest) SetPrioritiesShrink(v string) *ListTodosShrinkRequest {
	s.PrioritiesShrink = &v
	return s
}

func (s *ListTodosShrinkRequest) SetReadStatusShrink(v string) *ListTodosShrinkRequest {
	s.ReadStatusShrink = &v
	return s
}

func (s *ListTodosShrinkRequest) SetSortBy(v string) *ListTodosShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTodosShrinkRequest) SetSourceTypesShrink(v string) *ListTodosShrinkRequest {
	s.SourceTypesShrink = &v
	return s
}

func (s *ListTodosShrinkRequest) SetStatusesShrink(v string) *ListTodosShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListTodosShrinkRequest) SetTenantId(v string) *ListTodosShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListTodosShrinkRequest) Validate() error {
	return dara.Validate(s)
}
