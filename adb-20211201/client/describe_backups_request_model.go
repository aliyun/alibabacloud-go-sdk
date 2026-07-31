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
	SetRemote(v bool) *DescribeBackupsRequest
	GetRemote() *bool
	SetResourceOwnerAccount(v string) *DescribeBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeBackupsRequest
	GetStartTime() *string
}

type DescribeBackupsRequest struct {
	// The backup set ID.
	//
	// example:
	//
	// 1679758862
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all clusters in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp18934i73vb5****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the \\`yyyy-MM-ddTHH:mmZ\\` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-20T02:30Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the integer data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// - 30
	//
	// - 50
	//
	// - 100
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to query only cross-region backups. Set the value to true to query only cross-region backups. Set the value to false to query only local backups.
	//
	// example:
	//
	// false
	Remote               *bool   `json:"Remote,omitempty" xml:"Remote,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the \\`yyyy-MM-ddTHH:mmZ\\` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-06-01T16:00Z
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

func (s *DescribeBackupsRequest) GetRemote() *bool {
	return s.Remote
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

func (s *DescribeBackupsRequest) SetRemote(v bool) *DescribeBackupsRequest {
	s.Remote = &v
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
