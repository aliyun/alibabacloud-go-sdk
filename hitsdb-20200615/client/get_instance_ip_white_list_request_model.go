// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIpWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceIpWhiteListRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetInstanceIpWhiteListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetInstanceIpWhiteListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetInstanceIpWhiteListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetInstanceIpWhiteListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetInstanceIpWhiteListRequest
	GetSecurityToken() *string
}

type GetInstanceIpWhiteListRequest struct {
	// The ID of the instance whose whitelists you want to query. You can call the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426068.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetInstanceIpWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceIpWhiteListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetInstanceIpWhiteListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetInstanceIpWhiteListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetInstanceIpWhiteListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetInstanceIpWhiteListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetInstanceIpWhiteListRequest) SetInstanceId(v string) *GetInstanceIpWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetOwnerAccount(v string) *GetInstanceIpWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetOwnerId(v int64) *GetInstanceIpWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetResourceOwnerAccount(v string) *GetInstanceIpWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetResourceOwnerId(v int64) *GetInstanceIpWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetSecurityToken(v string) *GetInstanceIpWhiteListRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
