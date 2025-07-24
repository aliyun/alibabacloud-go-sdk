// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingGroupConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiskCategory(v string) *ScalingGroupConfig
	GetDataDiskCategory() *string
	SetDataDiskCount(v int32) *ScalingGroupConfig
	GetDataDiskCount() *int32
	SetDataDiskSize(v int64) *ScalingGroupConfig
	GetDataDiskSize() *int64
	SetDefaultCoolDownTime(v int64) *ScalingGroupConfig
	GetDefaultCoolDownTime() *int64
	SetInstanceTypeList(v []*ScalingGroupConfigInstanceTypeList) *ScalingGroupConfig
	GetInstanceTypeList() []*ScalingGroupConfigInstanceTypeList
	SetMultiAvailablePolicy(v *ScalingGroupConfigMultiAvailablePolicy) *ScalingGroupConfig
	GetMultiAvailablePolicy() *ScalingGroupConfigMultiAvailablePolicy
	SetNodeOfflinePolicy(v *ScalingGroupConfigNodeOfflinePolicy) *ScalingGroupConfig
	GetNodeOfflinePolicy() *ScalingGroupConfigNodeOfflinePolicy
	SetPrivatePoolOptions(v *ScalingGroupConfigPrivatePoolOptions) *ScalingGroupConfig
	GetPrivatePoolOptions() *ScalingGroupConfigPrivatePoolOptions
	SetScalingMaxSize(v int32) *ScalingGroupConfig
	GetScalingMaxSize() *int32
	SetScalingMinSize(v int32) *ScalingGroupConfig
	GetScalingMinSize() *int32
	SetSpotStrategy(v string) *ScalingGroupConfig
	GetSpotStrategy() *string
	SetSysDiskCategory(v string) *ScalingGroupConfig
	GetSysDiskCategory() *string
	SetSysDiskSize(v int64) *ScalingGroupConfig
	GetSysDiskSize() *int64
	SetTriggerMode(v string) *ScalingGroupConfig
	GetTriggerMode() *string
}

type ScalingGroupConfig struct {
	// 数据盘类型。
	//
	// example:
	//
	// cloud_essd
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// 数据盘个数。
	//
	// example:
	//
	// 4
	DataDiskCount *int32 `json:"DataDiskCount,omitempty" xml:"DataDiskCount,omitempty"`
	// 数据盘大小,单位GB。
	//
	// example:
	//
	// 40
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// 默认冷却时间。
	//
	// example:
	//
	// 0
	DefaultCoolDownTime *int64 `json:"DefaultCoolDownTime,omitempty" xml:"DefaultCoolDownTime,omitempty"`
	// 抢占实例列表。
	InstanceTypeList []*ScalingGroupConfigInstanceTypeList `json:"InstanceTypeList,omitempty" xml:"InstanceTypeList,omitempty" type:"Repeated"`
	// 资源可用性策略(成本优化参数)。
	MultiAvailablePolicy *ScalingGroupConfigMultiAvailablePolicy `json:"MultiAvailablePolicy,omitempty" xml:"MultiAvailablePolicy,omitempty" type:"Struct"`
	// 节点下线策略。
	NodeOfflinePolicy *ScalingGroupConfigNodeOfflinePolicy `json:"NodeOfflinePolicy,omitempty" xml:"NodeOfflinePolicy,omitempty" type:"Struct"`
	// 私有池选项	。
	PrivatePoolOptions *ScalingGroupConfigPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// 伸缩组节点最大个数。
	//
	// example:
	//
	// 10
	ScalingMaxSize *int32 `json:"ScalingMaxSize,omitempty" xml:"ScalingMaxSize,omitempty"`
	// 伸缩组节点最小个数。
	//
	// example:
	//
	// 1
	ScalingMinSize *int32 `json:"ScalingMinSize,omitempty" xml:"ScalingMinSize,omitempty"`
	// 抢占式Spot实例策略。
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// 系统盘类型。
	//
	// example:
	//
	// cloud_essd
	SysDiskCategory *string `json:"SysDiskCategory,omitempty" xml:"SysDiskCategory,omitempty"`
	// 系统盘大小,单位GB。
	//
	// example:
	//
	// 80
	SysDiskSize *int64 `json:"SysDiskSize,omitempty" xml:"SysDiskSize,omitempty"`
	// 伸缩活动触发模式。
	//
	// example:
	//
	// ByLoad
	TriggerMode *string `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
}

func (s ScalingGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s ScalingGroupConfig) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfig) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *ScalingGroupConfig) GetDataDiskCount() *int32 {
	return s.DataDiskCount
}

func (s *ScalingGroupConfig) GetDataDiskSize() *int64 {
	return s.DataDiskSize
}

func (s *ScalingGroupConfig) GetDefaultCoolDownTime() *int64 {
	return s.DefaultCoolDownTime
}

func (s *ScalingGroupConfig) GetInstanceTypeList() []*ScalingGroupConfigInstanceTypeList {
	return s.InstanceTypeList
}

func (s *ScalingGroupConfig) GetMultiAvailablePolicy() *ScalingGroupConfigMultiAvailablePolicy {
	return s.MultiAvailablePolicy
}

func (s *ScalingGroupConfig) GetNodeOfflinePolicy() *ScalingGroupConfigNodeOfflinePolicy {
	return s.NodeOfflinePolicy
}

func (s *ScalingGroupConfig) GetPrivatePoolOptions() *ScalingGroupConfigPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *ScalingGroupConfig) GetScalingMaxSize() *int32 {
	return s.ScalingMaxSize
}

func (s *ScalingGroupConfig) GetScalingMinSize() *int32 {
	return s.ScalingMinSize
}

func (s *ScalingGroupConfig) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *ScalingGroupConfig) GetSysDiskCategory() *string {
	return s.SysDiskCategory
}

func (s *ScalingGroupConfig) GetSysDiskSize() *int64 {
	return s.SysDiskSize
}

func (s *ScalingGroupConfig) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *ScalingGroupConfig) SetDataDiskCategory(v string) *ScalingGroupConfig {
	s.DataDiskCategory = &v
	return s
}

func (s *ScalingGroupConfig) SetDataDiskCount(v int32) *ScalingGroupConfig {
	s.DataDiskCount = &v
	return s
}

func (s *ScalingGroupConfig) SetDataDiskSize(v int64) *ScalingGroupConfig {
	s.DataDiskSize = &v
	return s
}

func (s *ScalingGroupConfig) SetDefaultCoolDownTime(v int64) *ScalingGroupConfig {
	s.DefaultCoolDownTime = &v
	return s
}

func (s *ScalingGroupConfig) SetInstanceTypeList(v []*ScalingGroupConfigInstanceTypeList) *ScalingGroupConfig {
	s.InstanceTypeList = v
	return s
}

func (s *ScalingGroupConfig) SetMultiAvailablePolicy(v *ScalingGroupConfigMultiAvailablePolicy) *ScalingGroupConfig {
	s.MultiAvailablePolicy = v
	return s
}

func (s *ScalingGroupConfig) SetNodeOfflinePolicy(v *ScalingGroupConfigNodeOfflinePolicy) *ScalingGroupConfig {
	s.NodeOfflinePolicy = v
	return s
}

func (s *ScalingGroupConfig) SetPrivatePoolOptions(v *ScalingGroupConfigPrivatePoolOptions) *ScalingGroupConfig {
	s.PrivatePoolOptions = v
	return s
}

func (s *ScalingGroupConfig) SetScalingMaxSize(v int32) *ScalingGroupConfig {
	s.ScalingMaxSize = &v
	return s
}

func (s *ScalingGroupConfig) SetScalingMinSize(v int32) *ScalingGroupConfig {
	s.ScalingMinSize = &v
	return s
}

func (s *ScalingGroupConfig) SetSpotStrategy(v string) *ScalingGroupConfig {
	s.SpotStrategy = &v
	return s
}

func (s *ScalingGroupConfig) SetSysDiskCategory(v string) *ScalingGroupConfig {
	s.SysDiskCategory = &v
	return s
}

func (s *ScalingGroupConfig) SetSysDiskSize(v int64) *ScalingGroupConfig {
	s.SysDiskSize = &v
	return s
}

func (s *ScalingGroupConfig) SetTriggerMode(v string) *ScalingGroupConfig {
	s.TriggerMode = &v
	return s
}

func (s *ScalingGroupConfig) Validate() error {
	return dara.Validate(s)
}

type ScalingGroupConfigInstanceTypeList struct {
	// Ecs类型。
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 抢占价格上限,可空。
	//
	// example:
	//
	// 0.79
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
}

func (s ScalingGroupConfigInstanceTypeList) String() string {
	return dara.Prettify(s)
}

func (s ScalingGroupConfigInstanceTypeList) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfigInstanceTypeList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ScalingGroupConfigInstanceTypeList) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *ScalingGroupConfigInstanceTypeList) SetInstanceType(v string) *ScalingGroupConfigInstanceTypeList {
	s.InstanceType = &v
	return s
}

func (s *ScalingGroupConfigInstanceTypeList) SetSpotPriceLimit(v float32) *ScalingGroupConfigInstanceTypeList {
	s.SpotPriceLimit = &v
	return s
}

func (s *ScalingGroupConfigInstanceTypeList) Validate() error {
	return dara.Validate(s)
}

type ScalingGroupConfigMultiAvailablePolicy struct {
	// 资源可用性策略(成本优化参数)。
	PolicyParam *ScalingGroupConfigMultiAvailablePolicyPolicyParam `json:"PolicyParam,omitempty" xml:"PolicyParam,omitempty" type:"Struct"`
	// 策略类型。
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ScalingGroupConfigMultiAvailablePolicy) String() string {
	return dara.Prettify(s)
}

func (s ScalingGroupConfigMultiAvailablePolicy) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfigMultiAvailablePolicy) GetPolicyParam() *ScalingGroupConfigMultiAvailablePolicyPolicyParam {
	return s.PolicyParam
}

func (s *ScalingGroupConfigMultiAvailablePolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ScalingGroupConfigMultiAvailablePolicy) SetPolicyParam(v *ScalingGroupConfigMultiAvailablePolicyPolicyParam) *ScalingGroupConfigMultiAvailablePolicy {
	s.PolicyParam = v
	return s
}

func (s *ScalingGroupConfigMultiAvailablePolicy) SetPolicyType(v string) *ScalingGroupConfigMultiAvailablePolicy {
	s.PolicyType = &v
	return s
}

func (s *ScalingGroupConfigMultiAvailablePolicy) Validate() error {
	return dara.Validate(s)
}

type ScalingGroupConfigMultiAvailablePolicyPolicyParam struct {
	// 按需实例最小个数。
	//
	// example:
	//
	// 1
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// 按需实例百分比。
	//
	// example:
	//
	// 10
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// 抢占实例类型池规模。
	//
	// example:
	//
	// 10
	SpotInstancePools *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	// 是否使用按量补偿。
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
}

func (s ScalingGroupConfigMultiAvailablePolicyPolicyParam) String() string {
	return dara.Prettify(s)
}

func (s ScalingGroupConfigMultiAvailablePolicyPolicyParam) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) GetOnDemandBaseCapacity() *int32 {
	return s.OnDemandBaseCapacity
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) GetOnDemandPercentageAboveBaseCapacity() *int32 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) GetSpotInstancePools() *int32 {
	return s.SpotInstancePools
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) SetOnDemandBaseCapacity(v int32) *ScalingGroupConfigMultiAvailablePolicyPolicyParam {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) SetOnDemandPercentageAboveBaseCapacity(v int32) *ScalingGroupConfigMultiAvailablePolicyPolicyParam {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) SetSpotInstancePools(v int32) *ScalingGroupConfigMultiAvailablePolicyPolicyParam {
	s.SpotInstancePools = &v
	return s
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) SetSpotInstanceRemedy(v bool) *ScalingGroupConfigMultiAvailablePolicyPolicyParam {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) Validate() error {
	return dara.Validate(s)
}

type ScalingGroupConfigNodeOfflinePolicy struct {
	// 下线模式,是否为优雅下线。
	//
	// example:
	//
	// DEFAULT
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// 下线超时时间,单位毫秒。
	//
	// example:
	//
	// 1000
	TimeoutMs *int64 `json:"TimeoutMs,omitempty" xml:"TimeoutMs,omitempty"`
}

func (s ScalingGroupConfigNodeOfflinePolicy) String() string {
	return dara.Prettify(s)
}

func (s ScalingGroupConfigNodeOfflinePolicy) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfigNodeOfflinePolicy) GetMode() *string {
	return s.Mode
}

func (s *ScalingGroupConfigNodeOfflinePolicy) GetTimeoutMs() *int64 {
	return s.TimeoutMs
}

func (s *ScalingGroupConfigNodeOfflinePolicy) SetMode(v string) *ScalingGroupConfigNodeOfflinePolicy {
	s.Mode = &v
	return s
}

func (s *ScalingGroupConfigNodeOfflinePolicy) SetTimeoutMs(v int64) *ScalingGroupConfigNodeOfflinePolicy {
	s.TimeoutMs = &v
	return s
}

func (s *ScalingGroupConfigNodeOfflinePolicy) Validate() error {
	return dara.Validate(s)
}

type ScalingGroupConfigPrivatePoolOptions struct {
	// 私有池id。
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 实例启动的私有池容量选项。。
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s ScalingGroupConfigPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ScalingGroupConfigPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfigPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *ScalingGroupConfigPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *ScalingGroupConfigPrivatePoolOptions) SetId(v string) *ScalingGroupConfigPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ScalingGroupConfigPrivatePoolOptions) SetMatchCriteria(v string) *ScalingGroupConfigPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *ScalingGroupConfigPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
