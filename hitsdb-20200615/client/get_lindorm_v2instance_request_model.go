// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormV2InstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormV2InstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormV2InstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetLindormV2InstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormV2InstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormV2InstanceRequest
	GetSecurityToken() *string
}

type GetLindormV2InstanceRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormV2InstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceRequest) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2InstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormV2InstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormV2InstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormV2InstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormV2InstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormV2InstanceRequest) SetInstanceId(v string) *GetLindormV2InstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceRequest) SetOwnerAccount(v string) *GetLindormV2InstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceRequest) SetOwnerId(v int64) *GetLindormV2InstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormV2InstanceRequest) SetResourceOwnerAccount(v string) *GetLindormV2InstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceRequest) SetResourceOwnerId(v int64) *GetLindormV2InstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormV2InstanceRequest) SetSecurityToken(v string) *GetLindormV2InstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormV2InstanceRequest) Validate() error {
	return dara.Validate(s)
}
