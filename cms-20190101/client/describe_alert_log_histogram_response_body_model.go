// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertLogHistogramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLogHistogramList(v []*DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) *DescribeAlertLogHistogramResponseBody
	GetAlertLogHistogramList() []*DescribeAlertLogHistogramResponseBodyAlertLogHistogramList
	SetCode(v string) *DescribeAlertLogHistogramResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeAlertLogHistogramResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertLogHistogramResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertLogHistogramResponseBody
	GetSuccess() *bool
}

type DescribeAlertLogHistogramResponseBody struct {
	// The number of alert logs that were generated during each interval of a time period.
	AlertLogHistogramList []*DescribeAlertLogHistogramResponseBodyAlertLogHistogramList `json:"AlertLogHistogramList,omitempty" xml:"AlertLogHistogramList,omitempty" type:"Repeated"`
	// The response code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1C4A3709-BF52-42EE-87B5-7435F0929585
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

func (s DescribeAlertLogHistogramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogHistogramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogHistogramResponseBody) GetAlertLogHistogramList() []*DescribeAlertLogHistogramResponseBodyAlertLogHistogramList {
	return s.AlertLogHistogramList
}

func (s *DescribeAlertLogHistogramResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAlertLogHistogramResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertLogHistogramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertLogHistogramResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertLogHistogramResponseBody) SetAlertLogHistogramList(v []*DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) *DescribeAlertLogHistogramResponseBody {
	s.AlertLogHistogramList = v
	return s
}

func (s *DescribeAlertLogHistogramResponseBody) SetCode(v string) *DescribeAlertLogHistogramResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBody) SetMessage(v string) *DescribeAlertLogHistogramResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBody) SetRequestId(v string) *DescribeAlertLogHistogramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBody) SetSuccess(v bool) *DescribeAlertLogHistogramResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBody) Validate() error {
	if s.AlertLogHistogramList != nil {
		for _, item := range s.AlertLogHistogramList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertLogHistogramResponseBodyAlertLogHistogramList struct {
	// The number of alert logs.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The start timestamp of the queried alert logs.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 1610074791
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// The end timestamp of the queried alert logs.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 1610074800
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) GetFrom() *int64 {
	return s.From
}

func (s *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) GetTo() *int64 {
	return s.To
}

func (s *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) SetCount(v int32) *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList {
	s.Count = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) SetFrom(v int64) *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList {
	s.From = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) SetTo(v int64) *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList {
	s.To = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) Validate() error {
	return dara.Validate(s)
}
