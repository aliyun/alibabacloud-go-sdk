// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerHTTPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetBandwidth() *int32
	SetCookie(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetCookie() *string
	SetCookieTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetCookieTimeout() *int32
	SetDomain(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetDomain() *string
	SetGzip(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetGzip() *string
	SetHealthCheck(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheckConnectPort() *int32
	SetHealthCheckHttpCode(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheckHttpCode() *string
	SetHealthCheckTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheckTimeout() *int32
	SetHealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthyThreshold() *int32
	SetInterval(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetInterval() *int32
	SetListenerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetListenerPort() *int32
	SetMaxConnection(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetMaxConnection() *int32
	SetRequestId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetRequestId() *string
	SetScheduler(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetScheduler() *string
	SetSecurityStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetSecurityStatus() *string
	SetStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetStatus() *string
	SetStickySession(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetStickySession() *string
	SetStickySessionType(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetStickySessionType() *string
	SetURI(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetURI() *string
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetVServerGroupId() *string
	SetXForwardedFor(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor() *string
	SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor_SLBID() *string
	SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor_SLBIP() *string
	SetXForwardedFor_proto(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor_proto() *string
}

type DescribeLoadBalancerHTTPListenerAttributeResponseBody struct {
	BackendServerPort      *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth              *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Cookie                 *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout          *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Domain                 *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Gzip                   *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	HealthCheck            *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckHttpCode    *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckTimeout     *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthyThreshold       *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	Interval               *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ListenerPort           *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	MaxConnection          *int32  `json:"MaxConnection,omitempty" xml:"MaxConnection,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scheduler              *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	SecurityStatus         *string `json:"SecurityStatus,omitempty" xml:"SecurityStatus,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StickySession          *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType      *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	URI                    *string `json:"URI,omitempty" xml:"URI,omitempty"`
	UnhealthyThreshold     *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId         *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	XForwardedFor          *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	XForwardedFor_SLBID    *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	XForwardedFor_SLBIP    *string `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	XForwardedFor_proto    *string `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetGzip() *string {
	return s.Gzip
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetMaxConnection() *int32 {
	return s.MaxConnection
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetSecurityStatus() *string {
	return s.SecurityStatus
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetStickySession() *string {
	return s.StickySession
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetURI() *string {
	return s.URI
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor_SLBID() *string {
	return s.XForwardedFor_SLBID
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor_SLBIP() *string {
	return s.XForwardedFor_SLBIP
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor_proto() *string {
	return s.XForwardedFor_proto
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetCookie(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Cookie = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetCookieTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetDomain(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetGzip(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Gzip = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetInterval(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetMaxConnection(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.MaxConnection = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetSecurityStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.SecurityStatus = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetStickySession(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.StickySession = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetStickySessionType(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.StickySessionType = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetURI(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.URI = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetVServerGroupId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_proto(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_proto = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
