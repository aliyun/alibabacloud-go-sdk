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
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
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
