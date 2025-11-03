// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeBackupsRequest
	GetBackupId() *string
	SetBackupJobId(v string) *DescribeBackupsRequest
	GetBackupJobId() *string
	SetDBInstanceId(v string) *DescribeBackupsRequest
	GetDBInstanceId() *string
	SetDestRegion(v string) *DescribeBackupsRequest
	GetDestRegion() *string
	SetEndTime(v string) *DescribeBackupsRequest
	GetEndTime() *string
	SetNodeId(v string) *DescribeBackupsRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeBackupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeBackupsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupsRequest
	GetResourceOwnerId() *int64
	SetSrcRegion(v string) *DescribeBackupsRequest
	GetSrcRegion() *string
	SetStartTime(v string) *DescribeBackupsRequest
	GetStartTime() *string
}

type DescribeBackupsRequest struct {
	// The ID of the backup set. You can call the [CreateBackup](https://help.aliyun.com/document_detail/62171.html) operation to query the backup set ID.
	//
	// If you set the **DBInstanceId*	- parameter to the ID of a sharded cluster instance, the number of backup IDs is the same as the number of shard nodes. Multiple backup IDs are separated with commas (,).
	//
	// example:
	//
	// 2072****,2072****,2072****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// 775051
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The instance ID.
	//
	// > If you set this parameter to the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1a7009eb24****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID of the Cross-regional backup.
	//
	// >  This parameter is required for the Cross-regional backup.
	//
	// example:
	//
	// cn-hangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// example:
	//
	// 2022-01-14T13:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the shard node in the sharded cluster instance.
	//
	// > This parameter takes effect only when you set the **DBInstanceId*	- parameter to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bp128a003436****
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The region ID of the instance.
	//
	// >- This parameter is required if you want to query the backup sets of a released instance.
	//
	// >-  This parameter is required if you want to query cross-region backups.
	//
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2022-01-13T13:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupsRequest) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *DescribeBackupsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupsRequest) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBackupsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeBackupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBackupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeBackupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupsRequest) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupJobId(v string) *DescribeBackupsRequest {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupsRequest) SetDBInstanceId(v string) *DescribeBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupsRequest) SetDestRegion(v string) *DescribeBackupsRequest {
	s.DestRegion = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetNodeId(v string) *DescribeBackupsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerAccount(v string) *DescribeBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerId(v int64) *DescribeBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v int32) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v int32) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceGroupId(v string) *DescribeBackupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerAccount(v string) *DescribeBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerId(v int64) *DescribeBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetSrcRegion(v string) *DescribeBackupsRequest {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) Validate() error {
	return dara.Validate(s)
}
