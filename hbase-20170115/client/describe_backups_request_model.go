// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v int32) *DescribeBackupsRequest
	GetBackupId() *int32
	SetClusterId(v string) *DescribeBackupsRequest
	GetClusterId() *string
	SetEndTime(v string) *DescribeBackupsRequest
	GetEndTime() *string
	SetEndTimeUTC(v string) *DescribeBackupsRequest
	GetEndTimeUTC() *string
	SetOwnerId(v int64) *DescribeBackupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBackupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeBackupsRequest
	GetStartTime() *string
	SetStartTimeUTC(v string) *DescribeBackupsRequest
	GetStartTimeUTC() *string
	SetZoneId(v string) *DescribeBackupsRequest
	GetZoneId() *string
}

type DescribeBackupsRequest struct {
	BackupId *int32 `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC           *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC         *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) GetBackupId() *int32 {
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

func (s *DescribeBackupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
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

func (s *DescribeBackupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupsRequest) GetStartTimeUTC() *string {
	return s.StartTimeUTC
}

func (s *DescribeBackupsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeBackupsRequest) SetBackupId(v int32) *DescribeBackupsRequest {
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

func (s *DescribeBackupsRequest) SetRegionId(v string) *DescribeBackupsRequest {
	s.RegionId = &v
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

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTimeUTC(v string) *DescribeBackupsRequest {
	s.StartTimeUTC = &v
	return s
}

func (s *DescribeBackupsRequest) SetZoneId(v string) *DescribeBackupsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeBackupsRequest) Validate() error {
	return dara.Validate(s)
}
