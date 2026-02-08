// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundCallNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOutboundCallNumbersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListOutboundCallNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListOutboundCallNumbersResponseBody
	GetMessage() *string
	SetOutboundCallNumbers(v *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) *ListOutboundCallNumbersResponseBody
	GetOutboundCallNumbers() *ListOutboundCallNumbersResponseBodyOutboundCallNumbers
	SetRequestId(v string) *ListOutboundCallNumbersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOutboundCallNumbersResponseBody
	GetSuccess() *bool
}

type ListOutboundCallNumbersResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Number list
	OutboundCallNumbers *ListOutboundCallNumbersResponseBodyOutboundCallNumbers `json:"OutboundCallNumbers,omitempty" xml:"OutboundCallNumbers,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOutboundCallNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundCallNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutboundCallNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOutboundCallNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListOutboundCallNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOutboundCallNumbersResponseBody) GetOutboundCallNumbers() *ListOutboundCallNumbersResponseBodyOutboundCallNumbers {
	return s.OutboundCallNumbers
}

func (s *ListOutboundCallNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOutboundCallNumbersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOutboundCallNumbersResponseBody) SetCode(v string) *ListOutboundCallNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBody) SetHttpStatusCode(v int32) *ListOutboundCallNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBody) SetMessage(v string) *ListOutboundCallNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBody) SetOutboundCallNumbers(v *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) *ListOutboundCallNumbersResponseBody {
	s.OutboundCallNumbers = v
	return s
}

func (s *ListOutboundCallNumbersResponseBody) SetRequestId(v string) *ListOutboundCallNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBody) SetSuccess(v bool) *ListOutboundCallNumbersResponseBody {
	s.Success = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBody) Validate() error {
	if s.OutboundCallNumbers != nil {
		if err := s.OutboundCallNumbers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOutboundCallNumbersResponseBodyOutboundCallNumbers struct {
	// List of number data
	List []*ListOutboundCallNumbersResponseBodyOutboundCallNumbersList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Count
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOutboundCallNumbersResponseBodyOutboundCallNumbers) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundCallNumbersResponseBodyOutboundCallNumbers) GoString() string {
	return s.String()
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) GetList() []*ListOutboundCallNumbersResponseBodyOutboundCallNumbersList {
	return s.List
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) SetList(v []*ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) *ListOutboundCallNumbersResponseBodyOutboundCallNumbers {
	s.List = v
	return s
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) SetPageNumber(v int32) *ListOutboundCallNumbersResponseBodyOutboundCallNumbers {
	s.PageNumber = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) SetPageSize(v int32) *ListOutboundCallNumbersResponseBodyOutboundCallNumbers {
	s.PageSize = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) SetTotalCount(v int32) *ListOutboundCallNumbersResponseBodyOutboundCallNumbers {
	s.TotalCount = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbers) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOutboundCallNumbersResponseBodyOutboundCallNumbersList struct {
	// Number
	//
	// example:
	//
	// 10088
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// Number ID
	//
	// example:
	//
	// fa0e21e9-caab-4629-9121-1e341243d599
	OutboundCallNumberId *string `json:"OutboundCallNumberId,omitempty" xml:"OutboundCallNumberId,omitempty"`
	// Time Range for Rate Limiting, in seconds
	//
	// example:
	//
	// 10
	RateLimitCount *string `json:"RateLimitCount,omitempty" xml:"RateLimitCount,omitempty"`
	// Rate Limiting count
	//
	// example:
	//
	// 100
	RateLimitPeriod *string `json:"RateLimitPeriod,omitempty" xml:"RateLimitPeriod,omitempty"`
}

func (s ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) GoString() string {
	return s.String()
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) GetNumber() *string {
	return s.Number
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) GetOutboundCallNumberId() *string {
	return s.OutboundCallNumberId
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) GetRateLimitCount() *string {
	return s.RateLimitCount
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) GetRateLimitPeriod() *string {
	return s.RateLimitPeriod
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) SetNumber(v string) *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList {
	s.Number = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) SetOutboundCallNumberId(v string) *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList {
	s.OutboundCallNumberId = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) SetRateLimitCount(v string) *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList {
	s.RateLimitCount = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) SetRateLimitPeriod(v string) *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList {
	s.RateLimitPeriod = &v
	return s
}

func (s *ListOutboundCallNumbersResponseBodyOutboundCallNumbersList) Validate() error {
	return dara.Validate(s)
}
