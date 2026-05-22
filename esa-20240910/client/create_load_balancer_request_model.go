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
	AdaptiveRouting *CreateLoadBalancerRequestAdaptiveRouting `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty" type:"Struct"`
	// This parameter is required.
	DefaultPools []*int64 `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty" type:"Repeated"`
	Description  *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled      *bool    `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	FallbackPool *int64 `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// This parameter is required.
	Monitor *CreateLoadBalancerRequestMonitor `json:"Monitor,omitempty" xml:"Monitor,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	RandomSteering  *CreateLoadBalancerRequestRandomSteering `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty" type:"Struct"`
	RegionPools     interface{}                              `json:"RegionPools,omitempty" xml:"RegionPools,omitempty"`
	Rules           []*CreateLoadBalancerRequestRules        `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	SessionAffinity *string                                  `json:"SessionAffinity,omitempty" xml:"SessionAffinity,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	SteeringPolicy *string     `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	Ttl            *int32      `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
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
	FailoverAcrossPools *bool `json:"FailoverAcrossPools,omitempty" xml:"FailoverAcrossPools,omitempty"`
	OriginLevelRetry    *bool `json:"OriginLevelRetry,omitempty" xml:"OriginLevelRetry,omitempty"`
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
	DefaultWeight *int32            `json:"DefaultWeight,omitempty" xml:"DefaultWeight,omitempty"`
	PoolWeights   map[string]*int32 `json:"PoolWeights,omitempty" xml:"PoolWeights,omitempty"`
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
	FixedResponse *CreateLoadBalancerRequestRulesFixedResponse `json:"FixedResponse,omitempty" xml:"FixedResponse,omitempty" type:"Struct"`
	Overrides     interface{}                                  `json:"Overrides,omitempty" xml:"Overrides,omitempty"`
	Rule          *string                                      `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable    *string                                      `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName      *string                                      `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence      *int32                                       `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	Terminates    *bool                                        `json:"Terminates,omitempty" xml:"Terminates,omitempty"`
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
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Location    *string `json:"Location,omitempty" xml:"Location,omitempty"`
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	StatusCode  *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
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
