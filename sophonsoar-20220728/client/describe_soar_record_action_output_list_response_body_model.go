// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarRecordActionOutputListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActionOutputs(v string) *DescribeSoarRecordActionOutputListResponseBody
	GetActionOutputs() *string
	SetPageNumber(v int32) *DescribeSoarRecordActionOutputListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSoarRecordActionOutputListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSoarRecordActionOutputListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSoarRecordActionOutputListResponseBody
	GetTotalCount() *int32
}

type DescribeSoarRecordActionOutputListResponseBody struct {
	// The data that is returned when the component action is performed. The value is a JSON array.
	//
	// >  The format of the output data is determined by the component that is configured when the playbook is written.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "a": "a",
	//
	//         "taskname": "92af3c79-1754-4646-9366-9ddbd1e45536_xxxx",
	//
	//         "log_time": 1699868849000
	//
	//     }
	//
	// ]
	ActionOutputs *string `json:"ActionOutputs,omitempty" xml:"ActionOutputs,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A2BF9CF-3E32-5E45-A79B-8F67E0A4FE90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSoarRecordActionOutputListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarRecordActionOutputListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordActionOutputListResponseBody) GetActionOutputs() *string {
	return s.ActionOutputs
}

func (s *DescribeSoarRecordActionOutputListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSoarRecordActionOutputListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarRecordActionOutputListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSoarRecordActionOutputListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetActionOutputs(v string) *DescribeSoarRecordActionOutputListResponseBody {
	s.ActionOutputs = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetPageNumber(v int32) *DescribeSoarRecordActionOutputListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetPageSize(v int32) *DescribeSoarRecordActionOutputListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetRequestId(v string) *DescribeSoarRecordActionOutputListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetTotalCount(v int32) *DescribeSoarRecordActionOutputListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) Validate() error {
	return dara.Validate(s)
}
