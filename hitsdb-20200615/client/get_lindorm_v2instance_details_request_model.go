// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormV2InstanceDetailsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormV2InstanceDetailsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormV2InstanceDetailsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetLindormV2InstanceDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormV2InstanceDetailsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormV2InstanceDetailsRequest
	GetSecurityToken() *string
}

type GetLindormV2InstanceDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ts-xxxxxxxxxxxxxxxxx
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormV2InstanceDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceDetailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2InstanceDetailsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormV2InstanceDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormV2InstanceDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormV2InstanceDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormV2InstanceDetailsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormV2InstanceDetailsRequest) SetInstanceId(v string) *GetLindormV2InstanceDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsRequest) SetOwnerAccount(v string) *GetLindormV2InstanceDetailsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceDetailsRequest) SetOwnerId(v int64) *GetLindormV2InstanceDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsRequest) SetResourceOwnerAccount(v string) *GetLindormV2InstanceDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceDetailsRequest) SetResourceOwnerId(v int64) *GetLindormV2InstanceDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsRequest) SetSecurityToken(v string) *GetLindormV2InstanceDetailsRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormV2InstanceDetailsRequest) Validate() error {
	return dara.Validate(s)
}
