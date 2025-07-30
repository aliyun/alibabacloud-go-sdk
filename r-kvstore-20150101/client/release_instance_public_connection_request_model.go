// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstancePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest
	GetCurrentConnectionString() *string
	SetInstanceId(v string) *ReleaseInstancePublicConnectionRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ReleaseInstancePublicConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseInstancePublicConnectionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReleaseInstancePublicConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseInstancePublicConnectionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ReleaseInstancePublicConnectionRequest
	GetSecurityToken() *string
}

type ReleaseInstancePublicConnectionRequest struct {
	// The public endpoint to be released.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****.redis.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The ID of the instance.
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

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ReleaseInstancePublicConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseInstancePublicConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseInstancePublicConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseInstancePublicConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseInstancePublicConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseInstancePublicConnectionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetInstanceId(v string) *ReleaseInstancePublicConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetOwnerAccount(v string) *ReleaseInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetOwnerId(v int64) *ReleaseInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *ReleaseInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *ReleaseInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetSecurityToken(v string) *ReleaseInstancePublicConnectionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
