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
	// The minimum number of pay-as-you-go instances that are required for the node group. Valid values: 0 to 1000. If the number of pay-as-you-go instances is less than the value of this parameter, pay-as-you-go instances are preferentially created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances among the instances that exceed the number specified by the OnDemandBaseCapacity parameter for the node group. Valid values: 0 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// The number of instance types that are specified. Preemptible instances of multiple instance types that are provided at the lowest price are evenly created in the scaling group. Valid values: 0 to 10.
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
