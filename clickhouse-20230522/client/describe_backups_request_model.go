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
	SetDBInstanceId(v string) *DescribeBackupsRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeBackupsRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBackupsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeBackupsRequest
	GetStartTime() *string
}

type DescribeBackupsRequest struct {
	// The backup record ID.
	//
	// example:
	//
	// 117403****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1v9kq45u0o80cvh
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. The time is in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-25T16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The time is in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-21T16:00Z
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

func (s *DescribeBackupsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetDBInstanceId(v string) *DescribeBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
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

func (s *DescribeBackupsRequest) SetRegionId(v string) *DescribeBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) Validate() error {
	return dara.Validate(s)
}
