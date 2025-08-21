// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQPSListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainQPSList(v []*DescribeDomainQPSListResponseBodyDomainQPSList) *DescribeDomainQPSListResponseBody
	GetDomainQPSList() []*DescribeDomainQPSListResponseBodyDomainQPSList
	SetRequestId(v string) *DescribeDomainQPSListResponseBody
	GetRequestId() *string
}

type DescribeDomainQPSListResponseBody struct {
	// An array that consists of the statistics on the QPS of the website.
	DomainQPSList []*DescribeDomainQPSListResponseBodyDomainQPSList `json:"DomainQPSList,omitempty" xml:"DomainQPSList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 327F2ABB-104D-437A-AAB5-D633E29A8C51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainQPSListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQPSListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQPSListResponseBody) GetDomainQPSList() []*DescribeDomainQPSListResponseBodyDomainQPSList {
	return s.DomainQPSList
}

func (s *DescribeDomainQPSListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainQPSListResponseBody) SetDomainQPSList(v []*DescribeDomainQPSListResponseBodyDomainQPSList) *DescribeDomainQPSListResponseBody {
	s.DomainQPSList = v
	return s
}

func (s *DescribeDomainQPSListResponseBody) SetRequestId(v string) *DescribeDomainQPSListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQPSListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainQPSListResponseBodyDomainQPSList struct {
	// The attack QPS.
	//
	// example:
	//
	// 1
	AttackQps *int64 `json:"AttackQps,omitempty" xml:"AttackQps,omitempty"`
	// The number of cache hits.
	//
	// example:
	//
	// 0
	CacheHits *int64 `json:"CacheHits,omitempty" xml:"CacheHits,omitempty"`
	// The index number of the returned data.
	//
	// example:
	//
	// 0
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The peak attack QPS.
	//
	// example:
	//
	// 37
	MaxAttackQps *int64 `json:"MaxAttackQps,omitempty" xml:"MaxAttackQps,omitempty"`
	// The peak of normal QPS.
	//
	// example:
	//
	// 93
	MaxNormalQps *int64 `json:"MaxNormalQps,omitempty" xml:"MaxNormalQps,omitempty"`
	// The peak of total QPS.
	//
	// example:
	//
	// 130
	MaxQps *int64 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	// The time when the data was collected. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1582992000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 20008
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total QPS.
	//
	// example:
	//
	// 1
	TotalQps *int64 `json:"TotalQps,omitempty" xml:"TotalQps,omitempty"`
}

func (s DescribeDomainQPSListResponseBodyDomainQPSList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQPSListResponseBodyDomainQPSList) GoString() string {
	return s.String()
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) GetAttackQps() *int64 {
	return s.AttackQps
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) GetCacheHits() *int64 {
	return s.CacheHits
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) GetIndex() *int64 {
	return s.Index
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) GetMaxAttackQps() *int64 {
	return s.MaxAttackQps
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) GetMaxNormalQps() *int64 {
	return s.MaxNormalQps
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) GetMaxQps() *int64 {
	return s.MaxQps
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) GetTotalQps() *int64 {
	return s.TotalQps
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetAttackQps(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.AttackQps = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetCacheHits(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.CacheHits = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetIndex(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.Index = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetMaxAttackQps(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.MaxAttackQps = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetMaxNormalQps(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.MaxNormalQps = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetMaxQps(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.MaxQps = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetTime(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.Time = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetTotalCount(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetTotalQps(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.TotalQps = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) Validate() error {
	return dara.Validate(s)
}
