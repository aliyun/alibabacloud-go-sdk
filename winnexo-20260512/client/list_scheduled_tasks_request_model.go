// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListScheduledTasksRequest
	GetCollaborationGroupId() *string
	SetCreatorOnly(v bool) *ListScheduledTasksRequest
	GetCreatorOnly() *bool
	SetKeyword(v string) *ListScheduledTasksRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListScheduledTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListScheduledTasksRequest
	GetNextToken() *string
	SetPage(v int64) *ListScheduledTasksRequest
	GetPage() *int64
	SetPageSize(v int64) *ListScheduledTasksRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListScheduledTasksRequest
	GetTenantId() *string
	SetVisibilities(v []*string) *ListScheduledTasksRequest
	GetVisibilities() []*string
}

type ListScheduledTasksRequest struct {
	// The ID of the collaboration group (such as cg_101). If specified, a group task is created (the caller must be a valid group member). If left empty, a personal task is created.
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// Specifies whether to return only tasks created by the caller. This parameter takes effect only in the group dimension (in the personal dimension, only the caller\\"s own tasks are returned). If not specified, no filtering is applied.
	//
	// example:
	//
	// true
	CreatorOnly *bool `json:"creatorOnly,omitempty" xml:"creatorOnly,omitempty"`
	// The keyword of the rule name for fuzzy match.
	//
	// example:
	//
	// SampleKeyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eHiB8vca1XDyBT0cNAmThA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// > The maximum number of entries per page is 30.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID that takes effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// Filters by visibility. Valid values:
	//
	// - PRIVATE: visible only to the creator and group owner.
	//
	// - COLLABORATIVE: visible to specified collaborators.
	//
	// - PUBLIC: visible to all group members.
	//
	// If not specified or an empty list is passed, no filtering is applied. This parameter takes effect only in the group dimension (when collaborationGroupId is specified) and is ignored in the personal dimension.
	//
	// example:
	//
	// PRIVATE
	Visibilities []*string `json:"visibilities,omitempty" xml:"visibilities,omitempty" type:"Repeated"`
}

func (s ListScheduledTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListScheduledTasksRequest) GetCreatorOnly() *bool {
	return s.CreatorOnly
}

func (s *ListScheduledTasksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListScheduledTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListScheduledTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListScheduledTasksRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListScheduledTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListScheduledTasksRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListScheduledTasksRequest) GetVisibilities() []*string {
	return s.Visibilities
}

func (s *ListScheduledTasksRequest) SetCollaborationGroupId(v string) *ListScheduledTasksRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListScheduledTasksRequest) SetCreatorOnly(v bool) *ListScheduledTasksRequest {
	s.CreatorOnly = &v
	return s
}

func (s *ListScheduledTasksRequest) SetKeyword(v string) *ListScheduledTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListScheduledTasksRequest) SetMaxResults(v int32) *ListScheduledTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListScheduledTasksRequest) SetNextToken(v string) *ListScheduledTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListScheduledTasksRequest) SetPage(v int64) *ListScheduledTasksRequest {
	s.Page = &v
	return s
}

func (s *ListScheduledTasksRequest) SetPageSize(v int64) *ListScheduledTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduledTasksRequest) SetTenantId(v string) *ListScheduledTasksRequest {
	s.TenantId = &v
	return s
}

func (s *ListScheduledTasksRequest) SetVisibilities(v []*string) *ListScheduledTasksRequest {
	s.Visibilities = v
	return s
}

func (s *ListScheduledTasksRequest) Validate() error {
	return dara.Validate(s)
}
