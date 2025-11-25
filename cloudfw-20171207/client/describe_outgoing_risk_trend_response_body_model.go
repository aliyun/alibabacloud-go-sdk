// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingRiskTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeOutgoingRiskTrendResponseBodyDataList) *DescribeOutgoingRiskTrendResponseBody
	GetDataList() []*DescribeOutgoingRiskTrendResponseBodyDataList
	SetInterval(v int32) *DescribeOutgoingRiskTrendResponseBody
	GetInterval() *int32
	SetRequestId(v string) *DescribeOutgoingRiskTrendResponseBody
	GetRequestId() *string
	SetTotalRiskDomain(v int32) *DescribeOutgoingRiskTrendResponseBody
	GetTotalRiskDomain() *int32
	SetTotalRiskIp(v int32) *DescribeOutgoingRiskTrendResponseBody
	GetTotalRiskIp() *int32
}

type DescribeOutgoingRiskTrendResponseBody struct {
	DataList []*DescribeOutgoingRiskTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 1CA8D98E-A71B-5856-A658-3E8B3152E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 7
	TotalRiskDomain *int32 `json:"TotalRiskDomain,omitempty" xml:"TotalRiskDomain,omitempty"`
	// example:
	//
	// 6
	TotalRiskIp *int32 `json:"TotalRiskIp,omitempty" xml:"TotalRiskIp,omitempty"`
}

func (s DescribeOutgoingRiskTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingRiskTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingRiskTrendResponseBody) GetDataList() []*DescribeOutgoingRiskTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeOutgoingRiskTrendResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeOutgoingRiskTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingRiskTrendResponseBody) GetTotalRiskDomain() *int32 {
	return s.TotalRiskDomain
}

func (s *DescribeOutgoingRiskTrendResponseBody) GetTotalRiskIp() *int32 {
	return s.TotalRiskIp
}

func (s *DescribeOutgoingRiskTrendResponseBody) SetDataList(v []*DescribeOutgoingRiskTrendResponseBodyDataList) *DescribeOutgoingRiskTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeOutgoingRiskTrendResponseBody) SetInterval(v int32) *DescribeOutgoingRiskTrendResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeOutgoingRiskTrendResponseBody) SetRequestId(v string) *DescribeOutgoingRiskTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingRiskTrendResponseBody) SetTotalRiskDomain(v int32) *DescribeOutgoingRiskTrendResponseBody {
	s.TotalRiskDomain = &v
	return s
}

func (s *DescribeOutgoingRiskTrendResponseBody) SetTotalRiskIp(v int32) *DescribeOutgoingRiskTrendResponseBody {
	s.TotalRiskIp = &v
	return s
}

func (s *DescribeOutgoingRiskTrendResponseBody) Validate() error {
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

type DescribeOutgoingRiskTrendResponseBodyDataList struct {
	// example:
	//
	// 20
	RiskDomain *int32 `json:"RiskDomain,omitempty" xml:"RiskDomain,omitempty"`
	// example:
	//
	// 5
	RiskIp *int32 `json:"RiskIp,omitempty" xml:"RiskIp,omitempty"`
	// example:
	//
	// 1659405600
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeOutgoingRiskTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingRiskTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingRiskTrendResponseBodyDataList) GetRiskDomain() *int32 {
	return s.RiskDomain
}

func (s *DescribeOutgoingRiskTrendResponseBodyDataList) GetRiskIp() *int32 {
	return s.RiskIp
}

func (s *DescribeOutgoingRiskTrendResponseBodyDataList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeOutgoingRiskTrendResponseBodyDataList) SetRiskDomain(v int32) *DescribeOutgoingRiskTrendResponseBodyDataList {
	s.RiskDomain = &v
	return s
}

func (s *DescribeOutgoingRiskTrendResponseBodyDataList) SetRiskIp(v int32) *DescribeOutgoingRiskTrendResponseBodyDataList {
	s.RiskIp = &v
	return s
}

func (s *DescribeOutgoingRiskTrendResponseBodyDataList) SetTime(v int64) *DescribeOutgoingRiskTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeOutgoingRiskTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
