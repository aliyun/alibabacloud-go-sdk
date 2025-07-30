// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartTairKVCacheCustomInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RestartTairKVCacheCustomInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RestartTairKVCacheCustomInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartTairKVCacheCustomInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestartTairKVCacheCustomInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartTairKVCacheCustomInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *RestartTairKVCacheCustomInstanceRequest
	GetSecurityToken() *string
}

type RestartTairKVCacheCustomInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RestartTairKVCacheCustomInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartTairKVCacheCustomInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartTairKVCacheCustomInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartTairKVCacheCustomInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartTairKVCacheCustomInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartTairKVCacheCustomInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartTairKVCacheCustomInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartTairKVCacheCustomInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RestartTairKVCacheCustomInstanceRequest) SetInstanceId(v string) *RestartTairKVCacheCustomInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartTairKVCacheCustomInstanceRequest) SetOwnerAccount(v string) *RestartTairKVCacheCustomInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartTairKVCacheCustomInstanceRequest) SetOwnerId(v int64) *RestartTairKVCacheCustomInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartTairKVCacheCustomInstanceRequest) SetResourceOwnerAccount(v string) *RestartTairKVCacheCustomInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartTairKVCacheCustomInstanceRequest) SetResourceOwnerId(v int64) *RestartTairKVCacheCustomInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartTairKVCacheCustomInstanceRequest) SetSecurityToken(v string) *RestartTairKVCacheCustomInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestartTairKVCacheCustomInstanceRequest) Validate() error {
	return dara.Validate(s)
}
