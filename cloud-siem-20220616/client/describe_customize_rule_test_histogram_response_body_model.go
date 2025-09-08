// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeRuleTestHistogramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeCustomizeRuleTestHistogramResponseBody
	GetCode() *int32
	SetData(v []*DescribeCustomizeRuleTestHistogramResponseBodyData) *DescribeCustomizeRuleTestHistogramResponseBody
	GetData() []*DescribeCustomizeRuleTestHistogramResponseBodyData
	SetMessage(v string) *DescribeCustomizeRuleTestHistogramResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCustomizeRuleTestHistogramResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCustomizeRuleTestHistogramResponseBody
	GetSuccess() *bool
}

type DescribeCustomizeRuleTestHistogramResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value for the request.
	//
	// example:
	//
	// 123456
	Data []*DescribeCustomizeRuleTestHistogramResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) GetData() []*DescribeCustomizeRuleTestHistogramResponseBodyData {
	return s.Data
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetCode(v int32) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetData(v []*DescribeCustomizeRuleTestHistogramResponseBodyData) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetMessage(v string) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetRequestId(v string) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetSuccess(v bool) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomizeRuleTestHistogramResponseBodyData struct {
	// The number of alerts that are generated in the query time range.
	//
	// example:
	//
	// 125
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The start of the time range for querying alerts. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1599897188
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// The end of the time range for querying alerts. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1599997188
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) GetFrom() *int64 {
	return s.From
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) GetTo() *int64 {
	return s.To
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) SetCount(v int64) *DescribeCustomizeRuleTestHistogramResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) SetFrom(v int64) *DescribeCustomizeRuleTestHistogramResponseBodyData {
	s.From = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) SetTo(v int64) *DescribeCustomizeRuleTestHistogramResponseBodyData {
	s.To = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) Validate() error {
	return dara.Validate(s)
}
