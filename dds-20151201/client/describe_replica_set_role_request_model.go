// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicaSetRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeReplicaSetRoleRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeReplicaSetRoleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeReplicaSetRoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeReplicaSetRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeReplicaSetRoleRequest
	GetResourceOwnerId() *int64
}

type DescribeReplicaSetRoleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeReplicaSetRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicaSetRoleRequest) GoString() string {
	return s.String()
}

func (s *DescribeReplicaSetRoleRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeReplicaSetRoleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeReplicaSetRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeReplicaSetRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeReplicaSetRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeReplicaSetRoleRequest) SetDBInstanceId(v string) *DescribeReplicaSetRoleRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeReplicaSetRoleRequest) SetOwnerAccount(v string) *DescribeReplicaSetRoleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeReplicaSetRoleRequest) SetOwnerId(v int64) *DescribeReplicaSetRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeReplicaSetRoleRequest) SetResourceOwnerAccount(v string) *DescribeReplicaSetRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeReplicaSetRoleRequest) SetResourceOwnerId(v int64) *DescribeReplicaSetRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeReplicaSetRoleRequest) Validate() error {
	return dara.Validate(s)
}
