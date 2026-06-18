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
	// Configures origin-pull behavior across address pools.
	AdaptiveRouting *UpdateLoadBalancerRequestAdaptiveRouting `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty" type:"Struct"`
	// A list of default address pool IDs.
	DefaultPools []*int64 `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty" type:"Repeated"`
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
	Monitor *UpdateLoadBalancerRequestMonitor `json:"Monitor,omitempty" xml:"Monitor,omitempty" type:"Struct"`
	// The configuration for weighted round-robin. This setting controls the weight of traffic distributed to different address pools.
	RandomSteering *UpdateLoadBalancerRequestRandomSteering `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty" type:"Struct"`
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
	Rules []*UpdateLoadBalancerRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
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
	// Specifies whether to perform origin-pull across address pools.
	//
	// - `true`: Enables origin-pull across address pools.
	//
	// - `false`: Disables origin-pull across address pools.
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
	// The number of consecutive failed health checks required to declare an origin server unhealthy. For example, `5`.
	//
	// example:
	//
	// 5
	ConsecutiveDown *int32 `json:"ConsecutiveDown,omitempty" xml:"ConsecutiveDown,omitempty"`
	// The number of consecutive successful health checks required to declare an origin server healthy. For example, `3`.
	//
	// example:
	//
	// 3
	ConsecutiveUp *int32 `json:"ConsecutiveUp,omitempty" xml:"ConsecutiveUp,omitempty"`
	// The expected HTTP status codes that indicate a healthy response. For example, `200,202`.
	//
	// example:
	//
	// 200,202
	ExpectedCodes *string `json:"ExpectedCodes,omitempty" xml:"ExpectedCodes,omitempty"`
	// Specifies whether the health check monitor follows HTTP redirections.
	//
	// - `true`: The monitor follows HTTP redirections.
	//
	// - `false`: The monitor does not follow HTTP redirections.
	//
	// example:
	//
	// true
	FollowRedirects *bool `json:"FollowRedirects,omitempty" xml:"FollowRedirects,omitempty"`
	// The HTTP request headers to send with each health check.
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
	// The interval in seconds between each health check. For example, `60`.
	//
	// example:
	//
	// 100
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The HTTP method to use for the health check. For example, `GET`.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The regions from which the health checks are performed. The default value is `Global`.
	//
	// - `Global`: From probe locations worldwide.
	//
	// - `ChineseMainland`: From probe locations within the Chinese mainland.
	//
	// - `OutsideChineseMainland`: From probe locations outside the Chinese mainland.
	//
	// example:
	//
	// Global
	MonitoringRegion *string `json:"MonitoringRegion,omitempty" xml:"MonitoringRegion,omitempty"`
	// The path on the origin server to request for the health check. For example, `/healthcheck`.
	//
	// example:
	//
	// /health
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port on the origin server to use for the health check.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The timeout for the health check, in seconds. The value must be between 1 and 10, inclusive.
	//
	// example:
	//
	// 5
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The protocol to use for the health check, such as `HTTP`. If you set this to `off`, no health check is performed.
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
	// The default weight applied to all address pools that do not have a specific weight defined. The value must be an integer from 0 to 100, inclusive.
	//
	// example:
	//
	// 50
	DefaultWeight *int32 `json:"DefaultWeight,omitempty" xml:"DefaultWeight,omitempty"`
	// A map of pool IDs to their corresponding weights. These weights determine the proportion of traffic routed to each backend server pool.
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
	// The fixed response to return when the rule\\"s condition is met.
	FixedResponse *UpdateLoadBalancerRequestRulesFixedResponse `json:"FixedResponse,omitempty" xml:"FixedResponse,omitempty" type:"Struct"`
	// The settings to override for requests that match this rule\\"s condition. These settings take precedence over the load balancer\\"s main configuration.
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
	// The content of the rule, specified as a conditional expression to match user requests. This parameter is not required when you configure global settings. Use cases:
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, set the value to a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// http.request.method eq "GET"
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. This parameter is not required when you configure global settings. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. This parameter is not required when you configure global settings.
	//
	// example:
	//
	// rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution priority of the rule. This parameter is optional. If you do not specify this parameter, rules are executed in the order they are listed. If specified, the value must be an integer greater than 0. A larger value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to stop processing subsequent rules after this rule is matched.
	//
	// - `true`: Stop processing subsequent rules.
	//
	// - `false`: Continue processing subsequent rules. This is the default value.
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
	// The value of the `Content-Type` field in the HTTP response header.
	//
	// example:
	//
	// application/json
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The value of the `Location` field in the HTTP response header. This is typically used for redirections.
	//
	// example:
	//
	// http://www.example.com/index.html
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The content of the response body.
	//
	// example:
	//
	// Hello World!
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// The HTTP status code of the response.
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
