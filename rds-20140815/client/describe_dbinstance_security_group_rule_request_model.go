// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceSecurityGroupRuleRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceSecurityGroupRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *DescribeDBInstanceSecurityGroupRuleRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstanceSecurityGroupRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceSecurityGroupRuleRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceSecurityGroupRuleRequest struct {
	// The ID of the instance. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/2628785.html) operation to query the IDs of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze202******
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceSecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) SetDBInstanceId(v string) *DescribeDBInstanceSecurityGroupRuleRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) SetOwnerAccount(v string) *DescribeDBInstanceSecurityGroupRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) SetOwnerId(v string) *DescribeDBInstanceSecurityGroupRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceSecurityGroupRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceSecurityGroupRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}
