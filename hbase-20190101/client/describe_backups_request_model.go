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
	SetClusterId(v string) *DescribeBackupsRequest
	GetClusterId() *string
	SetEndTime(v string) *DescribeBackupsRequest
	GetEndTime() *string
	SetEndTimeUTC(v string) *DescribeBackupsRequest
	GetEndTimeUTC() *string
	SetPageNumber(v string) *DescribeBackupsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeBackupsRequest
	GetPageSize() *string
	SetStartTime(v string) *DescribeBackupsRequest
	GetStartTime() *string
	SetStartTimeUTC(v string) *DescribeBackupsRequest
	GetStartTimeUTC() *string
}

type DescribeBackupsRequest struct {
	// example:
	//
	// job-xxxx
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 2020-12-23 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2020-12-23T15:59:59Z
	EndTimeUTC *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2020-12-13 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2020-12-12T16:00:00Z
	StartTimeUTC *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
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

func (s *DescribeBackupsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBackupsRequest) GetEndTimeUTC() *string {
	return s.EndTimeUTC
}

func (s *DescribeBackupsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeBackupsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupsRequest) GetStartTimeUTC() *string {
	return s.StartTimeUTC
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetClusterId(v string) *DescribeBackupsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTimeUTC(v string) *DescribeBackupsRequest {
	s.EndTimeUTC = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v string) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v string) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTimeUTC(v string) *DescribeBackupsRequest {
	s.StartTimeUTC = &v
	return s
}

func (s *DescribeBackupsRequest) Validate() error {
	return dara.Validate(s)
}
