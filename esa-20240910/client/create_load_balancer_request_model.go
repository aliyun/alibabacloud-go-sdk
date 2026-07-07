// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptiveRouting(v *CreateLoadBalancerRequestAdaptiveRouting) *CreateLoadBalancerRequest
	GetAdaptiveRouting() *CreateLoadBalancerRequestAdaptiveRouting
	SetDefaultPools(v []*int64) *CreateLoadBalancerRequest
	GetDefaultPools() []*int64
	SetDescription(v string) *CreateLoadBalancerRequest
	GetDescription() *string
	SetEnabled(v bool) *CreateLoadBalancerRequest
	GetEnabled() *bool
	SetFallbackPool(v int64) *CreateLoadBalancerRequest
	GetFallbackPool() *int64
	SetMonitor(v *CreateLoadBalancerRequestMonitor) *CreateLoadBalancerRequest
	GetMonitor() *CreateLoadBalancerRequestMonitor
	SetName(v string) *CreateLoadBalancerRequest
	GetName() *string
	SetRandomSteering(v *CreateLoadBalancerRequestRandomSteering) *CreateLoadBalancerRequest
	GetRandomSteering() *CreateLoadBalancerRequestRandomSteering
	SetRegionPools(v interface{}) *CreateLoadBalancerRequest
	GetRegionPools() interface{}
	SetRules(v []*CreateLoadBalancerRequestRules) *CreateLoadBalancerRequest
	GetRules() []*CreateLoadBalancerRequestRules
	SetSessionAffinity(v string) *CreateLoadBalancerRequest
	GetSessionAffinity() *string
	SetSiteId(v int64) *CreateLoadBalancerRequest
	GetSiteId() *int64
	SetSteeringPolicy(v string) *CreateLoadBalancerRequest
	GetSteeringPolicy() *string
	SetSubRegionPools(v interface{}) *CreateLoadBalancerRequest
	GetSubRegionPools() interface{}
	SetTtl(v int32) *CreateLoadBalancerRequest
	GetTtl() *int32
}

type CreateLoadBalancerRequest struct {
	// The cross-origin address pool back-to-origin configuration.
	//
	// example:
	//
	// true
	AdaptiveRouting *CreateLoadBalancerRequestAdaptiveRouting `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty" type:"Struct"`
	// The list of default address pool IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DefaultPools []*int64 `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty" type:"Repeated"`
	// The description of the load balancer for management and identification purposes.
	//
	// example:
	//
	// Load Balancer Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the load balancer is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The fallback address pool ID. Traffic is directed to this pool when all other pools are unavailable.
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
	Monitor *CreateLoadBalancerRequestMonitor `json:"Monitor,omitempty" xml:"Monitor,omitempty" type:"Struct"`
	// The name of the load balancer. The name must be in a valid domain name format and must be a subdomain of the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The weighted round-robin configuration that controls the traffic distribution weight across different address pools.
	//
	// example:
	//
	// 123
	RandomSteering *CreateLoadBalancerRequestRandomSteering `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty" type:"Struct"`
	// The address pools mapped to primary regions.
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
	Rules []*CreateLoadBalancerRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The session persistence mode. Valid values:
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
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The load balancing policy. Valid values:
	//
	// - geo: geo-based routing.
	//
	// - random: weighted round-robin.
	//
	// - order: primary/secondary mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// order
	SteeringPolicy *string `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	// The address pools mapped to secondary regions. If multiple secondary regions share the same set of address pools, you can concatenate the secondary region names with commas as the key.
	//
	// example:
	//
	// {"AL,MO": [92298024898****],"CN-SH,CN-SX,CN-SC":[92304347804****,92843536908****]}
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	// The TTL value, which specifies the time-to-live of the DNS record. Default value: 30 seconds. Valid values: 10 to 600.
	//
	// example:
	//
	// 300
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s CreateLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) GetAdaptiveRouting() *CreateLoadBalancerRequestAdaptiveRouting {
	return s.AdaptiveRouting
}

func (s *CreateLoadBalancerRequest) GetDefaultPools() []*int64 {
	return s.DefaultPools
}

func (s *CreateLoadBalancerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLoadBalancerRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateLoadBalancerRequest) GetFallbackPool() *int64 {
	return s.FallbackPool
}

func (s *CreateLoadBalancerRequest) GetMonitor() *CreateLoadBalancerRequestMonitor {
	return s.Monitor
}

func (s *CreateLoadBalancerRequest) GetName() *string {
	return s.Name
}

func (s *CreateLoadBalancerRequest) GetRandomSteering() *CreateLoadBalancerRequestRandomSteering {
	return s.RandomSteering
}

func (s *CreateLoadBalancerRequest) GetRegionPools() interface{} {
	return s.RegionPools
}

func (s *CreateLoadBalancerRequest) GetRules() []*CreateLoadBalancerRequestRules {
	return s.Rules
}

func (s *CreateLoadBalancerRequest) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *CreateLoadBalancerRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateLoadBalancerRequest) GetSteeringPolicy() *string {
	return s.SteeringPolicy
}

func (s *CreateLoadBalancerRequest) GetSubRegionPools() interface{} {
	return s.SubRegionPools
}

func (s *CreateLoadBalancerRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *CreateLoadBalancerRequest) SetAdaptiveRouting(v *CreateLoadBalancerRequestAdaptiveRouting) *CreateLoadBalancerRequest {
	s.AdaptiveRouting = v
	return s
}

func (s *CreateLoadBalancerRequest) SetDefaultPools(v []*int64) *CreateLoadBalancerRequest {
	s.DefaultPools = v
	return s
}

func (s *CreateLoadBalancerRequest) SetDescription(v string) *CreateLoadBalancerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetEnabled(v bool) *CreateLoadBalancerRequest {
	s.Enabled = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetFallbackPool(v int64) *CreateLoadBalancerRequest {
	s.FallbackPool = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetMonitor(v *CreateLoadBalancerRequestMonitor) *CreateLoadBalancerRequest {
	s.Monitor = v
	return s
}

func (s *CreateLoadBalancerRequest) SetName(v string) *CreateLoadBalancerRequest {
	s.Name = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetRandomSteering(v *CreateLoadBalancerRequestRandomSteering) *CreateLoadBalancerRequest {
	s.RandomSteering = v
	return s
}

func (s *CreateLoadBalancerRequest) SetRegionPools(v interface{}) *CreateLoadBalancerRequest {
	s.RegionPools = v
	return s
}

func (s *CreateLoadBalancerRequest) SetRules(v []*CreateLoadBalancerRequestRules) *CreateLoadBalancerRequest {
	s.Rules = v
	return s
}

func (s *CreateLoadBalancerRequest) SetSessionAffinity(v string) *CreateLoadBalancerRequest {
	s.SessionAffinity = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetSiteId(v int64) *CreateLoadBalancerRequest {
	s.SiteId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetSteeringPolicy(v string) *CreateLoadBalancerRequest {
	s.SteeringPolicy = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetSubRegionPools(v interface{}) *CreateLoadBalancerRequest {
	s.SubRegionPools = v
	return s
}

func (s *CreateLoadBalancerRequest) SetTtl(v int32) *CreateLoadBalancerRequest {
	s.Ttl = &v
	return s
}

func (s *CreateLoadBalancerRequest) Validate() error {
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

type CreateLoadBalancerRequestAdaptiveRouting struct {
	// Specifies whether to enable cross-origin address pool failover. Valid values:
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

func (s CreateLoadBalancerRequestAdaptiveRouting) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestAdaptiveRouting) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestAdaptiveRouting) GetFailoverAcrossPools() *bool {
	return s.FailoverAcrossPools
}

func (s *CreateLoadBalancerRequestAdaptiveRouting) GetOriginLevelRetry() *bool {
	return s.OriginLevelRetry
}

func (s *CreateLoadBalancerRequestAdaptiveRouting) SetFailoverAcrossPools(v bool) *CreateLoadBalancerRequestAdaptiveRouting {
	s.FailoverAcrossPools = &v
	return s
}

func (s *CreateLoadBalancerRequestAdaptiveRouting) SetOriginLevelRetry(v bool) *CreateLoadBalancerRequestAdaptiveRouting {
	s.OriginLevelRetry = &v
	return s
}

func (s *CreateLoadBalancerRequestAdaptiveRouting) Validate() error {
	return dara.Validate(s)
}

type CreateLoadBalancerRequestMonitor struct {
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
	// The expected status codes, such as `200,202`. These are the HTTP response codes that indicate success.
	//
	// example:
	//
	// 200
	ExpectedCodes *string `json:"ExpectedCodes,omitempty" xml:"ExpectedCodes,omitempty"`
	// Specifies whether to follow redirects. Valid values:
	//
	// - true: Follow redirects.
	//
	// - false: Do not follow redirects.
	//
	// example:
	//
	// true
	FollowRedirects *bool `json:"FollowRedirects,omitempty" xml:"FollowRedirects,omitempty"`
	// The header information included in the probe request. This is the HTTP header.
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
	// The monitoring interval in seconds, such as `60`. This specifies the check frequency.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The monitor request method, such as `GET`. This is the HTTP method used for health checks.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The region where the probe nodes are located. Default value: Global. Valid values:
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
	// The monitor check path, such as `/healthcheck`. This is the URI of the request.
	//
	// example:
	//
	// /health
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The origin server port.
	//
	// example:
	//
	// 1921
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The health check timeout period. Unit: seconds. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The monitor protocol type used for health checks. A value of off indicates that health checks are disabled. Valid values:
	//
	// - TCP
	//
	// - UDP
	//
	// - SMTP
	//
	// - HTTPS
	//
	// - HTTP
	//
	// - ICMP Ping
	//
	// - off
	//
	// example:
	//
	// HTTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateLoadBalancerRequestMonitor) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestMonitor) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestMonitor) GetConsecutiveDown() *int32 {
	return s.ConsecutiveDown
}

func (s *CreateLoadBalancerRequestMonitor) GetConsecutiveUp() *int32 {
	return s.ConsecutiveUp
}

func (s *CreateLoadBalancerRequestMonitor) GetExpectedCodes() *string {
	return s.ExpectedCodes
}

func (s *CreateLoadBalancerRequestMonitor) GetFollowRedirects() *bool {
	return s.FollowRedirects
}

func (s *CreateLoadBalancerRequestMonitor) GetHeader() interface{} {
	return s.Header
}

func (s *CreateLoadBalancerRequestMonitor) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateLoadBalancerRequestMonitor) GetMethod() *string {
	return s.Method
}

func (s *CreateLoadBalancerRequestMonitor) GetMonitoringRegion() *string {
	return s.MonitoringRegion
}

func (s *CreateLoadBalancerRequestMonitor) GetPath() *string {
	return s.Path
}

func (s *CreateLoadBalancerRequestMonitor) GetPort() *int32 {
	return s.Port
}

func (s *CreateLoadBalancerRequestMonitor) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateLoadBalancerRequestMonitor) GetType() *string {
	return s.Type
}

func (s *CreateLoadBalancerRequestMonitor) SetConsecutiveDown(v int32) *CreateLoadBalancerRequestMonitor {
	s.ConsecutiveDown = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetConsecutiveUp(v int32) *CreateLoadBalancerRequestMonitor {
	s.ConsecutiveUp = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetExpectedCodes(v string) *CreateLoadBalancerRequestMonitor {
	s.ExpectedCodes = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetFollowRedirects(v bool) *CreateLoadBalancerRequestMonitor {
	s.FollowRedirects = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetHeader(v interface{}) *CreateLoadBalancerRequestMonitor {
	s.Header = v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetInterval(v int32) *CreateLoadBalancerRequestMonitor {
	s.Interval = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetMethod(v string) *CreateLoadBalancerRequestMonitor {
	s.Method = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetMonitoringRegion(v string) *CreateLoadBalancerRequestMonitor {
	s.MonitoringRegion = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetPath(v string) *CreateLoadBalancerRequestMonitor {
	s.Path = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetPort(v int32) *CreateLoadBalancerRequestMonitor {
	s.Port = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetTimeout(v int32) *CreateLoadBalancerRequestMonitor {
	s.Timeout = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) SetType(v string) *CreateLoadBalancerRequestMonitor {
	s.Type = &v
	return s
}

func (s *CreateLoadBalancerRequestMonitor) Validate() error {
	return dara.Validate(s)
}

type CreateLoadBalancerRequestRandomSteering struct {
	// The default round-robin weight applied to all address pools that do not have an individually specified weight. Valid values: integers from 0 to 100.
	//
	// example:
	//
	// 50
	DefaultWeight *int32 `json:"DefaultWeight,omitempty" xml:"DefaultWeight,omitempty"`
	// The weight configuration for each backend server pool. The key is the pool ID and the value is the weight coefficient. The weight coefficient represents the relative proportion of traffic distribution.
	PoolWeights map[string]*int32 `json:"PoolWeights,omitempty" xml:"PoolWeights,omitempty"`
}

func (s CreateLoadBalancerRequestRandomSteering) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestRandomSteering) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestRandomSteering) GetDefaultWeight() *int32 {
	return s.DefaultWeight
}

func (s *CreateLoadBalancerRequestRandomSteering) GetPoolWeights() map[string]*int32 {
	return s.PoolWeights
}

func (s *CreateLoadBalancerRequestRandomSteering) SetDefaultWeight(v int32) *CreateLoadBalancerRequestRandomSteering {
	s.DefaultWeight = &v
	return s
}

func (s *CreateLoadBalancerRequestRandomSteering) SetPoolWeights(v map[string]*int32) *CreateLoadBalancerRequestRandomSteering {
	s.PoolWeights = v
	return s
}

func (s *CreateLoadBalancerRequestRandomSteering) Validate() error {
	return dara.Validate(s)
}

type CreateLoadBalancerRequestRules struct {
	// The fixed response content returned after a rule is matched.
	//
	// example:
	//
	// {"content_type": "application/json", "location": "www.example.com", "message_body": "Testing Hello", "status_code": 0}
	FixedResponse *CreateLoadBalancerRequestRulesFixedResponse `json:"FixedResponse,omitempty" xml:"FixedResponse,omitempty" type:"Struct"`
	// The load balancing configuration that overwrites the corresponding fields in the load balancer configuration when a rule is matched. The specified fields overwrite the corresponding fields in the load balancer configuration.
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
	//                 "AL,AT": [
	//
	//                     92298024898****,
	//
	//                     92304347804****
	//
	//                 ],
	//
	//                 "BG,BY": [
	//
	//                     92298024898****
	//
	//                 ]
	//
	//             },
	//
	//             "default_pools": [
	//
	//                 92298024898****,
	//
	//                 92304347804****
	//
	//             ],
	//
	//             "fallback_pool": 92298024898****,
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
	//                 "default_weight": 0.3,
	//
	//                 "pool_weights": {
	//
	//                     "92298024898****": 0.7,
	//
	//                     "92304347804****": 0.8
	//
	//                 }
	//
	//             },
	//
	//             "region_pools": {
	//
	//                 "CN,SEAS": [
	//
	//                     92298024898****,
	//
	//                     92304347804****
	//
	//                 ],
	//
	//                 "SAF,SAS": [
	//
	//                     92304347804****
	//
	//                 ]
	//
	//             },
	//
	//             "session_affinity": "ip",
	//
	//             "steering_policy": "geo",
	//
	//             "ttl": 30
	//
	//         }
	Overrides interface{} `json:"Overrides,omitempty" xml:"Overrides,omitempty"`
	// The rule content that uses conditional expressions to match user requests. This parameter is not required when you add a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specified requests: Set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.request.method eq "GET" and http.request.version eq "HTTP/1.0") or (ip.geoip.country eq "CN") or (http.host eq "www.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. This parameter is not required when you add a global configuration. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. This parameter is not required when you add a global configuration.
	//
	// example:
	//
	// rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution order. This parameter is optional. If not specified, rules are executed in list order. If specified, the value must be a positive integer. A larger value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to stop executing subsequent rules. Valid values:
	//
	// - true: Stop executing subsequent rules.
	//
	// - false: Continue executing subsequent rules. This is the default value.
	//
	// example:
	//
	// true
	Terminates *bool `json:"Terminates,omitempty" xml:"Terminates,omitempty"`
}

func (s CreateLoadBalancerRequestRules) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestRules) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestRules) GetFixedResponse() *CreateLoadBalancerRequestRulesFixedResponse {
	return s.FixedResponse
}

func (s *CreateLoadBalancerRequestRules) GetOverrides() interface{} {
	return s.Overrides
}

func (s *CreateLoadBalancerRequestRules) GetRule() *string {
	return s.Rule
}

func (s *CreateLoadBalancerRequestRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateLoadBalancerRequestRules) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateLoadBalancerRequestRules) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateLoadBalancerRequestRules) GetTerminates() *bool {
	return s.Terminates
}

func (s *CreateLoadBalancerRequestRules) SetFixedResponse(v *CreateLoadBalancerRequestRulesFixedResponse) *CreateLoadBalancerRequestRules {
	s.FixedResponse = v
	return s
}

func (s *CreateLoadBalancerRequestRules) SetOverrides(v interface{}) *CreateLoadBalancerRequestRules {
	s.Overrides = v
	return s
}

func (s *CreateLoadBalancerRequestRules) SetRule(v string) *CreateLoadBalancerRequestRules {
	s.Rule = &v
	return s
}

func (s *CreateLoadBalancerRequestRules) SetRuleEnable(v string) *CreateLoadBalancerRequestRules {
	s.RuleEnable = &v
	return s
}

func (s *CreateLoadBalancerRequestRules) SetRuleName(v string) *CreateLoadBalancerRequestRules {
	s.RuleName = &v
	return s
}

func (s *CreateLoadBalancerRequestRules) SetSequence(v int32) *CreateLoadBalancerRequestRules {
	s.Sequence = &v
	return s
}

func (s *CreateLoadBalancerRequestRules) SetTerminates(v bool) *CreateLoadBalancerRequestRules {
	s.Terminates = &v
	return s
}

func (s *CreateLoadBalancerRequestRules) Validate() error {
	if s.FixedResponse != nil {
		if err := s.FixedResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLoadBalancerRequestRulesFixedResponse struct {
	// The Content-Type field in the HTTP header.
	//
	// example:
	//
	// application/octet-stream
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
	// Hello World!
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CreateLoadBalancerRequestRulesFixedResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestRulesFixedResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestRulesFixedResponse) GetContentType() *string {
	return s.ContentType
}

func (s *CreateLoadBalancerRequestRulesFixedResponse) GetLocation() *string {
	return s.Location
}

func (s *CreateLoadBalancerRequestRulesFixedResponse) GetMessageBody() *string {
	return s.MessageBody
}

func (s *CreateLoadBalancerRequestRulesFixedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLoadBalancerRequestRulesFixedResponse) SetContentType(v string) *CreateLoadBalancerRequestRulesFixedResponse {
	s.ContentType = &v
	return s
}

func (s *CreateLoadBalancerRequestRulesFixedResponse) SetLocation(v string) *CreateLoadBalancerRequestRulesFixedResponse {
	s.Location = &v
	return s
}

func (s *CreateLoadBalancerRequestRulesFixedResponse) SetMessageBody(v string) *CreateLoadBalancerRequestRulesFixedResponse {
	s.MessageBody = &v
	return s
}

func (s *CreateLoadBalancerRequestRulesFixedResponse) SetStatusCode(v int32) *CreateLoadBalancerRequestRulesFixedResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerRequestRulesFixedResponse) Validate() error {
	return dara.Validate(s)
}
