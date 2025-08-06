// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRealtimeDeliveryAccResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRealTimeDeliveryAccData(v *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) *DescribeVodRealtimeDeliveryAccResponseBody
	GetRealTimeDeliveryAccData() *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData
	SetRequestId(v string) *DescribeVodRealtimeDeliveryAccResponseBody
	GetRequestId() *string
}

type DescribeVodRealtimeDeliveryAccResponseBody struct {
	// The information about real-time log deliveries.
	RealTimeDeliveryAccData *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData `json:"RealTimeDeliveryAccData,omitempty" xml:"RealTimeDeliveryAccData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8AA0364-0FDB-4AD5-****-D69FAB8924ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodRealtimeDeliveryAccResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRealtimeDeliveryAccResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodRealtimeDeliveryAccResponseBody) GetRealTimeDeliveryAccData() *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData {
	return s.RealTimeDeliveryAccData
}

func (s *DescribeVodRealtimeDeliveryAccResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodRealtimeDeliveryAccResponseBody) SetRealTimeDeliveryAccData(v *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) *DescribeVodRealtimeDeliveryAccResponseBody {
	s.RealTimeDeliveryAccData = v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccResponseBody) SetRequestId(v string) *DescribeVodRealtimeDeliveryAccResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData struct {
	AccData []*DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData `json:"AccData,omitempty" xml:"AccData,omitempty" type:"Repeated"`
}

func (s DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) GoString() string {
	return s.String()
}

func (s *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) GetAccData() []*DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData {
	return s.AccData
}

func (s *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) SetAccData(v []*DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData {
	s.AccData = v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData struct {
	// The number of failed real-time log deliveries.
	//
	// example:
	//
	// 1
	FailedNum *int32 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	// The number of successful real-time log deliveries.
	//
	// example:
	//
	// 1
	SuccessNum *int32 `json:"SuccessNum,omitempty" xml:"SuccessNum,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-10-20T04:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) GoString() string {
	return s.String()
}

func (s *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) GetFailedNum() *int32 {
	return s.FailedNum
}

func (s *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) GetSuccessNum() *int32 {
	return s.SuccessNum
}

func (s *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) SetFailedNum(v int32) *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData {
	s.FailedNum = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) SetSuccessNum(v int32) *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData {
	s.SuccessNum = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) SetTimeStamp(v string) *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) Validate() error {
	return dara.Validate(s)
}
