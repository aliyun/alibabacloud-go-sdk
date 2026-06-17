// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClassListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeClassListResponseBodyItems) *DescribeClassListResponseBody
	GetItems() []*DescribeClassListResponseBodyItems
	SetRegionId(v string) *DescribeClassListResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeClassListResponseBody
	GetRequestId() *string
}

type DescribeClassListResponseBody struct {
	// The list of cluster specifications.
	Items []*DescribeClassListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 69A85BAF-1089-4CDF-A82F-0A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClassListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClassListResponseBody) GetItems() []*DescribeClassListResponseBodyItems {
	return s.Items
}

func (s *DescribeClassListResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClassListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClassListResponseBody) SetItems(v []*DescribeClassListResponseBodyItems) *DescribeClassListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeClassListResponseBody) SetRegionId(v string) *DescribeClassListResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeClassListResponseBody) SetRequestId(v string) *DescribeClassListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClassListResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClassListResponseBodyItems struct {
	// The cluster specifications.
	//
	// example:
	//
	// polar.mysql.x4.medium
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The family of the cluster specifications. Valid values:
	//
	// - Exclusive package
	//
	// - Exclusive physical machine
	//
	// - Beginner
	//
	// - Historical specifications
	//
	// example:
	//
	// Exclusive package
	ClassGroup *string `json:"ClassGroup,omitempty" xml:"ClassGroup,omitempty"`
	// The specification type.
	//
	// example:
	//
	// enterprise
	ClassTypeLevel *string `json:"ClassTypeLevel,omitempty" xml:"ClassTypeLevel,omitempty"`
	// The number of CPU cores. Unit: cores.
	//
	// example:
	//
	// 8
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The maximum storage capacity of an ESSD. Unit: TB.
	//
	// example:
	//
	// 64
	EssdMaxStorageCapacity *string `json:"EssdMaxStorageCapacity,omitempty" xml:"EssdMaxStorageCapacity,omitempty"`
	// The maximum number of concurrent connections to the cluster.
	//
	// example:
	//
	// 8000
	MaxConnections *string `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum IOPS. Unit: IOPS.
	//
	// example:
	//
	// 32000
	MaxIOPS *string `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// The maximum storage capacity. Unit: TB.
	//
	// example:
	//
	// 20
	MaxStorageCapacity *string `json:"MaxStorageCapacity,omitempty" xml:"MaxStorageCapacity,omitempty"`
	// The memory capacity. Unit: GB.
	//
	// example:
	//
	// 32
	MemoryClass *string `json:"MemoryClass,omitempty" xml:"MemoryClass,omitempty"`
	// The maximum IOPS of an Enhanced SSD (ESSD) at performance level 1 (PL1). Unit: IOPS.
	//
	// example:
	//
	// 50000
	Pl1MaxIOPS *string `json:"Pl1MaxIOPS,omitempty" xml:"Pl1MaxIOPS,omitempty"`
	// The maximum IOPS of an ESSD at PL2. Unit: IOPS.
	//
	// example:
	//
	// 100000
	Pl2MaxIOPS *string `json:"Pl2MaxIOPS,omitempty" xml:"Pl2MaxIOPS,omitempty"`
	// The maximum IOPS of an ESSD at PL3. Unit: IOPS.
	//
	// example:
	//
	// 1000000
	Pl3MaxIOPS *string `json:"Pl3MaxIOPS,omitempty" xml:"Pl3MaxIOPS,omitempty"`
	// The maximum storage capacity of PSL4/PSL5. Unit: TB.
	//
	// example:
	//
	// 500
	PolarStoreMaxStorageCapacity *string `json:"PolarStoreMaxStorageCapacity,omitempty" xml:"PolarStoreMaxStorageCapacity,omitempty"`
	// The maximum input/output operations per second (IOPS) of PSL4. Unit: IOPS.
	//
	// example:
	//
	// 48000
	Psl4MaxIOPS *string `json:"Psl4MaxIOPS,omitempty" xml:"Psl4MaxIOPS,omitempty"`
	// The maximum IOPS of PSL5. Unit: IOPS.
	//
	// example:
	//
	// 96000
	Psl5MaxIOPS *string `json:"Psl5MaxIOPS,omitempty" xml:"Psl5MaxIOPS,omitempty"`
	// The additional price.
	//
	// <props="china">Unit: cents (CNY).
	//
	// <props="intl">Unit: cents (USD).
	//
	// > - If you set the MasterHa parameter to cluster or single, the value of this parameter is the same as the value of the ReferencePrice parameter.
	//
	// >
	//
	// > - If you set the MasterHa parameter to cluster or single, the price for a single-node commodity is returned.
	//
	// example:
	//
	// 200000
	ReferenceExtPrice *string `json:"ReferenceExtPrice,omitempty" xml:"ReferenceExtPrice,omitempty"`
	// The price.
	//
	// <props="china">Unit: cents (CNY).
	//
	// <props="intl">Unit: cents (USD).
	//
	// > - If you set the CommodityCode parameter to a pay-as-you-go commodity code, the hourly price is returned.
	//
	// >
	//
	// > - If you set the CommodityCode parameter to a subscription commodity code, the monthly price is returned.
	//
	// example:
	//
	// 200000
	ReferencePrice *string `json:"ReferencePrice,omitempty" xml:"ReferencePrice,omitempty"`
}

func (s DescribeClassListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeClassListResponseBodyItems) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeClassListResponseBodyItems) GetClassGroup() *string {
	return s.ClassGroup
}

func (s *DescribeClassListResponseBodyItems) GetClassTypeLevel() *string {
	return s.ClassTypeLevel
}

func (s *DescribeClassListResponseBodyItems) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeClassListResponseBodyItems) GetEssdMaxStorageCapacity() *string {
	return s.EssdMaxStorageCapacity
}

func (s *DescribeClassListResponseBodyItems) GetMaxConnections() *string {
	return s.MaxConnections
}

func (s *DescribeClassListResponseBodyItems) GetMaxIOPS() *string {
	return s.MaxIOPS
}

func (s *DescribeClassListResponseBodyItems) GetMaxStorageCapacity() *string {
	return s.MaxStorageCapacity
}

func (s *DescribeClassListResponseBodyItems) GetMemoryClass() *string {
	return s.MemoryClass
}

func (s *DescribeClassListResponseBodyItems) GetPl1MaxIOPS() *string {
	return s.Pl1MaxIOPS
}

func (s *DescribeClassListResponseBodyItems) GetPl2MaxIOPS() *string {
	return s.Pl2MaxIOPS
}

func (s *DescribeClassListResponseBodyItems) GetPl3MaxIOPS() *string {
	return s.Pl3MaxIOPS
}

func (s *DescribeClassListResponseBodyItems) GetPolarStoreMaxStorageCapacity() *string {
	return s.PolarStoreMaxStorageCapacity
}

func (s *DescribeClassListResponseBodyItems) GetPsl4MaxIOPS() *string {
	return s.Psl4MaxIOPS
}

func (s *DescribeClassListResponseBodyItems) GetPsl5MaxIOPS() *string {
	return s.Psl5MaxIOPS
}

func (s *DescribeClassListResponseBodyItems) GetReferenceExtPrice() *string {
	return s.ReferenceExtPrice
}

func (s *DescribeClassListResponseBodyItems) GetReferencePrice() *string {
	return s.ReferencePrice
}

func (s *DescribeClassListResponseBodyItems) SetClassCode(v string) *DescribeClassListResponseBodyItems {
	s.ClassCode = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetClassGroup(v string) *DescribeClassListResponseBodyItems {
	s.ClassGroup = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetClassTypeLevel(v string) *DescribeClassListResponseBodyItems {
	s.ClassTypeLevel = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetCpu(v string) *DescribeClassListResponseBodyItems {
	s.Cpu = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetEssdMaxStorageCapacity(v string) *DescribeClassListResponseBodyItems {
	s.EssdMaxStorageCapacity = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetMaxConnections(v string) *DescribeClassListResponseBodyItems {
	s.MaxConnections = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetMaxIOPS(v string) *DescribeClassListResponseBodyItems {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetMaxStorageCapacity(v string) *DescribeClassListResponseBodyItems {
	s.MaxStorageCapacity = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetMemoryClass(v string) *DescribeClassListResponseBodyItems {
	s.MemoryClass = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetPl1MaxIOPS(v string) *DescribeClassListResponseBodyItems {
	s.Pl1MaxIOPS = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetPl2MaxIOPS(v string) *DescribeClassListResponseBodyItems {
	s.Pl2MaxIOPS = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetPl3MaxIOPS(v string) *DescribeClassListResponseBodyItems {
	s.Pl3MaxIOPS = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetPolarStoreMaxStorageCapacity(v string) *DescribeClassListResponseBodyItems {
	s.PolarStoreMaxStorageCapacity = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetPsl4MaxIOPS(v string) *DescribeClassListResponseBodyItems {
	s.Psl4MaxIOPS = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetPsl5MaxIOPS(v string) *DescribeClassListResponseBodyItems {
	s.Psl5MaxIOPS = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetReferenceExtPrice(v string) *DescribeClassListResponseBodyItems {
	s.ReferenceExtPrice = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) SetReferencePrice(v string) *DescribeClassListResponseBodyItems {
	s.ReferencePrice = &v
	return s
}

func (s *DescribeClassListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
