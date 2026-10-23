// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatuses(v []*string) *ListWorkItemsRequest
	GetBusinessStatuses() []*string
	SetCollaborationGroupId(v string) *ListWorkItemsRequest
	GetCollaborationGroupId() *string
	SetItemTypes(v []*string) *ListWorkItemsRequest
	GetItemTypes() []*string
	SetKeyword(v string) *ListWorkItemsRequest
	GetKeyword() *string
	SetMaxResults(v int64) *ListWorkItemsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListWorkItemsRequest
	GetNextToken() *string
	SetPage(v int64) *ListWorkItemsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListWorkItemsRequest
	GetPageSize() *int64
	SetPriorities(v []*string) *ListWorkItemsRequest
	GetPriorities() []*string
	SetScope(v string) *ListWorkItemsRequest
	GetScope() *string
	SetTenantId(v string) *ListWorkItemsRequest
	GetTenantId() *string
}

type ListWorkItemsRequest struct {
	// The business statuses used for filtering.
	//
	// example:
	//
	// example
	BusinessStatuses []*string `json:"businessStatuses,omitempty" xml:"businessStatuses,omitempty" type:"Repeated"`
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The work item types used for filtering.
	//
	// example:
	//
	// example
	ItemTypes []*string `json:"itemTypes,omitempty" xml:"itemTypes,omitempty" type:"Repeated"`
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
	// The pagination token. Obtain the token from the nextToken parameter in the previous response. This parameter takes precedence over page if both are specified.
	//
	// example:
	//
	// example
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// The priorities used for filtering.
	//
	// example:
	//
	// example
	Priorities []*string `json:"priorities,omitempty" xml:"priorities,omitempty" type:"Repeated"`
	// The personal scope used for filtering.
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

func (s ListWorkItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkItemsRequest) GetBusinessStatuses() []*string {
	return s.BusinessStatuses
}

func (s *ListWorkItemsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkItemsRequest) GetItemTypes() []*string {
	return s.ItemTypes
}

func (s *ListWorkItemsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListWorkItemsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListWorkItemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkItemsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkItemsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkItemsRequest) GetPriorities() []*string {
	return s.Priorities
}

func (s *ListWorkItemsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListWorkItemsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkItemsRequest) SetBusinessStatuses(v []*string) *ListWorkItemsRequest {
	s.BusinessStatuses = v
	return s
}

func (s *ListWorkItemsRequest) SetCollaborationGroupId(v string) *ListWorkItemsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkItemsRequest) SetItemTypes(v []*string) *ListWorkItemsRequest {
	s.ItemTypes = v
	return s
}

func (s *ListWorkItemsRequest) SetKeyword(v string) *ListWorkItemsRequest {
	s.Keyword = &v
	return s
}

func (s *ListWorkItemsRequest) SetMaxResults(v int64) *ListWorkItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkItemsRequest) SetNextToken(v string) *ListWorkItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkItemsRequest) SetPage(v int64) *ListWorkItemsRequest {
	s.Page = &v
	return s
}

func (s *ListWorkItemsRequest) SetPageSize(v int64) *ListWorkItemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkItemsRequest) SetPriorities(v []*string) *ListWorkItemsRequest {
	s.Priorities = v
	return s
}

func (s *ListWorkItemsRequest) SetScope(v string) *ListWorkItemsRequest {
	s.Scope = &v
	return s
}

func (s *ListWorkItemsRequest) SetTenantId(v string) *ListWorkItemsRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkItemsRequest) Validate() error {
	return dara.Validate(s)
}
