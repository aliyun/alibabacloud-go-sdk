// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnprotectedPortTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeUnprotectedPortTrendResponseBodyDataList) *DescribeUnprotectedPortTrendResponseBody
	GetDataList() []*DescribeUnprotectedPortTrendResponseBodyDataList
	SetInterval(v int32) *DescribeUnprotectedPortTrendResponseBody
	GetInterval() *int32
	SetRequestId(v string) *DescribeUnprotectedPortTrendResponseBody
	GetRequestId() *string
}

type DescribeUnprotectedPortTrendResponseBody struct {
	DataList []*DescribeUnprotectedPortTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// E599A84E-CD22-5E42-A2A9-01A254AC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUnprotectedPortTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnprotectedPortTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUnprotectedPortTrendResponseBody) GetDataList() []*DescribeUnprotectedPortTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeUnprotectedPortTrendResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeUnprotectedPortTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUnprotectedPortTrendResponseBody) SetDataList(v []*DescribeUnprotectedPortTrendResponseBodyDataList) *DescribeUnprotectedPortTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeUnprotectedPortTrendResponseBody) SetInterval(v int32) *DescribeUnprotectedPortTrendResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeUnprotectedPortTrendResponseBody) SetRequestId(v string) *DescribeUnprotectedPortTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUnprotectedPortTrendResponseBody) Validate() error {
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

type DescribeUnprotectedPortTrendResponseBodyDataList struct {
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1659405600
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeUnprotectedPortTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnprotectedPortTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeUnprotectedPortTrendResponseBodyDataList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeUnprotectedPortTrendResponseBodyDataList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeUnprotectedPortTrendResponseBodyDataList) SetCount(v int32) *DescribeUnprotectedPortTrendResponseBodyDataList {
	s.Count = &v
	return s
}

func (s *DescribeUnprotectedPortTrendResponseBodyDataList) SetTime(v int64) *DescribeUnprotectedPortTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeUnprotectedPortTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
