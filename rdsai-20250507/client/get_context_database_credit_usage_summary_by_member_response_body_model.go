// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseCreditUsageSummaryByMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody
	GetEndTime() *string
	SetItems(v []*GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody
	GetItems() []*GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems
	SetPage(v int32) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody
	GetPage() *int32
	SetPageSize(v int32) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody
	GetStartTime() *string
	SetTotal(v int64) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody
	GetTotal() *int64
	SetWorkspaceId(v string) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody
	GetWorkspaceId() *string
}

type GetContextDatabaseCreditUsageSummaryByMemberResponseBody struct {
	// The end time of the query range passed in the request, exclusive.
	//
	// example:
	//
	// 2026-10-08T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The member summaries on the current page, sorted in descending order of total credit consumption. Each record covers the entire query range.
	Items []*GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The actual page number used for the query.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The actual number of items per page used for the query.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query range passed in the request, inclusive.
	//
	// example:
	//
	// 2026-10-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of members that match the filter criteria. This is not the number of metering records and is not affected by pagination.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The ID of the workspace queried.
	//
	// example:
	//
	// ws-x9z4ctzed9gfu2w
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetContextDatabaseCreditUsageSummaryByMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseCreditUsageSummaryByMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) GetItems() []*GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems {
	return s.Items
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) SetEndTime(v string) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) SetItems(v []*GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody {
	s.Items = v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) SetPage(v int32) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody {
	s.Page = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) SetPageSize(v int32) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) SetRequestId(v string) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) SetStartTime(v string) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) SetTotal(v int64) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody {
	s.Total = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) SetWorkspaceId(v string) *GetContextDatabaseCreditUsageSummaryByMemberResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems struct {
	// The total credit consumption of priced records for the member. Returned as a decimal string with up to 6 decimal places. 1 credit equals 1 CNY. The value is null if there are no priced records.
	//
	// example:
	//
	// 0.2513
	Credits *string `json:"Credits,omitempty" xml:"Credits,omitempty"`
	// The member ID. Historical or system records may not have a member ID.
	//
	// example:
	//
	// mb-abcdefghijklmno
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The total metered quantity of records that match the filter criteria for the member, including unpriced records. To view the usage of a specific billable item, specify Item in the request.
	//
	// example:
	//
	// 12
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The number of metering records that match the filter criteria for the member, including unpriced records.
	//
	// example:
	//
	// 3
	RecordCount *int64 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
}

func (s GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) GetCredits() *string {
	return s.Credits
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) GetMemberId() *string {
	return s.MemberId
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) GetQuantity() *int64 {
	return s.Quantity
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) SetCredits(v string) *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems {
	s.Credits = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) SetMemberId(v string) *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems {
	s.MemberId = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) SetQuantity(v int64) *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems {
	s.Quantity = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) SetRecordCount(v int64) *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems {
	s.RecordCount = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
