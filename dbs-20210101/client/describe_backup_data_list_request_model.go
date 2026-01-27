// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDataListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeBackupDataListRequest
	GetBackupId() *string
	SetBackupMethod(v string) *DescribeBackupDataListRequest
	GetBackupMethod() *string
	SetBackupMode(v string) *DescribeBackupDataListRequest
	GetBackupMode() *string
	SetBackupScale(v string) *DescribeBackupDataListRequest
	GetBackupScale() *string
	SetBackupStatus(v string) *DescribeBackupDataListRequest
	GetBackupStatus() *string
	SetBackupType(v string) *DescribeBackupDataListRequest
	GetBackupType() *string
	SetDataSourceId(v string) *DescribeBackupDataListRequest
	GetDataSourceId() *string
	SetEndTime(v string) *DescribeBackupDataListRequest
	GetEndTime() *string
	SetInstanceIsDeleted(v bool) *DescribeBackupDataListRequest
	GetInstanceIsDeleted() *bool
	SetInstanceName(v string) *DescribeBackupDataListRequest
	GetInstanceName() *string
	SetInstanceRegion(v string) *DescribeBackupDataListRequest
	GetInstanceRegion() *string
	SetPageNumber(v int32) *DescribeBackupDataListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupDataListRequest
	GetPageSize() *int32
	SetRegionCode(v string) *DescribeBackupDataListRequest
	GetRegionCode() *string
	SetSceneType(v string) *DescribeBackupDataListRequest
	GetSceneType() *string
	SetStartTime(v string) *DescribeBackupDataListRequest
	GetStartTime() *string
}

type DescribeBackupDataListRequest struct {
	// The ID of the backup set.
	//
	// example:
	//
	// 213064****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup method. Valid values:
	//
	// 	- **Physical**
	//
	// 	- **Logical**
	//
	// 	- **Snapshot**
	//
	// example:
	//
	// Snapshot
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The backup scale. Valid values:
	//
	// 	- **DBInstance**
	//
	// 	- **DBTable**
	//
	// example:
	//
	// DBInstance
	BackupScale *string `json:"BackupScale,omitempty" xml:"BackupScale,omitempty"`
	// The status of the backup set. Valid values:
	//
	// 	- **OK**: The backup succeeded.
	//
	// 	- **ERROR**: The backup failed.
	//
	// example:
	//
	// OK
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The backup type. Valid values:
	//
	// 	- **FullBackup**
	//
	// 	- **IncrementBackup**
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This is a reserved parameter.
	//
	// example:
	//
	// test****
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The time follows the yyyy-MM-ddTHH:mm:ssZ format and is displayed in UTC.
	//
	// example:
	//
	// 2024-04-17T17:00:32Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether the cluster is deleted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	InstanceIsDeleted *bool `json:"InstanceIsDeleted,omitempty" xml:"InstanceIsDeleted,omitempty"`
	// The ID of the PolarDB for MySQL cluster.
	//
	// example:
	//
	// pc-2ze3nrr64c5******
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region in which the original cluster resides.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegion *string `json:"InstanceRegion,omitempty" xml:"InstanceRegion,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the backup set resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The type of the backup scenario. Set the value to **LEVEL_1**, which indicates the level-1 backup of the region in which the PolarDB for MySQL cluster resides.
	//
	// example:
	//
	// LEVEL_1
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2024-04-17T17:00:16Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupDataListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDataListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupDataListRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupDataListRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupDataListRequest) GetBackupScale() *string {
	return s.BackupScale
}

func (s *DescribeBackupDataListRequest) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupDataListRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupDataListRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeBackupDataListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBackupDataListRequest) GetInstanceIsDeleted() *bool {
	return s.InstanceIsDeleted
}

func (s *DescribeBackupDataListRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeBackupDataListRequest) GetInstanceRegion() *string {
	return s.InstanceRegion
}

func (s *DescribeBackupDataListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupDataListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupDataListRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeBackupDataListRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *DescribeBackupDataListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupDataListRequest) SetBackupId(v string) *DescribeBackupDataListRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetBackupMethod(v string) *DescribeBackupDataListRequest {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetBackupMode(v string) *DescribeBackupDataListRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetBackupScale(v string) *DescribeBackupDataListRequest {
	s.BackupScale = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetBackupStatus(v string) *DescribeBackupDataListRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetBackupType(v string) *DescribeBackupDataListRequest {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetDataSourceId(v string) *DescribeBackupDataListRequest {
	s.DataSourceId = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetEndTime(v string) *DescribeBackupDataListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetInstanceIsDeleted(v bool) *DescribeBackupDataListRequest {
	s.InstanceIsDeleted = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetInstanceName(v string) *DescribeBackupDataListRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetInstanceRegion(v string) *DescribeBackupDataListRequest {
	s.InstanceRegion = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetPageNumber(v int32) *DescribeBackupDataListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetPageSize(v int32) *DescribeBackupDataListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetRegionCode(v string) *DescribeBackupDataListRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetSceneType(v string) *DescribeBackupDataListRequest {
	s.SceneType = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetStartTime(v string) *DescribeBackupDataListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupDataListRequest) Validate() error {
	return dara.Validate(s)
}
