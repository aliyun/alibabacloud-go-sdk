// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSharedBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeSharedBackupsRequest
	GetBackupId() *string
	SetDBClusterId(v string) *DescribeSharedBackupsRequest
	GetDBClusterId() *string
	SetDBType(v string) *DescribeSharedBackupsRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeSharedBackupsRequest
	GetDBVersion() *string
	SetOwnerAccount(v string) *DescribeSharedBackupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSharedBackupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSharedBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSharedBackupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSharedBackupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSharedBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSharedBackupsRequest
	GetResourceOwnerId() *int64
	SetShareType(v string) *DescribeSharedBackupsRequest
	GetShareType() *string
}

type DescribeSharedBackupsRequest struct {
	// example:
	//
	// 111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// MySQL,PostgreSQL,Oracle
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 8.0
	DBVersion    *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ShareIncoming
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s DescribeSharedBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSharedBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSharedBackupsRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeSharedBackupsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSharedBackupsRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeSharedBackupsRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeSharedBackupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSharedBackupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSharedBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSharedBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSharedBackupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSharedBackupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSharedBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSharedBackupsRequest) GetShareType() *string {
	return s.ShareType
}

func (s *DescribeSharedBackupsRequest) SetBackupId(v string) *DescribeSharedBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetDBClusterId(v string) *DescribeSharedBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetDBType(v string) *DescribeSharedBackupsRequest {
	s.DBType = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetDBVersion(v string) *DescribeSharedBackupsRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetOwnerAccount(v string) *DescribeSharedBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetOwnerId(v int64) *DescribeSharedBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetPageNumber(v int32) *DescribeSharedBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetPageSize(v int32) *DescribeSharedBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetRegionId(v string) *DescribeSharedBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetResourceOwnerAccount(v string) *DescribeSharedBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetResourceOwnerId(v int64) *DescribeSharedBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSharedBackupsRequest) SetShareType(v string) *DescribeSharedBackupsRequest {
	s.ShareType = &v
	return s
}

func (s *DescribeSharedBackupsRequest) Validate() error {
	return dara.Validate(s)
}
