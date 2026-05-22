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
	AdaptiveRouting *GetLoadBalancerResponseBodyAdaptiveRouting `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty" type:"Struct"`
	DefaultPools    []*int64                                    `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty" type:"Repeated"`
	Description     *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled         *bool                                       `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FallbackPool    *int64                                      `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	Id              *int64                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Monitor         *GetLoadBalancerResponseBodyMonitor         `json:"Monitor,omitempty" xml:"Monitor,omitempty" type:"Struct"`
	Name            *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RandomSteering  *GetLoadBalancerResponseBodyRandomSteering  `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty" type:"Struct"`
	RegionPools     interface{}                                 `json:"RegionPools,omitempty" xml:"RegionPools,omitempty"`
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules           []*GetLoadBalancerResponseBodyRules         `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	SessionAffinity *string                                     `json:"SessionAffinity,omitempty" xml:"SessionAffinity,omitempty"`
	SiteId          *int64                                      `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	Status          *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	SteeringPolicy  *string                                     `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	SubRegionPools  interface{}                                 `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	Ttl             *int32                                      `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
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
	FailoverAcrossPools *bool `json:"FailoverAcrossPools,omitempty" xml:"FailoverAcrossPools,omitempty"`
	OriginLevelRetry    *bool `json:"OriginLevelRetry,omitempty" xml:"OriginLevelRetry,omitempty"`
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
	ConsecutiveDown  *int32      `json:"ConsecutiveDown,omitempty" xml:"ConsecutiveDown,omitempty"`
	ConsecutiveUp    *int32      `json:"ConsecutiveUp,omitempty" xml:"ConsecutiveUp,omitempty"`
	ExpectedCodes    *string     `json:"ExpectedCodes,omitempty" xml:"ExpectedCodes,omitempty"`
	FollowRedirects  *bool       `json:"FollowRedirects,omitempty" xml:"FollowRedirects,omitempty"`
	Header           interface{} `json:"Header,omitempty" xml:"Header,omitempty"`
	Interval         *int32      `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Method           *string     `json:"Method,omitempty" xml:"Method,omitempty"`
	MonitoringRegion *string     `json:"MonitoringRegion,omitempty" xml:"MonitoringRegion,omitempty"`
	Path             *string     `json:"Path,omitempty" xml:"Path,omitempty"`
	Port             *int32      `json:"Port,omitempty" xml:"Port,omitempty"`
	Timeout          *int32      `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Type             *string     `json:"Type,omitempty" xml:"Type,omitempty"`
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
	DefaultWeight *int32            `json:"DefaultWeight,omitempty" xml:"DefaultWeight,omitempty"`
	PoolWeights   map[string]*int32 `json:"PoolWeights,omitempty" xml:"PoolWeights,omitempty"`
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
	FixedResponse *GetLoadBalancerResponseBodyRulesFixedResponse `json:"FixedResponse,omitempty" xml:"FixedResponse,omitempty" type:"Struct"`
	Overrides     interface{}                                    `json:"Overrides,omitempty" xml:"Overrides,omitempty"`
	Rule          *string                                        `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable    *string                                        `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName      *string                                        `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence      *int32                                         `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	Terminates    *bool                                          `json:"Terminates,omitempty" xml:"Terminates,omitempty"`
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
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Location    *string `json:"Location,omitempty" xml:"Location,omitempty"`
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	StatusCode  *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
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
