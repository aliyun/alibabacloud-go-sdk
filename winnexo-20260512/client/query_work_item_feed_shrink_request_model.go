// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkItemFeedShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatusesShrink(v string) *QueryWorkItemFeedShrinkRequest
	GetBusinessStatusesShrink() *string
	SetCollaborationGroupId(v string) *QueryWorkItemFeedShrinkRequest
	GetCollaborationGroupId() *string
	SetCursor(v string) *QueryWorkItemFeedShrinkRequest
	GetCursor() *string
	SetItemTypesShrink(v string) *QueryWorkItemFeedShrinkRequest
	GetItemTypesShrink() *string
	SetKeyword(v string) *QueryWorkItemFeedShrinkRequest
	GetKeyword() *string
	SetPageSize(v int64) *QueryWorkItemFeedShrinkRequest
	GetPageSize() *int64
	SetPrioritiesShrink(v string) *QueryWorkItemFeedShrinkRequest
	GetPrioritiesShrink() *string
	SetScope(v string) *QueryWorkItemFeedShrinkRequest
	GetScope() *string
	SetSourceTaskId(v string) *QueryWorkItemFeedShrinkRequest
	GetSourceTaskId() *string
	SetTenantId(v string) *QueryWorkItemFeedShrinkRequest
	GetTenantId() *string
	SetView(v string) *QueryWorkItemFeedShrinkRequest
	GetView() *string
}

type QueryWorkItemFeedShrinkRequest struct {
	// The business status filter (seven states).
	//
	// example:
	//
	// example
	BusinessStatusesShrink *string `json:"businessStatuses,omitempty" xml:"businessStatuses,omitempty"`
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
	ItemTypesShrink *string `json:"itemTypes,omitempty" xml:"itemTypes,omitempty"`
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
	PrioritiesShrink *string `json:"priorities,omitempty" xml:"priorities,omitempty"`
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

func (s QueryWorkItemFeedShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedShrinkRequest) GetBusinessStatusesShrink() *string {
	return s.BusinessStatusesShrink
}

func (s *QueryWorkItemFeedShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *QueryWorkItemFeedShrinkRequest) GetCursor() *string {
	return s.Cursor
}

func (s *QueryWorkItemFeedShrinkRequest) GetItemTypesShrink() *string {
	return s.ItemTypesShrink
}

func (s *QueryWorkItemFeedShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryWorkItemFeedShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryWorkItemFeedShrinkRequest) GetPrioritiesShrink() *string {
	return s.PrioritiesShrink
}

func (s *QueryWorkItemFeedShrinkRequest) GetScope() *string {
	return s.Scope
}

func (s *QueryWorkItemFeedShrinkRequest) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *QueryWorkItemFeedShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryWorkItemFeedShrinkRequest) GetView() *string {
	return s.View
}

func (s *QueryWorkItemFeedShrinkRequest) SetBusinessStatusesShrink(v string) *QueryWorkItemFeedShrinkRequest {
	s.BusinessStatusesShrink = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetCollaborationGroupId(v string) *QueryWorkItemFeedShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetCursor(v string) *QueryWorkItemFeedShrinkRequest {
	s.Cursor = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetItemTypesShrink(v string) *QueryWorkItemFeedShrinkRequest {
	s.ItemTypesShrink = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetKeyword(v string) *QueryWorkItemFeedShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetPageSize(v int64) *QueryWorkItemFeedShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetPrioritiesShrink(v string) *QueryWorkItemFeedShrinkRequest {
	s.PrioritiesShrink = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetScope(v string) *QueryWorkItemFeedShrinkRequest {
	s.Scope = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetSourceTaskId(v string) *QueryWorkItemFeedShrinkRequest {
	s.SourceTaskId = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetTenantId(v string) *QueryWorkItemFeedShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) SetView(v string) *QueryWorkItemFeedShrinkRequest {
	s.View = &v
	return s
}

func (s *QueryWorkItemFeedShrinkRequest) Validate() error {
	return dara.Validate(s)
}
