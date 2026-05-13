// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairSkvDdbTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTairSkvDdbTableRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeTairSkvDdbTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTairSkvDdbTableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeTairSkvDdbTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTairSkvDdbTableRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeTairSkvDdbTableRequest
	GetSecurityToken() *string
}

type DescribeTairSkvDdbTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tt-bp1zxszhcgatnx\\*\\*\\*\\*
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeTairSkvDdbTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairSkvDdbTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTairSkvDdbTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTairSkvDdbTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTairSkvDdbTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTairSkvDdbTableRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairSkvDdbTableRequest) SetInstanceId(v string) *DescribeTairSkvDdbTableRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairSkvDdbTableRequest) SetOwnerAccount(v string) *DescribeTairSkvDdbTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTairSkvDdbTableRequest) SetOwnerId(v int64) *DescribeTairSkvDdbTableRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTairSkvDdbTableRequest) SetResourceOwnerAccount(v string) *DescribeTairSkvDdbTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTairSkvDdbTableRequest) SetResourceOwnerId(v int64) *DescribeTairSkvDdbTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTairSkvDdbTableRequest) SetSecurityToken(v string) *DescribeTairSkvDdbTableRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairSkvDdbTableRequest) Validate() error {
	return dara.Validate(s)
}
