// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerHTTPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookie(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetCookie() *string
	SetCookieTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetCookieTimeout() *int32
	SetDomain(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetDomain() *string
	SetHealthCheck(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthCheck() *string
	SetHealthCheckTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthCheckTimeout() *int32
	SetHealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthyThreshold() *int32
	SetHostId(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHostId() *string
	SetInterval(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetInterval() *int32
	SetListenerPort(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetOwnerId() *string
	SetScheduler(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetScheduler() *string
	SetStickySession(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetStickySession() *string
	SetStickySessionType(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetStickySessionType() *string
	SetURI(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetURI() *string
	SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetUnhealthyThreshold() *int32
	SetXForwardedFor(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetXForwardedFor() *string
}

type SetLoadBalancerHTTPListenerAttributeRequest struct {
	Cookie             *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout      *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Domain             *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HealthCheck        *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckTimeout *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthyThreshold   *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	HostId             *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Interval           *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ListenerPort       *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId     *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Scheduler          *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	StickySession      *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType  *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	URI                *string `json:"URI,omitempty" xml:"URI,omitempty"`
	UnhealthyThreshold *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	XForwardedFor      *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
}

func (s SetLoadBalancerHTTPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerHTTPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetCookie() *string {
	return s.Cookie
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetDomain() *string {
	return s.Domain
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHostId() *string {
	return s.HostId
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetStickySession() *string {
	return s.StickySession
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetURI() *string {
	return s.URI
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetCookie(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Cookie = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetCookieTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.CookieTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetDomain(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Domain = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheck(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheck = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHostId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HostId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetInterval(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Interval = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetOwnerAccount(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetOwnerId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetStickySession(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.StickySession = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetStickySessionType(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.StickySessionType = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetURI(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.URI = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
