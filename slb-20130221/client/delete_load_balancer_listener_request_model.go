// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoadBalancerListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DeleteLoadBalancerListenerRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DeleteLoadBalancerListenerRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DeleteLoadBalancerListenerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteLoadBalancerListenerRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteLoadBalancerListenerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteLoadBalancerListenerRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *DeleteLoadBalancerListenerRequest
	GetTags() *string
	SetAccessKeyId(v string) *DeleteLoadBalancerListenerRequest
	GetAccessKeyId() *string
}

type DeleteLoadBalancerListenerRequest struct {
	// This parameter is required.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// This parameter is required.
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	AccessKeyId          *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
}

func (s DeleteLoadBalancerListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoadBalancerListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DeleteLoadBalancerListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DeleteLoadBalancerListenerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteLoadBalancerListenerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLoadBalancerListenerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteLoadBalancerListenerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteLoadBalancerListenerRequest) GetTags() *string {
	return s.Tags
}

func (s *DeleteLoadBalancerListenerRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *DeleteLoadBalancerListenerRequest) SetListenerPort(v int32) *DeleteLoadBalancerListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetOwnerAccount(v string) *DeleteLoadBalancerListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetOwnerId(v int64) *DeleteLoadBalancerListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetResourceOwnerAccount(v string) *DeleteLoadBalancerListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetResourceOwnerId(v int64) *DeleteLoadBalancerListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetTags(v string) *DeleteLoadBalancerListenerRequest {
	s.Tags = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetAccessKeyId(v string) *DeleteLoadBalancerListenerRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) Validate() error {
	return dara.Validate(s)
}
