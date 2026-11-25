// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventOperationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAiEngineEventOperationsResponseBodyItems) *DescribeAiEngineEventOperationsResponseBody
	GetItems() []*DescribeAiEngineEventOperationsResponseBodyItems
	SetMaxResults(v int32) *DescribeAiEngineEventOperationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAiEngineEventOperationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAiEngineEventOperationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAiEngineEventOperationsResponseBody
	GetTotalCount() *int64
}

type DescribeAiEngineEventOperationsResponseBody struct {
	// The list of event status change records, sorted in descending order by operation time.
	Items []*DescribeAiEngineEventOperationsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries to return on each page during paging in a paged query. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. If a next page exists, this field returns a value.
	//
	// > If this parameter returns a value, it indicates that a next page exists. You can use the returned **NextToken*	- as a request parameter to obtain the data on the next page until no value is returned, which indicates that all data is obtained.
	//
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01A1****-1219-019A-8B44-*****9C2C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of operation records for the event.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAiEngineEventOperationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventOperationsResponseBody) GetItems() []*DescribeAiEngineEventOperationsResponseBodyItems {
	return s.Items
}

func (s *DescribeAiEngineEventOperationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAiEngineEventOperationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAiEngineEventOperationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAiEngineEventOperationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAiEngineEventOperationsResponseBody) SetItems(v []*DescribeAiEngineEventOperationsResponseBodyItems) *DescribeAiEngineEventOperationsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAiEngineEventOperationsResponseBody) SetMaxResults(v int32) *DescribeAiEngineEventOperationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeAiEngineEventOperationsResponseBody) SetNextToken(v string) *DescribeAiEngineEventOperationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAiEngineEventOperationsResponseBody) SetRequestId(v string) *DescribeAiEngineEventOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAiEngineEventOperationsResponseBody) SetTotalCount(v int64) *DescribeAiEngineEventOperationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAiEngineEventOperationsResponseBody) Validate() error {
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

type DescribeAiEngineEventOperationsResponseBodyItems struct {
	// The event processing status before the modification. Valid values:
	//
	// - **unhandled**: Unprocessed.
	//
	// - **handling**: Processing.
	//
	// - **handled**: Processed.
	//
	// example:
	//
	// handling
	FromStatus *string `json:"FromStatus,omitempty" xml:"FromStatus,omitempty"`
	// The event operation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1792973400000
	OperatedAt *int64 `json:"OperatedAt,omitempty" xml:"OperatedAt,omitempty"`
	// The ID of the operation record.
	//
	// example:
	//
	// op109102
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The event processing status after the modification. Valid values:
	//
	// - **unhandled**: Unprocessed.
	//
	// - **handling**: Processing.
	//
	// - **handled**: Processed.
	//
	// example:
	//
	// handled
	ToStatus *string `json:"ToStatus,omitempty" xml:"ToStatus,omitempty"`
}

func (s DescribeAiEngineEventOperationsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventOperationsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventOperationsResponseBodyItems) GetFromStatus() *string {
	return s.FromStatus
}

func (s *DescribeAiEngineEventOperationsResponseBodyItems) GetOperatedAt() *int64 {
	return s.OperatedAt
}

func (s *DescribeAiEngineEventOperationsResponseBodyItems) GetOperationId() *string {
	return s.OperationId
}

func (s *DescribeAiEngineEventOperationsResponseBodyItems) GetToStatus() *string {
	return s.ToStatus
}

func (s *DescribeAiEngineEventOperationsResponseBodyItems) SetFromStatus(v string) *DescribeAiEngineEventOperationsResponseBodyItems {
	s.FromStatus = &v
	return s
}

func (s *DescribeAiEngineEventOperationsResponseBodyItems) SetOperatedAt(v int64) *DescribeAiEngineEventOperationsResponseBodyItems {
	s.OperatedAt = &v
	return s
}

func (s *DescribeAiEngineEventOperationsResponseBodyItems) SetOperationId(v string) *DescribeAiEngineEventOperationsResponseBodyItems {
	s.OperationId = &v
	return s
}

func (s *DescribeAiEngineEventOperationsResponseBodyItems) SetToStatus(v string) *DescribeAiEngineEventOperationsResponseBodyItems {
	s.ToStatus = &v
	return s
}

func (s *DescribeAiEngineEventOperationsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
