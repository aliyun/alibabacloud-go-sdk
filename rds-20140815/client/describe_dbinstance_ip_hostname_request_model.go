// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceIpHostnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceIpHostnameRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceIpHostnameRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceIpHostnameRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBInstanceIpHostnameRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstanceIpHostnameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceIpHostnameRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDBInstanceIpHostnameRequest
	GetSecurityToken() *string
}

type DescribeDBInstanceIpHostnameRequest struct {
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/2628785.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/2628783.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDBInstanceIpHostnameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIpHostnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIpHostnameRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceIpHostnameRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceIpHostnameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceIpHostnameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceIpHostnameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceIpHostnameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceIpHostnameRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDBInstanceIpHostnameRequest) SetDBInstanceId(v string) *DescribeDBInstanceIpHostnameRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameRequest) SetOwnerAccount(v string) *DescribeDBInstanceIpHostnameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameRequest) SetOwnerId(v int64) *DescribeDBInstanceIpHostnameRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameRequest) SetRegionId(v string) *DescribeDBInstanceIpHostnameRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceIpHostnameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceIpHostnameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameRequest) SetSecurityToken(v string) *DescribeDBInstanceIpHostnameRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameRequest) Validate() error {
	return dara.Validate(s)
}
