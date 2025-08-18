// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLoadBalancersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancers(v []*ListLoadBalancersResponseBodyLoadBalancers) *ListLoadBalancersResponseBody
	GetLoadBalancers() []*ListLoadBalancersResponseBodyLoadBalancers
	SetPageNumber(v int32) *ListLoadBalancersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLoadBalancersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLoadBalancersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLoadBalancersResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListLoadBalancersResponseBody
	GetTotalPage() *int32
}

type ListLoadBalancersResponseBody struct {
	// An array format that returns the list of load balancers.
	LoadBalancers []*ListLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	// Page number, same as the PageNumber in the request parameters.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 10
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListLoadBalancersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBody) GetLoadBalancers() []*ListLoadBalancersResponseBodyLoadBalancers {
	return s.LoadBalancers
}

func (s *ListLoadBalancersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLoadBalancersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLoadBalancersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLoadBalancersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLoadBalancersResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListLoadBalancersResponseBody) SetLoadBalancers(v []*ListLoadBalancersResponseBodyLoadBalancers) *ListLoadBalancersResponseBody {
	s.LoadBalancers = v
	return s
}

func (s *ListLoadBalancersResponseBody) SetPageNumber(v int32) *ListLoadBalancersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetPageSize(v int32) *ListLoadBalancersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetRequestId(v string) *ListLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetTotalCount(v int32) *ListLoadBalancersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetTotalPage(v int32) *ListLoadBalancersResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListLoadBalancersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancers struct {
	// Cross-pool failover configuration.
	AdaptiveRouting *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty" type:"Struct"`
	// List of default address pool IDs.
	DefaultPools []*int64 `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty" type:"Repeated"`
	// The description of the load balancer.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether the load balancer is enabled.
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The fallback pool ID, to which traffic will be redirected if all other pools are unavailable.
	//
	// example:
	//
	// 96228666776****
	FallbackPool *int64 `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// The unique identifier ID of the load balancer.
	//
	// example:
	//
	// 998676487607104
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Monitor configuration.
	Monitor *ListLoadBalancersResponseBodyLoadBalancersMonitor `json:"Monitor,omitempty" xml:"Monitor,omitempty" type:"Struct"`
	// The name of the load balancer.
	//
	// example:
	//
	// lb.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Weighted round-robin configuration, used to control the traffic distribution weights among different pools.
	RandomSteering *ListLoadBalancersResponseBodyLoadBalancersRandomSteering `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty" type:"Struct"`
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
	// List of rule configurations, used to define behaviors under specific conditions.
	Rules []*ListLoadBalancersResponseBodyLoadBalancersRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
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
	// The site ID to which the load balancer belongs.
	//
	// example:
	//
	// 1159101787****
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
	// Address pools corresponding to secondary regions. When multiple secondary regions share a set of address pools, the keys can be concatenated with commas.
	//
	// example:
	//
	// {"AL,MO": [92298024898****],"CN-SH,CN-SX,CN-SC":[92304347804****,92843536908****]}
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	// The TTL value, which is the DNS record\\"s time to live, with a default value of 30.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancers) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancers) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetAdaptiveRouting() *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting {
	return s.AdaptiveRouting
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetDefaultPools() []*int64 {
	return s.DefaultPools
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetDescription() *string {
	return s.Description
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetFallbackPool() *int64 {
	return s.FallbackPool
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetId() *int64 {
	return s.Id
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetMonitor() *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	return s.Monitor
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetName() *string {
	return s.Name
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetRandomSteering() *ListLoadBalancersResponseBodyLoadBalancersRandomSteering {
	return s.RandomSteering
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetRegionPools() interface{} {
	return s.RegionPools
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetRules() []*ListLoadBalancersResponseBodyLoadBalancersRules {
	return s.Rules
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetStatus() *string {
	return s.Status
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetSteeringPolicy() *string {
	return s.SteeringPolicy
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetSubRegionPools() interface{} {
	return s.SubRegionPools
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetTtl() *int32 {
	return s.Ttl
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAdaptiveRouting(v *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AdaptiveRouting = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetDefaultPools(v []*int64) *ListLoadBalancersResponseBodyLoadBalancers {
	s.DefaultPools = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetDescription(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Description = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetEnabled(v bool) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Enabled = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetFallbackPool(v int64) *ListLoadBalancersResponseBodyLoadBalancers {
	s.FallbackPool = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetId(v int64) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Id = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetMonitor(v *ListLoadBalancersResponseBodyLoadBalancersMonitor) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Monitor = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetName(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Name = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetRandomSteering(v *ListLoadBalancersResponseBodyLoadBalancersRandomSteering) *ListLoadBalancersResponseBodyLoadBalancers {
	s.RandomSteering = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetRegionPools(v interface{}) *ListLoadBalancersResponseBodyLoadBalancers {
	s.RegionPools = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetRules(v []*ListLoadBalancersResponseBodyLoadBalancersRules) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Rules = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetSessionAffinity(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.SessionAffinity = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetSiteId(v int64) *ListLoadBalancersResponseBodyLoadBalancers {
	s.SiteId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Status = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetSteeringPolicy(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.SteeringPolicy = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetSubRegionPools(v interface{}) *ListLoadBalancersResponseBodyLoadBalancers {
	s.SubRegionPools = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetTtl(v int32) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Ttl = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting struct {
	// Whether to fail over across pools.
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

func (s ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting) GetFailoverAcrossPools() *bool {
	return s.FailoverAcrossPools
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting) SetFailoverAcrossPools(v bool) *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting {
	s.FailoverAcrossPools = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersMonitor struct {
	// The number of consecutive failed probes required to consider the target unhealthy, such as 5.
	//
	// example:
	//
	// 5
	ConsecutiveDown *int32 `json:"ConsecutiveDown,omitempty" xml:"ConsecutiveDown,omitempty"`
	// The number of consecutive successful probes required to consider the target healthy, such as 3.
	//
	// example:
	//
	// 3
	ConsecutiveUp *int32 `json:"ConsecutiveUp,omitempty" xml:"ConsecutiveUp,omitempty"`
	// The expected status codes, such as 200,202, indicating successful HTTP responses.
	//
	// example:
	//
	// 200,202
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
	// The header information included in the probe, such as HTTP headers.
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
	// The interval for the health check, in seconds.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The method for the health check.
	//
	// example:
	//
	// GET
	Method           *string `json:"Method,omitempty" xml:"Method,omitempty"`
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
	// Application health check timeout, in seconds.
	//
	// example:
	//
	// 5
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of monitor protocol, such as HTTP, used for health checks. When the value is `off`, it indicates that no check will be performed.
	//
	// example:
	//
	// HTTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersMonitor) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersMonitor) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetConsecutiveDown() *int32 {
	return s.ConsecutiveDown
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetConsecutiveUp() *int32 {
	return s.ConsecutiveUp
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetExpectedCodes() *string {
	return s.ExpectedCodes
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetFollowRedirects() *bool {
	return s.FollowRedirects
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetHeader() interface{} {
	return s.Header
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetInterval() *int32 {
	return s.Interval
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetMethod() *string {
	return s.Method
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetMonitoringRegion() *string {
	return s.MonitoringRegion
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetPath() *string {
	return s.Path
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetPort() *int32 {
	return s.Port
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) GetType() *string {
	return s.Type
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetConsecutiveDown(v int32) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.ConsecutiveDown = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetConsecutiveUp(v int32) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.ConsecutiveUp = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetExpectedCodes(v string) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.ExpectedCodes = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetFollowRedirects(v bool) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.FollowRedirects = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetHeader(v interface{}) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.Header = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetInterval(v int32) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.Interval = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetMethod(v string) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.Method = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetMonitoringRegion(v string) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.MonitoringRegion = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetPath(v string) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.Path = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetPort(v int32) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.Port = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetTimeout(v int32) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.Timeout = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) SetType(v string) *ListLoadBalancersResponseBodyLoadBalancersMonitor {
	s.Type = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersMonitor) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersRandomSteering struct {
	// Weight configuration for each backend server pool, where the key is the pool ID and the value is the weight coefficient. The weight coefficient represents the proportion of relative traffic distribution.
	//
	// example:
	//
	// 50
	DefaultWeight *int32 `json:"DefaultWeight,omitempty" xml:"DefaultWeight,omitempty"`
	// Weight configuration for each backend server pool, where the key is the pool ID and the value is the weight coefficient.
	PoolWeights map[string]*int32 `json:"PoolWeights,omitempty" xml:"PoolWeights,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersRandomSteering) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersRandomSteering) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRandomSteering) GetDefaultWeight() *int32 {
	return s.DefaultWeight
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRandomSteering) GetPoolWeights() map[string]*int32 {
	return s.PoolWeights
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRandomSteering) SetDefaultWeight(v int32) *ListLoadBalancersResponseBodyLoadBalancersRandomSteering {
	s.DefaultWeight = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRandomSteering) SetPoolWeights(v map[string]*int32) *ListLoadBalancersResponseBodyLoadBalancersRandomSteering {
	s.PoolWeights = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRandomSteering) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersRules struct {
	// Executes a specified response after matching the rule.
	FixedResponse *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse `json:"FixedResponse,omitempty" xml:"FixedResponse,omitempty" type:"Struct"`
	// Modifies the corresponding load balancer configuration after matching the rule. The fields in this configuration will override the corresponding fields in the load balancer configuration.
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
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding global configurations. There are two usage scenarios:
	//
	// - Match all incoming requests: set the value to true
	//
	// - Match specific requests: set the value to a custom expression, for example: (http.host eq "video.example.com")
	//
	// example:
	//
	// http.request.uri.path contains "/testing"
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The switch for the rule. This parameter is not required when adding a global configuration. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// r2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule. The higher the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Whether to terminate the execution of subsequent rules.
	//
	// - true: Yes.
	//
	// - false: No, which is the default value.
	//
	// example:
	//
	// true
	Terminates *bool `json:"Terminates,omitempty" xml:"Terminates,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersRules) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersRules) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) GetFixedResponse() *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse {
	return s.FixedResponse
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) GetOverrides() interface{} {
	return s.Overrides
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) GetRule() *string {
	return s.Rule
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) GetTerminates() *bool {
	return s.Terminates
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) SetFixedResponse(v *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) *ListLoadBalancersResponseBodyLoadBalancersRules {
	s.FixedResponse = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) SetOverrides(v interface{}) *ListLoadBalancersResponseBodyLoadBalancersRules {
	s.Overrides = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) SetRule(v string) *ListLoadBalancersResponseBodyLoadBalancersRules {
	s.Rule = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) SetRuleEnable(v string) *ListLoadBalancersResponseBodyLoadBalancersRules {
	s.RuleEnable = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) SetRuleName(v string) *ListLoadBalancersResponseBodyLoadBalancersRules {
	s.RuleName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) SetSequence(v int32) *ListLoadBalancersResponseBodyLoadBalancersRules {
	s.Sequence = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) SetTerminates(v bool) *ListLoadBalancersResponseBodyLoadBalancersRules {
	s.Terminates = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRules) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse struct {
	// The Content-Type field in the HTTP Header.
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
	// Status code.
	//
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) GetContentType() *string {
	return s.ContentType
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) GetLocation() *string {
	return s.Location
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) GetMessageBody() *string {
	return s.MessageBody
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) SetContentType(v string) *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse {
	s.ContentType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) SetLocation(v string) *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse {
	s.Location = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) SetMessageBody(v string) *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse {
	s.MessageBody = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) SetStatusCode(v int32) *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse) Validate() error {
	return dara.Validate(s)
}
