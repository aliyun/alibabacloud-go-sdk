// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodEditingUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEditingData(v []*DescribeVodEditingUsageDataResponseBodyEditingData) *DescribeVodEditingUsageDataResponseBody
	GetEditingData() []*DescribeVodEditingUsageDataResponseBodyEditingData
	SetEndTime(v string) *DescribeVodEditingUsageDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodEditingUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodEditingUsageDataResponseBody
	GetStartTime() *string
}

type DescribeVodEditingUsageDataResponseBody struct {
	EditingData []*DescribeVodEditingUsageDataResponseBodyEditingData `json:"EditingData,omitempty" xml:"EditingData,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-11-07T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 6CB55B62-6E5F-54D1-80BF-DFA3DE9F0***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2024-11-06T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodEditingUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodEditingUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodEditingUsageDataResponseBody) GetEditingData() []*DescribeVodEditingUsageDataResponseBodyEditingData {
	return s.EditingData
}

func (s *DescribeVodEditingUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodEditingUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodEditingUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodEditingUsageDataResponseBody) SetEditingData(v []*DescribeVodEditingUsageDataResponseBodyEditingData) *DescribeVodEditingUsageDataResponseBody {
	s.EditingData = v
	return s
}

func (s *DescribeVodEditingUsageDataResponseBody) SetEndTime(v string) *DescribeVodEditingUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodEditingUsageDataResponseBody) SetRequestId(v string) *DescribeVodEditingUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodEditingUsageDataResponseBody) SetStartTime(v string) *DescribeVodEditingUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodEditingUsageDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodEditingUsageDataResponseBodyEditingData struct {
	// example:
	//
	// 123
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// H264.SD
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// example:
	//
	// 2024-11-06T16:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodEditingUsageDataResponseBodyEditingData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodEditingUsageDataResponseBodyEditingData) GoString() string {
	return s.String()
}

func (s *DescribeVodEditingUsageDataResponseBodyEditingData) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeVodEditingUsageDataResponseBodyEditingData) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodEditingUsageDataResponseBodyEditingData) GetSpecification() *string {
	return s.Specification
}

func (s *DescribeVodEditingUsageDataResponseBodyEditingData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodEditingUsageDataResponseBodyEditingData) SetDuration(v int64) *DescribeVodEditingUsageDataResponseBodyEditingData {
	s.Duration = &v
	return s
}

func (s *DescribeVodEditingUsageDataResponseBodyEditingData) SetRegion(v string) *DescribeVodEditingUsageDataResponseBodyEditingData {
	s.Region = &v
	return s
}

func (s *DescribeVodEditingUsageDataResponseBodyEditingData) SetSpecification(v string) *DescribeVodEditingUsageDataResponseBodyEditingData {
	s.Specification = &v
	return s
}

func (s *DescribeVodEditingUsageDataResponseBodyEditingData) SetTimeStamp(v string) *DescribeVodEditingUsageDataResponseBodyEditingData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodEditingUsageDataResponseBodyEditingData) Validate() error {
	return dara.Validate(s)
}
