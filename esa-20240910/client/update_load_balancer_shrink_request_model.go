// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptiveRoutingShrink(v string) *UpdateLoadBalancerShrinkRequest
	GetAdaptiveRoutingShrink() *string
	SetDefaultPoolsShrink(v string) *UpdateLoadBalancerShrinkRequest
	GetDefaultPoolsShrink() *string
	SetDescription(v string) *UpdateLoadBalancerShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdateLoadBalancerShrinkRequest
	GetEnabled() *bool
	SetFallbackPool(v int64) *UpdateLoadBalancerShrinkRequest
	GetFallbackPool() *int64
	SetId(v int64) *UpdateLoadBalancerShrinkRequest
	GetId() *int64
	SetMonitorShrink(v string) *UpdateLoadBalancerShrinkRequest
	GetMonitorShrink() *string
	SetRandomSteeringShrink(v string) *UpdateLoadBalancerShrinkRequest
	GetRandomSteeringShrink() *string
	SetRegionPools(v interface{}) *UpdateLoadBalancerShrinkRequest
	GetRegionPools() interface{}
	SetRulesShrink(v string) *UpdateLoadBalancerShrinkRequest
	GetRulesShrink() *string
	SetSessionAffinity(v string) *UpdateLoadBalancerShrinkRequest
	GetSessionAffinity() *string
	SetSiteId(v int64) *UpdateLoadBalancerShrinkRequest
	GetSiteId() *int64
	SetSteeringPolicy(v string) *UpdateLoadBalancerShrinkRequest
	GetSteeringPolicy() *string
	SetSubRegionPools(v interface{}) *UpdateLoadBalancerShrinkRequest
	GetSubRegionPools() interface{}
	SetTtl(v int32) *UpdateLoadBalancerShrinkRequest
	GetTtl() *int32
}

type UpdateLoadBalancerShrinkRequest struct {
	// Configuration for fallback across pools.
	AdaptiveRoutingShrink *string `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty"`
	// List of default pool IDs.
	DefaultPoolsShrink *string `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty"`
	// Detailed description of the load balancer, for easier management and identification.
	//
	// example:
	//
	// Load balancer description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether the load balancer is enabled.
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Fallback pool ID, where traffic will be directed when all other pools are unavailable.
	//
	// example:
	//
	// 96228666776****
	FallbackPool *int64 `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// Load balancer ID, which can be obtained by calling the [ListLoadBalancers](https://help.aliyun.com/document_detail/2868897.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95913670174****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Monitor configuration for health checks.
	MonitorShrink *string `json:"Monitor,omitempty" xml:"Monitor,omitempty"`
	// Weighted round-robin configuration, used to control the traffic distribution weights among different pools.
	RandomSteeringShrink *string `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty"`
	// Address pool corresponding to the primary region.
	//
	// example:
	//
	// {
	//
	//   "ENAM": [
	//
	//     12345678****
	//
	//   ],
	//
	//   "WNAM": [
	//
	//     23456789****,
	//
	//     23456789****
	//
	//   ]
	//
	// }
	RegionPools interface{} `json:"RegionPools,omitempty" xml:"RegionPools,omitempty"`
	// Rule configuration list, used to define behavior overrides under specific conditions.
	//
	// if can be null:
	// false
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// Session persistence, with possible values:
	//
	// - off: Not enabled.
	//
	// - ip: Session persistence by IP.
	//
	// - cookie: Session persistence by cookie.
	//
	// example:
	//
	// ip
	SessionAffinity *string `json:"SessionAffinity,omitempty" xml:"SessionAffinity,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1159101787****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Load balancing policy.
	//
	// example:
	//
	// order
	SteeringPolicy *string `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	// Address pool corresponding to the secondary region. When multiple secondary regions share the same address pool, the regions can be concatenated with commas as the key.
	//
	// example:
	//
	// {"AL,MO": [92298024898****],"CN-SH,CN-SX,CN-SC":[92304347804****,92843536908****]}
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	// TTL value, the time-to-live for DNS records, with a default of 30 and a range of 10-600.
	//
	// example:
	//
	// 300
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateLoadBalancerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerShrinkRequest) GetAdaptiveRoutingShrink() *string {
	return s.AdaptiveRoutingShrink
}

func (s *UpdateLoadBalancerShrinkRequest) GetDefaultPoolsShrink() *string {
	return s.DefaultPoolsShrink
}

func (s *UpdateLoadBalancerShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLoadBalancerShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateLoadBalancerShrinkRequest) GetFallbackPool() *int64 {
	return s.FallbackPool
}

func (s *UpdateLoadBalancerShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateLoadBalancerShrinkRequest) GetMonitorShrink() *string {
	return s.MonitorShrink
}

func (s *UpdateLoadBalancerShrinkRequest) GetRandomSteeringShrink() *string {
	return s.RandomSteeringShrink
}

func (s *UpdateLoadBalancerShrinkRequest) GetRegionPools() interface{} {
	return s.RegionPools
}

func (s *UpdateLoadBalancerShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *UpdateLoadBalancerShrinkRequest) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *UpdateLoadBalancerShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateLoadBalancerShrinkRequest) GetSteeringPolicy() *string {
	return s.SteeringPolicy
}

func (s *UpdateLoadBalancerShrinkRequest) GetSubRegionPools() interface{} {
	return s.SubRegionPools
}

func (s *UpdateLoadBalancerShrinkRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateLoadBalancerShrinkRequest) SetAdaptiveRoutingShrink(v string) *UpdateLoadBalancerShrinkRequest {
	s.AdaptiveRoutingShrink = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetDefaultPoolsShrink(v string) *UpdateLoadBalancerShrinkRequest {
	s.DefaultPoolsShrink = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetDescription(v string) *UpdateLoadBalancerShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetEnabled(v bool) *UpdateLoadBalancerShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetFallbackPool(v int64) *UpdateLoadBalancerShrinkRequest {
	s.FallbackPool = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetId(v int64) *UpdateLoadBalancerShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetMonitorShrink(v string) *UpdateLoadBalancerShrinkRequest {
	s.MonitorShrink = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetRandomSteeringShrink(v string) *UpdateLoadBalancerShrinkRequest {
	s.RandomSteeringShrink = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetRegionPools(v interface{}) *UpdateLoadBalancerShrinkRequest {
	s.RegionPools = v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetRulesShrink(v string) *UpdateLoadBalancerShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetSessionAffinity(v string) *UpdateLoadBalancerShrinkRequest {
	s.SessionAffinity = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetSiteId(v int64) *UpdateLoadBalancerShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetSteeringPolicy(v string) *UpdateLoadBalancerShrinkRequest {
	s.SteeringPolicy = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetSubRegionPools(v interface{}) *UpdateLoadBalancerShrinkRequest {
	s.SubRegionPools = v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetTtl(v int32) *UpdateLoadBalancerShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
