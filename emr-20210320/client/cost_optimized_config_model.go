// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostOptimizedConfig interface {
	dara.Model
	String() string
	GoString() string
	SetOnDemandBaseCapacity(v int32) *CostOptimizedConfig
	GetOnDemandBaseCapacity() *int32
	SetOnDemandPercentageAboveBaseCapacity(v int32) *CostOptimizedConfig
	GetOnDemandPercentageAboveBaseCapacity() *int32
	SetSpotInstancePools(v int32) *CostOptimizedConfig
	GetSpotInstancePools() *int32
}

type CostOptimizedConfig struct {
	// 按量实例个数的最小值。节点组所需要按量实例个数的最小值，取值范围：0~1000。当按量实例个数少于该值时，将优先创建按量实例。
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// 节点组满足最小按量实例OnDemandBaseCapacity要求后，超出的实例中按量实例应占的比例，取值范围：0～100。
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// 指定可用实例规格的个数，伸缩组将按成本最低的多个规格均衡创建抢占式实例。取值范围：0~10。
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	SpotInstancePools *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
}

func (s CostOptimizedConfig) String() string {
	return dara.Prettify(s)
}

func (s CostOptimizedConfig) GoString() string {
	return s.String()
}

func (s *CostOptimizedConfig) GetOnDemandBaseCapacity() *int32 {
	return s.OnDemandBaseCapacity
}

func (s *CostOptimizedConfig) GetOnDemandPercentageAboveBaseCapacity() *int32 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *CostOptimizedConfig) GetSpotInstancePools() *int32 {
	return s.SpotInstancePools
}

func (s *CostOptimizedConfig) SetOnDemandBaseCapacity(v int32) *CostOptimizedConfig {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *CostOptimizedConfig) SetOnDemandPercentageAboveBaseCapacity(v int32) *CostOptimizedConfig {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *CostOptimizedConfig) SetSpotInstancePools(v int32) *CostOptimizedConfig {
	s.SpotInstancePools = &v
	return s
}

func (s *CostOptimizedConfig) Validate() error {
	return dara.Validate(s)
}
