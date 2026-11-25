// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseCreditUsageSummaryByMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetContextDatabaseCreditUsageSummaryByMemberRequest
	GetEndTime() *string
	SetItem(v string) *GetContextDatabaseCreditUsageSummaryByMemberRequest
	GetItem() *string
	SetMemberId(v string) *GetContextDatabaseCreditUsageSummaryByMemberRequest
	GetMemberId() *string
	SetPage(v int32) *GetContextDatabaseCreditUsageSummaryByMemberRequest
	GetPage() *int32
	SetPageSize(v int32) *GetContextDatabaseCreditUsageSummaryByMemberRequest
	GetPageSize() *int32
	SetStartTime(v string) *GetContextDatabaseCreditUsageSummaryByMemberRequest
	GetStartTime() *string
	SetWorkspaceId(v string) *GetContextDatabaseCreditUsageSummaryByMemberRequest
	GetWorkspaceId() *string
}

type GetContextDatabaseCreditUsageSummaryByMemberRequest struct {
	// The end time of the query interval, exclusive. It must be later than `StartTime`. Use the ISO 8601 time format, such as `2026-10-08T00:00:00Z` (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-08T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The billing item used to filter records. Supported values include KnowledgeBaseRead (knowledge base read), MemoryRead (memory read), MemoryWrite (memory write), KnowledgeBaseWriteToken (knowledge base write token), and ContextGrowth (context growth). If omitted, all billing items that consume credits are included. Specifying storage items or unlisted values returns an `InvalidParameter` error.
	//
	// example:
	//
	// MemoryWrite
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The member ID used to filter the aggregated results. If specified, only this member is analyzed without changing the caller identity. An empty list is returned if no matching records are found.
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
	// The maximum number of members to return per page. Valid values: 1 to 200. Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the query interval, inclusive. Use the ISO 8601 time format, such as `2026-10-01T00:00:00Z` (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the workspace to be analyzed. This parameter is required. The workspace must belong to the caller\"s account. Locked workspaces can still be queried. Member IDs are unique only within their respective workspaces, and cross-workspace ranking is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-x9z4ctzed9gfu2w
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetContextDatabaseCreditUsageSummaryByMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseCreditUsageSummaryByMemberRequest) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) GetItem() *string {
	return s.Item
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) SetEndTime(v string) *GetContextDatabaseCreditUsageSummaryByMemberRequest {
	s.EndTime = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) SetItem(v string) *GetContextDatabaseCreditUsageSummaryByMemberRequest {
	s.Item = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) SetMemberId(v string) *GetContextDatabaseCreditUsageSummaryByMemberRequest {
	s.MemberId = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) SetPage(v int32) *GetContextDatabaseCreditUsageSummaryByMemberRequest {
	s.Page = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) SetPageSize(v int32) *GetContextDatabaseCreditUsageSummaryByMemberRequest {
	s.PageSize = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) SetStartTime(v string) *GetContextDatabaseCreditUsageSummaryByMemberRequest {
	s.StartTime = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) SetWorkspaceId(v string) *GetContextDatabaseCreditUsageSummaryByMemberRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberRequest) Validate() error {
	return dara.Validate(s)
}
