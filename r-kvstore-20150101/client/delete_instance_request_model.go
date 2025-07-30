// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalInstanceId(v string) *DeleteInstanceRequest
	GetGlobalInstanceId() *string
	SetInstanceId(v string) *DeleteInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeleteInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DeleteInstanceRequest
	GetSecurityToken() *string
}

type DeleteInstanceRequest struct {
	// The ID of the distributed instance to which the instance belongs. This parameter is applicable to only China site (aliyun.com).
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The ID of the instance that you want to release.
	//
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

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetGlobalInstanceId() *string {
	return s.GlobalInstanceId
}

func (s *DeleteInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteInstanceRequest) SetGlobalInstanceId(v string) *DeleteInstanceRequest {
	s.GlobalInstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetOwnerAccount(v string) *DeleteInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteInstanceRequest) SetOwnerId(v int64) *DeleteInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerAccount(v string) *DeleteInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerId(v int64) *DeleteInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteInstanceRequest) SetSecurityToken(v string) *DeleteInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
