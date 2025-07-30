// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTairKVCacheCustomInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopTairKVCacheCustomInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *StopTairKVCacheCustomInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StopTairKVCacheCustomInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StopTairKVCacheCustomInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopTairKVCacheCustomInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *StopTairKVCacheCustomInstanceRequest
	GetSecurityToken() *string
}

type StopTairKVCacheCustomInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StopTairKVCacheCustomInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTairKVCacheCustomInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopTairKVCacheCustomInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopTairKVCacheCustomInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StopTairKVCacheCustomInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopTairKVCacheCustomInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopTairKVCacheCustomInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopTairKVCacheCustomInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StopTairKVCacheCustomInstanceRequest) SetInstanceId(v string) *StopTairKVCacheCustomInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopTairKVCacheCustomInstanceRequest) SetOwnerAccount(v string) *StopTairKVCacheCustomInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StopTairKVCacheCustomInstanceRequest) SetOwnerId(v int64) *StopTairKVCacheCustomInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *StopTairKVCacheCustomInstanceRequest) SetResourceOwnerAccount(v string) *StopTairKVCacheCustomInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopTairKVCacheCustomInstanceRequest) SetResourceOwnerId(v int64) *StopTairKVCacheCustomInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopTairKVCacheCustomInstanceRequest) SetSecurityToken(v string) *StopTairKVCacheCustomInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *StopTairKVCacheCustomInstanceRequest) Validate() error {
	return dara.Validate(s)
}
