// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeDomainQpsListResponseBodyDataList) *DescribeDomainQpsListResponseBody
	GetDataList() []*DescribeDomainQpsListResponseBodyDataList
	SetRequestId(v string) *DescribeDomainQpsListResponseBody
	GetRequestId() *string
}

type DescribeDomainQpsListResponseBody struct {
	DataList  []*DescribeDomainQpsListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainQpsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsListResponseBody) GetDataList() []*DescribeDomainQpsListResponseBodyDataList {
	return s.DataList
}

func (s *DescribeDomainQpsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainQpsListResponseBody) SetDataList(v []*DescribeDomainQpsListResponseBodyDataList) *DescribeDomainQpsListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeDomainQpsListResponseBody) SetRequestId(v string) *DescribeDomainQpsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainQpsListResponseBodyDataList struct {
	AttackQps    *int64 `json:"AttackQps,omitempty" xml:"AttackQps,omitempty"`
	CacheHits    *int64 `json:"CacheHits,omitempty" xml:"CacheHits,omitempty"`
	Index        *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	MaxAttackQps *int64 `json:"MaxAttackQps,omitempty" xml:"MaxAttackQps,omitempty"`
	MaxNormalQps *int64 `json:"MaxNormalQps,omitempty" xml:"MaxNormalQps,omitempty"`
	MaxQps       *int64 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalQps     *int64 `json:"TotalQps,omitempty" xml:"TotalQps,omitempty"`
}

func (s DescribeDomainQpsListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsListResponseBodyDataList) GetAttackQps() *int64 {
	return s.AttackQps
}

func (s *DescribeDomainQpsListResponseBodyDataList) GetCacheHits() *int64 {
	return s.CacheHits
}

func (s *DescribeDomainQpsListResponseBodyDataList) GetIndex() *int64 {
	return s.Index
}

func (s *DescribeDomainQpsListResponseBodyDataList) GetMaxAttackQps() *int64 {
	return s.MaxAttackQps
}

func (s *DescribeDomainQpsListResponseBodyDataList) GetMaxNormalQps() *int64 {
	return s.MaxNormalQps
}

func (s *DescribeDomainQpsListResponseBodyDataList) GetMaxQps() *int64 {
	return s.MaxQps
}

func (s *DescribeDomainQpsListResponseBodyDataList) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDomainQpsListResponseBodyDataList) GetTotalQps() *int64 {
	return s.TotalQps
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetAttackQps(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.AttackQps = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetCacheHits(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.CacheHits = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetIndex(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.Index = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetMaxAttackQps(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.MaxAttackQps = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetMaxNormalQps(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.MaxNormalQps = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetMaxQps(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.MaxQps = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetTotalCount(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetTotalQps(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.TotalQps = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
