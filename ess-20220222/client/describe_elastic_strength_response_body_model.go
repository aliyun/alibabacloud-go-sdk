// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticStrengthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticStrength(v string) *DescribeElasticStrengthResponseBody
	GetElasticStrength() *string
	SetElasticStrengthModels(v []*DescribeElasticStrengthResponseBodyElasticStrengthModels) *DescribeElasticStrengthResponseBody
	GetElasticStrengthModels() []*DescribeElasticStrengthResponseBodyElasticStrengthModels
	SetRequestId(v string) *DescribeElasticStrengthResponseBody
	GetRequestId() *string
	SetResourcePools(v []*DescribeElasticStrengthResponseBodyResourcePools) *DescribeElasticStrengthResponseBody
	GetResourcePools() []*DescribeElasticStrengthResponseBodyResourcePools
	SetTotalStrength(v float64) *DescribeElasticStrengthResponseBody
	GetTotalStrength() *float64
}

type DescribeElasticStrengthResponseBody struct {
	// The scaling strength level of the scaling group. Valid values:
	//
	// 	- Strong
	//
	// 	- Medium
	//
	// 	- Weak
	//
	// example:
	//
	// Strong
	ElasticStrength *string `json:"ElasticStrength,omitempty" xml:"ElasticStrength,omitempty"`
	// The scaling strength models.
	ElasticStrengthModels []*DescribeElasticStrengthResponseBodyElasticStrengthModels `json:"ElasticStrengthModels,omitempty" xml:"ElasticStrengthModels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 73469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource pools.
	ResourcePools []*DescribeElasticStrengthResponseBodyResourcePools `json:"ResourcePools,omitempty" xml:"ResourcePools,omitempty" type:"Repeated"`
	// The scaling strength score of the scaling group. Each combination of instance type + zone is scored from 0 to 1 based on its availability, with 0 being the weakest scaling strength and 1 being the strongest. The scaling strength score of the scaling group is measured by the combined scores of all the combinations of instance type + zone.
	//
	// **
	//
	// **Warning*	- This parameter is deprecated.
	//
	// example:
	//
	// 1.5
	TotalStrength *float64 `json:"TotalStrength,omitempty" xml:"TotalStrength,omitempty"`
}

func (s DescribeElasticStrengthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticStrengthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticStrengthResponseBody) GetElasticStrength() *string {
	return s.ElasticStrength
}

func (s *DescribeElasticStrengthResponseBody) GetElasticStrengthModels() []*DescribeElasticStrengthResponseBodyElasticStrengthModels {
	return s.ElasticStrengthModels
}

func (s *DescribeElasticStrengthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticStrengthResponseBody) GetResourcePools() []*DescribeElasticStrengthResponseBodyResourcePools {
	return s.ResourcePools
}

func (s *DescribeElasticStrengthResponseBody) GetTotalStrength() *float64 {
	return s.TotalStrength
}

func (s *DescribeElasticStrengthResponseBody) SetElasticStrength(v string) *DescribeElasticStrengthResponseBody {
	s.ElasticStrength = &v
	return s
}

func (s *DescribeElasticStrengthResponseBody) SetElasticStrengthModels(v []*DescribeElasticStrengthResponseBodyElasticStrengthModels) *DescribeElasticStrengthResponseBody {
	s.ElasticStrengthModels = v
	return s
}

func (s *DescribeElasticStrengthResponseBody) SetRequestId(v string) *DescribeElasticStrengthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticStrengthResponseBody) SetResourcePools(v []*DescribeElasticStrengthResponseBodyResourcePools) *DescribeElasticStrengthResponseBody {
	s.ResourcePools = v
	return s
}

func (s *DescribeElasticStrengthResponseBody) SetTotalStrength(v float64) *DescribeElasticStrengthResponseBody {
	s.TotalStrength = &v
	return s
}

func (s *DescribeElasticStrengthResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticStrengthResponseBodyElasticStrengthModels struct {
	// The scaling strength level of the scaling group. Valid values:
	//
	// 	- Strong
	//
	// 	- Medium
	//
	// 	- Weak
	//
	// example:
	//
	// Strong
	ElasticStrength *string `json:"ElasticStrength,omitempty" xml:"ElasticStrength,omitempty"`
	// The resource pools.
	ResourcePools []*DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools `json:"ResourcePools,omitempty" xml:"ResourcePools,omitempty" type:"Repeated"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-wz98mnj7nblv9gc****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The scaling strength score of the scaling group. Each combination of instance type + zone is scored from 0 to 1 based on its availability, with 0 being the weakest scaling strength and 1 being the strongest. The scaling strength score of the scaling group is measured by the combined scores of all the combinations of instance type + zone.
	//
	// **
	//
	// **Warning*	- This parameter is deprecated.
	//
	// example:
	//
	// 1.5
	TotalStrength *float64 `json:"TotalStrength,omitempty" xml:"TotalStrength,omitempty"`
}

func (s DescribeElasticStrengthResponseBodyElasticStrengthModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticStrengthResponseBodyElasticStrengthModels) GoString() string {
	return s.String()
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModels) GetElasticStrength() *string {
	return s.ElasticStrength
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModels) GetResourcePools() []*DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools {
	return s.ResourcePools
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModels) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModels) GetTotalStrength() *float64 {
	return s.TotalStrength
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModels) SetElasticStrength(v string) *DescribeElasticStrengthResponseBodyElasticStrengthModels {
	s.ElasticStrength = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModels) SetResourcePools(v []*DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) *DescribeElasticStrengthResponseBodyElasticStrengthModels {
	s.ResourcePools = v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModels) SetScalingGroupId(v string) *DescribeElasticStrengthResponseBodyElasticStrengthModels {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModels) SetTotalStrength(v float64) *DescribeElasticStrengthResponseBodyElasticStrengthModels {
	s.TotalStrength = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModels) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools struct {
	// The error code returned when the scaling strength is the weakest.
	//
	// example:
	//
	// InstanceTypesOrDiskTypesNotSupported
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance type of the resource pool.
	//
	// example:
	//
	// ecs.r7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The inventory health.
	InventoryHealth *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth `json:"InventoryHealth,omitempty" xml:"InventoryHealth,omitempty" type:"Struct"`
	// The error message returned when the scaling strength is the weakest.
	//
	// example:
	//
	// The instanceTypes or diskTypes are not supported.
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Indicates whether the resource pool is available. Valid values:
	//
	// 	- Available
	//
	// 	- Unavailable (If a constraint is not provided, the instance type is not deployed, or the instance type is out of stock, the resource pool becomes unavailable.)
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The scaling strength of the resource pool.
	//
	// **
	//
	// **Warning*	- This parameter is deprecated.
	//
	// example:
	//
	// 0.6
	Strength *float64 `json:"Strength,omitempty" xml:"Strength,omitempty"`
	// The IDs of the vSwitches in the zones of the resource pool.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The zone ID of the resource pool.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) GoString() string {
	return s.String()
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) GetCode() *string {
	return s.Code
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) GetInventoryHealth() *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth {
	return s.InventoryHealth
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) GetMsg() *string {
	return s.Msg
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) GetStatus() *string {
	return s.Status
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) GetStrength() *float64 {
	return s.Strength
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) SetCode(v string) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools {
	s.Code = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) SetInstanceType(v string) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools {
	s.InstanceType = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) SetInventoryHealth(v *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools {
	s.InventoryHealth = v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) SetMsg(v string) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools {
	s.Msg = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) SetStatus(v string) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools {
	s.Status = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) SetStrength(v float64) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools {
	s.Strength = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) SetVSwitchIds(v []*string) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools {
	s.VSwitchIds = v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) SetZoneId(v string) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools {
	s.ZoneId = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePools) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth struct {
	// The adequacy score.
	//
	// Valid values: 0 to 3.
	//
	// example:
	//
	// 3
	AdequacyScore *int32 `json:"AdequacyScore,omitempty" xml:"AdequacyScore,omitempty"`
	// The score of the inventory health.
	//
	// 	- A score between 5 and 6 indicates a sufficient inventory.
	//
	// 	- A score between 1 and 4 indicates that there is no guarantee of a sufficient inventory. Select a reservation as necessary.
	//
	// 	- A score between -3 and 0 indicates that the inventory is sufficient, and an alert is triggered. Select another instance type.
	//
	// Calculation formula: `HealthScore` = `AdequacyScore` + `SupplyScore` - `HotScore`.
	//
	// example:
	//
	// 3
	HealthScore *int32 `json:"HealthScore,omitempty" xml:"HealthScore,omitempty"`
	// The popularity score.
	//
	// Valid values: 0 to 3.
	//
	// example:
	//
	// 0
	HotScore *int32 `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	// The score of the replenishment capability.
	//
	// Valid values: 0 to 3.
	//
	// example:
	//
	// 2
	SupplyScore *int32 `json:"SupplyScore,omitempty" xml:"SupplyScore,omitempty"`
}

func (s DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) GoString() string {
	return s.String()
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) GetAdequacyScore() *int32 {
	return s.AdequacyScore
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) GetHealthScore() *int32 {
	return s.HealthScore
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) GetHotScore() *int32 {
	return s.HotScore
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) GetSupplyScore() *int32 {
	return s.SupplyScore
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) SetAdequacyScore(v int32) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth {
	s.AdequacyScore = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) SetHealthScore(v int32) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth {
	s.HealthScore = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) SetHotScore(v int32) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth {
	s.HotScore = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) SetSupplyScore(v int32) *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth {
	s.SupplyScore = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyElasticStrengthModelsResourcePoolsInventoryHealth) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticStrengthResponseBodyResourcePools struct {
	// The error code returned when the scaling strength is the weakest.
	//
	// example:
	//
	// IMG_NOT_SUPPORTED
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance type of the resource pool.
	//
	// example:
	//
	// ecs.c7t.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The inventory health.
	InventoryHealth *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth `json:"InventoryHealth,omitempty" xml:"InventoryHealth,omitempty" type:"Struct"`
	// The error message returned when the scaling strength is the weakest.
	//
	// example:
	//
	// The instanceType does not support the image in the configuration.
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Indicates whether the resource pool is available. Valid values:
	//
	// 	- Available
	//
	// 	- Unavailable (If a constraint is not provided, the instance type is not deployed, or the instance type is out of stock, the resource pool becomes unavailable.)
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The scaling strength of the resource pool.
	//
	// example:
	//
	// 0.6
	Strength *float64 `json:"Strength,omitempty" xml:"Strength,omitempty"`
	// The IDs of the vSwitches in the zones of the resource pool.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The zone ID of the resource pool.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeElasticStrengthResponseBodyResourcePools) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticStrengthResponseBodyResourcePools) GoString() string {
	return s.String()
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) GetCode() *string {
	return s.Code
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) GetInventoryHealth() *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth {
	return s.InventoryHealth
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) GetMsg() *string {
	return s.Msg
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) GetStatus() *string {
	return s.Status
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) GetStrength() *float64 {
	return s.Strength
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) SetCode(v string) *DescribeElasticStrengthResponseBodyResourcePools {
	s.Code = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) SetInstanceType(v string) *DescribeElasticStrengthResponseBodyResourcePools {
	s.InstanceType = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) SetInventoryHealth(v *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) *DescribeElasticStrengthResponseBodyResourcePools {
	s.InventoryHealth = v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) SetMsg(v string) *DescribeElasticStrengthResponseBodyResourcePools {
	s.Msg = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) SetStatus(v string) *DescribeElasticStrengthResponseBodyResourcePools {
	s.Status = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) SetStrength(v float64) *DescribeElasticStrengthResponseBodyResourcePools {
	s.Strength = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) SetVSwitchIds(v []*string) *DescribeElasticStrengthResponseBodyResourcePools {
	s.VSwitchIds = v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) SetZoneId(v string) *DescribeElasticStrengthResponseBodyResourcePools {
	s.ZoneId = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePools) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth struct {
	// The adequacy score.
	//
	// Valid values: 0 to 3.
	//
	// example:
	//
	// 3
	AdequacyScore *int32 `json:"AdequacyScore,omitempty" xml:"AdequacyScore,omitempty"`
	// The inventory health score.
	//
	// 	- A score between 5 and 6 indicates a sufficient inventory.
	//
	// 	- A score between 1 and 4 indicates that there is no guarantee of a sufficient inventory. Select a reservation as necessary.
	//
	// 	- A score between -3 and 0 indicates that the inventory is sufficient, and an alert is triggered. Select another instance type.
	//
	// Calculation formula: `HealthScore` = `AdequacyScore` + `SupplyScore` - `HotScore`.
	//
	// example:
	//
	// 3
	HealthScore *int32 `json:"HealthScore,omitempty" xml:"HealthScore,omitempty"`
	// The popularity score.
	//
	// Valid values: 0 to 3.
	//
	// example:
	//
	// 3
	HotScore *int32 `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	// The replenishment capability score.
	//
	// Valid values: 0 to 3.
	//
	// example:
	//
	// 3
	SupplyScore *int32 `json:"SupplyScore,omitempty" xml:"SupplyScore,omitempty"`
}

func (s DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) GoString() string {
	return s.String()
}

func (s *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) GetAdequacyScore() *int32 {
	return s.AdequacyScore
}

func (s *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) GetHealthScore() *int32 {
	return s.HealthScore
}

func (s *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) GetHotScore() *int32 {
	return s.HotScore
}

func (s *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) GetSupplyScore() *int32 {
	return s.SupplyScore
}

func (s *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) SetAdequacyScore(v int32) *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth {
	s.AdequacyScore = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) SetHealthScore(v int32) *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth {
	s.HealthScore = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) SetHotScore(v int32) *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth {
	s.HotScore = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) SetSupplyScore(v int32) *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth {
	s.SupplyScore = &v
	return s
}

func (s *DescribeElasticStrengthResponseBodyResourcePoolsInventoryHealth) Validate() error {
	return dara.Validate(s)
}
