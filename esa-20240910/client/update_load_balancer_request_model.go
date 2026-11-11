// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptiveRouting(v *UpdateLoadBalancerRequestAdaptiveRouting) *UpdateLoadBalancerRequest
	GetAdaptiveRouting() *UpdateLoadBalancerRequestAdaptiveRouting
	SetDefaultPools(v []*int64) *UpdateLoadBalancerRequest
	GetDefaultPools() []*int64
	SetDescription(v string) *UpdateLoadBalancerRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdateLoadBalancerRequest
	GetEnabled() *bool
	SetFallbackPool(v int64) *UpdateLoadBalancerRequest
	GetFallbackPool() *int64
	SetId(v int64) *UpdateLoadBalancerRequest
	GetId() *int64
	SetMonitor(v *UpdateLoadBalancerRequestMonitor) *UpdateLoadBalancerRequest
	GetMonitor() *UpdateLoadBalancerRequestMonitor
	SetRandomSteering(v *UpdateLoadBalancerRequestRandomSteering) *UpdateLoadBalancerRequest
	GetRandomSteering() *UpdateLoadBalancerRequestRandomSteering
	SetRegionPools(v interface{}) *UpdateLoadBalancerRequest
	GetRegionPools() interface{}
	SetRules(v []*UpdateLoadBalancerRequestRules) *UpdateLoadBalancerRequest
	GetRules() []*UpdateLoadBalancerRequestRules
	SetSessionAffinity(v string) *UpdateLoadBalancerRequest
	GetSessionAffinity() *string
	SetSiteId(v int64) *UpdateLoadBalancerRequest
	GetSiteId() *int64
	SetSteeringPolicy(v string) *UpdateLoadBalancerRequest
	GetSteeringPolicy() *string
	SetSubRegionPools(v interface{}) *UpdateLoadBalancerRequest
	GetSubRegionPools() interface{}
	SetTtl(v int32) *UpdateLoadBalancerRequest
	GetTtl() *int32
}

type UpdateLoadBalancerRequest struct {
	// The cross-pool back-to-origin configuration.
	AdaptiveRouting *UpdateLoadBalancerRequestAdaptiveRouting `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty" type:"Struct"`
	// The list of default pool IDs.
	DefaultPools []*int64 `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty" type:"Repeated"`
	// The detailed description of the load balancing instance. This description facilitates management and identification.
	//
	// example:
	//
	// Load balancer description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the load balancing instance is enabled.
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The fallback pool ID. Traffic is routed to this pool when all other pools are unavailable.
	//
	// example:
	//
	// 96228666776****
	FallbackPool *int64 `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// The load balancing instance ID. You can invoke the [ListLoadBalancers](https://help.aliyun.com/document_detail/2868897.html) operation to obtain the load balancing instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95913670174****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The monitor configuration for health checks.
	Monitor *UpdateLoadBalancerRequestMonitor `json:"Monitor,omitempty" xml:"Monitor,omitempty" type:"Struct"`
	// The weighted round-robin configuration. This configuration controls the traffic distribution weights for different pools.
	RandomSteering *UpdateLoadBalancerRequestRandomSteering `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty" type:"Struct"`
	// The pools that correspond to the primary regions.
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
	// The list of rule configurations. This list is used to define behavior overrides under specific conditions.
	//
	// if can be null:
	// false
	Rules []*UpdateLoadBalancerRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The session persistence policy. Valid values:
	//
	// - off: Disabled.
	//
	// - ip: Session persistence based on IP address.
	//
	// - cookie: Session persistence based on cookie.
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
	// 1159101787****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The load balancing policy.
	//
	// example:
	//
	// order
	SteeringPolicy *string `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	// The pools that correspond to the secondary regions. When multiple secondary regions share the same set of pools, you can concatenate multiple secondary regions with commas (,) to use as the key.
	//
	// example:
	//
	// {"AL,MO": [92298024898****],"CN-SH,CN-SX,CN-SC":[92304347804****,92843536908****]}
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	// The time-to-live (TTL) value of the DNS record. Default value: 30. Valid values: 10 to 600.
	//
	// example:
	//
	// 300
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerRequest) GetAdaptiveRouting() *UpdateLoadBalancerRequestAdaptiveRouting {
	return s.AdaptiveRouting
}

func (s *UpdateLoadBalancerRequest) GetDefaultPools() []*int64 {
	return s.DefaultPools
}

func (s *UpdateLoadBalancerRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLoadBalancerRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateLoadBalancerRequest) GetFallbackPool() *int64 {
	return s.FallbackPool
}

func (s *UpdateLoadBalancerRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateLoadBalancerRequest) GetMonitor() *UpdateLoadBalancerRequestMonitor {
	return s.Monitor
}

func (s *UpdateLoadBalancerRequest) GetRandomSteering() *UpdateLoadBalancerRequestRandomSteering {
	return s.RandomSteering
}

func (s *UpdateLoadBalancerRequest) GetRegionPools() interface{} {
	return s.RegionPools
}

func (s *UpdateLoadBalancerRequest) GetRules() []*UpdateLoadBalancerRequestRules {
	return s.Rules
}

func (s *UpdateLoadBalancerRequest) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *UpdateLoadBalancerRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateLoadBalancerRequest) GetSteeringPolicy() *string {
	return s.SteeringPolicy
}

func (s *UpdateLoadBalancerRequest) GetSubRegionPools() interface{} {
	return s.SubRegionPools
}

func (s *UpdateLoadBalancerRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateLoadBalancerRequest) SetAdaptiveRouting(v *UpdateLoadBalancerRequestAdaptiveRouting) *UpdateLoadBalancerRequest {
	s.AdaptiveRouting = v
	return s
}

func (s *UpdateLoadBalancerRequest) SetDefaultPools(v []*int64) *UpdateLoadBalancerRequest {
	s.DefaultPools = v
	return s
}

func (s *UpdateLoadBalancerRequest) SetDescription(v string) *UpdateLoadBalancerRequest {
	s.Description = &v
	return s
}

func (s *UpdateLoadBalancerRequest) SetEnabled(v bool) *UpdateLoadBalancerRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateLoadBalancerRequest) SetFallbackPool(v int64) *UpdateLoadBalancerRequest {
	s.FallbackPool = &v
	return s
}

func (s *UpdateLoadBalancerRequest) SetId(v int64) *UpdateLoadBalancerRequest {
	s.Id = &v
	return s
}

func (s *UpdateLoadBalancerRequest) SetMonitor(v *UpdateLoadBalancerRequestMonitor) *UpdateLoadBalancerRequest {
	s.Monitor = v
	return s
}

func (s *UpdateLoadBalancerRequest) SetRandomSteering(v *UpdateLoadBalancerRequestRandomSteering) *UpdateLoadBalancerRequest {
	s.RandomSteering = v
	return s
}

func (s *UpdateLoadBalancerRequest) SetRegionPools(v interface{}) *UpdateLoadBalancerRequest {
	s.RegionPools = v
	return s
}

func (s *UpdateLoadBalancerRequest) SetRules(v []*UpdateLoadBalancerRequestRules) *UpdateLoadBalancerRequest {
	s.Rules = v
	return s
}

func (s *UpdateLoadBalancerRequest) SetSessionAffinity(v string) *UpdateLoadBalancerRequest {
	s.SessionAffinity = &v
	return s
}

func (s *UpdateLoadBalancerRequest) SetSiteId(v int64) *UpdateLoadBalancerRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateLoadBalancerRequest) SetSteeringPolicy(v string) *UpdateLoadBalancerRequest {
	s.SteeringPolicy = &v
	return s
}

func (s *UpdateLoadBalancerRequest) SetSubRegionPools(v interface{}) *UpdateLoadBalancerRequest {
	s.SubRegionPools = v
	return s
}

func (s *UpdateLoadBalancerRequest) SetTtl(v int32) *UpdateLoadBalancerRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateLoadBalancerRequest) Validate() error {
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

type UpdateLoadBalancerRequestAdaptiveRouting struct {
	// Specifies whether to enable cross-pool back-to-origin routing.
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// false
	FailoverAcrossPools *bool `json:"FailoverAcrossPools,omitempty" xml:"FailoverAcrossPools,omitempty"`
	OriginLevelRetry    *bool `json:"OriginLevelRetry,omitempty" xml:"OriginLevelRetry,omitempty"`
}

func (s UpdateLoadBalancerRequestAdaptiveRouting) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerRequestAdaptiveRouting) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerRequestAdaptiveRouting) GetFailoverAcrossPools() *bool {
	return s.FailoverAcrossPools
}

func (s *UpdateLoadBalancerRequestAdaptiveRouting) GetOriginLevelRetry() *bool {
	return s.OriginLevelRetry
}

func (s *UpdateLoadBalancerRequestAdaptiveRouting) SetFailoverAcrossPools(v bool) *UpdateLoadBalancerRequestAdaptiveRouting {
	s.FailoverAcrossPools = &v
	return s
}

func (s *UpdateLoadBalancerRequestAdaptiveRouting) SetOriginLevelRetry(v bool) *UpdateLoadBalancerRequestAdaptiveRouting {
	s.OriginLevelRetry = &v
	return s
}

func (s *UpdateLoadBalancerRequestAdaptiveRouting) Validate() error {
	return dara.Validate(s)
}

type UpdateLoadBalancerRequestMonitor struct {
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
	// The expected status codes for a successful HTTP response, such as 200 and 202.
	//
	// example:
	//
	// 200,202
	ExpectedCodes *string `json:"ExpectedCodes,omitempty" xml:"ExpectedCodes,omitempty"`
	// Specifies whether to follow redirects.
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// true
	FollowRedirects *bool `json:"FollowRedirects,omitempty" xml:"FollowRedirects,omitempty"`
	// The monitor request header configuration.
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
	// The monitor interval, such as 60 seconds. This parameter specifies the check frequency.
	//
	// example:
	//
	// 100
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The monitor request method, such as GET in the HTTP protocol.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The region where the probe nodes are located. Default value: Global. Valid values:
	//
	// - Global: Global.
	//
	// - ChineseMainland: The Chinese mainland.
	//
	// - OutsideChineseMainland: Global (excluding the Chinese mainland).
	//
	// example:
	//
	// Global
	MonitoringRegion *string `json:"MonitoringRegion,omitempty" xml:"MonitoringRegion,omitempty"`
	// The health check path of the monitor, such as /healthcheck, which is the URI of the request.
	//
	// example:
	//
	// /health
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The origin server port.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The health check timeout period, in seconds. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The monitor protocol type for health checks, such as HTTP. A value of off indicates that health checks are disabled.
	//
	// example:
	//
	// HTTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateLoadBalancerRequestMonitor) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerRequestMonitor) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerRequestMonitor) GetConsecutiveDown() *int32 {
	return s.ConsecutiveDown
}

func (s *UpdateLoadBalancerRequestMonitor) GetConsecutiveUp() *int32 {
	return s.ConsecutiveUp
}

func (s *UpdateLoadBalancerRequestMonitor) GetExpectedCodes() *string {
	return s.ExpectedCodes
}

func (s *UpdateLoadBalancerRequestMonitor) GetFollowRedirects() *bool {
	return s.FollowRedirects
}

func (s *UpdateLoadBalancerRequestMonitor) GetHeader() interface{} {
	return s.Header
}

func (s *UpdateLoadBalancerRequestMonitor) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateLoadBalancerRequestMonitor) GetMethod() *string {
	return s.Method
}

func (s *UpdateLoadBalancerRequestMonitor) GetMonitoringRegion() *string {
	return s.MonitoringRegion
}

func (s *UpdateLoadBalancerRequestMonitor) GetPath() *string {
	return s.Path
}

func (s *UpdateLoadBalancerRequestMonitor) GetPort() *int32 {
	return s.Port
}

func (s *UpdateLoadBalancerRequestMonitor) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateLoadBalancerRequestMonitor) GetType() *string {
	return s.Type
}

func (s *UpdateLoadBalancerRequestMonitor) SetConsecutiveDown(v int32) *UpdateLoadBalancerRequestMonitor {
	s.ConsecutiveDown = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetConsecutiveUp(v int32) *UpdateLoadBalancerRequestMonitor {
	s.ConsecutiveUp = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetExpectedCodes(v string) *UpdateLoadBalancerRequestMonitor {
	s.ExpectedCodes = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetFollowRedirects(v bool) *UpdateLoadBalancerRequestMonitor {
	s.FollowRedirects = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetHeader(v interface{}) *UpdateLoadBalancerRequestMonitor {
	s.Header = v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetInterval(v int32) *UpdateLoadBalancerRequestMonitor {
	s.Interval = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetMethod(v string) *UpdateLoadBalancerRequestMonitor {
	s.Method = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetMonitoringRegion(v string) *UpdateLoadBalancerRequestMonitor {
	s.MonitoringRegion = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetPath(v string) *UpdateLoadBalancerRequestMonitor {
	s.Path = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetPort(v int32) *UpdateLoadBalancerRequestMonitor {
	s.Port = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetTimeout(v int32) *UpdateLoadBalancerRequestMonitor {
	s.Timeout = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) SetType(v string) *UpdateLoadBalancerRequestMonitor {
	s.Type = &v
	return s
}

func (s *UpdateLoadBalancerRequestMonitor) Validate() error {
	return dara.Validate(s)
}

type UpdateLoadBalancerRequestRandomSteering struct {
	// The default round-robin weight for all pools that do not have a specific weight configured. Valid values: integers from 0 to 100.
	//
	// example:
	//
	// 50
	DefaultWeight *int32 `json:"DefaultWeight,omitempty" xml:"DefaultWeight,omitempty"`
	// The weight configuration for each backend server pool. The key is the pool ID and the value is the weight coefficient. The weight coefficient indicates the proportion of relative traffic distribution.
	PoolWeights map[string]*int32 `json:"PoolWeights,omitempty" xml:"PoolWeights,omitempty"`
}

func (s UpdateLoadBalancerRequestRandomSteering) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerRequestRandomSteering) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerRequestRandomSteering) GetDefaultWeight() *int32 {
	return s.DefaultWeight
}

func (s *UpdateLoadBalancerRequestRandomSteering) GetPoolWeights() map[string]*int32 {
	return s.PoolWeights
}

func (s *UpdateLoadBalancerRequestRandomSteering) SetDefaultWeight(v int32) *UpdateLoadBalancerRequestRandomSteering {
	s.DefaultWeight = &v
	return s
}

func (s *UpdateLoadBalancerRequestRandomSteering) SetPoolWeights(v map[string]*int32) *UpdateLoadBalancerRequestRandomSteering {
	s.PoolWeights = v
	return s
}

func (s *UpdateLoadBalancerRequestRandomSteering) Validate() error {
	return dara.Validate(s)
}

type UpdateLoadBalancerRequestRules struct {
	// The specified content to return in the response after a rule is matched.
	FixedResponse *UpdateLoadBalancerRequestRulesFixedResponse `json:"FixedResponse,omitempty" xml:"FixedResponse,omitempty" type:"Struct"`
	// The load balancing configuration to modify for request rewriting after a rule is matched. The configured fields overwrite the corresponding fields in the load balancing instance configuration.
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
	// The rule content. This parameter uses conditional expressions to match user requests. This parameter is not required when you add a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// http.request.method eq "GET"
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
	// The rule execution sequence. This parameter is optional. If you do not specify this parameter, rules are executed in the order they appear in the list. If you specify this parameter, the value must be an integer greater than 0. A larger value indicates a higher execution priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to stop executing subsequent rules.
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

func (s UpdateLoadBalancerRequestRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerRequestRules) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerRequestRules) GetFixedResponse() *UpdateLoadBalancerRequestRulesFixedResponse {
	return s.FixedResponse
}

func (s *UpdateLoadBalancerRequestRules) GetOverrides() interface{} {
	return s.Overrides
}

func (s *UpdateLoadBalancerRequestRules) GetRule() *string {
	return s.Rule
}

func (s *UpdateLoadBalancerRequestRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateLoadBalancerRequestRules) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateLoadBalancerRequestRules) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateLoadBalancerRequestRules) GetTerminates() *bool {
	return s.Terminates
}

func (s *UpdateLoadBalancerRequestRules) SetFixedResponse(v *UpdateLoadBalancerRequestRulesFixedResponse) *UpdateLoadBalancerRequestRules {
	s.FixedResponse = v
	return s
}

func (s *UpdateLoadBalancerRequestRules) SetOverrides(v interface{}) *UpdateLoadBalancerRequestRules {
	s.Overrides = v
	return s
}

func (s *UpdateLoadBalancerRequestRules) SetRule(v string) *UpdateLoadBalancerRequestRules {
	s.Rule = &v
	return s
}

func (s *UpdateLoadBalancerRequestRules) SetRuleEnable(v string) *UpdateLoadBalancerRequestRules {
	s.RuleEnable = &v
	return s
}

func (s *UpdateLoadBalancerRequestRules) SetRuleName(v string) *UpdateLoadBalancerRequestRules {
	s.RuleName = &v
	return s
}

func (s *UpdateLoadBalancerRequestRules) SetSequence(v int32) *UpdateLoadBalancerRequestRules {
	s.Sequence = &v
	return s
}

func (s *UpdateLoadBalancerRequestRules) SetTerminates(v bool) *UpdateLoadBalancerRequestRules {
	s.Terminates = &v
	return s
}

func (s *UpdateLoadBalancerRequestRules) Validate() error {
	if s.FixedResponse != nil {
		if err := s.FixedResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLoadBalancerRequestRulesFixedResponse struct {
	// The Content-Type field in the HTTP header.
	//
	// example:
	//
	// application/json
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The Location field in the HTTP response.
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

func (s UpdateLoadBalancerRequestRulesFixedResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerRequestRulesFixedResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerRequestRulesFixedResponse) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateLoadBalancerRequestRulesFixedResponse) GetLocation() *string {
	return s.Location
}

func (s *UpdateLoadBalancerRequestRulesFixedResponse) GetMessageBody() *string {
	return s.MessageBody
}

func (s *UpdateLoadBalancerRequestRulesFixedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLoadBalancerRequestRulesFixedResponse) SetContentType(v string) *UpdateLoadBalancerRequestRulesFixedResponse {
	s.ContentType = &v
	return s
}

func (s *UpdateLoadBalancerRequestRulesFixedResponse) SetLocation(v string) *UpdateLoadBalancerRequestRulesFixedResponse {
	s.Location = &v
	return s
}

func (s *UpdateLoadBalancerRequestRulesFixedResponse) SetMessageBody(v string) *UpdateLoadBalancerRequestRulesFixedResponse {
	s.MessageBody = &v
	return s
}

func (s *UpdateLoadBalancerRequestRulesFixedResponse) SetStatusCode(v int32) *UpdateLoadBalancerRequestRulesFixedResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerRequestRulesFixedResponse) Validate() error {
	return dara.Validate(s)
}
