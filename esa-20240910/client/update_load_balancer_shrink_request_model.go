// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptiveRoutingShrink(v string) *UpdateLoadBalancerShrinkRequest
	GetAdaptiveRoutingShrink() *string
	SetDefaultPoolsShrink(v string) *UpdateLoadBalancerShrinkRequest
	GetDefaultPoolsShrink() *string
	SetDescription(v string) *UpdateLoadBalancerShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdateLoadBalancerShrinkRequest
	GetEnabled() *bool
	SetFallbackPool(v int64) *UpdateLoadBalancerShrinkRequest
	GetFallbackPool() *int64
	SetId(v int64) *UpdateLoadBalancerShrinkRequest
	GetId() *int64
	SetMonitorShrink(v string) *UpdateLoadBalancerShrinkRequest
	GetMonitorShrink() *string
	SetRandomSteeringShrink(v string) *UpdateLoadBalancerShrinkRequest
	GetRandomSteeringShrink() *string
	SetRegionPools(v interface{}) *UpdateLoadBalancerShrinkRequest
	GetRegionPools() interface{}
	SetRulesShrink(v string) *UpdateLoadBalancerShrinkRequest
	GetRulesShrink() *string
	SetSessionAffinity(v string) *UpdateLoadBalancerShrinkRequest
	GetSessionAffinity() *string
	SetSiteId(v int64) *UpdateLoadBalancerShrinkRequest
	GetSiteId() *int64
	SetSteeringPolicy(v string) *UpdateLoadBalancerShrinkRequest
	GetSteeringPolicy() *string
	SetSubRegionPools(v interface{}) *UpdateLoadBalancerShrinkRequest
	GetSubRegionPools() interface{}
	SetTtl(v int32) *UpdateLoadBalancerShrinkRequest
	GetTtl() *int32
}

type UpdateLoadBalancerShrinkRequest struct {
	AdaptiveRoutingShrink *string `json:"AdaptiveRouting,omitempty" xml:"AdaptiveRouting,omitempty"`
	DefaultPoolsShrink    *string `json:"DefaultPools,omitempty" xml:"DefaultPools,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled               *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FallbackPool          *int64  `json:"FallbackPool,omitempty" xml:"FallbackPool,omitempty"`
	// This parameter is required.
	Id                   *int64      `json:"Id,omitempty" xml:"Id,omitempty"`
	MonitorShrink        *string     `json:"Monitor,omitempty" xml:"Monitor,omitempty"`
	RandomSteeringShrink *string     `json:"RandomSteering,omitempty" xml:"RandomSteering,omitempty"`
	RegionPools          interface{} `json:"RegionPools,omitempty" xml:"RegionPools,omitempty"`
	// if can be null:
	// false
	RulesShrink     *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	SessionAffinity *string `json:"SessionAffinity,omitempty" xml:"SessionAffinity,omitempty"`
	// This parameter is required.
	SiteId         *int64      `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SteeringPolicy *string     `json:"SteeringPolicy,omitempty" xml:"SteeringPolicy,omitempty"`
	SubRegionPools interface{} `json:"SubRegionPools,omitempty" xml:"SubRegionPools,omitempty"`
	Ttl            *int32      `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateLoadBalancerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerShrinkRequest) GetAdaptiveRoutingShrink() *string {
	return s.AdaptiveRoutingShrink
}

func (s *UpdateLoadBalancerShrinkRequest) GetDefaultPoolsShrink() *string {
	return s.DefaultPoolsShrink
}

func (s *UpdateLoadBalancerShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLoadBalancerShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateLoadBalancerShrinkRequest) GetFallbackPool() *int64 {
	return s.FallbackPool
}

func (s *UpdateLoadBalancerShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateLoadBalancerShrinkRequest) GetMonitorShrink() *string {
	return s.MonitorShrink
}

func (s *UpdateLoadBalancerShrinkRequest) GetRandomSteeringShrink() *string {
	return s.RandomSteeringShrink
}

func (s *UpdateLoadBalancerShrinkRequest) GetRegionPools() interface{} {
	return s.RegionPools
}

func (s *UpdateLoadBalancerShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *UpdateLoadBalancerShrinkRequest) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *UpdateLoadBalancerShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateLoadBalancerShrinkRequest) GetSteeringPolicy() *string {
	return s.SteeringPolicy
}

func (s *UpdateLoadBalancerShrinkRequest) GetSubRegionPools() interface{} {
	return s.SubRegionPools
}

func (s *UpdateLoadBalancerShrinkRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateLoadBalancerShrinkRequest) SetAdaptiveRoutingShrink(v string) *UpdateLoadBalancerShrinkRequest {
	s.AdaptiveRoutingShrink = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetDefaultPoolsShrink(v string) *UpdateLoadBalancerShrinkRequest {
	s.DefaultPoolsShrink = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetDescription(v string) *UpdateLoadBalancerShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetEnabled(v bool) *UpdateLoadBalancerShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetFallbackPool(v int64) *UpdateLoadBalancerShrinkRequest {
	s.FallbackPool = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetId(v int64) *UpdateLoadBalancerShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetMonitorShrink(v string) *UpdateLoadBalancerShrinkRequest {
	s.MonitorShrink = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetRandomSteeringShrink(v string) *UpdateLoadBalancerShrinkRequest {
	s.RandomSteeringShrink = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetRegionPools(v interface{}) *UpdateLoadBalancerShrinkRequest {
	s.RegionPools = v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetRulesShrink(v string) *UpdateLoadBalancerShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetSessionAffinity(v string) *UpdateLoadBalancerShrinkRequest {
	s.SessionAffinity = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetSiteId(v int64) *UpdateLoadBalancerShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetSteeringPolicy(v string) *UpdateLoadBalancerShrinkRequest {
	s.SteeringPolicy = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetSubRegionPools(v interface{}) *UpdateLoadBalancerShrinkRequest {
	s.SubRegionPools = v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) SetTtl(v int32) *UpdateLoadBalancerShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateLoadBalancerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
