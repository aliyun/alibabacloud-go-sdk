// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkItemFeedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatuses(v []*string) *QueryWorkItemFeedRequest
	GetBusinessStatuses() []*string
	SetCollaborationGroupId(v string) *QueryWorkItemFeedRequest
	GetCollaborationGroupId() *string
	SetCursor(v string) *QueryWorkItemFeedRequest
	GetCursor() *string
	SetItemTypes(v []*string) *QueryWorkItemFeedRequest
	GetItemTypes() []*string
	SetKeyword(v string) *QueryWorkItemFeedRequest
	GetKeyword() *string
	SetPageSize(v int64) *QueryWorkItemFeedRequest
	GetPageSize() *int64
	SetPriorities(v []*string) *QueryWorkItemFeedRequest
	GetPriorities() []*string
	SetScope(v string) *QueryWorkItemFeedRequest
	GetScope() *string
	SetSourceTaskId(v string) *QueryWorkItemFeedRequest
	GetSourceTaskId() *string
	SetTenantId(v string) *QueryWorkItemFeedRequest
	GetTenantId() *string
	SetView(v string) *QueryWorkItemFeedRequest
	GetView() *string
}

type QueryWorkItemFeedRequest struct {
	// The business status filter (seven states).
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
	// The pagination cursor. Use the nextCursor value from the previous response.
	//
	// example:
	//
	// example
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// The work item type filter.
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
	// example
	Priorities []*string `json:"priorities,omitempty" xml:"priorities,omitempty" type:"Repeated"`
	// The personal scope filter.
	//
	// example:
	//
	// ASSIGNED_TO_ME
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The ID of the source task. If specified, the flat sub-items of the task are returned.
	//
	// example:
	//
	// example
	SourceTaskId *string `json:"sourceTaskId,omitempty" xml:"sourceTaskId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The view type. Valid values: FLAT (flat view) and TREE (task grouping).
	//
	// example:
	//
	// FLAT
	View *string `json:"view,omitempty" xml:"view,omitempty"`
}

func (s QueryWorkItemFeedRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedRequest) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedRequest) GetBusinessStatuses() []*string {
	return s.BusinessStatuses
}

func (s *QueryWorkItemFeedRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *QueryWorkItemFeedRequest) GetCursor() *string {
	return s.Cursor
}

func (s *QueryWorkItemFeedRequest) GetItemTypes() []*string {
	return s.ItemTypes
}

func (s *QueryWorkItemFeedRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryWorkItemFeedRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryWorkItemFeedRequest) GetPriorities() []*string {
	return s.Priorities
}

func (s *QueryWorkItemFeedRequest) GetScope() *string {
	return s.Scope
}

func (s *QueryWorkItemFeedRequest) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *QueryWorkItemFeedRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryWorkItemFeedRequest) GetView() *string {
	return s.View
}

func (s *QueryWorkItemFeedRequest) SetBusinessStatuses(v []*string) *QueryWorkItemFeedRequest {
	s.BusinessStatuses = v
	return s
}

func (s *QueryWorkItemFeedRequest) SetCollaborationGroupId(v string) *QueryWorkItemFeedRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetCursor(v string) *QueryWorkItemFeedRequest {
	s.Cursor = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetItemTypes(v []*string) *QueryWorkItemFeedRequest {
	s.ItemTypes = v
	return s
}

func (s *QueryWorkItemFeedRequest) SetKeyword(v string) *QueryWorkItemFeedRequest {
	s.Keyword = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetPageSize(v int64) *QueryWorkItemFeedRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetPriorities(v []*string) *QueryWorkItemFeedRequest {
	s.Priorities = v
	return s
}

func (s *QueryWorkItemFeedRequest) SetScope(v string) *QueryWorkItemFeedRequest {
	s.Scope = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetSourceTaskId(v string) *QueryWorkItemFeedRequest {
	s.SourceTaskId = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetTenantId(v string) *QueryWorkItemFeedRequest {
	s.TenantId = &v
	return s
}

func (s *QueryWorkItemFeedRequest) SetView(v string) *QueryWorkItemFeedRequest {
	s.View = &v
	return s
}

func (s *QueryWorkItemFeedRequest) Validate() error {
	return dara.Validate(s)
}
