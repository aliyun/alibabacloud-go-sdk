// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeClusterBackupsRequest
	GetBackupId() *string
	SetDBInstanceId(v string) *DescribeClusterBackupsRequest
	GetDBInstanceId() *string
	SetDestRegion(v string) *DescribeClusterBackupsRequest
	GetDestRegion() *string
	SetEndTime(v string) *DescribeClusterBackupsRequest
	GetEndTime() *string
	SetIsOnlyGetClusterBackUp(v bool) *DescribeClusterBackupsRequest
	GetIsOnlyGetClusterBackUp() *bool
	SetOwnerAccount(v string) *DescribeClusterBackupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeClusterBackupsRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *DescribeClusterBackupsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeClusterBackupsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeClusterBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClusterBackupsRequest
	GetResourceOwnerId() *int64
	SetSrcRegion(v string) *DescribeClusterBackupsRequest
	GetSrcRegion() *string
	SetStartTime(v string) *DescribeClusterBackupsRequest
	GetStartTime() *string
}

type DescribeClusterBackupsRequest struct {
	// The ID of the cluster backup set.
	//
	// example:
	//
	// 5664****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp16cb162771****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region where cross-region backups reside.
	//
	// >  This parameter is required if you want to query cross-region backups.
	//
	// example:
	//
	// cn-shanghai
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. The end time must be later than the start time.
	//
	// example:
	//
	// 2019-03-14T13:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to query information about child nodes in the cluster backup. Valid values:
	//
	// 	- **true**: The system returns only the basic information of the cluster backup.
	//
	// 	- **false*	- (default): The system returns the backup information of all child nodes.
	//
	// example:
	//
	// true
	IsOnlyGetClusterBackUp *bool   `json:"IsOnlyGetClusterBackUp,omitempty" xml:"IsOnlyGetClusterBackUp,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**. The page number must be a positive integer.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Valid values:
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
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The region ID of the instance.
	//
	// >
	//
	// 	- This parameter is required if you want to query the backup sets of a released instance.
	//
	// 	- This parameter is required if you want to query cross-region backups.
	//
	// example:
	//
	// cn-hangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2019-03-13T12:11:14Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeClusterBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupsRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeClusterBackupsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeClusterBackupsRequest) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeClusterBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeClusterBackupsRequest) GetIsOnlyGetClusterBackUp() *bool {
	return s.IsOnlyGetClusterBackUp
}

func (s *DescribeClusterBackupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeClusterBackupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClusterBackupsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeClusterBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterBackupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClusterBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterBackupsRequest) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeClusterBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeClusterBackupsRequest) SetBackupId(v string) *DescribeClusterBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetDBInstanceId(v string) *DescribeClusterBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetDestRegion(v string) *DescribeClusterBackupsRequest {
	s.DestRegion = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetEndTime(v string) *DescribeClusterBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetIsOnlyGetClusterBackUp(v bool) *DescribeClusterBackupsRequest {
	s.IsOnlyGetClusterBackUp = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetOwnerAccount(v string) *DescribeClusterBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetOwnerId(v int64) *DescribeClusterBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetPageNo(v int32) *DescribeClusterBackupsRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetPageSize(v int32) *DescribeClusterBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetResourceOwnerAccount(v string) *DescribeClusterBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetResourceOwnerId(v int64) *DescribeClusterBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetSrcRegion(v string) *DescribeClusterBackupsRequest {
	s.SrcRegion = &v
	return s
}

func (s *DescribeClusterBackupsRequest) SetStartTime(v string) *DescribeClusterBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeClusterBackupsRequest) Validate() error {
	return dara.Validate(s)
}
