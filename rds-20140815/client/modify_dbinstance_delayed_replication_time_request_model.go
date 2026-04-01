// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDelayedReplicationTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceDelayedReplicationTimeRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ModifyDBInstanceDelayedReplicationTimeRequest
	GetOwnerId() *int64
	SetReadSQLReplicationTime(v string) *ModifyDBInstanceDelayedReplicationTimeRequest
	GetReadSQLReplicationTime() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceDelayedReplicationTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceDelayedReplicationTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceDelayedReplicationTimeRequest struct {
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The replication latency of the read-only instance. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ReadSQLReplicationTime *string `json:"ReadSQLReplicationTime,omitempty" xml:"ReadSQLReplicationTime,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceDelayedReplicationTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDelayedReplicationTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) GetReadSQLReplicationTime() *string {
	return s.ReadSQLReplicationTime
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceDelayedReplicationTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) SetOwnerId(v int64) *ModifyDBInstanceDelayedReplicationTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) SetReadSQLReplicationTime(v string) *ModifyDBInstanceDelayedReplicationTimeRequest {
	s.ReadSQLReplicationTime = &v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceDelayedReplicationTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceDelayedReplicationTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeRequest) Validate() error {
	return dara.Validate(s)
}
