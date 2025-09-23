// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAttackEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeDomainAttackEventListResponseBodyDataList) *DescribeDomainAttackEventListResponseBody
	GetDataList() []*DescribeDomainAttackEventListResponseBodyDataList
	SetRequestId(v string) *DescribeDomainAttackEventListResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeDomainAttackEventListResponseBody
	GetTotal() *int64
}

type DescribeDomainAttackEventListResponseBody struct {
	DataList []*DescribeDomainAttackEventListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDomainAttackEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventListResponseBody) GetDataList() []*DescribeDomainAttackEventListResponseBodyDataList {
	return s.DataList
}

func (s *DescribeDomainAttackEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainAttackEventListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDomainAttackEventListResponseBody) SetDataList(v []*DescribeDomainAttackEventListResponseBodyDataList) *DescribeDomainAttackEventListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeDomainAttackEventListResponseBody) SetRequestId(v string) *DescribeDomainAttackEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAttackEventListResponseBody) SetTotal(v int64) *DescribeDomainAttackEventListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDomainAttackEventListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainAttackEventListResponseBodyDataList struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1670918400
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 300
	MaxQps *int64 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	// example:
	//
	// 1666083600
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainAttackEventListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackEventListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) GetMaxQps() *int64 {
	return s.MaxQps
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) SetDomain(v string) *DescribeDomainAttackEventListResponseBodyDataList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) SetEndTime(v int64) *DescribeDomainAttackEventListResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) SetMaxQps(v int64) *DescribeDomainAttackEventListResponseBodyDataList {
	s.MaxQps = &v
	return s
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) SetStartTime(v int64) *DescribeDomainAttackEventListResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
