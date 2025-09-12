// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2StorageUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormV2StorageUsageRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormV2StorageUsageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormV2StorageUsageRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetLindormV2StorageUsageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormV2StorageUsageRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormV2StorageUsageRequest
	GetSecurityToken() *string
}

type GetLindormV2StorageUsageRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormV2StorageUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2StorageUsageRequest) GoString() string {
	return s.String()
}

func (s *GetLindormV2StorageUsageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2StorageUsageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormV2StorageUsageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormV2StorageUsageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormV2StorageUsageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormV2StorageUsageRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormV2StorageUsageRequest) SetInstanceId(v string) *GetLindormV2StorageUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) SetOwnerAccount(v string) *GetLindormV2StorageUsageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) SetOwnerId(v int64) *GetLindormV2StorageUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) SetResourceOwnerAccount(v string) *GetLindormV2StorageUsageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) SetResourceOwnerId(v int64) *GetLindormV2StorageUsageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) SetSecurityToken(v string) *GetLindormV2StorageUsageRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) Validate() error {
	return dara.Validate(s)
}
