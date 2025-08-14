// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsMediaConvertUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeMeterImsMediaConvertUsageResponseBodyData) *DescribeMeterImsMediaConvertUsageResponseBody
	GetData() []*DescribeMeterImsMediaConvertUsageResponseBodyData
	SetRequestId(v string) *DescribeMeterImsMediaConvertUsageResponseBody
	GetRequestId() *string
}

type DescribeMeterImsMediaConvertUsageResponseBody struct {
	// The usage statistics of IMS on VOD transcoding.
	Data []*DescribeMeterImsMediaConvertUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// FBBB5210-2B78-58FB-A6FE-9DD887BB2C61
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterImsMediaConvertUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMediaConvertUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMediaConvertUsageResponseBody) GetData() []*DescribeMeterImsMediaConvertUsageResponseBodyData {
	return s.Data
}

func (s *DescribeMeterImsMediaConvertUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMeterImsMediaConvertUsageResponseBody) SetData(v []*DescribeMeterImsMediaConvertUsageResponseBodyData) *DescribeMeterImsMediaConvertUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterImsMediaConvertUsageResponseBody) SetRequestId(v string) *DescribeMeterImsMediaConvertUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUsageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMeterImsMediaConvertUsageResponseBodyData struct {
	// The usage duration, in minutes.
	//
	// example:
	//
	// 20
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The transcoding specifications.
	//
	// example:
	//
	// H264.HD
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The beginning time of usage. The value is a 10-digit timestamp.
	//
	// example:
	//
	// 1656950400
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeMeterImsMediaConvertUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMediaConvertUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMediaConvertUsageResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeMeterImsMediaConvertUsageResponseBodyData) GetSpecification() *string {
	return s.Specification
}

func (s *DescribeMeterImsMediaConvertUsageResponseBodyData) GetTime() *int64 {
	return s.Time
}

func (s *DescribeMeterImsMediaConvertUsageResponseBodyData) SetDuration(v int64) *DescribeMeterImsMediaConvertUsageResponseBodyData {
	s.Duration = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUsageResponseBodyData) SetSpecification(v string) *DescribeMeterImsMediaConvertUsageResponseBodyData {
	s.Specification = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUsageResponseBodyData) SetTime(v int64) *DescribeMeterImsMediaConvertUsageResponseBodyData {
	s.Time = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
