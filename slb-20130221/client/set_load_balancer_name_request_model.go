// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *SetLoadBalancerNameRequest
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *SetLoadBalancerNameRequest
	GetLoadBalancerName() *string
	SetOwnerAccount(v string) *SetLoadBalancerNameRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetLoadBalancerNameRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SetLoadBalancerNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetLoadBalancerNameRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *SetLoadBalancerNameRequest
	GetTags() *string
	SetAccessKeyId(v string) *SetLoadBalancerNameRequest
	GetAccessKeyId() *string
}

type SetLoadBalancerNameRequest struct {
	// This parameter is required.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// This parameter is required.
	LoadBalancerName     *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	AccessKeyId          *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
}

func (s SetLoadBalancerNameRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerNameRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerNameRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerNameRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *SetLoadBalancerNameRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLoadBalancerNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetLoadBalancerNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetLoadBalancerNameRequest) GetTags() *string {
	return s.Tags
}

func (s *SetLoadBalancerNameRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *SetLoadBalancerNameRequest) SetLoadBalancerId(v string) *SetLoadBalancerNameRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetLoadBalancerName(v string) *SetLoadBalancerNameRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetOwnerAccount(v string) *SetLoadBalancerNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetOwnerId(v int64) *SetLoadBalancerNameRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetResourceOwnerId(v int64) *SetLoadBalancerNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetTags(v string) *SetLoadBalancerNameRequest {
	s.Tags = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetAccessKeyId(v string) *SetLoadBalancerNameRequest {
	s.AccessKeyId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) Validate() error {
	return dara.Validate(s)
}
