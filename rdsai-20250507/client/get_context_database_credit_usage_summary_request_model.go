// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseCreditUsageSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetContextDatabaseCreditUsageSummaryRequest
	GetEndTime() *string
	SetInterval(v string) *GetContextDatabaseCreditUsageSummaryRequest
	GetInterval() *string
	SetItem(v string) *GetContextDatabaseCreditUsageSummaryRequest
	GetItem() *string
	SetMemberId(v string) *GetContextDatabaseCreditUsageSummaryRequest
	GetMemberId() *string
	SetStartTime(v string) *GetContextDatabaseCreditUsageSummaryRequest
	GetStartTime() *string
	SetSubItem(v string) *GetContextDatabaseCreditUsageSummaryRequest
	GetSubItem() *string
	SetWorkspaceId(v string) *GetContextDatabaseCreditUsageSummaryRequest
	GetWorkspaceId() *string
}

type GetContextDatabaseCreditUsageSummaryRequest struct {
	// The end time of the query interval, exclusive. It must be later than StartTime. Use the ISO 8601 time format, such as `2026-10-15T00:00:00Z` (UTC). The maximum interval is 7 days for hourly statistics and 31 days for daily statistics.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-15T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The bucketing granularity. Valid values: Hour (hourly) and Day (daily). The values are case-insensitive. When omitted, Hour is used by default for intervals up to 7 days, and Day is used for intervals longer than 7 days. The maximum interval for Hour is 7 days, and the maximum interval for Day is 31 days. Exceeding the limit returns InvalidParameter.
	//
	// example:
	//
	// Day
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The billing item used to filter records. Valid values: KnowledgeBaseRead (knowledge base read), MemoryRead (memory read), MemoryWrite (memory write), KnowledgeBaseWriteToken (knowledge base write token), and ContextGrowth (context growth). When omitted, all billing items that consume credits are included. Storage items or unlisted values return InvalidParameter.
	//
	// example:
	//
	// ContextGrowth
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The member ID used to filter records. When specified, WorkspaceId must also be provided. Otherwise, InvalidParameter is returned. A member ID is unique only within its workspace. This parameter only filters data and does not change the caller identity. Records without a member ID are not matched. The results are still bucketed by time and billing item, and are not split by member.
	//
	// example:
	//
	// mb-abcdefghijklmno
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The start time of the query interval, inclusive. Use the ISO 8601 time format, such as `2026-09-15T00:00:00Z` (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-15T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The context evolution sub-item used to filter records. Valid values: KnowledgeOrganization (knowledge organization), MemoryPromotion (memory promotion), KnowledgeReview (knowledge review), KnowledgeEvolution (knowledge evolution), and MemoryEvolution (memory evolution). The values are case-sensitive. When only SubItem is provided, the billing item is automatically restricted to ContextGrowth. If other Item values are provided simultaneously or the sub-item is not in the list, InvalidParameter is returned.
	//
	// example:
	//
	// MemoryEvolution
	SubItem *string `json:"SubItem,omitempty" xml:"SubItem,omitempty"`
	// The ID of the workspace to query. When specified, only the data of this workspace under the caller account is queried, and locked workspaces can also be queried. When omitted, all workspaces under the caller account are queried.
	//
	// example:
	//
	// ws-x9z4ctzed9gfu2w
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetContextDatabaseCreditUsageSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseCreditUsageSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) GetInterval() *string {
	return s.Interval
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) GetItem() *string {
	return s.Item
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) GetSubItem() *string {
	return s.SubItem
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) SetEndTime(v string) *GetContextDatabaseCreditUsageSummaryRequest {
	s.EndTime = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) SetInterval(v string) *GetContextDatabaseCreditUsageSummaryRequest {
	s.Interval = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) SetItem(v string) *GetContextDatabaseCreditUsageSummaryRequest {
	s.Item = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) SetMemberId(v string) *GetContextDatabaseCreditUsageSummaryRequest {
	s.MemberId = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) SetStartTime(v string) *GetContextDatabaseCreditUsageSummaryRequest {
	s.StartTime = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) SetSubItem(v string) *GetContextDatabaseCreditUsageSummaryRequest {
	s.SubItem = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) SetWorkspaceId(v string) *GetContextDatabaseCreditUsageSummaryRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryRequest) Validate() error {
	return dara.Validate(s)
}
