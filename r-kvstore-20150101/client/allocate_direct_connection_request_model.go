// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateDirectConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *AllocateDirectConnectionRequest
	GetConnectionString() *string
	SetInstanceId(v string) *AllocateDirectConnectionRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *AllocateDirectConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateDirectConnectionRequest
	GetOwnerId() *int64
	SetPort(v string) *AllocateDirectConnectionRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *AllocateDirectConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateDirectConnectionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *AllocateDirectConnectionRequest
	GetSecurityToken() *string
}

type AllocateDirectConnectionRequest struct {
	// The prefix of the private endpoint. The prefix must start with a lowercase letter and can contain lowercase letters and digits. The prefix must be 8 to 40 characters in length.
	//
	// example:
	//
	// redisdirect123
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
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
	// The port number of the instance. Valid values: **1024*	- to **65535**. Default value: **6379**.
	//
	// example:
	//
	// 6379
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AllocateDirectConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateDirectConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateDirectConnectionRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *AllocateDirectConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AllocateDirectConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateDirectConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateDirectConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateDirectConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateDirectConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateDirectConnectionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AllocateDirectConnectionRequest) SetConnectionString(v string) *AllocateDirectConnectionRequest {
	s.ConnectionString = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetInstanceId(v string) *AllocateDirectConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetOwnerAccount(v string) *AllocateDirectConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetOwnerId(v int64) *AllocateDirectConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetPort(v string) *AllocateDirectConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetResourceOwnerAccount(v string) *AllocateDirectConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetResourceOwnerId(v int64) *AllocateDirectConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateDirectConnectionRequest) SetSecurityToken(v string) *AllocateDirectConnectionRequest {
	s.SecurityToken = &v
	return s
}

func (s *AllocateDirectConnectionRequest) Validate() error {
	return dara.Validate(s)
}
