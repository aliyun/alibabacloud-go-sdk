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
	AdaptiveRouting *UpdateLoadBalancerRequestAdaptiveRouting `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty" type:"Struct"`
	DefaultPools    []*int64                                  `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty" type:"Repeated"`
	Description     *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled         *bool                                     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FallbackPool    *int64                                    `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// This parameter is required.
	Id             *int64                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	Monitor        *UpdateLoadBalancerRequestMonitor        `json:"Monitor,omitempty" xml:"Monitor,omitempty" type:"Struct"`
	RandomSteering *UpdateLoadBalancerRequestRandomSteering `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty" type:"Struct"`
	RegionPools    interface{}                              `json:"RegionPools,omitempty" xml:"RegionPools,omitempty"`
	// if can be null:
	// false
	Rules           []*UpdateLoadBalancerRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	SessionAffinity *string                           `json:"SessionAffinity,omitempty" xml:"SessionAffinity,omitempty"`
	// This parameter is required.
	SiteId         *int64      `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SteeringPolicy *string     `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	Ttl            *int32      `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
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
	DefaultWeight *int32            `json:"DefaultWeight,omitempty" xml:"DefaultWeight,omitempty"`
	PoolWeights   map[string]*int32 `json:"PoolWeights,omitempty" xml:"PoolWeights,omitempty"`
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
	FixedResponse *UpdateLoadBalancerRequestRulesFixedResponse `json:"FixedResponse,omitempty" xml:"FixedResponse,omitempty" type:"Struct"`
	Overrides     interface{}                                  `json:"Overrides,omitempty" xml:"Overrides,omitempty"`
	Rule          *string                                      `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable    *string                                      `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName      *string                                      `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence      *int32                                       `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	Terminates    *bool                                        `json:"Terminates,omitempty" xml:"Terminates,omitempty"`
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
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Location    *string `json:"Location,omitempty" xml:"Location,omitempty"`
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	StatusCode  *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
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
