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
	// Configuration for failover across pools.
	//
	// example:
	//
	// true
	AdaptiveRouting *CreateLoadBalancerRequestAdaptiveRouting `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty" type:"Struct"`
	// List of default pools.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DefaultPools []*int64 `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty" type:"Repeated"`
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
	// This parameter is required.
	//
	// example:
	//
	// 123
	FallbackPool *int64 `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// Monitor configuration, used for health checks.
	//
	// This parameter is required.
	//
	// example:
	//
	// order
	Monitor *CreateLoadBalancerRequestMonitor `json:"Monitor,omitempty" xml:"Monitor,omitempty" type:"Struct"`
	// The name of the load balancer, which must meet the domain name format validation and be a subdomain under the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Weighted round-robin configuration, used to control the traffic distribution weights among different pools.
	//
	// example:
	//
	// 123
	RandomSteering *CreateLoadBalancerRequestRandomSteering `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty" type:"Struct"`
	// Address pools corresponding to primary regions.
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
	// Rule information.
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
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Load balancing strategy.
	//
	// - geo: Geographical strategy.
	//
	// - random: Weighted round-robin.
	//
	// - order: Primary and backup method.
	//
	// This parameter is required.
	//
	// example:
	//
	// order
	SteeringPolicy *string `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	// Address pools corresponding to secondary regions. When multiple secondary regions share the same set of address pools, the keys can be concatenated with commas.
	//
	// example:
	//
	// {"AL,MO": [92298024898****],"CN-SH,CN-SX,CN-SC":[92304347804****,92843536908****]}
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	// TTL value, the time-to-live for DNS records, with a default of 30 seconds. The value range is 10-600.
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
	// Whether to failover across pools.
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// true
	FailoverAcrossPools *bool `json:"FailoverAcrossPools,omitempty" xml:"FailoverAcrossPools,omitempty"`
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

func (s *CreateLoadBalancerRequestAdaptiveRouting) SetFailoverAcrossPools(v bool) *CreateLoadBalancerRequestAdaptiveRouting {
	s.FailoverAcrossPools = &v
	return s
}

func (s *CreateLoadBalancerRequestAdaptiveRouting) Validate() error {
	return dara.Validate(s)
}

type CreateLoadBalancerRequestMonitor struct {
	// Number of consecutive failed probes required to consider the target unhealthy, such as `5`.
	//
	// example:
	//
	// 5
	ConsecutiveDown *int32 `json:"ConsecutiveDown,omitempty" xml:"ConsecutiveDown,omitempty"`
	// Number of consecutive successful probes required to consider the target healthy, such as `3`.
	//
	// example:
	//
	// 3
	ConsecutiveUp *int32 `json:"ConsecutiveUp,omitempty" xml:"ConsecutiveUp,omitempty"`
	// Expected status codes, such as `200,202`, which are successful HTTP responses.
	//
	// example:
	//
	// 200
	ExpectedCodes *string `json:"ExpectedCodes,omitempty" xml:"ExpectedCodes,omitempty"`
	// Whether to follow redirects.
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// true
	FollowRedirects *bool `json:"FollowRedirects,omitempty" xml:"FollowRedirects,omitempty"`
	// Header information included in the probe, which is an HTTP header.
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
	// Monitoring interval, such as `60` seconds, which is the frequency of checks.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Monitor request method, such as `GET`, which is a method in the HTTP protocol.
	//
	// example:
	//
	// GET
	Method           *string `json:"Method,omitempty" xml:"Method,omitempty"`
	MonitoringRegion *string `json:"MonitoringRegion,omitempty" xml:"MonitoringRegion,omitempty"`
	// Monitor check path, such as `/healthcheck`, which is an HTTP request path.
	//
	// example:
	//
	// /health
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Origin server port.
	//
	// example:
	//
	// 1921
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Application health check timeout, in seconds, with a value range of 1-10.
	//
	// example:
	//
	// 5
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Monitor protocol type, such as HTTP, used for health checks. When set to `off`, no check is performed.
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
	// Default weight for all pools that do not have individual weights specified. The value range is an integer between 0 and 100.
	//
	// example:
	//
	// 50
	DefaultWeight *int32 `json:"DefaultWeight,omitempty" xml:"DefaultWeight,omitempty"`
	// Weight configuration for each backend server pool, with the key being the pool ID and the value being the weight coefficient. The weight coefficient represents the proportion of relative traffic distribution.
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
	// Execute a specified response after matching the rule.
	//
	// example:
	//
	// {"content_type": "application/json", "location": "www.example.com", "message_body": "Testing Hello", "status_code": 0}
	FixedResponse *CreateLoadBalancerRequestRulesFixedResponse `json:"FixedResponse,omitempty" xml:"FixedResponse,omitempty" type:"Struct"`
	// Modify the corresponding load balancing configuration after matching the rule. The fields in the configuration will override the corresponding fields in the load balancer configuration.
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
	// Rule content, using conditional expressions to match user requests. This parameter does not need to be set when adding global configurations. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.request.method eq "GET" and http.request.version eq "HTTP/1.0") or (ip.geoip.country eq "CN") or (http.host eq "www.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter does not need to be set when adding global configurations. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. This parameter does not need to be set when adding global configurations.
	//
	// example:
	//
	// rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule. It can be left blank, in which case the rules will be executed in the order they appear in the list. If specified, it must be a positive integer, with higher values indicating higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Whether to terminate the execution of subsequent rules.
	//
	// - true: Yes.
	//
	// - false: No, default value.
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
	// Content-Type field in the HTTP Header.
	//
	// example:
	//
	// application/octet-stream
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Location field in the HTTP response.
	//
	// example:
	//
	// http://www.example.com/index.html
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// Response body value.
	//
	// example:
	//
	// Hello World!
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// Response status code.
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
