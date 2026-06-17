// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeMetaListRequest
	GetBackupId() *string
	SetDBClusterId(v string) *DescribeMetaListRequest
	GetDBClusterId() *string
	SetGetDbName(v string) *DescribeMetaListRequest
	GetGetDbName() *string
	SetOwnerAccount(v string) *DescribeMetaListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeMetaListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeMetaListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMetaListRequest
	GetPageSize() *int32
	SetRegionCode(v string) *DescribeMetaListRequest
	GetRegionCode() *string
	SetResourceOwnerAccount(v string) *DescribeMetaListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMetaListRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *DescribeMetaListRequest
	GetRestoreTime() *string
	SetSecurityToken(v string) *DescribeMetaListRequest
	GetSecurityToken() *string
}

type DescribeMetaListRequest struct {
	// The ID of the backup set.
	//
	// > - You must specify either the `BackupId` or `RestoreTime` parameter.
	//
	// >
	//
	// > - Call the [DescribeBackups](https://help.aliyun.com/document_detail/98102.html) operation to query the backup set ID.
	//
	// example:
	//
	// 111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The cluster ID.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters in your account.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database, such as `test_db`. If you specify this parameter, the names of all tables that can be recovered from the specified database are returned.
	//
	// > - You can specify only one database name at a time.
	//
	// >
	//
	// > - If you do not specify this parameter, the names of all databases that can be recovered from the current backup set are returned. The names of tables in each database are not returned.
	//
	// example:
	//
	// test_db
	GetDbName    *string `json:"GetDbName,omitempty" xml:"GetDbName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer that is greater than 0 and does not exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100*	- Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. You can call [DescribeDBClusterAttribute](https://help.aliyun.com/document_detail/2319132.html) to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode           *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore data. Specify the time in the YYYY-MM-DDThh:mmZ format. The time must be in UTC.
	//
	// > You must specify either the `BackupId` or `RestoreTime` parameter. Call the [DescribeBackups](https://help.aliyun.com/document_detail/98102.html) operation to query the point in time for restoration.
	//
	// example:
	//
	// 2020-10-04T01:40:00Z
	RestoreTime   *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeMetaListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetaListRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeMetaListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeMetaListRequest) GetGetDbName() *string {
	return s.GetDbName
}

func (s *DescribeMetaListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeMetaListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMetaListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMetaListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetaListRequest) GetRegionCode() *string {
	return s.RegionCode
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

func (s *DescribeMetaListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeMetaListRequest) SetBackupId(v string) *DescribeMetaListRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeMetaListRequest) SetDBClusterId(v string) *DescribeMetaListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMetaListRequest) SetGetDbName(v string) *DescribeMetaListRequest {
	s.GetDbName = &v
	return s
}

func (s *DescribeMetaListRequest) SetOwnerAccount(v string) *DescribeMetaListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMetaListRequest) SetOwnerId(v int64) *DescribeMetaListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMetaListRequest) SetPageNumber(v int32) *DescribeMetaListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetaListRequest) SetPageSize(v int32) *DescribeMetaListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetaListRequest) SetRegionCode(v string) *DescribeMetaListRequest {
	s.RegionCode = &v
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

func (s *DescribeMetaListRequest) SetSecurityToken(v string) *DescribeMetaListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeMetaListRequest) Validate() error {
	return dara.Validate(s)
}
