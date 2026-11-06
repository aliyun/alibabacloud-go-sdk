// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationBackupsRequest
	GetApplicationId() *string
	SetApplicationType(v string) *DescribeApplicationBackupsRequest
	GetApplicationType() *string
	SetBackupStatus(v string) *DescribeApplicationBackupsRequest
	GetBackupStatus() *string
	SetEndTime(v string) *DescribeApplicationBackupsRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeApplicationBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApplicationBackupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeApplicationBackupsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeApplicationBackupsRequest
	GetStartTime() *string
}

type DescribeApplicationBackupsRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type, such as polarclaw.
	//
	// example:
	//
	// polarclaw
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The filter condition for the backup status. A value of Success indicates that the backup is completed, and a value of Failed indicates that the backup failed.
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The end time of the query. This parameter is required and cannot be earlier than the start time. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-11-16T00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. The value must be an integer greater than 0 and less than or equal to 2147483647. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 30, 50, and 100. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the application.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query. This parameter is required. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time must be in UTC and cannot be later than the current time.
	//
	// example:
	//
	// 2020-11-14T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApplicationBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationBackupsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationBackupsRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeApplicationBackupsRequest) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeApplicationBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApplicationBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApplicationBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationBackupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApplicationBackupsRequest) SetApplicationId(v string) *DescribeApplicationBackupsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationBackupsRequest) SetApplicationType(v string) *DescribeApplicationBackupsRequest {
	s.ApplicationType = &v
	return s
}

func (s *DescribeApplicationBackupsRequest) SetBackupStatus(v string) *DescribeApplicationBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeApplicationBackupsRequest) SetEndTime(v string) *DescribeApplicationBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApplicationBackupsRequest) SetPageNumber(v int32) *DescribeApplicationBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationBackupsRequest) SetPageSize(v int32) *DescribeApplicationBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationBackupsRequest) SetRegionId(v string) *DescribeApplicationBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationBackupsRequest) SetStartTime(v string) *DescribeApplicationBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApplicationBackupsRequest) Validate() error {
	return dara.Validate(s)
}
