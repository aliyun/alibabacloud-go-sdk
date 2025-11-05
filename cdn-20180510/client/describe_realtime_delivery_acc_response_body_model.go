// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRealtimeDeliveryAccResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReatTimeDeliveryAccData(v *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) *DescribeRealtimeDeliveryAccResponseBody
	GetReatTimeDeliveryAccData() *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData
	SetRequestId(v string) *DescribeRealtimeDeliveryAccResponseBody
	GetRequestId() *string
}

type DescribeRealtimeDeliveryAccResponseBody struct {
	// The statistics about real-time log deliveries.
	ReatTimeDeliveryAccData *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData `json:"ReatTimeDeliveryAccData,omitempty" xml:"ReatTimeDeliveryAccData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 684306D2-2511-4977-991D-CE97E91FD7C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRealtimeDeliveryAccResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRealtimeDeliveryAccResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRealtimeDeliveryAccResponseBody) GetReatTimeDeliveryAccData() *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData {
	return s.ReatTimeDeliveryAccData
}

func (s *DescribeRealtimeDeliveryAccResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRealtimeDeliveryAccResponseBody) SetReatTimeDeliveryAccData(v *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) *DescribeRealtimeDeliveryAccResponseBody {
	s.ReatTimeDeliveryAccData = v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponseBody) SetRequestId(v string) *DescribeRealtimeDeliveryAccResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponseBody) Validate() error {
	if s.ReatTimeDeliveryAccData != nil {
		if err := s.ReatTimeDeliveryAccData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData struct {
	AccData []*DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData `json:"AccData,omitempty" xml:"AccData,omitempty" type:"Repeated"`
}

func (s DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) GoString() string {
	return s.String()
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) GetAccData() []*DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData {
	return s.AccData
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) SetAccData(v []*DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData {
	s.AccData = v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) Validate() error {
	if s.AccData != nil {
		for _, item := range s.AccData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData struct {
	// The number of failed attempts to deliver log data to Log Service.
	//
	// example:
	//
	// 2
	FailedNum *int32 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	// The number of successful deliveries of log data to Log Service.
	//
	// example:
	//
	// 2
	SuccessNum *int32 `json:"SuccessNum,omitempty" xml:"SuccessNum,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2018-09-03T06:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) GoString() string {
	return s.String()
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) GetFailedNum() *int32 {
	return s.FailedNum
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) GetSuccessNum() *int32 {
	return s.SuccessNum
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) SetFailedNum(v int32) *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData {
	s.FailedNum = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) SetSuccessNum(v int32) *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData {
	s.SuccessNum = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) SetTimeStamp(v string) *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) Validate() error {
	return dara.Validate(s)
}
