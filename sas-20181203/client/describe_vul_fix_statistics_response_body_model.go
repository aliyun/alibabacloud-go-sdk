// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulFixStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFixStat(v []*DescribeVulFixStatisticsResponseBodyFixStat) *DescribeVulFixStatisticsResponseBody
	GetFixStat() []*DescribeVulFixStatisticsResponseBodyFixStat
	SetFixTotal(v *DescribeVulFixStatisticsResponseBodyFixTotal) *DescribeVulFixStatisticsResponseBody
	GetFixTotal() *DescribeVulFixStatisticsResponseBodyFixTotal
	SetRequestId(v string) *DescribeVulFixStatisticsResponseBody
	GetRequestId() *string
}

type DescribeVulFixStatisticsResponseBody struct {
	// An array that consists of the statistics of vulnerability fixes by vulnerability type.
	FixStat []*DescribeVulFixStatisticsResponseBodyFixStat `json:"FixStat,omitempty" xml:"FixStat,omitempty" type:"Repeated"`
	// The total statistics of vulnerability fixes.
	FixTotal *DescribeVulFixStatisticsResponseBodyFixTotal `json:"FixTotal,omitempty" xml:"FixTotal,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F3B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVulFixStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulFixStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulFixStatisticsResponseBody) GetFixStat() []*DescribeVulFixStatisticsResponseBodyFixStat {
	return s.FixStat
}

func (s *DescribeVulFixStatisticsResponseBody) GetFixTotal() *DescribeVulFixStatisticsResponseBodyFixTotal {
	return s.FixTotal
}

func (s *DescribeVulFixStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulFixStatisticsResponseBody) SetFixStat(v []*DescribeVulFixStatisticsResponseBodyFixStat) *DescribeVulFixStatisticsResponseBody {
	s.FixStat = v
	return s
}

func (s *DescribeVulFixStatisticsResponseBody) SetFixTotal(v *DescribeVulFixStatisticsResponseBodyFixTotal) *DescribeVulFixStatisticsResponseBody {
	s.FixTotal = v
	return s
}

func (s *DescribeVulFixStatisticsResponseBody) SetRequestId(v string) *DescribeVulFixStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulFixStatisticsResponseBody) Validate() error {
	if s.FixStat != nil {
		for _, item := range s.FixStat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FixTotal != nil {
		if err := s.FixTotal.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVulFixStatisticsResponseBodyFixStat struct {
	// The number of vulnerabilities that are fixed on the current day.
	//
	// example:
	//
	// 10
	FixedTodayNum *int32 `json:"FixedTodayNum,omitempty" xml:"FixedTodayNum,omitempty"`
	// The total number of fixed vulnerabilities.
	//
	// example:
	//
	// 22
	FixedTotalNum *int32 `json:"FixedTotalNum,omitempty" xml:"FixedTotalNum,omitempty"`
	// The number of vulnerabilities that are being fixed.
	//
	// example:
	//
	// 17
	FixingNum *int32 `json:"FixingNum,omitempty" xml:"FixingNum,omitempty"`
	// The number of unfixed vulnerabilities.
	//
	// example:
	//
	// 8
	NeedFixNum *int32 `json:"NeedFixNum,omitempty" xml:"NeedFixNum,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **app**: application vulnerability
	//
	// 	- **emg**: urgent vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulFixStatisticsResponseBodyFixStat) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulFixStatisticsResponseBodyFixStat) GoString() string {
	return s.String()
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) GetFixedTodayNum() *int32 {
	return s.FixedTodayNum
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) GetFixedTotalNum() *int32 {
	return s.FixedTotalNum
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) GetFixingNum() *int32 {
	return s.FixingNum
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) GetNeedFixNum() *int32 {
	return s.NeedFixNum
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) GetType() *string {
	return s.Type
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) SetFixedTodayNum(v int32) *DescribeVulFixStatisticsResponseBodyFixStat {
	s.FixedTodayNum = &v
	return s
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) SetFixedTotalNum(v int32) *DescribeVulFixStatisticsResponseBodyFixStat {
	s.FixedTotalNum = &v
	return s
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) SetFixingNum(v int32) *DescribeVulFixStatisticsResponseBodyFixStat {
	s.FixingNum = &v
	return s
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) SetNeedFixNum(v int32) *DescribeVulFixStatisticsResponseBodyFixStat {
	s.NeedFixNum = &v
	return s
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) SetType(v string) *DescribeVulFixStatisticsResponseBodyFixStat {
	s.Type = &v
	return s
}

func (s *DescribeVulFixStatisticsResponseBodyFixStat) Validate() error {
	return dara.Validate(s)
}

type DescribeVulFixStatisticsResponseBodyFixTotal struct {
	// The number of vulnerabilities that are fixed on the current day.
	//
	// example:
	//
	// 15
	FixedTodayNum *int32 `json:"FixedTodayNum,omitempty" xml:"FixedTodayNum,omitempty"`
	// The total number of fixed vulnerabilities.
	//
	// example:
	//
	// 47
	FixedTotalNum *int32 `json:"FixedTotalNum,omitempty" xml:"FixedTotalNum,omitempty"`
	// The number of vulnerabilities that are being fixed.
	//
	// example:
	//
	// 22
	FixingNum *int32 `json:"FixingNum,omitempty" xml:"FixingNum,omitempty"`
	// The number of unfixed vulnerabilities.
	//
	// example:
	//
	// 62
	NeedFixNum *int32 `json:"NeedFixNum,omitempty" xml:"NeedFixNum,omitempty"`
}

func (s DescribeVulFixStatisticsResponseBodyFixTotal) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulFixStatisticsResponseBodyFixTotal) GoString() string {
	return s.String()
}

func (s *DescribeVulFixStatisticsResponseBodyFixTotal) GetFixedTodayNum() *int32 {
	return s.FixedTodayNum
}

func (s *DescribeVulFixStatisticsResponseBodyFixTotal) GetFixedTotalNum() *int32 {
	return s.FixedTotalNum
}

func (s *DescribeVulFixStatisticsResponseBodyFixTotal) GetFixingNum() *int32 {
	return s.FixingNum
}

func (s *DescribeVulFixStatisticsResponseBodyFixTotal) GetNeedFixNum() *int32 {
	return s.NeedFixNum
}

func (s *DescribeVulFixStatisticsResponseBodyFixTotal) SetFixedTodayNum(v int32) *DescribeVulFixStatisticsResponseBodyFixTotal {
	s.FixedTodayNum = &v
	return s
}

func (s *DescribeVulFixStatisticsResponseBodyFixTotal) SetFixedTotalNum(v int32) *DescribeVulFixStatisticsResponseBodyFixTotal {
	s.FixedTotalNum = &v
	return s
}

func (s *DescribeVulFixStatisticsResponseBodyFixTotal) SetFixingNum(v int32) *DescribeVulFixStatisticsResponseBodyFixTotal {
	s.FixingNum = &v
	return s
}

func (s *DescribeVulFixStatisticsResponseBodyFixTotal) SetNeedFixNum(v int32) *DescribeVulFixStatisticsResponseBodyFixTotal {
	s.NeedFixNum = &v
	return s
}

func (s *DescribeVulFixStatisticsResponseBodyFixTotal) Validate() error {
	return dara.Validate(s)
}
