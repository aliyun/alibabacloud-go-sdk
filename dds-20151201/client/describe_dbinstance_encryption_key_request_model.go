// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceEncryptionKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceEncryptionKeyRequest
	GetDBInstanceId() *string
	SetEncryptionKey(v string) *DescribeDBInstanceEncryptionKeyRequest
	GetEncryptionKey() *string
	SetOwnerAccount(v string) *DescribeDBInstanceEncryptionKeyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceEncryptionKeyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceEncryptionKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceEncryptionKeyRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceEncryptionKeyRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp2235****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The custom key for the instance. You can call the [DescribeUserEncryptionKeyList](https://help.aliyun.com/document_detail/151729.html) operation to query the list of custom keys for an ApsaraDB for MongoDB instance.
	//
	// example:
	//
	// 2axxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	EncryptionKey        *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceEncryptionKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEncryptionKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEncryptionKeyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceEncryptionKeyRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBInstanceEncryptionKeyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceEncryptionKeyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceEncryptionKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceEncryptionKeyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetDBInstanceId(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetEncryptionKey(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetOwnerAccount(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetOwnerId(v int64) *DescribeDBInstanceEncryptionKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceEncryptionKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) Validate() error {
	return dara.Validate(s)
}
