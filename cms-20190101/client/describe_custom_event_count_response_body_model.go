// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEventCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCustomEventCountResponseBody
	GetCode() *string
	SetCustomEventCounts(v *DescribeCustomEventCountResponseBodyCustomEventCounts) *DescribeCustomEventCountResponseBody
	GetCustomEventCounts() *DescribeCustomEventCountResponseBodyCustomEventCounts
	SetMessage(v string) *DescribeCustomEventCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCustomEventCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCustomEventCountResponseBody
	GetSuccess() *bool
}

type DescribeCustomEventCountResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the custom event.
	CustomEventCounts *DescribeCustomEventCountResponseBodyCustomEventCounts `json:"CustomEventCounts,omitempty" xml:"CustomEventCounts,omitempty" type:"Struct"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 60912C8D-B340-4253-ADE7-61ACDFD25CFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomEventCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCustomEventCountResponseBody) GetCustomEventCounts() *DescribeCustomEventCountResponseBodyCustomEventCounts {
	return s.CustomEventCounts
}

func (s *DescribeCustomEventCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCustomEventCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomEventCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCustomEventCountResponseBody) SetCode(v string) *DescribeCustomEventCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomEventCountResponseBody) SetCustomEventCounts(v *DescribeCustomEventCountResponseBodyCustomEventCounts) *DescribeCustomEventCountResponseBody {
	s.CustomEventCounts = v
	return s
}

func (s *DescribeCustomEventCountResponseBody) SetMessage(v string) *DescribeCustomEventCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomEventCountResponseBody) SetRequestId(v string) *DescribeCustomEventCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomEventCountResponseBody) SetSuccess(v bool) *DescribeCustomEventCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomEventCountResponseBody) Validate() error {
	if s.CustomEventCounts != nil {
		if err := s.CustomEventCounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomEventCountResponseBodyCustomEventCounts struct {
	CustomEventCount []*DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount `json:"CustomEventCount,omitempty" xml:"CustomEventCount,omitempty" type:"Repeated"`
}

func (s DescribeCustomEventCountResponseBodyCustomEventCounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventCountResponseBodyCustomEventCounts) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCounts) GetCustomEventCount() []*DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount {
	return s.CustomEventCount
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCounts) SetCustomEventCount(v []*DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) *DescribeCustomEventCountResponseBodyCustomEventCounts {
	s.CustomEventCount = v
	return s
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCounts) Validate() error {
	if s.CustomEventCount != nil {
		for _, item := range s.CustomEventCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount struct {
	// The event name.
	//
	// example:
	//
	// BABEL_BUY
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of times that the custom event occurred.
	//
	// example:
	//
	// 20
	Num *int32 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The time when the event occurred.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552267615000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) GetName() *string {
	return s.Name
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) GetNum() *int32 {
	return s.Num
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) GetTime() *int64 {
	return s.Time
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) SetName(v string) *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount {
	s.Name = &v
	return s
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) SetNum(v int32) *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount {
	s.Num = &v
	return s
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) SetTime(v int64) *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount {
	s.Time = &v
	return s
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) Validate() error {
	return dara.Validate(s)
}
