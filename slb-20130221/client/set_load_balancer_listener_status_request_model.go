// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerListenerStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostId(v string) *SetLoadBalancerListenerStatusRequest
	GetHostId() *string
	SetListenerPort(v int32) *SetLoadBalancerListenerStatusRequest
	GetListenerPort() *int32
	SetListenerStatus(v string) *SetLoadBalancerListenerStatusRequest
	GetListenerStatus() *string
	SetLoadBalancerId(v string) *SetLoadBalancerListenerStatusRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *SetLoadBalancerListenerStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SetLoadBalancerListenerStatusRequest
	GetOwnerId() *string
}

type SetLoadBalancerListenerStatusRequest struct {
	HostId         *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	ListenerPort   *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SetLoadBalancerListenerStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerListenerStatusRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerListenerStatusRequest) GetHostId() *string {
	return s.HostId
}

func (s *SetLoadBalancerListenerStatusRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetLoadBalancerListenerStatusRequest) GetListenerStatus() *string {
	return s.ListenerStatus
}

func (s *SetLoadBalancerListenerStatusRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerListenerStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerListenerStatusRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetLoadBalancerListenerStatusRequest) SetHostId(v string) *SetLoadBalancerListenerStatusRequest {
	s.HostId = &v
	return s
}

func (s *SetLoadBalancerListenerStatusRequest) SetListenerPort(v int32) *SetLoadBalancerListenerStatusRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerListenerStatusRequest) SetListenerStatus(v string) *SetLoadBalancerListenerStatusRequest {
	s.ListenerStatus = &v
	return s
}

func (s *SetLoadBalancerListenerStatusRequest) SetLoadBalancerId(v string) *SetLoadBalancerListenerStatusRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerListenerStatusRequest) SetOwnerAccount(v string) *SetLoadBalancerListenerStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerListenerStatusRequest) SetOwnerId(v string) *SetLoadBalancerListenerStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerListenerStatusRequest) Validate() error {
	return dara.Validate(s)
}
