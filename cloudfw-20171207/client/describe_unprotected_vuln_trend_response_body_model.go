// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnprotectedVulnTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurProtectedCnt(v int32) *DescribeUnprotectedVulnTrendResponseBody
	GetCurProtectedCnt() *int32
	SetCurUnprotectedCnt(v int32) *DescribeUnprotectedVulnTrendResponseBody
	GetCurUnprotectedCnt() *int32
	SetDataList(v []*DescribeUnprotectedVulnTrendResponseBodyDataList) *DescribeUnprotectedVulnTrendResponseBody
	GetDataList() []*DescribeUnprotectedVulnTrendResponseBodyDataList
	SetEndTime(v int64) *DescribeUnprotectedVulnTrendResponseBody
	GetEndTime() *int64
	SetInterval(v int32) *DescribeUnprotectedVulnTrendResponseBody
	GetInterval() *int32
	SetRequestId(v string) *DescribeUnprotectedVulnTrendResponseBody
	GetRequestId() *string
	SetStartTime(v int64) *DescribeUnprotectedVulnTrendResponseBody
	GetStartTime() *int64
}

type DescribeUnprotectedVulnTrendResponseBody struct {
	// example:
	//
	// 7
	CurProtectedCnt *int32 `json:"CurProtectedCnt,omitempty" xml:"CurProtectedCnt,omitempty"`
	// example:
	//
	// 8
	CurUnprotectedCnt *int32                                              `json:"CurUnprotectedCnt,omitempty" xml:"CurUnprotectedCnt,omitempty"`
	DataList          []*DescribeUnprotectedVulnTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 1731551104
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 05BEF2B5-EAAA-509D-9824-E3C7DC17****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1749434787
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUnprotectedVulnTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnprotectedVulnTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUnprotectedVulnTrendResponseBody) GetCurProtectedCnt() *int32 {
	return s.CurProtectedCnt
}

func (s *DescribeUnprotectedVulnTrendResponseBody) GetCurUnprotectedCnt() *int32 {
	return s.CurUnprotectedCnt
}

func (s *DescribeUnprotectedVulnTrendResponseBody) GetDataList() []*DescribeUnprotectedVulnTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeUnprotectedVulnTrendResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeUnprotectedVulnTrendResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeUnprotectedVulnTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUnprotectedVulnTrendResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeUnprotectedVulnTrendResponseBody) SetCurProtectedCnt(v int32) *DescribeUnprotectedVulnTrendResponseBody {
	s.CurProtectedCnt = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponseBody) SetCurUnprotectedCnt(v int32) *DescribeUnprotectedVulnTrendResponseBody {
	s.CurUnprotectedCnt = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponseBody) SetDataList(v []*DescribeUnprotectedVulnTrendResponseBodyDataList) *DescribeUnprotectedVulnTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponseBody) SetEndTime(v int64) *DescribeUnprotectedVulnTrendResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponseBody) SetInterval(v int32) *DescribeUnprotectedVulnTrendResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponseBody) SetRequestId(v string) *DescribeUnprotectedVulnTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponseBody) SetStartTime(v int64) *DescribeUnprotectedVulnTrendResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponseBody) Validate() error {
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

type DescribeUnprotectedVulnTrendResponseBodyDataList struct {
	// example:
	//
	// 3
	ProtectedVulnCnt *int32 `json:"ProtectedVulnCnt,omitempty" xml:"ProtectedVulnCnt,omitempty"`
	// example:
	//
	// 1525833105
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 2
	UnprotectedVulnCnt *int32 `json:"UnprotectedVulnCnt,omitempty" xml:"UnprotectedVulnCnt,omitempty"`
}

func (s DescribeUnprotectedVulnTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnprotectedVulnTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeUnprotectedVulnTrendResponseBodyDataList) GetProtectedVulnCnt() *int32 {
	return s.ProtectedVulnCnt
}

func (s *DescribeUnprotectedVulnTrendResponseBodyDataList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeUnprotectedVulnTrendResponseBodyDataList) GetUnprotectedVulnCnt() *int32 {
	return s.UnprotectedVulnCnt
}

func (s *DescribeUnprotectedVulnTrendResponseBodyDataList) SetProtectedVulnCnt(v int32) *DescribeUnprotectedVulnTrendResponseBodyDataList {
	s.ProtectedVulnCnt = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponseBodyDataList) SetTime(v int64) *DescribeUnprotectedVulnTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponseBodyDataList) SetUnprotectedVulnCnt(v int32) *DescribeUnprotectedVulnTrendResponseBodyDataList {
	s.UnprotectedVulnCnt = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
