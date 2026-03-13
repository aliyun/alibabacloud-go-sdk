// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *SetLoadBalancerStatusRequest
	GetLoadBalancerId() *string
	SetLoadBalancerStatus(v string) *SetLoadBalancerStatusRequest
	GetLoadBalancerStatus() *string
	SetOwnerAccount(v string) *SetLoadBalancerStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetLoadBalancerStatusRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SetLoadBalancerStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetLoadBalancerStatusRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *SetLoadBalancerStatusRequest
	GetTags() *string
	SetAccessKeyId(v string) *SetLoadBalancerStatusRequest
	GetAccessKeyId() *string
}

type SetLoadBalancerStatusRequest struct {
	// This parameter is required.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// This parameter is required.
	LoadBalancerStatus   *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	AccessKeyId          *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
}

func (s SetLoadBalancerStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerStatusRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerStatusRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerStatusRequest) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *SetLoadBalancerStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLoadBalancerStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetLoadBalancerStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetLoadBalancerStatusRequest) GetTags() *string {
	return s.Tags
}

func (s *SetLoadBalancerStatusRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *SetLoadBalancerStatusRequest) SetLoadBalancerId(v string) *SetLoadBalancerStatusRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetLoadBalancerStatus(v string) *SetLoadBalancerStatusRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetOwnerAccount(v string) *SetLoadBalancerStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetOwnerId(v int64) *SetLoadBalancerStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetResourceOwnerId(v int64) *SetLoadBalancerStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetTags(v string) *SetLoadBalancerStatusRequest {
	s.Tags = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetAccessKeyId(v string) *SetLoadBalancerStatusRequest {
	s.AccessKeyId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) Validate() error {
	return dara.Validate(s)
}
