// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseCreditUsageRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDimension(v string) *ListContextDatabaseCreditUsageRecordsResponseBody
	GetDimension() *string
	SetEndTime(v string) *ListContextDatabaseCreditUsageRecordsResponseBody
	GetEndTime() *string
	SetPage(v int32) *ListContextDatabaseCreditUsageRecordsResponseBody
	GetPage() *int32
	SetPageSize(v int32) *ListContextDatabaseCreditUsageRecordsResponseBody
	GetPageSize() *int32
	SetRecords(v []*ListContextDatabaseCreditUsageRecordsResponseBodyRecords) *ListContextDatabaseCreditUsageRecordsResponseBody
	GetRecords() []*ListContextDatabaseCreditUsageRecordsResponseBodyRecords
	SetRequestId(v string) *ListContextDatabaseCreditUsageRecordsResponseBody
	GetRequestId() *string
	SetStartTime(v string) *ListContextDatabaseCreditUsageRecordsResponseBody
	GetStartTime() *string
	SetTotal(v int64) *ListContextDatabaseCreditUsageRecordsResponseBody
	GetTotal() *int64
	SetWorkspaceId(v string) *ListContextDatabaseCreditUsageRecordsResponseBody
	GetWorkspaceId() *string
}

type ListContextDatabaseCreditUsageRecordsResponseBody struct {
	// The query dimension. Valid values: Workspace indicates a single workspace, and Account indicates all workspaces under the caller account.
	//
	// example:
	//
	// Workspace
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The end time of the query range passed in the request, exclusive.
	//
	// example:
	//
	// 2026-10-15T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The actual page number used for this query.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The actual page size used for this query.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The metering records on the current page. If SortBy is not specified, the records are sorted in descending order of request time.
	Records []*ListContextDatabaseCreditUsageRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The ID of the OpenAPI request, which is used to troubleshoot call issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query range passed in the request, inclusive.
	//
	// example:
	//
	// 2026-09-15T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of metering records that match the filter conditions, which is not affected by pagination.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The ID of the queried workspace. This parameter has a value only when Dimension is Workspace. It is empty for the account dimension.
	//
	// example:
	//
	// ws-x9z4ctzed9gfu2w
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseCreditUsageRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseCreditUsageRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) GetDimension() *string {
	return s.Dimension
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) GetRecords() []*ListContextDatabaseCreditUsageRecordsResponseBodyRecords {
	return s.Records
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) SetDimension(v string) *ListContextDatabaseCreditUsageRecordsResponseBody {
	s.Dimension = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) SetEndTime(v string) *ListContextDatabaseCreditUsageRecordsResponseBody {
	s.EndTime = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) SetPage(v int32) *ListContextDatabaseCreditUsageRecordsResponseBody {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) SetPageSize(v int32) *ListContextDatabaseCreditUsageRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) SetRecords(v []*ListContextDatabaseCreditUsageRecordsResponseBodyRecords) *ListContextDatabaseCreditUsageRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) SetRequestId(v string) *ListContextDatabaseCreditUsageRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) SetStartTime(v string) *ListContextDatabaseCreditUsageRecordsResponseBody {
	s.StartTime = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) SetTotal(v int64) *ListContextDatabaseCreditUsageRecordsResponseBody {
	s.Total = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) SetWorkspaceId(v string) *ListContextDatabaseCreditUsageRecordsResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseCreditUsageRecordsResponseBodyRecords struct {
	// The credit consumption of this record. It is returned as a decimal string with up to 6 decimal places. 1 credit = 1 CNY. The amount is recorded at the time of metering and does not change with subsequent price adjustments. It is null when unpriced, which does not indicate zero consumption.
	//
	// example:
	//
	// 0.0742
	Credits *string `json:"Credits,omitempty" xml:"Credits,omitempty"`
	// The billing item of this record. Valid values: KnowledgeBaseRead (knowledge base read), MemoryRead (memory base read), MemoryWrite (memory base write), KnowledgeBaseWriteToken (knowledge base write token), and ContextGrowth (context growth).
	//
	// example:
	//
	// ContextGrowth
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The ID of the member that generated this record. Historical or system records may not have a member ID, in which case null is returned.
	//
	// example:
	//
	// mb-abcdefghijklmno
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The usage of this record. The unit is specified by Unit in the same record.
	//
	// example:
	//
	// 74200
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The request time when this metering record was generated, in ISO 8601 UTC time format ending with Z.
	//
	// example:
	//
	// 2026-10-14T08:30:12Z
	RequestTime *string `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// The context growth sub-item. Recognized values include KnowledgeOrganization (knowledge organization), MemoryPromotion (memory promotion), KnowledgeReview (knowledge review), KnowledgeEvolution (knowledge evolution), and MemoryEvolution (memory evolution). null is returned for other billing items, missing sub-items, or unrecognized sub-items.
	//
	// example:
	//
	// KnowledgeReview
	SubItem *string `json:"SubItem,omitempty" xml:"SubItem,omitempty"`
	// The metering unit. Valid values: Count indicates the number of times, Token indicates the number of tokens, and MicroCredit indicates one millionth of a credit.
	//
	// example:
	//
	// MicroCredit
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The ID of the workspace to which this record belongs. Even if queried by account, each record still indicates its workspace.
	//
	// example:
	//
	// ws-x9z4ctzed9gfu2w
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseCreditUsageRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseCreditUsageRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) GetCredits() *string {
	return s.Credits
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) GetItem() *string {
	return s.Item
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) GetMemberId() *string {
	return s.MemberId
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) GetQuantity() *int64 {
	return s.Quantity
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) GetRequestTime() *string {
	return s.RequestTime
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) GetSubItem() *string {
	return s.SubItem
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) GetUnit() *string {
	return s.Unit
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) SetCredits(v string) *ListContextDatabaseCreditUsageRecordsResponseBodyRecords {
	s.Credits = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) SetItem(v string) *ListContextDatabaseCreditUsageRecordsResponseBodyRecords {
	s.Item = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) SetMemberId(v string) *ListContextDatabaseCreditUsageRecordsResponseBodyRecords {
	s.MemberId = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) SetQuantity(v int64) *ListContextDatabaseCreditUsageRecordsResponseBodyRecords {
	s.Quantity = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) SetRequestTime(v string) *ListContextDatabaseCreditUsageRecordsResponseBodyRecords {
	s.RequestTime = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) SetSubItem(v string) *ListContextDatabaseCreditUsageRecordsResponseBodyRecords {
	s.SubItem = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) SetUnit(v string) *ListContextDatabaseCreditUsageRecordsResponseBodyRecords {
	s.Unit = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) SetWorkspaceId(v string) *ListContextDatabaseCreditUsageRecordsResponseBodyRecords {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
