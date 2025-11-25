// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEcsTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeInvadeEcsTrendResponseBodyDataList) *DescribeInvadeEcsTrendResponseBody
	GetDataList() []*DescribeInvadeEcsTrendResponseBodyDataList
	SetEndTime(v int64) *DescribeInvadeEcsTrendResponseBody
	GetEndTime() *int64
	SetInterval(v int32) *DescribeInvadeEcsTrendResponseBody
	GetInterval() *int32
	SetInvadeEcsCount(v int32) *DescribeInvadeEcsTrendResponseBody
	GetInvadeEcsCount() *int32
	SetRequestId(v string) *DescribeInvadeEcsTrendResponseBody
	GetRequestId() *string
	SetStartTime(v int64) *DescribeInvadeEcsTrendResponseBody
	GetStartTime() *int64
}

type DescribeInvadeEcsTrendResponseBody struct {
	DataList []*DescribeInvadeEcsTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 1736820365
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 2
	InvadeEcsCount *int32 `json:"InvadeEcsCount,omitempty" xml:"InvadeEcsCount,omitempty"`
	// example:
	//
	// F90E816D-BEE7-5BD6-95ED-474F54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1742177725
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvadeEcsTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEcsTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEcsTrendResponseBody) GetDataList() []*DescribeInvadeEcsTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeInvadeEcsTrendResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeInvadeEcsTrendResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeInvadeEcsTrendResponseBody) GetInvadeEcsCount() *int32 {
	return s.InvadeEcsCount
}

func (s *DescribeInvadeEcsTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvadeEcsTrendResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeInvadeEcsTrendResponseBody) SetDataList(v []*DescribeInvadeEcsTrendResponseBodyDataList) *DescribeInvadeEcsTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInvadeEcsTrendResponseBody) SetEndTime(v int64) *DescribeInvadeEcsTrendResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeInvadeEcsTrendResponseBody) SetInterval(v int32) *DescribeInvadeEcsTrendResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeInvadeEcsTrendResponseBody) SetInvadeEcsCount(v int32) *DescribeInvadeEcsTrendResponseBody {
	s.InvadeEcsCount = &v
	return s
}

func (s *DescribeInvadeEcsTrendResponseBody) SetRequestId(v string) *DescribeInvadeEcsTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvadeEcsTrendResponseBody) SetStartTime(v int64) *DescribeInvadeEcsTrendResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeInvadeEcsTrendResponseBody) Validate() error {
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

type DescribeInvadeEcsTrendResponseBodyDataList struct {
	// example:
	//
	// 27
	EcsCount *int32 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// example:
	//
	// 1659405600
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeInvadeEcsTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEcsTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEcsTrendResponseBodyDataList) GetEcsCount() *int32 {
	return s.EcsCount
}

func (s *DescribeInvadeEcsTrendResponseBodyDataList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeInvadeEcsTrendResponseBodyDataList) SetEcsCount(v int32) *DescribeInvadeEcsTrendResponseBodyDataList {
	s.EcsCount = &v
	return s
}

func (s *DescribeInvadeEcsTrendResponseBodyDataList) SetTime(v int64) *DescribeInvadeEcsTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeInvadeEcsTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
