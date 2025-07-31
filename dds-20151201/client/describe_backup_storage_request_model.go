// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeBackupStorageRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *DescribeBackupStorageRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeBackupStorageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupStorageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeBackupStorageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeBackupStorageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupStorageRequest
	GetResourceOwnerId() *int64
}

type DescribeBackupStorageRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-2zeb2d64cb46xxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the shard node in the sharded cluster instance.
	//
	// >  This parameter is required only when the **DBInstanceId*	- parameter is set to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-2zee48956b4axxxx
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-bejing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBackupStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupStorageRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupStorageRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeBackupStorageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBackupStorageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackupStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupStorageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupStorageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupStorageRequest) SetDBInstanceId(v string) *DescribeBackupStorageRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupStorageRequest) SetNodeId(v string) *DescribeBackupStorageRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeBackupStorageRequest) SetOwnerAccount(v string) *DescribeBackupStorageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupStorageRequest) SetOwnerId(v int64) *DescribeBackupStorageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupStorageRequest) SetRegionId(v string) *DescribeBackupStorageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupStorageRequest) SetResourceOwnerAccount(v string) *DescribeBackupStorageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupStorageRequest) SetResourceOwnerId(v int64) *DescribeBackupStorageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupStorageRequest) Validate() error {
	return dara.Validate(s)
}
