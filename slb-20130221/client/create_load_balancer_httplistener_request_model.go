// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerHTTPListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetBandwidth() *int32
	SetCookie(v string) *CreateLoadBalancerHTTPListenerRequest
	GetCookie() *string
	SetCookieTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetCookieTimeout() *int32
	SetDomain(v string) *CreateLoadBalancerHTTPListenerRequest
	GetDomain() *string
	SetHealthCheck(v string) *CreateLoadBalancerHTTPListenerRequest
	GetHealthCheck() *string
	SetHealthCheckTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetHealthCheckTimeout() *int32
	SetHealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetHealthyThreshold() *int32
	SetHostId(v string) *CreateLoadBalancerHTTPListenerRequest
	GetHostId() *string
	SetInterval(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetInterval() *int32
	SetListenerPort(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetListenerPort() *int32
	SetListenerStatus(v string) *CreateLoadBalancerHTTPListenerRequest
	GetListenerStatus() *string
	SetLoadBalancerId(v string) *CreateLoadBalancerHTTPListenerRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *CreateLoadBalancerHTTPListenerRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *CreateLoadBalancerHTTPListenerRequest
	GetOwnerId() *string
	SetResourceOwnerId(v int64) *CreateLoadBalancerHTTPListenerRequest
	GetResourceOwnerId() *int64
	SetScheduler(v string) *CreateLoadBalancerHTTPListenerRequest
	GetScheduler() *string
	SetStickySession(v string) *CreateLoadBalancerHTTPListenerRequest
	GetStickySession() *string
	SetStickySessionType(v string) *CreateLoadBalancerHTTPListenerRequest
	GetStickySessionType() *string
	SetURI(v string) *CreateLoadBalancerHTTPListenerRequest
	GetURI() *string
	SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetUnhealthyThreshold() *int32
	SetXForwardedFor(v string) *CreateLoadBalancerHTTPListenerRequest
	GetXForwardedFor() *string
}

type CreateLoadBalancerHTTPListenerRequest struct {
	BackendServerPort  *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Cookie             *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout      *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Domain             *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HealthCheck        *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckTimeout *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthyThreshold   *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	HostId             *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Interval           *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ListenerPort       *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerStatus     *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	LoadBalancerId     *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerId    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scheduler          *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	StickySession      *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType  *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	URI                *string `json:"URI,omitempty" xml:"URI,omitempty"`
	UnhealthyThreshold *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	XForwardedFor      *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
}

func (s CreateLoadBalancerHTTPListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetCookie() *string {
	return s.Cookie
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHostId() *string {
	return s.HostId
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetListenerStatus() *string {
	return s.ListenerStatus
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetStickySession() *string {
	return s.StickySession
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetURI() *string {
	return s.URI
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetBandwidth(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetCookie(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Cookie = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetCookieTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.CookieTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetDomain(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Domain = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheck(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheck = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHostId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HostId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetInterval(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.Interval = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetListenerStatus(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.ListenerStatus = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetOwnerAccount(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetOwnerId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerHTTPListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetScheduler(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetStickySession(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.StickySession = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetStickySessionType(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.StickySessionType = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetURI(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.URI = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) Validate() error {
	return dara.Validate(s)
}
