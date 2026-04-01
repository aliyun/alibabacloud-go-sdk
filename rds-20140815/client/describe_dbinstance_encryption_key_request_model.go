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
	SetRegionId(v string) *DescribeDBInstanceEncryptionKeyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstanceEncryptionKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceEncryptionKeyRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDBInstanceEncryptionKeyRequest
	GetSecurityToken() *string
	SetTargetRegionId(v string) *DescribeDBInstanceEncryptionKeyRequest
	GetTargetRegionId() *string
}

type DescribeDBInstanceEncryptionKeyRequest struct {
	// The ID of the instance You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the IDs of instances.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the custom key.
	//
	// example:
	//
	// 749c1df7-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the destination region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-qingdao
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
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

func (s *DescribeDBInstanceEncryptionKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceEncryptionKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceEncryptionKeyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceEncryptionKeyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDBInstanceEncryptionKeyRequest) GetTargetRegionId() *string {
	return s.TargetRegionId
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

func (s *DescribeDBInstanceEncryptionKeyRequest) SetRegionId(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.RegionId = &v
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

func (s *DescribeDBInstanceEncryptionKeyRequest) SetSecurityToken(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) SetTargetRegionId(v string) *DescribeDBInstanceEncryptionKeyRequest {
	s.TargetRegionId = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyRequest) Validate() error {
	return dara.Validate(s)
}
