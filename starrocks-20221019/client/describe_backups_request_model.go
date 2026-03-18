// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupTaskId(v string) *DescribeBackupsRequest
	GetBackupTaskId() *string
	SetInstanceId(v string) *DescribeBackupsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBackupsRequest
	GetRegionId() *string
	SetStatuses(v string) *DescribeBackupsRequest
	GetStatuses() *string
	SetTimePeriodEndTime(v int64) *DescribeBackupsRequest
	GetTimePeriodEndTime() *int64
	SetTimePeriodStartTime(v int64) *DescribeBackupsRequest
	GetTimePeriodStartTime() *int64
}

type DescribeBackupsRequest struct {
	// example:
	//
	// bt-2389hsdui12m
	BackupTaskId *string `json:"BackupTaskId,omitempty" xml:"BackupTaskId,omitempty"`
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// COMPLETED
	Statuses *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
	// example:
	//
	// 1747728000
	TimePeriodEndTime *int64 `json:"TimePeriodEndTime,omitempty" xml:"TimePeriodEndTime,omitempty"`
	// example:
	//
	// 1747708000
	TimePeriodStartTime *int64 `json:"TimePeriodStartTime,omitempty" xml:"TimePeriodStartTime,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) GetBackupTaskId() *string {
	return s.BackupTaskId
}

func (s *DescribeBackupsRequest) GetInstanceId() *string {
	return s.InstanceId
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

func (s *DescribeBackupsRequest) GetStatuses() *string {
	return s.Statuses
}

func (s *DescribeBackupsRequest) GetTimePeriodEndTime() *int64 {
	return s.TimePeriodEndTime
}

func (s *DescribeBackupsRequest) GetTimePeriodStartTime() *int64 {
	return s.TimePeriodStartTime
}

func (s *DescribeBackupsRequest) SetBackupTaskId(v string) *DescribeBackupsRequest {
	s.BackupTaskId = &v
	return s
}

func (s *DescribeBackupsRequest) SetInstanceId(v string) *DescribeBackupsRequest {
	s.InstanceId = &v
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

func (s *DescribeBackupsRequest) SetStatuses(v string) *DescribeBackupsRequest {
	s.Statuses = &v
	return s
}

func (s *DescribeBackupsRequest) SetTimePeriodEndTime(v int64) *DescribeBackupsRequest {
	s.TimePeriodEndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetTimePeriodStartTime(v int64) *DescribeBackupsRequest {
	s.TimePeriodStartTime = &v
	return s
}

func (s *DescribeBackupsRequest) Validate() error {
	return dara.Validate(s)
}
