// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetLindormInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormInstanceRequest
	GetSecurityToken() *string
}

type GetLindormInstanceRequest struct {
	// The instance ID. Call [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormInstanceRequest) SetInstanceId(v string) *GetLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceRequest) SetOwnerAccount(v string) *GetLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormInstanceRequest) SetOwnerId(v int64) *GetLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormInstanceRequest) SetResourceOwnerAccount(v string) *GetLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormInstanceRequest) SetResourceOwnerId(v int64) *GetLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormInstanceRequest) SetSecurityToken(v string) *GetLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormInstanceRequest) Validate() error {
	return dara.Validate(s)
}
