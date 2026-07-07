// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoadBalancerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptiveRouting(v *GetLoadBalancerResponseBodyAdaptiveRouting) *GetLoadBalancerResponseBody
	GetAdaptiveRouting() *GetLoadBalancerResponseBodyAdaptiveRouting
	SetDefaultPools(v []*int64) *GetLoadBalancerResponseBody
	GetDefaultPools() []*int64
	SetDescription(v string) *GetLoadBalancerResponseBody
	GetDescription() *string
	SetEnabled(v bool) *GetLoadBalancerResponseBody
	GetEnabled() *bool
	SetFallbackPool(v int64) *GetLoadBalancerResponseBody
	GetFallbackPool() *int64
	SetId(v int64) *GetLoadBalancerResponseBody
	GetId() *int64
	SetMonitor(v *GetLoadBalancerResponseBodyMonitor) *GetLoadBalancerResponseBody
	GetMonitor() *GetLoadBalancerResponseBodyMonitor
	SetName(v string) *GetLoadBalancerResponseBody
	GetName() *string
	SetRandomSteering(v *GetLoadBalancerResponseBodyRandomSteering) *GetLoadBalancerResponseBody
	GetRandomSteering() *GetLoadBalancerResponseBodyRandomSteering
	SetRegionPools(v interface{}) *GetLoadBalancerResponseBody
	GetRegionPools() interface{}
	SetRequestId(v string) *GetLoadBalancerResponseBody
	GetRequestId() *string
	SetRules(v []*GetLoadBalancerResponseBodyRules) *GetLoadBalancerResponseBody
	GetRules() []*GetLoadBalancerResponseBodyRules
	SetSessionAffinity(v string) *GetLoadBalancerResponseBody
	GetSessionAffinity() *string
	SetSiteId(v int64) *GetLoadBalancerResponseBody
	GetSiteId() *int64
	SetStatus(v string) *GetLoadBalancerResponseBody
	GetStatus() *string
	SetSteeringPolicy(v string) *GetLoadBalancerResponseBody
	GetSteeringPolicy() *string
	SetSubRegionPools(v interface{}) *GetLoadBalancerResponseBody
	GetSubRegionPools() interface{}
	SetTtl(v int32) *GetLoadBalancerResponseBody
	GetTtl() *int32
}

type GetLoadBalancerResponseBody struct {
	// The cross-origin pool back-to-origin configuration.
	AdaptiveRouting *GetLoadBalancerResponseBodyAdaptiveRouting `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty" type:"Struct"`
	// The list of default pool IDs.
	DefaultPools []*int64 `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty" type:"Repeated"`
	// The description of the load balancer.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the load balancer is enabled.
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The fallback pool ID. When all other pools are unavailable, traffic is redirected to this pool.
	//
	// example:
	//
	// 96228666776****
	FallbackPool *int64 `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// The unique ID of the load balancer.
	//
	// example:
	//
	// 99867648760****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The monitor configuration.
	Monitor *GetLoadBalancerResponseBodyMonitor `json:"Monitor,omitempty" xml:"Monitor,omitempty" type:"Struct"`
	// The name of the load balancer.
	//
	// example:
	//
	// lb.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The weighted round-robin configuration, which controls the traffic distribution weight across different pools.
	RandomSteering *GetLoadBalancerResponseBodyRandomSteering `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty" type:"Struct"`
	// The pools mapped to primary regions.
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
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of rule configurations that define behaviors under specific conditions.
	Rules []*GetLoadBalancerResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The session persistence setting. Valid values:
	//
	// - off: disabled.
	//
	// - ip: IP-based session persistence.
	//
	// - cookie: cookie-based session persistence.
	//
	// - http_header: HTTP header-based session persistence.
	//
	// example:
	//
	// ip
	SessionAffinity *string `json:"SessionAffinity,omitempty" xml:"SessionAffinity,omitempty"`
	// The ID of the site to which the load balancer belongs.
	//
	// example:
	//
	// 11591017874****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The status of the load balancer.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The load balancing policy.
	//
	// example:
	//
	// order
	SteeringPolicy *string `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	// The pools mapped to secondary regions. When multiple secondary regions share the same set of pools, you can concatenate multiple secondary regions with commas as the key.
	//
	// example:
	//
	// {"AL,MO": [92298024898****],"CN-SH,CN-SX,CN-SC":[92304347804****,92843536908****]}
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	// The TTL value, which specifies the time-to-live of the DNS record. Default value: 30 seconds.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s GetLoadBalancerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerResponseBody) GetAdaptiveRouting() *GetLoadBalancerResponseBodyAdaptiveRouting {
	return s.AdaptiveRouting
}

func (s *GetLoadBalancerResponseBody) GetDefaultPools() []*int64 {
	return s.DefaultPools
}

func (s *GetLoadBalancerResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetLoadBalancerResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetLoadBalancerResponseBody) GetFallbackPool() *int64 {
	return s.FallbackPool
}

func (s *GetLoadBalancerResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetLoadBalancerResponseBody) GetMonitor() *GetLoadBalancerResponseBodyMonitor {
	return s.Monitor
}

func (s *GetLoadBalancerResponseBody) GetName() *string {
	return s.Name
}

func (s *GetLoadBalancerResponseBody) GetRandomSteering() *GetLoadBalancerResponseBodyRandomSteering {
	return s.RandomSteering
}

func (s *GetLoadBalancerResponseBody) GetRegionPools() interface{} {
	return s.RegionPools
}

func (s *GetLoadBalancerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLoadBalancerResponseBody) GetRules() []*GetLoadBalancerResponseBodyRules {
	return s.Rules
}

func (s *GetLoadBalancerResponseBody) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *GetLoadBalancerResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetLoadBalancerResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetLoadBalancerResponseBody) GetSteeringPolicy() *string {
	return s.SteeringPolicy
}

func (s *GetLoadBalancerResponseBody) GetSubRegionPools() interface{} {
	return s.SubRegionPools
}

func (s *GetLoadBalancerResponseBody) GetTtl() *int32 {
	return s.Ttl
}

func (s *GetLoadBalancerResponseBody) SetAdaptiveRouting(v *GetLoadBalancerResponseBodyAdaptiveRouting) *GetLoadBalancerResponseBody {
	s.AdaptiveRouting = v
	return s
}

func (s *GetLoadBalancerResponseBody) SetDefaultPools(v []*int64) *GetLoadBalancerResponseBody {
	s.DefaultPools = v
	return s
}

func (s *GetLoadBalancerResponseBody) SetDescription(v string) *GetLoadBalancerResponseBody {
	s.Description = &v
	return s
}

func (s *GetLoadBalancerResponseBody) SetEnabled(v bool) *GetLoadBalancerResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetLoadBalancerResponseBody) SetFallbackPool(v int64) *GetLoadBalancerResponseBody {
	s.FallbackPool = &v
	return s
}

func (s *GetLoadBalancerResponseBody) SetId(v int64) *GetLoadBalancerResponseBody {
	s.Id = &v
	return s
}

func (s *GetLoadBalancerResponseBody) SetMonitor(v *GetLoadBalancerResponseBodyMonitor) *GetLoadBalancerResponseBody {
	s.Monitor = v
	return s
}

func (s *GetLoadBalancerResponseBody) SetName(v string) *GetLoadBalancerResponseBody {
	s.Name = &v
	return s
}

func (s *GetLoadBalancerResponseBody) SetRandomSteering(v *GetLoadBalancerResponseBodyRandomSteering) *GetLoadBalancerResponseBody {
	s.RandomSteering = v
	return s
}

func (s *GetLoadBalancerResponseBody) SetRegionPools(v interface{}) *GetLoadBalancerResponseBody {
	s.RegionPools = v
	return s
}

func (s *GetLoadBalancerResponseBody) SetRequestId(v string) *GetLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalancerResponseBody) SetRules(v []*GetLoadBalancerResponseBodyRules) *GetLoadBalancerResponseBody {
	s.Rules = v
	return s
}

func (s *GetLoadBalancerResponseBody) SetSessionAffinity(v string) *GetLoadBalancerResponseBody {
	s.SessionAffinity = &v
	return s
}

func (s *GetLoadBalancerResponseBody) SetSiteId(v int64) *GetLoadBalancerResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetLoadBalancerResponseBody) SetStatus(v string) *GetLoadBalancerResponseBody {
	s.Status = &v
	return s
}

func (s *GetLoadBalancerResponseBody) SetSteeringPolicy(v string) *GetLoadBalancerResponseBody {
	s.SteeringPolicy = &v
	return s
}

func (s *GetLoadBalancerResponseBody) SetSubRegionPools(v interface{}) *GetLoadBalancerResponseBody {
	s.SubRegionPools = v
	return s
}

func (s *GetLoadBalancerResponseBody) SetTtl(v int32) *GetLoadBalancerResponseBody {
	s.Ttl = &v
	return s
}

func (s *GetLoadBalancerResponseBody) Validate() error {
	if s.AdaptiveRouting != nil {
		if err := s.AdaptiveRouting.Validate(); err != nil {
			return err
		}
	}
	if s.Monitor != nil {
		if err := s.Monitor.Validate(); err != nil {
			return err
		}
	}
	if s.RandomSteering != nil {
		if err := s.RandomSteering.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLoadBalancerResponseBodyAdaptiveRouting struct {
	// Indicates whether failover across origin pools is enabled.
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// true
	FailoverAcrossPools *bool `json:"FailoverAcrossPools,omitempty" xml:"FailoverAcrossPools,omitempty"`
	// Specifies whether to retry the next IP address when back-to-origin fails and the origin server is a domain name that resolves to multiple IP addresses.
	//
	// example:
	//
	// false
	OriginLevelRetry *bool `json:"OriginLevelRetry,omitempty" xml:"OriginLevelRetry,omitempty"`
}

func (s GetLoadBalancerResponseBodyAdaptiveRouting) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerResponseBodyAdaptiveRouting) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerResponseBodyAdaptiveRouting) GetFailoverAcrossPools() *bool {
	return s.FailoverAcrossPools
}

func (s *GetLoadBalancerResponseBodyAdaptiveRouting) GetOriginLevelRetry() *bool {
	return s.OriginLevelRetry
}

func (s *GetLoadBalancerResponseBodyAdaptiveRouting) SetFailoverAcrossPools(v bool) *GetLoadBalancerResponseBodyAdaptiveRouting {
	s.FailoverAcrossPools = &v
	return s
}

func (s *GetLoadBalancerResponseBodyAdaptiveRouting) SetOriginLevelRetry(v bool) *GetLoadBalancerResponseBodyAdaptiveRouting {
	s.OriginLevelRetry = &v
	return s
}

func (s *GetLoadBalancerResponseBodyAdaptiveRouting) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerResponseBodyMonitor struct {
	// The number of consecutive failed probes required to consider the check failed, such as `5`.
	//
	// example:
	//
	// 5
	ConsecutiveDown *int32 `json:"ConsecutiveDown,omitempty" xml:"ConsecutiveDown,omitempty"`
	// The number of consecutive successful probes required to consider the check successful, such as `3`.
	//
	// example:
	//
	// 3
	ConsecutiveUp *int32 `json:"ConsecutiveUp,omitempty" xml:"ConsecutiveUp,omitempty"`
	// The expected status codes for a successful HTTP response, such as 200,202.
	//
	// example:
	//
	// 200,202
	ExpectedCodes *string `json:"ExpectedCodes,omitempty" xml:"ExpectedCodes,omitempty"`
	// Indicates whether to follow redirects.
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// true
	FollowRedirects *bool `json:"FollowRedirects,omitempty" xml:"FollowRedirects,omitempty"`
	// The HTTP header information included in the probe request.
	//
	// example:
	//
	// {
	//
	//         "host": [
	//
	//             "example1.com",
	//
	//             "example2.com"
	//
	//         ]
	//
	//     }
	Header interface{} `json:"Header,omitempty" xml:"Header,omitempty"`
	// The health check interval. Unit: seconds.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The health check method.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The region where the probe nodes are located. Default value: Global.
	//
	// - Global: worldwide.
	//
	// - ChineseMainland: the Chinese mainland.
	//
	// - OutsideChineseMainland: worldwide (excluding the Chinese mainland).
	//
	// example:
	//
	// Global
	MonitoringRegion *string `json:"MonitoringRegion,omitempty" xml:"MonitoringRegion,omitempty"`
	// The path.
	//
	// example:
	//
	// /
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The target port.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The health check timeout period. Unit: seconds.
	//
	// example:
	//
	// 5
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The monitor protocol type, such as HTTP, used for health checks. A value of off indicates that no health check is performed.
	//
	// example:
	//
	// HTTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLoadBalancerResponseBodyMonitor) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerResponseBodyMonitor) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerResponseBodyMonitor) GetConsecutiveDown() *int32 {
	return s.ConsecutiveDown
}

func (s *GetLoadBalancerResponseBodyMonitor) GetConsecutiveUp() *int32 {
	return s.ConsecutiveUp
}

func (s *GetLoadBalancerResponseBodyMonitor) GetExpectedCodes() *string {
	return s.ExpectedCodes
}

func (s *GetLoadBalancerResponseBodyMonitor) GetFollowRedirects() *bool {
	return s.FollowRedirects
}

func (s *GetLoadBalancerResponseBodyMonitor) GetHeader() interface{} {
	return s.Header
}

func (s *GetLoadBalancerResponseBodyMonitor) GetInterval() *int32 {
	return s.Interval
}

func (s *GetLoadBalancerResponseBodyMonitor) GetMethod() *string {
	return s.Method
}

func (s *GetLoadBalancerResponseBodyMonitor) GetMonitoringRegion() *string {
	return s.MonitoringRegion
}

func (s *GetLoadBalancerResponseBodyMonitor) GetPath() *string {
	return s.Path
}

func (s *GetLoadBalancerResponseBodyMonitor) GetPort() *int32 {
	return s.Port
}

func (s *GetLoadBalancerResponseBodyMonitor) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetLoadBalancerResponseBodyMonitor) GetType() *string {
	return s.Type
}

func (s *GetLoadBalancerResponseBodyMonitor) SetConsecutiveDown(v int32) *GetLoadBalancerResponseBodyMonitor {
	s.ConsecutiveDown = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetConsecutiveUp(v int32) *GetLoadBalancerResponseBodyMonitor {
	s.ConsecutiveUp = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetExpectedCodes(v string) *GetLoadBalancerResponseBodyMonitor {
	s.ExpectedCodes = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetFollowRedirects(v bool) *GetLoadBalancerResponseBodyMonitor {
	s.FollowRedirects = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetHeader(v interface{}) *GetLoadBalancerResponseBodyMonitor {
	s.Header = v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetInterval(v int32) *GetLoadBalancerResponseBodyMonitor {
	s.Interval = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetMethod(v string) *GetLoadBalancerResponseBodyMonitor {
	s.Method = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetMonitoringRegion(v string) *GetLoadBalancerResponseBodyMonitor {
	s.MonitoringRegion = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetPath(v string) *GetLoadBalancerResponseBodyMonitor {
	s.Path = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetPort(v int32) *GetLoadBalancerResponseBodyMonitor {
	s.Port = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetTimeout(v int32) *GetLoadBalancerResponseBodyMonitor {
	s.Timeout = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) SetType(v string) *GetLoadBalancerResponseBodyMonitor {
	s.Type = &v
	return s
}

func (s *GetLoadBalancerResponseBodyMonitor) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerResponseBodyRandomSteering struct {
	// The default round-robin weight applied to all pools that do not have an individually specified weight. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	DefaultWeight *int32 `json:"DefaultWeight,omitempty" xml:"DefaultWeight,omitempty"`
	// The weight configuration for each backend server pool. The key is the pool ID and the value is the weight coefficient. The weight coefficient represents the relative proportion of traffic distribution.
	PoolWeights map[string]*int32 `json:"PoolWeights,omitempty" xml:"PoolWeights,omitempty"`
}

func (s GetLoadBalancerResponseBodyRandomSteering) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerResponseBodyRandomSteering) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerResponseBodyRandomSteering) GetDefaultWeight() *int32 {
	return s.DefaultWeight
}

func (s *GetLoadBalancerResponseBodyRandomSteering) GetPoolWeights() map[string]*int32 {
	return s.PoolWeights
}

func (s *GetLoadBalancerResponseBodyRandomSteering) SetDefaultWeight(v int32) *GetLoadBalancerResponseBodyRandomSteering {
	s.DefaultWeight = &v
	return s
}

func (s *GetLoadBalancerResponseBodyRandomSteering) SetPoolWeights(v map[string]*int32) *GetLoadBalancerResponseBodyRandomSteering {
	s.PoolWeights = v
	return s
}

func (s *GetLoadBalancerResponseBodyRandomSteering) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerResponseBodyRules struct {
	// The fixed response content returned when the rule is matched.
	FixedResponse *GetLoadBalancerResponseBodyRulesFixedResponse `json:"FixedResponse,omitempty" xml:"FixedResponse,omitempty" type:"Struct"`
	// The load balancing configuration overrides applied to matching requests. The configured fields override the corresponding fields in the load balancer configuration.
	//
	// example:
	//
	// {
	//
	//             "adaptive_routing": {
	//
	//                 "failover_across_pools": true
	//
	//             },
	//
	//             "sub_region_pools": {
	//
	//                 "GB": [
	//
	//                     96228666776****
	//
	//                 ],
	//
	//                 "US": [
	//
	//                     96228666776****
	//
	//                 ]
	//
	//             },
	//
	//             "default_pools": [
	//
	//                 96228666776****,
	//
	//                 96228666776****
	//
	//             ],
	//
	//             "fallback_pool": 96228666776****,
	//
	//             "location_strategy": {
	//
	//                 "mode": "resolver_ip",
	//
	//                 "prefer_ecs": "always"
	//
	//             },
	//
	//             "random_steering": {
	//
	//                 "default_weight": 30,
	//
	//                 "pool_weights": {
	//
	//                     "96228666776****": 70,
	//
	//                     "96228666776****": 80
	//
	//                 }
	//
	//             },
	//
	//             "region_pools": {
	//
	//                 "ENAM": [
	//
	//                     96228666776****,
	//
	//                     92843536908****
	//
	//                 ],
	//
	//                 "WNAM": [
	//
	//                     92843536908****
	//
	//                 ]
	//
	//             },
	//
	//             "session_affinity": "cookie",
	//
	//             "session_affinity_attributes": {
	//
	//                 "drain_duration": 100,
	//
	//                 "headers": ["none"],
	//
	//                 "require_all_headers": false,
	//
	//                 "samesite": "Auto",
	//
	//                 "secure": "Auto",
	//
	//                 "zero_downtime_failover": "sticky"
	//
	//             },
	//
	//             "session_affinity_ttl": 1800,
	//
	//             "steering_policy": "dynamic_latency",
	//
	//             "ttl": 30
	//
	//         }
	Overrides interface{} `json:"Overrides,omitempty" xml:"Overrides,omitempty"`
	// The rule content, which uses conditional expressions to match user requests. This parameter is not required when you add a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// http.request.uri.path contains "/testing"
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. This parameter is not required when you add a global configuration. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// off
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. This parameter is not required when you add a global configuration.
	//
	// example:
	//
	// r2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution order. A larger value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Indicates whether to terminate the execution of subsequent rules.
	//
	// - true: Yes.
	//
	// - false: No. This is the default value.
	//
	// example:
	//
	// true
	Terminates *bool `json:"Terminates,omitempty" xml:"Terminates,omitempty"`
}

func (s GetLoadBalancerResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerResponseBodyRules) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerResponseBodyRules) GetFixedResponse() *GetLoadBalancerResponseBodyRulesFixedResponse {
	return s.FixedResponse
}

func (s *GetLoadBalancerResponseBodyRules) GetOverrides() interface{} {
	return s.Overrides
}

func (s *GetLoadBalancerResponseBodyRules) GetRule() *string {
	return s.Rule
}

func (s *GetLoadBalancerResponseBodyRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetLoadBalancerResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *GetLoadBalancerResponseBodyRules) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetLoadBalancerResponseBodyRules) GetTerminates() *bool {
	return s.Terminates
}

func (s *GetLoadBalancerResponseBodyRules) SetFixedResponse(v *GetLoadBalancerResponseBodyRulesFixedResponse) *GetLoadBalancerResponseBodyRules {
	s.FixedResponse = v
	return s
}

func (s *GetLoadBalancerResponseBodyRules) SetOverrides(v interface{}) *GetLoadBalancerResponseBodyRules {
	s.Overrides = v
	return s
}

func (s *GetLoadBalancerResponseBodyRules) SetRule(v string) *GetLoadBalancerResponseBodyRules {
	s.Rule = &v
	return s
}

func (s *GetLoadBalancerResponseBodyRules) SetRuleEnable(v string) *GetLoadBalancerResponseBodyRules {
	s.RuleEnable = &v
	return s
}

func (s *GetLoadBalancerResponseBodyRules) SetRuleName(v string) *GetLoadBalancerResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *GetLoadBalancerResponseBodyRules) SetSequence(v int32) *GetLoadBalancerResponseBodyRules {
	s.Sequence = &v
	return s
}

func (s *GetLoadBalancerResponseBodyRules) SetTerminates(v bool) *GetLoadBalancerResponseBodyRules {
	s.Terminates = &v
	return s
}

func (s *GetLoadBalancerResponseBodyRules) Validate() error {
	if s.FixedResponse != nil {
		if err := s.FixedResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLoadBalancerResponseBodyRulesFixedResponse struct {
	// The Content-Type field in the HTTP header.
	//
	// example:
	//
	// application/json
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The location field in the HTTP response.
	//
	// example:
	//
	// http://www.example.com/index.html
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The response body value.
	//
	// example:
	//
	// Hello World.
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetLoadBalancerResponseBodyRulesFixedResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerResponseBodyRulesFixedResponse) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerResponseBodyRulesFixedResponse) GetContentType() *string {
	return s.ContentType
}

func (s *GetLoadBalancerResponseBodyRulesFixedResponse) GetLocation() *string {
	return s.Location
}

func (s *GetLoadBalancerResponseBodyRulesFixedResponse) GetMessageBody() *string {
	return s.MessageBody
}

func (s *GetLoadBalancerResponseBodyRulesFixedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLoadBalancerResponseBodyRulesFixedResponse) SetContentType(v string) *GetLoadBalancerResponseBodyRulesFixedResponse {
	s.ContentType = &v
	return s
}

func (s *GetLoadBalancerResponseBodyRulesFixedResponse) SetLocation(v string) *GetLoadBalancerResponseBodyRulesFixedResponse {
	s.Location = &v
	return s
}

func (s *GetLoadBalancerResponseBodyRulesFixedResponse) SetMessageBody(v string) *GetLoadBalancerResponseBodyRulesFixedResponse {
	s.MessageBody = &v
	return s
}

func (s *GetLoadBalancerResponseBodyRulesFixedResponse) SetStatusCode(v int32) *GetLoadBalancerResponseBodyRulesFixedResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoadBalancerResponseBodyRulesFixedResponse) Validate() error {
	return dara.Validate(s)
}
