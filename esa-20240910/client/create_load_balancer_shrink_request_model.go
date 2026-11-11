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
	// The back-to-origin configuration across origin address pools.
	//
	// example:
	//
	// true
	AdaptiveRoutingShrink *string `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty"`
	// The list of default address pools.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DefaultPoolsShrink *string `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty"`
	// The detailed description of the load balancer for easy management and identification.
	//
	// example:
	//
	// Load Balancer Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the load balancer is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the fallback address pool. Traffic is routed to this pool when all other pools are unavailable.
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
	// The name of the load balancer. The name must pass the domain name format validation and be a subdomain of the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The weighted round-robin configuration used to control the traffic distribution weights of different address pools.
	//
	// example:
	//
	// 123
	RandomSteeringShrink *string `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty"`
	// The address pools corresponding to primary regions.
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
	// The rule information.
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
	// The session persistence configuration. Valid values:
	//
	// - off: Disabled.
	//
	// - ip: Session persistence based on IP addresses.
	//
	// - cookie: Session persistence based on cookies.
	//
	// - http_header: Session persistence based on HTTP headers.
	//
	// example:
	//
	// ip
	SessionAffinity *string `json:"SessionAffinity,omitempty" xml:"SessionAffinity,omitempty"`
	// The site ID. You can obtain the site ID by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The load balancing policy. Valid values:
	//
	// - geo: Geolocation-based policy.
	//
	// - random: Weighted round-robin.
	//
	// - order: Active/standby mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// order
	SteeringPolicy *string `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	// The address pools corresponding to secondary regions. When multiple secondary regions share the same set of address pools, you can concatenate the secondary regions with commas (,) to use as the key.
	//
	// example:
	//
	// {"AL,MO": [92298024898****],"CN-SH,CN-SX,CN-SC":[92304347804****,92843536908****]}
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	// The time-to-live (TTL) value of the DNS record. Default value: 30. Unit: seconds. Valid values: 10 to 600.
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
