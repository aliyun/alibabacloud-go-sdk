// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptionKey(v string) *DescribeEncryptionKeyRequest
	GetEncryptionKey() *string
	SetInstanceId(v string) *DescribeEncryptionKeyRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeEncryptionKeyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEncryptionKeyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeEncryptionKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEncryptionKeyRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeEncryptionKeyRequest
	GetSecurityToken() *string
}

type DescribeEncryptionKeyRequest struct {
	// The ID of the custom key. You can call the [DescribeEncryptionKeyList](https://help.aliyun.com/document_detail/473860.html) operation to query the ID of the key.
	//
	// example:
	//
	// ad463061-992d-4195-8a94-ed63********
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of the instance.
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

func (s DescribeEncryptionKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionKeyRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeEncryptionKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEncryptionKeyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEncryptionKeyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEncryptionKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEncryptionKeyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEncryptionKeyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeEncryptionKeyRequest) SetEncryptionKey(v string) *DescribeEncryptionKeyRequest {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeEncryptionKeyRequest) SetInstanceId(v string) *DescribeEncryptionKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEncryptionKeyRequest) SetOwnerAccount(v string) *DescribeEncryptionKeyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEncryptionKeyRequest) SetOwnerId(v int64) *DescribeEncryptionKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEncryptionKeyRequest) SetResourceOwnerAccount(v string) *DescribeEncryptionKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEncryptionKeyRequest) SetResourceOwnerId(v int64) *DescribeEncryptionKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEncryptionKeyRequest) SetSecurityToken(v string) *DescribeEncryptionKeyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeEncryptionKeyRequest) Validate() error {
	return dara.Validate(s)
}
