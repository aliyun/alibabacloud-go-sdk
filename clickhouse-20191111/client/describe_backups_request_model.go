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
	SetDBClusterId(v string) *DescribeBackupsRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeBackupsRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeBackupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeBackupsRequest
	GetStartTime() *string
}

type DescribeBackupsRequest struct {
	// The backup task ID.
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
	// cc-bp1qx68m06981****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-25T16:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
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
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
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

func (s *DescribeBackupsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeBackupsRequest) GetEndTime() *string {
	return s.EndTime
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

func (s *DescribeBackupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetDBClusterId(v string) *DescribeBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
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

func (s *DescribeBackupsRequest) Validate() error {
	return dara.Validate(s)
}
