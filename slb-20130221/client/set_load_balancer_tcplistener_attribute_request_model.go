// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerTCPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetConnectPort() *int32
	SetConnectTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetConnectTimeout() *int32
	SetHealthCheck(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheck() *string
	SetHealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthyThreshold() *int32
	SetHostId(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHostId() *string
	SetInterval(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetInterval() *int32
	SetListenerPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetOwnerId() *string
	SetPersistenceTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetPersistenceTimeout() *int32
	SetScheduler(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetScheduler() *string
	SetUnhealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetUnhealthyThreshold() *int32
}

type SetLoadBalancerTCPListenerAttributeRequest struct {
	ConnectPort        *int32  `json:"ConnectPort,omitempty" xml:"ConnectPort,omitempty"`
	ConnectTimeout     *int32  `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	HealthCheck        *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthyThreshold   *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	HostId             *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Interval           *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ListenerPort       *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId     *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PersistenceTimeout *int32  `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	Scheduler          *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	UnhealthyThreshold *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s SetLoadBalancerTCPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerTCPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetConnectPort() *int32 {
	return s.ConnectPort
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetConnectTimeout() *int32 {
	return s.ConnectTimeout
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHostId() *string {
	return s.HostId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetConnectPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ConnectPort = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetConnectTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ConnectTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheck(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheck = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHostId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HostId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetInterval(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.Interval = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetOwnerAccount(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetOwnerId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetPersistenceTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.PersistenceTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
