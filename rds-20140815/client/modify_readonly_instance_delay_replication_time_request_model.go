// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReadonlyInstanceDelayReplicationTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyReadonlyInstanceDelayReplicationTimeRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ModifyReadonlyInstanceDelayReplicationTimeRequest
	GetOwnerId() *int64
	SetReadSQLReplicationTime(v string) *ModifyReadonlyInstanceDelayReplicationTimeRequest
	GetReadSQLReplicationTime() *string
	SetResourceGroupId(v string) *ModifyReadonlyInstanceDelayReplicationTimeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyReadonlyInstanceDelayReplicationTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyReadonlyInstanceDelayReplicationTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyReadonlyInstanceDelayReplicationTimeRequest struct {
	// The ID of the read-only instance. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rr-bpxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The replication latency of the data replication. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ReadSQLReplicationTime *string `json:"ReadSQLReplicationTime,omitempty" xml:"ReadSQLReplicationTime,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyReadonlyInstanceDelayReplicationTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyReadonlyInstanceDelayReplicationTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) GetReadSQLReplicationTime() *string {
	return s.ReadSQLReplicationTime
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) SetDBInstanceId(v string) *ModifyReadonlyInstanceDelayReplicationTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) SetOwnerId(v int64) *ModifyReadonlyInstanceDelayReplicationTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) SetReadSQLReplicationTime(v string) *ModifyReadonlyInstanceDelayReplicationTimeRequest {
	s.ReadSQLReplicationTime = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) SetResourceGroupId(v string) *ModifyReadonlyInstanceDelayReplicationTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) SetResourceOwnerAccount(v string) *ModifyReadonlyInstanceDelayReplicationTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) SetResourceOwnerId(v int64) *ModifyReadonlyInstanceDelayReplicationTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyReadonlyInstanceDelayReplicationTimeRequest) Validate() error {
	return dara.Validate(s)
}
