// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEventHistogramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCustomEventHistogramResponseBody
	GetCode() *string
	SetEventHistograms(v *DescribeCustomEventHistogramResponseBodyEventHistograms) *DescribeCustomEventHistogramResponseBody
	GetEventHistograms() *DescribeCustomEventHistogramResponseBodyEventHistograms
	SetMessage(v string) *DescribeCustomEventHistogramResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCustomEventHistogramResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCustomEventHistogramResponseBody
	GetSuccess() *string
}

type DescribeCustomEventHistogramResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the number of times that the custom event occurred during each interval of the specified time period.
	EventHistograms *DescribeCustomEventHistogramResponseBodyEventHistograms `json:"EventHistograms,omitempty" xml:"EventHistograms,omitempty" type:"Struct"`
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
	// 486029C9-53E1-44B4-85A8-16A571A043FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomEventHistogramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventHistogramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventHistogramResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCustomEventHistogramResponseBody) GetEventHistograms() *DescribeCustomEventHistogramResponseBodyEventHistograms {
	return s.EventHistograms
}

func (s *DescribeCustomEventHistogramResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCustomEventHistogramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomEventHistogramResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCustomEventHistogramResponseBody) SetCode(v string) *DescribeCustomEventHistogramResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBody) SetEventHistograms(v *DescribeCustomEventHistogramResponseBodyEventHistograms) *DescribeCustomEventHistogramResponseBody {
	s.EventHistograms = v
	return s
}

func (s *DescribeCustomEventHistogramResponseBody) SetMessage(v string) *DescribeCustomEventHistogramResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBody) SetRequestId(v string) *DescribeCustomEventHistogramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBody) SetSuccess(v string) *DescribeCustomEventHistogramResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomEventHistogramResponseBodyEventHistograms struct {
	EventHistogram []*DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram `json:"EventHistogram,omitempty" xml:"EventHistogram,omitempty" type:"Repeated"`
}

func (s DescribeCustomEventHistogramResponseBodyEventHistograms) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventHistogramResponseBodyEventHistograms) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistograms) GetEventHistogram() []*DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram {
	return s.EventHistogram
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistograms) SetEventHistogram(v []*DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) *DescribeCustomEventHistogramResponseBodyEventHistograms {
	s.EventHistogram = v
	return s
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistograms) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram struct {
	// The information about the number of times that the custom event occurred during an interval of the specified time period.
	//
	// example:
	//
	// 3
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The end time.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552226750000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552226740000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) SetCount(v int64) *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram {
	s.Count = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) SetEndTime(v int64) *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram {
	s.EndTime = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) SetStartTime(v int64) *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) Validate() error {
	return dara.Validate(s)
}
