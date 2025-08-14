// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsMediaConvertUHDUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeMeterImsMediaConvertUHDUsageResponseBodyData) *DescribeMeterImsMediaConvertUHDUsageResponseBody
	GetData() []*DescribeMeterImsMediaConvertUHDUsageResponseBodyData
	SetRequestId(v string) *DescribeMeterImsMediaConvertUHDUsageResponseBody
	GetRequestId() *string
}

type DescribeMeterImsMediaConvertUHDUsageResponseBody struct {
	// The usage statistics of IMS on UHD transcoding of MPS.
	Data []*DescribeMeterImsMediaConvertUHDUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BEA98A0C-7870-15FE-B96F-8880BB600A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterImsMediaConvertUHDUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMediaConvertUHDUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBody) GetData() []*DescribeMeterImsMediaConvertUHDUsageResponseBodyData {
	return s.Data
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBody) SetData(v []*DescribeMeterImsMediaConvertUHDUsageResponseBodyData) *DescribeMeterImsMediaConvertUHDUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBody) SetRequestId(v string) *DescribeMeterImsMediaConvertUHDUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMeterImsMediaConvertUHDUsageResponseBodyData struct {
	// The usage duration, in minutes.
	//
	// example:
	//
	// 308028
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The transcoding specifications.
	//
	// example:
	//
	// SuperResolution.Standard.1080P
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The beginning time of usage. The value is a 10-digit timestamp.
	//
	// example:
	//
	// 1656950400
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeMeterImsMediaConvertUHDUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMediaConvertUHDUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBodyData) GetSpecification() *string {
	return s.Specification
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBodyData) GetTime() *int64 {
	return s.Time
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBodyData) SetDuration(v int64) *DescribeMeterImsMediaConvertUHDUsageResponseBodyData {
	s.Duration = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBodyData) SetSpecification(v string) *DescribeMeterImsMediaConvertUHDUsageResponseBodyData {
	s.Specification = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBodyData) SetTime(v int64) *DescribeMeterImsMediaConvertUHDUsageResponseBodyData {
	s.Time = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
