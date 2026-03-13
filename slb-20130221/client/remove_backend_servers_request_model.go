// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v string) *RemoveBackendServersRequest
	GetBackendServers() *string
	SetLoadBalancerId(v string) *RemoveBackendServersRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *RemoveBackendServersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveBackendServersRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RemoveBackendServersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveBackendServersRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *RemoveBackendServersRequest
	GetTags() *string
	SetAccessKeyId(v string) *RemoveBackendServersRequest
	GetAccessKeyId() *string
}

type RemoveBackendServersRequest struct {
	// This parameter is required.
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

func (s RemoveBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackendServersRequest) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersRequest) GetBackendServers() *string {
	return s.BackendServers
}

func (s *RemoveBackendServersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *RemoveBackendServersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveBackendServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveBackendServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveBackendServersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveBackendServersRequest) GetTags() *string {
	return s.Tags
}

func (s *RemoveBackendServersRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *RemoveBackendServersRequest) SetBackendServers(v string) *RemoveBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *RemoveBackendServersRequest) SetLoadBalancerId(v string) *RemoveBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveBackendServersRequest) SetOwnerAccount(v string) *RemoveBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveBackendServersRequest) SetOwnerId(v int64) *RemoveBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveBackendServersRequest) SetResourceOwnerAccount(v string) *RemoveBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveBackendServersRequest) SetResourceOwnerId(v int64) *RemoveBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveBackendServersRequest) SetTags(v string) *RemoveBackendServersRequest {
	s.Tags = &v
	return s
}

func (s *RemoveBackendServersRequest) SetAccessKeyId(v string) *RemoveBackendServersRequest {
	s.AccessKeyId = &v
	return s
}

func (s *RemoveBackendServersRequest) Validate() error {
	return dara.Validate(s)
}
