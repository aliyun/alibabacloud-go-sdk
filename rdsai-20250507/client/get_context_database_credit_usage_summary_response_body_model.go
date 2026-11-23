// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseCreditUsageSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuckets(v []*GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) *GetContextDatabaseCreditUsageSummaryResponseBody
	GetBuckets() []*GetContextDatabaseCreditUsageSummaryResponseBodyBuckets
	SetDimension(v string) *GetContextDatabaseCreditUsageSummaryResponseBody
	GetDimension() *string
	SetEndTime(v string) *GetContextDatabaseCreditUsageSummaryResponseBody
	GetEndTime() *string
	SetInterval(v string) *GetContextDatabaseCreditUsageSummaryResponseBody
	GetInterval() *string
	SetRequestId(v string) *GetContextDatabaseCreditUsageSummaryResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetContextDatabaseCreditUsageSummaryResponseBody
	GetStartTime() *string
	SetTotalCredits(v string) *GetContextDatabaseCreditUsageSummaryResponseBody
	GetTotalCredits() *string
	SetWorkspaceId(v string) *GetContextDatabaseCreditUsageSummaryResponseBody
	GetWorkspaceId() *string
}

type GetContextDatabaseCreditUsageSummaryResponseBody struct {
	// The bucketing results, sorted by BucketStart, Item, and SubItem in sequence. Unpriced records are included in Quantity and RecordCount, but not in Credits.
	Buckets []*GetContextDatabaseCreditUsageSummaryResponseBodyBuckets `json:"Buckets,omitempty" xml:"Buckets,omitempty" type:"Repeated"`
	// The query dimension. Valid values: Workspace (a single workspace) and Account (all workspaces under the caller account).
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
	// The actual bucketing granularity used. Valid values: Hour (hourly) and Day (daily). If the request parameter is omitted, the automatically selected granularity is returned.
	//
	// example:
	//
	// Day
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
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
	// The total credit consumption of priced records that meet the filter conditions within the query range. The value is returned as a decimal string with up to 6 decimal places. 1 credit = 1 CNY. If there are no priced records, the string 0 is returned. Unpriced records are not included in the total amount. Do not treat this value as the final bill amount.
	//
	// example:
	//
	// 0.2513
	TotalCredits *string `json:"TotalCredits,omitempty" xml:"TotalCredits,omitempty"`
	// The ID of the queried workspace. This parameter has a value only when Dimension is Workspace. It is empty for the account dimension.
	//
	// example:
	//
	// ws-x9z4ctzed9gfu2w
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetContextDatabaseCreditUsageSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseCreditUsageSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) GetBuckets() []*GetContextDatabaseCreditUsageSummaryResponseBodyBuckets {
	return s.Buckets
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) GetDimension() *string {
	return s.Dimension
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) GetInterval() *string {
	return s.Interval
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) GetTotalCredits() *string {
	return s.TotalCredits
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) SetBuckets(v []*GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) *GetContextDatabaseCreditUsageSummaryResponseBody {
	s.Buckets = v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) SetDimension(v string) *GetContextDatabaseCreditUsageSummaryResponseBody {
	s.Dimension = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) SetEndTime(v string) *GetContextDatabaseCreditUsageSummaryResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) SetInterval(v string) *GetContextDatabaseCreditUsageSummaryResponseBody {
	s.Interval = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) SetRequestId(v string) *GetContextDatabaseCreditUsageSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) SetStartTime(v string) *GetContextDatabaseCreditUsageSummaryResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) SetTotalCredits(v string) *GetContextDatabaseCreditUsageSummaryResponseBody {
	s.TotalCredits = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) SetWorkspaceId(v string) *GetContextDatabaseCreditUsageSummaryResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBody) Validate() error {
	if s.Buckets != nil {
		for _, item := range s.Buckets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetContextDatabaseCreditUsageSummaryResponseBodyBuckets struct {
	// The start time of the bucket, in ISO 8601 UTC time format ending with Z.
	//
	// example:
	//
	// 2026-10-01T00:00:00Z
	BucketStart *string `json:"BucketStart,omitempty" xml:"BucketStart,omitempty"`
	// The total credit consumption of priced records in the bucket. The value is returned as a decimal string with up to 6 decimal places. 1 credit = 1 CNY. If all records are unpriced, the string 0 is returned.
	//
	// example:
	//
	// 0.2513
	Credits *string `json:"Credits,omitempty" xml:"Credits,omitempty"`
	// The billing item corresponding to the bucket. Valid values: KnowledgeBaseRead (knowledge base read), MemoryRead (memory read), MemoryWrite (memory write), KnowledgeBaseWriteToken (knowledge base write token), and ContextGrowth (context evolution).
	//
	// example:
	//
	// ContextGrowth
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The total usage of all records in the bucket, including unpriced records. The unit is specified by Unit.
	//
	// example:
	//
	// 251300
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The number of metering records in the bucket, including unpriced records.
	//
	// example:
	//
	// 3
	RecordCount *int64 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The context evolution sub-item. Recognized values include KnowledgeOrganization (knowledge organization), MemoryPromotion (memory promotion), KnowledgeReview (knowledge review), KnowledgeEvolution (knowledge evolution), and MemoryEvolution (memory evolution). For other billing items, missing sub-items, or unrecognized sub-items, null is returned.
	//
	// example:
	//
	// MemoryEvolution
	SubItem *string `json:"SubItem,omitempty" xml:"SubItem,omitempty"`
	// The unit of measurement. Valid values: Count (number of times), Token (number of tokens), and MicroCredit (one millionth of a credit).
	//
	// example:
	//
	// MicroCredit
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) GetBucketStart() *string {
	return s.BucketStart
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) GetCredits() *string {
	return s.Credits
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) GetItem() *string {
	return s.Item
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) GetQuantity() *int64 {
	return s.Quantity
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) GetSubItem() *string {
	return s.SubItem
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) GetUnit() *string {
	return s.Unit
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) SetBucketStart(v string) *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets {
	s.BucketStart = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) SetCredits(v string) *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets {
	s.Credits = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) SetItem(v string) *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets {
	s.Item = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) SetQuantity(v int64) *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets {
	s.Quantity = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) SetRecordCount(v int64) *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets {
	s.RecordCount = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) SetSubItem(v string) *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets {
	s.SubItem = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) SetUnit(v string) *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets {
	s.Unit = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponseBodyBuckets) Validate() error {
	return dara.Validate(s)
}
