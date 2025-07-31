// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest
	GetDBInstanceId() *string
	SetEngine(v string) *DescribeDBInstanceAttributeRequest
	GetEngine() *string
	SetIsDelete(v bool) *DescribeDBInstanceAttributeRequest
	GetIsDelete() *bool
	SetOwnerAccount(v string) *DescribeDBInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceAttributeRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeDBInstanceAttributeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDBInstanceAttributeRequest
	GetSecurityToken() *string
}

type DescribeDBInstanceAttributeRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp11483712c1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database engine. Set the value to **MongoDB**.
	//
	// example:
	//
	// MongoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Specifies whether to query instances that are deleted. Valid values:
	//
	// 	- **false**: queries instances that are running.
	//
	// 	- **true**: queries instance that are deleted.
	//
	// example:
	//
	// false
	IsDelete     *bool   `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group. For more information, see [View the basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// > This parameter is available only if you use the China site (aliyun.com).
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstanceAttributeRequest) GetIsDelete() *bool {
	return s.IsDelete
}

func (s *DescribeDBInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceAttributeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetEngine(v string) *DescribeDBInstanceAttributeRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetIsDelete(v bool) *DescribeDBInstanceAttributeRequest {
	s.IsDelete = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetOwnerAccount(v string) *DescribeDBInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceGroupId(v string) *DescribeDBInstanceAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetSecurityToken(v string) *DescribeDBInstanceAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
