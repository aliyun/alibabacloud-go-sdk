// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsWithCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody
	GetBlocks() []*string
	SetCacheHits(v []*string) *DescribeDomainQpsWithCacheResponseBody
	GetCacheHits() []*string
	SetCcBlockQps(v []*string) *DescribeDomainQpsWithCacheResponseBody
	GetCcBlockQps() []*string
	SetCcJsQps(v []*string) *DescribeDomainQpsWithCacheResponseBody
	GetCcJsQps() []*string
	SetInterval(v int32) *DescribeDomainQpsWithCacheResponseBody
	GetInterval() *int32
	SetIpBlockQps(v []*string) *DescribeDomainQpsWithCacheResponseBody
	GetIpBlockQps() []*string
	SetPreciseBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody
	GetPreciseBlocks() []*string
	SetPreciseJsQps(v []*string) *DescribeDomainQpsWithCacheResponseBody
	GetPreciseJsQps() []*string
	SetRegionBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody
	GetRegionBlocks() []*string
	SetRequestId(v string) *DescribeDomainQpsWithCacheResponseBody
	GetRequestId() *string
	SetStartTime(v int64) *DescribeDomainQpsWithCacheResponseBody
	GetStartTime() *int64
	SetTotals(v []*string) *DescribeDomainQpsWithCacheResponseBody
	GetTotals() []*string
}

type DescribeDomainQpsWithCacheResponseBody struct {
	Blocks     []*string `json:"Blocks,omitempty" xml:"Blocks,omitempty" type:"Repeated"`
	CacheHits  []*string `json:"CacheHits,omitempty" xml:"CacheHits,omitempty" type:"Repeated"`
	CcBlockQps []*string `json:"CcBlockQps,omitempty" xml:"CcBlockQps,omitempty" type:"Repeated"`
	CcJsQps    []*string `json:"CcJsQps,omitempty" xml:"CcJsQps,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	Interval      *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IpBlockQps    []*string `json:"IpBlockQps,omitempty" xml:"IpBlockQps,omitempty" type:"Repeated"`
	PreciseBlocks []*string `json:"PreciseBlocks,omitempty" xml:"PreciseBlocks,omitempty" type:"Repeated"`
	PreciseJsQps  []*string `json:"PreciseJsQps,omitempty" xml:"PreciseJsQps,omitempty" type:"Repeated"`
	RegionBlocks  []*string `json:"RegionBlocks,omitempty" xml:"RegionBlocks,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1577794500
	StartTime *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Totals    []*string `json:"Totals,omitempty" xml:"Totals,omitempty" type:"Repeated"`
}

func (s DescribeDomainQpsWithCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsWithCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetBlocks() []*string {
	return s.Blocks
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetCacheHits() []*string {
	return s.CacheHits
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetCcBlockQps() []*string {
	return s.CcBlockQps
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetCcJsQps() []*string {
	return s.CcJsQps
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetIpBlockQps() []*string {
	return s.IpBlockQps
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetPreciseBlocks() []*string {
	return s.PreciseBlocks
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetPreciseJsQps() []*string {
	return s.PreciseJsQps
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetRegionBlocks() []*string {
	return s.RegionBlocks
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainQpsWithCacheResponseBody) GetTotals() []*string {
	return s.Totals
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.Blocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCacheHits(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CacheHits = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCcBlockQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CcBlockQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCcJsQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CcJsQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetInterval(v int32) *DescribeDomainQpsWithCacheResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetIpBlockQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.IpBlockQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetPreciseBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.PreciseBlocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetPreciseJsQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.PreciseJsQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetRegionBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.RegionBlocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetRequestId(v string) *DescribeDomainQpsWithCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetStartTime(v int64) *DescribeDomainQpsWithCacheResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetTotals(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.Totals = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
