// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePlansModificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlans(v []*ListInstancePlansModificationResponseBodyPlans) *ListInstancePlansModificationResponseBody
	GetPlans() []*ListInstancePlansModificationResponseBodyPlans
	SetRequestId(v string) *ListInstancePlansModificationResponseBody
	GetRequestId() *string
}

type ListInstancePlansModificationResponseBody struct {
	// The array of plans.
	Plans []*ListInstancePlansModificationResponseBodyPlans `json:"Plans,omitempty" xml:"Plans,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancePlansModificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePlansModificationResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationResponseBody) GetPlans() []*ListInstancePlansModificationResponseBodyPlans {
	return s.Plans
}

func (s *ListInstancePlansModificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancePlansModificationResponseBody) SetPlans(v []*ListInstancePlansModificationResponseBodyPlans) *ListInstancePlansModificationResponseBody {
	s.Plans = v
	return s
}

func (s *ListInstancePlansModificationResponseBody) SetRequestId(v string) *ListInstancePlansModificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancePlansModificationResponseBody) Validate() error {
	if s.Plans != nil {
		for _, item := range s.Plans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancePlansModificationResponseBodyPlans struct {
	// The peak bandwidth. Unit: Mbps.
	//
	// example:
	//
	// 200
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of vCPU cores.
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The unit of the price. Valid values:
	//
	// - CNY: Chinese Yuan
	//
	// - USD: US Dollar
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The disk capacity. Unit: GiB.
	//
	// example:
	//
	// 40
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// - SSD: a standard SSD.
	//
	// - ESSD: an Enhanced SSD (ESSD) at performance level (PL) 0.
	//
	// - CLOUD_EFFICIENCY: an ultra disk.
	//
	// example:
	//
	// ESSD
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The monthly traffic. Unit: GB.
	//
	// example:
	//
	// 400
	Flow *int32 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The memory size. Unit: GiB.
	//
	// example:
	//
	// 1
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The price of the plan.
	//
	// example:
	//
	// 60
	OriginPrice *float64 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	// The plan ID.
	//
	// example:
	//
	// swas.s2.c2m1s40b3t04
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The operating systems supported by the plan.
	//
	// example:
	//
	// ["Linux","Windows"]
	SupportPlatform *string `json:"SupportPlatform,omitempty" xml:"SupportPlatform,omitempty"`
}

func (s ListInstancePlansModificationResponseBodyPlans) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePlansModificationResponseBodyPlans) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationResponseBodyPlans) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListInstancePlansModificationResponseBodyPlans) GetCore() *int32 {
	return s.Core
}

func (s *ListInstancePlansModificationResponseBodyPlans) GetCurrency() *string {
	return s.Currency
}

func (s *ListInstancePlansModificationResponseBodyPlans) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *ListInstancePlansModificationResponseBodyPlans) GetDiskType() *string {
	return s.DiskType
}

func (s *ListInstancePlansModificationResponseBodyPlans) GetFlow() *int32 {
	return s.Flow
}

func (s *ListInstancePlansModificationResponseBodyPlans) GetMemory() *int32 {
	return s.Memory
}

func (s *ListInstancePlansModificationResponseBodyPlans) GetOriginPrice() *float64 {
	return s.OriginPrice
}

func (s *ListInstancePlansModificationResponseBodyPlans) GetPlanId() *string {
	return s.PlanId
}

func (s *ListInstancePlansModificationResponseBodyPlans) GetSupportPlatform() *string {
	return s.SupportPlatform
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetBandwidth(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Bandwidth = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetCore(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Core = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetCurrency(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.Currency = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetDiskSize(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.DiskSize = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetDiskType(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.DiskType = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetFlow(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Flow = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetMemory(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Memory = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetOriginPrice(v float64) *ListInstancePlansModificationResponseBodyPlans {
	s.OriginPrice = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetPlanId(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.PlanId = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetSupportPlatform(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.SupportPlatform = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) Validate() error {
	return dara.Validate(s)
}
