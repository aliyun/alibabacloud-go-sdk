// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRealtimeDeliveryAccResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRealTimeDeliveryAccData(v *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) *DescribeLiveRealtimeDeliveryAccResponseBody
	GetRealTimeDeliveryAccData() *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData
	SetRequestId(v string) *DescribeLiveRealtimeDeliveryAccResponseBody
	GetRequestId() *string
}

type DescribeLiveRealtimeDeliveryAccResponseBody struct {
	// The information about real-time log deliveries.
	RealTimeDeliveryAccData *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData `json:"RealTimeDeliveryAccData,omitempty" xml:"RealTimeDeliveryAccData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 684306D2-2511-4977-991D-CE97E91FD7C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveRealtimeDeliveryAccResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRealtimeDeliveryAccResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBody) GetRealTimeDeliveryAccData() *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData {
	return s.RealTimeDeliveryAccData
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBody) SetRealTimeDeliveryAccData(v *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) *DescribeLiveRealtimeDeliveryAccResponseBody {
	s.RealTimeDeliveryAccData = v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBody) SetRequestId(v string) *DescribeLiveRealtimeDeliveryAccResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData struct {
	AccData []*DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData `json:"AccData,omitempty" xml:"AccData,omitempty" type:"Repeated"`
}

func (s DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) GetAccData() []*DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData {
	return s.AccData
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) SetAccData(v []*DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData {
	s.AccData = v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccData) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData struct {
	// The number of failed real-time log deliveries.
	//
	// example:
	//
	// 0
	FailedNum *int32 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	// The number of successful real-time log deliveries.
	//
	// example:
	//
	// 321321
	SuccessNum *int32 `json:"SuccessNum,omitempty" xml:"SuccessNum,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) GetFailedNum() *int32 {
	return s.FailedNum
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) GetSuccessNum() *int32 {
	return s.SuccessNum
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) SetFailedNum(v int32) *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData {
	s.FailedNum = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) SetSuccessNum(v int32) *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData {
	s.SuccessNum = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) SetTimeStamp(v string) *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponseBodyRealTimeDeliveryAccDataAccData) Validate() error {
	return dara.Validate(s)
}
