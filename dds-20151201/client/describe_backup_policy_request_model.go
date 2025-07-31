// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeBackupPolicyRequest
	GetDBInstanceId() *string
	SetInstanceType(v string) *DescribeBackupPolicyRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *DescribeBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeBackupPolicyRequest
	GetSecurityToken() *string
	SetSrcRegion(v string) *DescribeBackupPolicyRequest
	GetSrcRegion() *string
}

type DescribeBackupPolicyRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp16cb162771****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **sharding**: sharded cluster instance
	//
	// 	- **replicate**: replica set or standalone instance
	//
	// example:
	//
	// sharding
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupPolicyRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupPolicyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeBackupPolicyRequest) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeBackupPolicyRequest) SetDBInstanceId(v string) *DescribeBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetInstanceType(v string) *DescribeBackupPolicyRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetSecurityToken(v string) *DescribeBackupPolicyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetSrcRegion(v string) *DescribeBackupPolicyRequest {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
