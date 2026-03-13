// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerTCPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetBandwidth() *int32
	SetConnectPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetConnectPort() *int32
	SetConnectTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetConnectTimeout() *int32
	SetEstablishedTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetEstablishedTimeout() *int32
	SetHealthCheck(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheck() *string
	SetHealthCheckDomain(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckHttpCode() *string
	SetHealthCheckType(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckType() *string
	SetHealthCheckURI(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthyThreshold() *int32
	SetInterval(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetInterval() *int32
	SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetListenerPort() *int32
	SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetMasterSlaveServerGroupId() *string
	SetMaxConnection(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetMaxConnection() *int32
	SetPersistenceTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetPersistenceTimeout() *int32
	SetRequestId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetRequestId() *string
	SetScheduler(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetScheduler() *string
	SetStatus(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetStatus() *string
	SetSynProxy(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetSynProxy() *string
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetVServerGroupId() *string
}

type DescribeLoadBalancerTCPListenerAttributeResponseBody struct {
	BackendServerPort        *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth                *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectPort              *int32  `json:"ConnectPort,omitempty" xml:"ConnectPort,omitempty"`
	ConnectTimeout           *int32  `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	EstablishedTimeout       *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	HealthCheck              *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckDomain        *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode      *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckType          *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	HealthCheckURI           *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold         *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	Interval                 *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ListenerPort             *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	MaxConnection            *int32  `json:"MaxConnection,omitempty" xml:"MaxConnection,omitempty"`
	PersistenceTimeout       *int32  `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scheduler                *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SynProxy                 *string `json:"SynProxy,omitempty" xml:"SynProxy,omitempty"`
	UnhealthyThreshold       *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId           *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetConnectPort() *int32 {
	return s.ConnectPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetConnectTimeout() *int32 {
	return s.ConnectTimeout
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetMaxConnection() *int32 {
	return s.MaxConnection
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetSynProxy() *string {
	return s.SynProxy
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetConnectPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetConnectTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetEstablishedTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.EstablishedTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckType(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckURI(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetInterval(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetMaxConnection(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.MaxConnection = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetPersistenceTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetSynProxy(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.SynProxy = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetVServerGroupId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
