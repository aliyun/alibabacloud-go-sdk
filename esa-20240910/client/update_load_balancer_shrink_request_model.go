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
	// Configures origin-pull behavior across address pools.
	AdaptiveRoutingShrink *string `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty"`
	// A list of default address pool IDs.
	DefaultPoolsShrink *string `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty"`
	// An optional description of the load balancer for easier identification and management.
	//
	// example:
	//
	// Load balancer description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the load balancer is enabled.
	//
	// - `true`: The load balancer is enabled.
	//
	// - `false`: The load balancer is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the fallback address pool. Traffic is routed to this pool when all other address pools are unavailable.
	//
	// example:
	//
	// 96228666776****
	FallbackPool *int64 `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// The ID of the load balancer. You can obtain this ID by calling the [ListLoadBalancers](https://help.aliyun.com/document_detail/2868897.html) API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95913670174****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The health check monitor configuration.
	MonitorShrink *string `json:"Monitor,omitempty" xml:"Monitor,omitempty"`
	// The configuration for weighted round-robin. This setting controls the weight of traffic distributed to different address pools.
	RandomSteeringShrink *string `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty"`
	// A map of primary regions to their corresponding address pools.
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
	// A list of rules that define behavior overrides for specific conditions.
	//
	// if can be null:
	// false
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The method for session affinity, which ensures that requests from the same client are routed to the same origin server. Valid values:
	//
	// - `off`: Disables session affinity.
	//
	// - `ip`: Enables session affinity based on the client IP address.
	//
	// - `cookie`: Enables session affinity based on a cookie.
	//
	// example:
	//
	// ip
	SessionAffinity *string `json:"SessionAffinity,omitempty" xml:"SessionAffinity,omitempty"`
	// The ID of the Site. You can obtain this ID by calling the [ListSites](~~ListSites~~) API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1159101787****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The traffic steering policy, which determines how traffic is distributed among the address pools.
	//
	// example:
	//
	// order
	SteeringPolicy *string `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	// A map of secondary regions to their corresponding address pools. To assign the same address pools to multiple secondary regions, combine their codes into a single, comma-separated key.
	//
	// example:
	//
	// {"AL,MO": [92298024898****],"CN-SH,CN-SX,CN-SC":[92304347804****,92843536908****]}
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	// The Time to Live (TTL) for the DNS record, in seconds. The default is 30. The value must be between 10 and 600, inclusive.
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
