// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKPlayFailStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBody
	GetDataInterval() *string
	SetEndTime(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBody
	GetEndTime() *string
	SetPlayFailStatus(v []*DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) *DescribeRTSNativeSDKPlayFailStatusResponseBody
	GetPlayFailStatus() []*DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus
	SetRequestId(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBody
	GetStartTime() *string
}

type DescribeRTSNativeSDKPlayFailStatusResponseBody struct {
	// The time granularity.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The end of the time range for which the data was queried.
	//
	// example:
	//
	// 2021-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of error status codes at each interval.
	PlayFailStatus []*DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus `json:"PlayFailStatus,omitempty" xml:"PlayFailStatus,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range for which the data was queried.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRTSNativeSDKPlayFailStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKPlayFailStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) GetPlayFailStatus() []*DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus {
	return s.PlayFailStatus
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) SetDataInterval(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) SetEndTime(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) SetPlayFailStatus(v []*DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) *DescribeRTSNativeSDKPlayFailStatusResponseBody {
	s.PlayFailStatus = v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) SetRequestId(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) SetStartTime(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The status code that indicates failed DNS resolution.
	//
	// example:
	//
	// 1
	V20001 *string `json:"V20001,omitempty" xml:"V20001,omitempty"`
	// The status code that indicates failed authentication.
	//
	// example:
	//
	// 2
	V20002 *string `json:"V20002,omitempty" xml:"V20002,omitempty"`
	// The status code that indicates a connection signaling timeout.
	//
	// example:
	//
	// 3
	V20011 *string `json:"V20011,omitempty" xml:"V20011,omitempty"`
	// The status code that indicates a subscription signaling error.
	//
	// example:
	//
	// 4
	V20012 *string `json:"V20012,omitempty" xml:"V20012,omitempty"`
	// The status code indicating that the stream to subscribe to does not exist.
	//
	// example:
	//
	// 5
	V20013 *string `json:"V20013,omitempty" xml:"V20013,omitempty"`
	// The status code that indicates a media packet collection timeout.
	//
	// example:
	//
	// 6
	V20052 *string `json:"V20052,omitempty" xml:"V20052,omitempty"`
}

func (s DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) GetV20001() *string {
	return s.V20001
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) GetV20002() *string {
	return s.V20002
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) GetV20011() *string {
	return s.V20011
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) GetV20012() *string {
	return s.V20012
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) GetV20013() *string {
	return s.V20013
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) GetV20052() *string {
	return s.V20052
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) SetTimeStamp(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) SetV20001(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus {
	s.V20001 = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) SetV20002(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus {
	s.V20002 = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) SetV20011(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus {
	s.V20011 = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) SetV20012(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus {
	s.V20012 = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) SetV20013(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus {
	s.V20013 = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) SetV20052(v string) *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus {
	s.V20052 = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponseBodyPlayFailStatus) Validate() error {
	return dara.Validate(s)
}
