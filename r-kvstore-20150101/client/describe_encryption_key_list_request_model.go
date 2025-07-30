// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionKeyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeEncryptionKeyListRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeEncryptionKeyListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEncryptionKeyListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeEncryptionKeyListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEncryptionKeyListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeEncryptionKeyListRequest
	GetSecurityToken() *string
}

type DescribeEncryptionKeyListRequest struct {
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

func (s DescribeEncryptionKeyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionKeyListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEncryptionKeyListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEncryptionKeyListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEncryptionKeyListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEncryptionKeyListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEncryptionKeyListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeEncryptionKeyListRequest) SetInstanceId(v string) *DescribeEncryptionKeyListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEncryptionKeyListRequest) SetOwnerAccount(v string) *DescribeEncryptionKeyListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEncryptionKeyListRequest) SetOwnerId(v int64) *DescribeEncryptionKeyListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEncryptionKeyListRequest) SetResourceOwnerAccount(v string) *DescribeEncryptionKeyListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEncryptionKeyListRequest) SetResourceOwnerId(v int64) *DescribeEncryptionKeyListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEncryptionKeyListRequest) SetSecurityToken(v string) *DescribeEncryptionKeyListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeEncryptionKeyListRequest) Validate() error {
	return dara.Validate(s)
}
