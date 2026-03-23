// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetID(v int64) *DescribeMetaListRequest
	GetBackupSetID() *int64
	SetClientToken(v string) *DescribeMetaListRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeMetaListRequest
	GetDBInstanceId() *string
	SetGetDbName(v string) *DescribeMetaListRequest
	GetGetDbName() *string
	SetOwnerId(v int64) *DescribeMetaListRequest
	GetOwnerId() *int64
	SetPageIndex(v int32) *DescribeMetaListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *DescribeMetaListRequest
	GetPageSize() *int32
	SetPattern(v string) *DescribeMetaListRequest
	GetPattern() *string
	SetResourceGroupId(v string) *DescribeMetaListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeMetaListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMetaListRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *DescribeMetaListRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *DescribeMetaListRequest
	GetRestoreType() *string
}

type DescribeMetaListRequest struct {
	BackupSetID *int64  `json:"BackupSetID,omitempty" xml:"BackupSetID,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	GetDbName            *string `json:"GetDbName,omitempty" xml:"GetDbName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageIndex            *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pattern              *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	RestoreType          *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
}

func (s DescribeMetaListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetaListRequest) GetBackupSetID() *int64 {
	return s.BackupSetID
}

func (s *DescribeMetaListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeMetaListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeMetaListRequest) GetGetDbName() *string {
	return s.GetDbName
}

func (s *DescribeMetaListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMetaListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeMetaListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetaListRequest) GetPattern() *string {
	return s.Pattern
}

func (s *DescribeMetaListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMetaListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMetaListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMetaListRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *DescribeMetaListRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *DescribeMetaListRequest) SetBackupSetID(v int64) *DescribeMetaListRequest {
	s.BackupSetID = &v
	return s
}

func (s *DescribeMetaListRequest) SetClientToken(v string) *DescribeMetaListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMetaListRequest) SetDBInstanceId(v string) *DescribeMetaListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeMetaListRequest) SetGetDbName(v string) *DescribeMetaListRequest {
	s.GetDbName = &v
	return s
}

func (s *DescribeMetaListRequest) SetOwnerId(v int64) *DescribeMetaListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMetaListRequest) SetPageIndex(v int32) *DescribeMetaListRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeMetaListRequest) SetPageSize(v int32) *DescribeMetaListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetaListRequest) SetPattern(v string) *DescribeMetaListRequest {
	s.Pattern = &v
	return s
}

func (s *DescribeMetaListRequest) SetResourceGroupId(v string) *DescribeMetaListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMetaListRequest) SetResourceOwnerAccount(v string) *DescribeMetaListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMetaListRequest) SetResourceOwnerId(v int64) *DescribeMetaListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMetaListRequest) SetRestoreTime(v string) *DescribeMetaListRequest {
	s.RestoreTime = &v
	return s
}

func (s *DescribeMetaListRequest) SetRestoreType(v string) *DescribeMetaListRequest {
	s.RestoreType = &v
	return s
}

func (s *DescribeMetaListRequest) Validate() error {
	return dara.Validate(s)
}
