// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseCreditUsageRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListContextDatabaseCreditUsageRecordsRequest
	GetEndTime() *string
	SetItem(v string) *ListContextDatabaseCreditUsageRecordsRequest
	GetItem() *string
	SetMemberId(v string) *ListContextDatabaseCreditUsageRecordsRequest
	GetMemberId() *string
	SetPage(v int32) *ListContextDatabaseCreditUsageRecordsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListContextDatabaseCreditUsageRecordsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListContextDatabaseCreditUsageRecordsRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListContextDatabaseCreditUsageRecordsRequest
	GetSortOrder() *string
	SetStartTime(v string) *ListContextDatabaseCreditUsageRecordsRequest
	GetStartTime() *string
	SetSubItem(v string) *ListContextDatabaseCreditUsageRecordsRequest
	GetSubItem() *string
	SetWorkspaceId(v string) *ListContextDatabaseCreditUsageRecordsRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseCreditUsageRecordsRequest struct {
	// The end time of the query range, exclusive. It must be later than StartTime. Use the ISO 8601 time format, such as `2026-10-15T00:00:00Z` (UTC). The maximum length of the query range is 31 days. If the limit is exceeded, InvalidParameter is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-15T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The billing item used to filter records. Valid values: KnowledgeBaseRead (knowledge base read), MemoryRead (memory base read), MemoryWrite (memory base write), KnowledgeBaseWriteToken (knowledge base write token), and ContextGrowth (context growth). If omitted, all billing items that consume credits are included. Storage items or unlisted values return InvalidParameter.
	//
	// example:
	//
	// ContextGrowth
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The member ID used to filter records. If specified, WorkspaceId must also be provided. Otherwise, InvalidParameter is returned. A member ID is unique only within the workspace to which it belongs. This parameter only filters data and does not change the caller identity. Records without a member ID are not matched.
	//
	// example:
	//
	// mb-abcdefghijklmno
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The maximum number of records to return per page. Valid values: 1 to 200. Default value: 50. The pagination offset, calculated as (Page - 1) × PageSize, must not exceed 10,000.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field by which to sort the records. Only credits is supported. If omitted, the records are sorted in descending order of request time. Specifying a sort field does not change the total number of filtered records.
	//
	// example:
	//
	// credits
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Valid values: asc (ascending order) and desc (descending order). The values are case-insensitive. If SortBy is specified, the default value is desc. If only SortOrder is provided, InvalidParameter is returned. When sorting by credit, unpriced records are always placed at the end, and records with the same amount are sorted in descending order of request time.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The start time of the query range, inclusive. Use the ISO 8601 time format, such as `2026-09-15T00:00:00Z` (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-15T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The context growth sub-item used to filter records. Valid values: KnowledgeOrganization (knowledge organization), MemoryPromotion (memory promotion), KnowledgeReview (knowledge review), KnowledgeEvolution (knowledge evolution), and MemoryEvolution (memory evolution). The values are case-sensitive. If only SubItem is provided, the billing item is automatically restricted to ContextGrowth. If other Item values are provided simultaneously or the sub-item is not in the list, InvalidParameter is returned.
	//
	// example:
	//
	// KnowledgeReview
	SubItem *string `json:"SubItem,omitempty" xml:"SubItem,omitempty"`
	// The ID of the workspace to query. If specified, only the data of this workspace under the caller account is queried, and locked workspaces can also be queried. If omitted, all workspaces under the caller account are queried.
	//
	// example:
	//
	// ws-x9z4ctzed9gfu2w
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseCreditUsageRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseCreditUsageRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) GetItem() *string {
	return s.Item
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) GetSubItem() *string {
	return s.SubItem
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) SetEndTime(v string) *ListContextDatabaseCreditUsageRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) SetItem(v string) *ListContextDatabaseCreditUsageRecordsRequest {
	s.Item = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) SetMemberId(v string) *ListContextDatabaseCreditUsageRecordsRequest {
	s.MemberId = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) SetPage(v int32) *ListContextDatabaseCreditUsageRecordsRequest {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) SetPageSize(v int32) *ListContextDatabaseCreditUsageRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) SetSortBy(v string) *ListContextDatabaseCreditUsageRecordsRequest {
	s.SortBy = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) SetSortOrder(v string) *ListContextDatabaseCreditUsageRecordsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) SetStartTime(v string) *ListContextDatabaseCreditUsageRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) SetSubItem(v string) *ListContextDatabaseCreditUsageRecordsRequest {
	s.SubItem = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) SetWorkspaceId(v string) *ListContextDatabaseCreditUsageRecordsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsRequest) Validate() error {
	return dara.Validate(s)
}
