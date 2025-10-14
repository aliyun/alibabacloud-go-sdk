// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventHistogramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSystemEventHistogramResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeSystemEventHistogramResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSystemEventHistogramResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSystemEventHistogramResponseBody
	GetSuccess() *string
	SetSystemEventHistograms(v *DescribeSystemEventHistogramResponseBodySystemEventHistograms) *DescribeSystemEventHistogramResponseBody
	GetSystemEventHistograms() *DescribeSystemEventHistogramResponseBodySystemEventHistograms
}

type DescribeSystemEventHistogramResponseBody struct {
	// The response code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 486029C9-53E1-44B4-85A8-16A571A043FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true: The request was successful. false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information about the number of times the system event occurred during each interval of a time period.
	SystemEventHistograms *DescribeSystemEventHistogramResponseBodySystemEventHistograms `json:"SystemEventHistograms,omitempty" xml:"SystemEventHistograms,omitempty" type:"Struct"`
}

func (s DescribeSystemEventHistogramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventHistogramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventHistogramResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSystemEventHistogramResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSystemEventHistogramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSystemEventHistogramResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSystemEventHistogramResponseBody) GetSystemEventHistograms() *DescribeSystemEventHistogramResponseBodySystemEventHistograms {
	return s.SystemEventHistograms
}

func (s *DescribeSystemEventHistogramResponseBody) SetCode(v string) *DescribeSystemEventHistogramResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBody) SetMessage(v string) *DescribeSystemEventHistogramResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBody) SetRequestId(v string) *DescribeSystemEventHistogramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBody) SetSuccess(v string) *DescribeSystemEventHistogramResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBody) SetSystemEventHistograms(v *DescribeSystemEventHistogramResponseBodySystemEventHistograms) *DescribeSystemEventHistogramResponseBody {
	s.SystemEventHistograms = v
	return s
}

func (s *DescribeSystemEventHistogramResponseBody) Validate() error {
	if s.SystemEventHistograms != nil {
		if err := s.SystemEventHistograms.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSystemEventHistogramResponseBodySystemEventHistograms struct {
	SystemEventHistogram []*DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram `json:"SystemEventHistogram,omitempty" xml:"SystemEventHistogram,omitempty" type:"Repeated"`
}

func (s DescribeSystemEventHistogramResponseBodySystemEventHistograms) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventHistogramResponseBodySystemEventHistograms) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistograms) GetSystemEventHistogram() []*DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram {
	return s.SystemEventHistogram
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistograms) SetSystemEventHistogram(v []*DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) *DescribeSystemEventHistogramResponseBodySystemEventHistograms {
	s.SystemEventHistogram = v
	return s
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistograms) Validate() error {
	if s.SystemEventHistogram != nil {
		for _, item := range s.SystemEventHistogram {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram struct {
	// The number of times the system event occurred.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The end time.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552225753000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552225770000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) GetCount() *int64 {
	return s.Count
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) SetCount(v int64) *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram {
	s.Count = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) SetEndTime(v int64) *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram {
	s.EndTime = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) SetStartTime(v int64) *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram {
	s.StartTime = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) Validate() error {
	return dara.Validate(s)
}
