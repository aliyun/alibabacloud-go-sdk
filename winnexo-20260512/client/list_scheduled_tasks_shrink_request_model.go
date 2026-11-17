// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListScheduledTasksShrinkRequest
	GetCollaborationGroupId() *string
	SetCreatorOnly(v bool) *ListScheduledTasksShrinkRequest
	GetCreatorOnly() *bool
	SetKeyword(v string) *ListScheduledTasksShrinkRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListScheduledTasksShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListScheduledTasksShrinkRequest
	GetNextToken() *string
	SetPage(v int64) *ListScheduledTasksShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListScheduledTasksShrinkRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListScheduledTasksShrinkRequest
	GetTenantId() *string
	SetVisibilitiesShrink(v string) *ListScheduledTasksShrinkRequest
	GetVisibilitiesShrink() *string
}

type ListScheduledTasksShrinkRequest struct {
	// The ID of the collaboration group to which the task belongs, such as cg_101. If this parameter is specified, a group space task is created, and the caller must be a valid group member. If this parameter is left empty, a personal task is created.
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// Specifies whether to view only tasks created by yourself. This parameter takes effect only at the group level (at the personal level, tasks are always your own). If this parameter is not specified, no filtering is applied.
	//
	// example:
	//
	// true
	CreatorOnly *bool `json:"creatorOnly,omitempty" xml:"creatorOnly,omitempty"`
	// The keyword in the rule name for fuzzy matching.
	//
	// example:
	//
	// Sample keyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of records to return for this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token used to retrieve the next page.
	//
	// example:
	//
	// eHiB8vca1XDyBT0cNAmThA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number for pagination. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries to return per page.
	//
	// > The maximum number of entries per page is 30.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// Filters tasks by visibility scope. Multiple values can be selected. Valid values: PRIVATE (visible only to the creator and group owner), COLLABORATIVE (visible to specified collaboration members), and PUBLIC (visible to all group members). If this parameter is not specified or an empty list is passed, no filtering is applied. This parameter takes effect only at the group level (when collaborationGroupId is specified) and is ignored at the personal level.
	//
	// example:
	//
	// ["PRIVATE"]
	VisibilitiesShrink *string `json:"visibilities,omitempty" xml:"visibilities,omitempty"`
}

func (s ListScheduledTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListScheduledTasksShrinkRequest) GetCreatorOnly() *bool {
	return s.CreatorOnly
}

func (s *ListScheduledTasksShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListScheduledTasksShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListScheduledTasksShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListScheduledTasksShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListScheduledTasksShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListScheduledTasksShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListScheduledTasksShrinkRequest) GetVisibilitiesShrink() *string {
	return s.VisibilitiesShrink
}

func (s *ListScheduledTasksShrinkRequest) SetCollaborationGroupId(v string) *ListScheduledTasksShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListScheduledTasksShrinkRequest) SetCreatorOnly(v bool) *ListScheduledTasksShrinkRequest {
	s.CreatorOnly = &v
	return s
}

func (s *ListScheduledTasksShrinkRequest) SetKeyword(v string) *ListScheduledTasksShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListScheduledTasksShrinkRequest) SetMaxResults(v int32) *ListScheduledTasksShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListScheduledTasksShrinkRequest) SetNextToken(v string) *ListScheduledTasksShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListScheduledTasksShrinkRequest) SetPage(v int64) *ListScheduledTasksShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListScheduledTasksShrinkRequest) SetPageSize(v int64) *ListScheduledTasksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduledTasksShrinkRequest) SetTenantId(v string) *ListScheduledTasksShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListScheduledTasksShrinkRequest) SetVisibilitiesShrink(v string) *ListScheduledTasksShrinkRequest {
	s.VisibilitiesShrink = &v
	return s
}

func (s *ListScheduledTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
