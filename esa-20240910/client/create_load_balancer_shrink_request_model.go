// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptiveRoutingShrink(v string) *CreateLoadBalancerShrinkRequest
	GetAdaptiveRoutingShrink() *string
	SetDefaultPoolsShrink(v string) *CreateLoadBalancerShrinkRequest
	GetDefaultPoolsShrink() *string
	SetDescription(v string) *CreateLoadBalancerShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *CreateLoadBalancerShrinkRequest
	GetEnabled() *bool
	SetFallbackPool(v int64) *CreateLoadBalancerShrinkRequest
	GetFallbackPool() *int64
	SetMonitorShrink(v string) *CreateLoadBalancerShrinkRequest
	GetMonitorShrink() *string
	SetName(v string) *CreateLoadBalancerShrinkRequest
	GetName() *string
	SetRandomSteeringShrink(v string) *CreateLoadBalancerShrinkRequest
	GetRandomSteeringShrink() *string
	SetRegionPools(v interface{}) *CreateLoadBalancerShrinkRequest
	GetRegionPools() interface{}
	SetRulesShrink(v string) *CreateLoadBalancerShrinkRequest
	GetRulesShrink() *string
	SetSessionAffinity(v string) *CreateLoadBalancerShrinkRequest
	GetSessionAffinity() *string
	SetSiteId(v int64) *CreateLoadBalancerShrinkRequest
	GetSiteId() *int64
	SetSteeringPolicy(v string) *CreateLoadBalancerShrinkRequest
	GetSteeringPolicy() *string
	SetSubRegionPools(v interface{}) *CreateLoadBalancerShrinkRequest
	GetSubRegionPools() interface{}
	SetTtl(v int32) *CreateLoadBalancerShrinkRequest
	GetTtl() *int32
}

type CreateLoadBalancerShrinkRequest struct {
	// The configuration for failover across address pools.
	//
	// example:
	//
	// true
	AdaptiveRoutingShrink *string `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty"`
	// A list of default address pool IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DefaultPoolsShrink *string `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty"`
	// A description of the Server Load Balancer.
	//
	// example:
	//
	// Test load balancer description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the Server Load Balancer.
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the fallback pool. The system directs traffic to this pool when all other pools are unavailable.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	FallbackPool *int64 `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// The monitor configuration for health checks.
	//
	// This parameter is required.
	//
	// example:
	//
	// order
	MonitorShrink *string `json:"Monitor,omitempty" xml:"Monitor,omitempty"`
	// The name of the Server Load Balancer. It must be a valid domain name and a subdomain of the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration for weighted round-robin steering. This setting controls how the system distributes traffic across different address pools based on their weights.
	//
	// example:
	//
	// 123
	RandomSteeringShrink *string `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty"`
	// The mapping of primary regions to address pools.
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
	// A list of rules to override the default traffic steering policy for specific requests.
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
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// Specifies the session affinity policy, which consistently routes requests from the same client to the same origin server. Valid values:
	//
	// - `off`: Disables session affinity.
	//
	// - `ip`: Routes requests based on the client\\"s IP address.
	//
	// - `cookie`: Uses a cookie to maintain session affinity.
	//
	// example:
	//
	// ip
	SessionAffinity *string `json:"SessionAffinity,omitempty" xml:"SessionAffinity,omitempty"`
	// The site ID. Call the [ListSites](~~ListSites~~) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The traffic steering policy, which determines how the system distributes traffic among the address pools. Valid values:
	//
	// - `geo`: Geographic routing.
	//
	// - `random`: Weighted round-robin.
	//
	// - `order`: Primary/standby.
	//
	// This parameter is required.
	//
	// example:
	//
	// order
	SteeringPolicy *string `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	// The mapping of secondary regions to address pools. To map multiple secondary regions to the same address pools, combine their region codes with commas to form the key.
	//
	// example:
	//
	// {"AL,MO": [92298024898****],"CN-SH,CN-SX,CN-SC":[92304347804****,92843536908****]}
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	// The time to live (TTL) for the DNS record, in seconds. The default value is 30. The value must be between 10 and 600.
	//
	// example:
	//
	// 300
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s CreateLoadBalancerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerShrinkRequest) GetAdaptiveRoutingShrink() *string {
	return s.AdaptiveRoutingShrink
}

func (s *CreateLoadBalancerShrinkRequest) GetDefaultPoolsShrink() *string {
	return s.DefaultPoolsShrink
}

func (s *CreateLoadBalancerShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLoadBalancerShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateLoadBalancerShrinkRequest) GetFallbackPool() *int64 {
	return s.FallbackPool
}

func (s *CreateLoadBalancerShrinkRequest) GetMonitorShrink() *string {
	return s.MonitorShrink
}

func (s *CreateLoadBalancerShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateLoadBalancerShrinkRequest) GetRandomSteeringShrink() *string {
	return s.RandomSteeringShrink
}

func (s *CreateLoadBalancerShrinkRequest) GetRegionPools() interface{} {
	return s.RegionPools
}

func (s *CreateLoadBalancerShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *CreateLoadBalancerShrinkRequest) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *CreateLoadBalancerShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateLoadBalancerShrinkRequest) GetSteeringPolicy() *string {
	return s.SteeringPolicy
}

func (s *CreateLoadBalancerShrinkRequest) GetSubRegionPools() interface{} {
	return s.SubRegionPools
}

func (s *CreateLoadBalancerShrinkRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *CreateLoadBalancerShrinkRequest) SetAdaptiveRoutingShrink(v string) *CreateLoadBalancerShrinkRequest {
	s.AdaptiveRoutingShrink = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetDefaultPoolsShrink(v string) *CreateLoadBalancerShrinkRequest {
	s.DefaultPoolsShrink = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetDescription(v string) *CreateLoadBalancerShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetEnabled(v bool) *CreateLoadBalancerShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetFallbackPool(v int64) *CreateLoadBalancerShrinkRequest {
	s.FallbackPool = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetMonitorShrink(v string) *CreateLoadBalancerShrinkRequest {
	s.MonitorShrink = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetName(v string) *CreateLoadBalancerShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetRandomSteeringShrink(v string) *CreateLoadBalancerShrinkRequest {
	s.RandomSteeringShrink = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetRegionPools(v interface{}) *CreateLoadBalancerShrinkRequest {
	s.RegionPools = v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetRulesShrink(v string) *CreateLoadBalancerShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetSessionAffinity(v string) *CreateLoadBalancerShrinkRequest {
	s.SessionAffinity = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetSiteId(v int64) *CreateLoadBalancerShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetSteeringPolicy(v string) *CreateLoadBalancerShrinkRequest {
	s.SteeringPolicy = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetSubRegionPools(v interface{}) *CreateLoadBalancerShrinkRequest {
	s.SubRegionPools = v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) SetTtl(v int32) *CreateLoadBalancerShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *CreateLoadBalancerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
