// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v string) *AddBackendServersRequest
	GetBackendServers() *string
	SetLoadBalancerId(v string) *AddBackendServersRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *AddBackendServersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddBackendServersRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddBackendServersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddBackendServersRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *AddBackendServersRequest
	GetTags() *string
	SetAccessKeyId(v string) *AddBackendServersRequest
	GetAccessKeyId() *string
}

type AddBackendServersRequest struct {
	BackendServers *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	// This parameter is required.
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	AccessKeyId          *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
}

func (s AddBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBackendServersRequest) GoString() string {
	return s.String()
}

func (s *AddBackendServersRequest) GetBackendServers() *string {
	return s.BackendServers
}

func (s *AddBackendServersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *AddBackendServersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddBackendServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddBackendServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddBackendServersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddBackendServersRequest) GetTags() *string {
	return s.Tags
}

func (s *AddBackendServersRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *AddBackendServersRequest) SetBackendServers(v string) *AddBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *AddBackendServersRequest) SetLoadBalancerId(v string) *AddBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AddBackendServersRequest) SetOwnerAccount(v string) *AddBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddBackendServersRequest) SetOwnerId(v int64) *AddBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBackendServersRequest) SetResourceOwnerAccount(v string) *AddBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddBackendServersRequest) SetResourceOwnerId(v int64) *AddBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddBackendServersRequest) SetTags(v string) *AddBackendServersRequest {
	s.Tags = &v
	return s
}

func (s *AddBackendServersRequest) SetAccessKeyId(v string) *AddBackendServersRequest {
	s.AccessKeyId = &v
	return s
}

func (s *AddBackendServersRequest) Validate() error {
	return dara.Validate(s)
}
