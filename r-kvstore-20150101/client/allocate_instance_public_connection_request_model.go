// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateInstancePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetInstanceId(v string) *AllocateInstancePublicConnectionRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetOwnerId() *int64
	SetPort(v string) *AllocateInstancePublicConnectionRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *AllocateInstancePublicConnectionRequest
	GetSecurityToken() *string
}

type AllocateInstancePublicConnectionRequest struct {
	// The prefix of the public endpoint. The prefix must start with a lowercase letter and can contain lowercase letters and digits. The prefix can be 8 to 40 characters in length.
	//
	// >  The endpoint is in the `<prefix>.redis.rds.aliyuncs.com` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The service port number of the instance. Valid values: **1024*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6379
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateInstancePublicConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AllocateInstancePublicConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateInstancePublicConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateInstancePublicConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateInstancePublicConnectionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetSecurityToken(v string) *AllocateInstancePublicConnectionRequest {
	s.SecurityToken = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
