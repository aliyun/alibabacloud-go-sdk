// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogBackupFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLogBackupFilesRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeLogBackupFilesRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeLogBackupFilesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLogBackupFilesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLogBackupFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLogBackupFilesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeLogBackupFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLogBackupFilesRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeLogBackupFilesRequest
	GetStartTime() *string
}

type DescribeLogBackupFilesRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-10-31T08:40Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: any non-zero positive integer.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30*	- to **1000**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-10-01T08:40Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLogBackupFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupFilesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLogBackupFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLogBackupFilesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLogBackupFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLogBackupFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLogBackupFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLogBackupFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLogBackupFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLogBackupFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLogBackupFilesRequest) SetDBInstanceId(v string) *DescribeLogBackupFilesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLogBackupFilesRequest) SetEndTime(v string) *DescribeLogBackupFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLogBackupFilesRequest) SetOwnerAccount(v string) *DescribeLogBackupFilesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLogBackupFilesRequest) SetOwnerId(v int64) *DescribeLogBackupFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLogBackupFilesRequest) SetPageNumber(v int32) *DescribeLogBackupFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogBackupFilesRequest) SetPageSize(v int32) *DescribeLogBackupFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogBackupFilesRequest) SetResourceOwnerAccount(v string) *DescribeLogBackupFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLogBackupFilesRequest) SetResourceOwnerId(v int64) *DescribeLogBackupFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLogBackupFilesRequest) SetStartTime(v string) *DescribeLogBackupFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLogBackupFilesRequest) Validate() error {
	return dara.Validate(s)
}
