// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetDropTrafficTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeInternetDropTrafficTrendResponseBodyDataList) *DescribeInternetDropTrafficTrendResponseBody
	GetDataList() []*DescribeInternetDropTrafficTrendResponseBodyDataList
	SetDropSessionMax(v int64) *DescribeInternetDropTrafficTrendResponseBody
	GetDropSessionMax() *int64
	SetRatioAverage(v string) *DescribeInternetDropTrafficTrendResponseBody
	GetRatioAverage() *string
	SetRequestId(v string) *DescribeInternetDropTrafficTrendResponseBody
	GetRequestId() *string
	SetRingRatioAverage(v string) *DescribeInternetDropTrafficTrendResponseBody
	GetRingRatioAverage() *string
}

type DescribeInternetDropTrafficTrendResponseBody struct {
	DataList []*DescribeInternetDropTrafficTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 8090
	DropSessionMax *int64 `json:"DropSessionMax,omitempty" xml:"DropSessionMax,omitempty"`
	// example:
	//
	// 12.34
	RatioAverage *string `json:"RatioAverage,omitempty" xml:"RatioAverage,omitempty"`
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1.23
	RingRatioAverage *string `json:"RingRatioAverage,omitempty" xml:"RingRatioAverage,omitempty"`
}

func (s DescribeInternetDropTrafficTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetDropTrafficTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetDropTrafficTrendResponseBody) GetDataList() []*DescribeInternetDropTrafficTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeInternetDropTrafficTrendResponseBody) GetDropSessionMax() *int64 {
	return s.DropSessionMax
}

func (s *DescribeInternetDropTrafficTrendResponseBody) GetRatioAverage() *string {
	return s.RatioAverage
}

func (s *DescribeInternetDropTrafficTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetDropTrafficTrendResponseBody) GetRingRatioAverage() *string {
	return s.RingRatioAverage
}

func (s *DescribeInternetDropTrafficTrendResponseBody) SetDataList(v []*DescribeInternetDropTrafficTrendResponseBodyDataList) *DescribeInternetDropTrafficTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBody) SetDropSessionMax(v int64) *DescribeInternetDropTrafficTrendResponseBody {
	s.DropSessionMax = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBody) SetRatioAverage(v string) *DescribeInternetDropTrafficTrendResponseBody {
	s.RatioAverage = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBody) SetRequestId(v string) *DescribeInternetDropTrafficTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBody) SetRingRatioAverage(v string) *DescribeInternetDropTrafficTrendResponseBody {
	s.RingRatioAverage = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInternetDropTrafficTrendResponseBodyDataList struct {
	// example:
	//
	// 12
	AclDrop *int64 `json:"AclDrop,omitempty" xml:"AclDrop,omitempty"`
	// example:
	//
	// 2018-08-25 12:00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// example:
	//
	// 1
	DropRatio *string `json:"DropRatio,omitempty" xml:"DropRatio,omitempty"`
	// example:
	//
	// 0
	DropRing *int64 `json:"DropRing,omitempty" xml:"DropRing,omitempty"`
	// example:
	//
	// 1
	DropRingRatio *string `json:"DropRingRatio,omitempty" xml:"DropRingRatio,omitempty"`
	// example:
	//
	// 12
	DropSession *int64 `json:"DropSession,omitempty" xml:"DropSession,omitempty"`
	// example:
	//
	// 5
	IpsDrop *int64 `json:"IpsDrop,omitempty" xml:"IpsDrop,omitempty"`
	// example:
	//
	// 2018-08-25 12:00:00
	RingDataTime *string `json:"RingDataTime,omitempty" xml:"RingDataTime,omitempty"`
	// example:
	//
	// 1724982259
	RingTime *int64 `json:"RingTime,omitempty" xml:"RingTime,omitempty"`
	// example:
	//
	// 1659405600
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 10
	TotalSession *int64 `json:"TotalSession,omitempty" xml:"TotalSession,omitempty"`
}

func (s DescribeInternetDropTrafficTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetDropTrafficTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetAclDrop() *int64 {
	return s.AclDrop
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetDataTime() *string {
	return s.DataTime
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetDropRatio() *string {
	return s.DropRatio
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetDropRing() *int64 {
	return s.DropRing
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetDropRingRatio() *string {
	return s.DropRingRatio
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetDropSession() *int64 {
	return s.DropSession
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetIpsDrop() *int64 {
	return s.IpsDrop
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetRingDataTime() *string {
	return s.RingDataTime
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetRingTime() *int64 {
	return s.RingTime
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) GetTotalSession() *int64 {
	return s.TotalSession
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetAclDrop(v int64) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.AclDrop = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetDataTime(v string) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.DataTime = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetDropRatio(v string) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.DropRatio = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetDropRing(v int64) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.DropRing = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetDropRingRatio(v string) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.DropRingRatio = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetDropSession(v int64) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.DropSession = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetIpsDrop(v int64) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.IpsDrop = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetRingDataTime(v string) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.RingDataTime = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetRingTime(v int64) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.RingTime = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetTime(v int64) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) SetTotalSession(v int64) *DescribeInternetDropTrafficTrendResponseBodyDataList {
	s.TotalSession = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
