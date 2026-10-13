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
	// The ID of the backup set that is used to query data when the restore method is based on a backup set. You can call the DescribeBackups operation to query the ID.
	//
	// > This parameter is required when **RestoreType*	- is set to **BackupSetID**.
	//
	// example:
	//
	// 14***
	BackupSetID *int64 `json:"BackupSetID,omitempty" xml:"BackupSetID,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The value can be up to 64 ASCII characters in length and cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Instance ID of the instance. You can call the DescribeDBInstances operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database to query. Exact match is used. The specific database name and all tables in the database are returned.
	//
	// > If this parameter is left empty, a list of all databases is returned.
	//
	// example:
	//
	// testdb1
	GetDbName *string `json:"GetDbName,omitempty" xml:"GetDbName,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: any integer greater than **0*	- and less than or equal to the maximum value of the Integer data type. Default value: **1**.
	//
	// > This parameter takes effect only when it is used together with **PageSize**.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page. Default value: **1**.
	//
	// > This parameter takes effect only when it is used together with **PageIndex**.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the database to query. Fuzzy match is used. Only the matched database names are returned, and table names are not returned.
	//
	// > For example, if you specify `test`, `testdb1` and `testdb2` are matched. After you determine the name of the target database, you can specify the exact database name in the **GetDbName*	- parameter to view the information about all tables in the target database.
	//
	// example:
	//
	// test
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore data. The time must be earlier than the current time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC. You can call the DescribeBackups operation to query the time.
	//
	// > This parameter is required when **RestoreType*	- is set to **RestoreTime**.
	//
	// example:
	//
	// 2019-05-30T03:29:10Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The restore method. Valid values:
	//
	// 	- **BackupSetID**: restores data from a backup set. If you set this parameter to BackupSetID, you must also specify the **BackupSetID*	- parameter.
	//
	// 	- **RestoreTime**: restores data to a specific point in time. If you set this parameter to RestoreTime, you must also specify the **RestoreTime*	- parameter.
	//
	// Default value: **BackupSetID**.
	//
	// example:
	//
	// BackupSetID
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
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
