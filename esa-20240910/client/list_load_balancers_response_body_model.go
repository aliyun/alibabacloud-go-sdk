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
	// An array of load balancers.
	LoadBalancers []*ListLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	// The page number. This value matches the `PageNumber` request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
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
	// The total number of entries found.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
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
	if s.LoadBalancers != nil {
		for _, item := range s.LoadBalancers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLoadBalancersResponseBodyLoadBalancers struct {
	// The configuration for failover across pools.
	AdaptiveRouting *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty" type:"Struct"`
	// A list of default pool IDs.
	DefaultPools []*int64 `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty" type:"Repeated"`
	// The description of the load balancer.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the load balancer is enabled.
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the fallback pool. The load balancer routes traffic to this pool when all other pools are unavailable.
	//
	// example:
	//
	// 96228666776****
	FallbackPool *int64 `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// The unique ID of the load balancer.
	//
	// example:
	//
	// 998676487607104
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The health check configuration.
	Monitor *ListLoadBalancersResponseBodyLoadBalancersMonitor `json:"Monitor,omitempty" xml:"Monitor,omitempty" type:"Struct"`
	// The name of the load balancer.
	//
	// example:
	//
	// lb.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration for weighted round robin, which controls traffic distribution among pools.
	RandomSteering *ListLoadBalancersResponseBodyLoadBalancersRandomSteering `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty" type:"Struct"`
	// The pools that correspond to regions.
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
	// A list of rule configurations that define behavior for specific conditions.
	Rules []*ListLoadBalancersResponseBodyLoadBalancersRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The session affinity setting. Valid values:
	//
	// - `off`: Session affinity is disabled.
	//
	// - `ip`: Enables session affinity based on the client\\"s IP address.
	//
	// - `cookie`: Enables session affinity based on a cookie.
	//
	// example:
	//
	// ip
	SessionAffinity *string `json:"SessionAffinity,omitempty" xml:"SessionAffinity,omitempty"`
	// The site ID of the load balancer.
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
	// The pools that correspond to sub-regions. If multiple sub-regions share the same set of pools, you can use a comma-separated list of sub-region codes as the key.
	//
	// example:
	//
	// {"AL,MO": [92298024898****],"CN-SH,CN-SX,CN-SC":[92304347804****,92843536908****]}
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	// The Time to Live (TTL) for the DNS record, in seconds. The default value is 30.
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

type ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting struct {
	// Indicates whether to enable failover across pools.
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// true
	FailoverAcrossPools *bool `json:"FailoverAcrossPools,omitempty" xml:"FailoverAcrossPools,omitempty"`
	OriginLevelRetry    *bool `json:"OriginLevelRetry,omitempty" xml:"OriginLevelRetry,omitempty"`
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

func (s *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting) GetOriginLevelRetry() *bool {
	return s.OriginLevelRetry
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting) SetFailoverAcrossPools(v bool) *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting {
	s.FailoverAcrossPools = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting) SetOriginLevelRetry(v bool) *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting {
	s.OriginLevelRetry = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAdaptiveRouting) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersMonitor struct {
	// The number of consecutive failed health checks required to consider a pool unhealthy. For example, `5`.
	//
	// example:
	//
	// 5
	ConsecutiveDown *int32 `json:"ConsecutiveDown,omitempty" xml:"ConsecutiveDown,omitempty"`
	// The number of consecutive successful health checks required to consider a pool healthy. For example, `3`.
	//
	// example:
	//
	// 3
	ConsecutiveUp *int32 `json:"ConsecutiveUp,omitempty" xml:"ConsecutiveUp,omitempty"`
	// The expected HTTP status codes that indicate a successful health check, such as `200` or `202`.
	//
	// example:
	//
	// 200,202
	ExpectedCodes *string `json:"ExpectedCodes,omitempty" xml:"ExpectedCodes,omitempty"`
	// Indicates whether the health check should follow redirects.
	//
	// - `true`: Follows redirects.
	//
	// - `false`: Does not follow redirects.
	//
	// example:
	//
	// true
	FollowRedirects *bool `json:"FollowRedirects,omitempty" xml:"FollowRedirects,omitempty"`
	// The HTTP headers to include in the health check probe.
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
	// The interval between health checks, in seconds.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The method used for the health check.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The probe locations for health checks. The default is `Global`. Valid values:
	//
	// - `Global`: Sends probes from global locations.
	//
	// - `ChineseMainland`: Sends probes from locations within the Chinese mainland.
	//
	// - `OutsideChineseMainland`: Sends probes from global locations outside the Chinese mainland.
	//
	// example:
	//
	// Global
	MonitoringRegion *string `json:"MonitoringRegion,omitempty" xml:"MonitoringRegion,omitempty"`
	// The path to request for the health check.
	//
	// example:
	//
	// /
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The destination port for the health check.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The timeout for a single health check, in seconds.
	//
	// example:
	//
	// 5
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The protocol for the health check, such as `HTTP`. Set to `off` to disable health checks.
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
	// The default weight applied to any pool not defined in `PoolWeights`. The weight determines the proportion of traffic sent to the pool.
	//
	// example:
	//
	// 50
	DefaultWeight *int32 `json:"DefaultWeight,omitempty" xml:"DefaultWeight,omitempty"`
	// The weight configuration for each backend pool, where the key is the pool ID and the value is the weight.
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
	// Specifies the response to return when a rule matches.
	FixedResponse *ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse `json:"FixedResponse,omitempty" xml:"FixedResponse,omitempty" type:"Struct"`
	// The load balancer settings to override when a rule matches. Any field you specify overwrites the load balancer\\"s corresponding setting.
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
	// The rule expression that matches user requests. This parameter is not required for global configurations. Use cases:
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, set the value to a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// http.request.uri.path contains "/testing"
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Indicates whether the rule is enabled. This parameter is not required for global configurations. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. This parameter is not required for global configurations.
	//
	// example:
	//
	// r2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution priority of the rule. Higher values indicate higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to stop executing subsequent rules after this rule matches.
	//
	// - `true`: Stops executing subsequent rules.
	//
	// - `false`: Continues to execute subsequent rules. This is the default value.
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
	if s.FixedResponse != nil {
		if err := s.FixedResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLoadBalancersResponseBodyLoadBalancersRulesFixedResponse struct {
	// The `Content-Type` header.
	//
	// example:
	//
	// application/json
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The `Location` response header.
	//
	// example:
	//
	// http://www.example.com/index.html
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The content of the response body.
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
