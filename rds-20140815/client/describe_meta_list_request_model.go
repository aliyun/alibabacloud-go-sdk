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
	// The ID of the backup set from which you want to restore data. You can call the DescribeBackups operation to query the IDs of data backup files.
	//
	// >  This parameter is required when you set the **RestoreType*	- parameter to **BackupSetID**.
	//
	// example:
	//
	// 14358
	BackupSetID *int64 `json:"BackupSetID,omitempty" xml:"BackupSetID,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database to query. The system implements exact match based on the value of this parameter and returns the name of the matched database and the names of all tables contained in the database.
	//
	// > If you leave this parameter empty, the system returns all databases that are created on the instance.
	//
	// example:
	//
	// testdb1
	GetDbName *string `json:"GetDbName,omitempty" xml:"GetDbName,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Valid values: any non-zero positive integer.***	- Default value: **1**.
	//
	// > This parameter only takes effect when you specify the **PageSize*	- parameter.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries to return on each page. Default value: **1**.
	//
	// > This parameter only takes effect when you specify the **PageIndex*	- parameter.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the database to query. The system implements fuzzy match based on the value of this parameter and returns only the name of the matched database.
	//
	// > For example, if you set the value to `test`, the system returns `testdb1` and `testdb2`. Then, you can specify the **GetDbName*	- parameter to query tables in the required database.
	//
	// example:
	//
	// test
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore data. The specified point in time must be earlier than the current time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC. You can call the DescribeBackups operation to query the restorable time range.
	//
	// >  This parameter must be specified when the **RestoreType*	- parameter is set to **RestoreTime**.
	//
	// example:
	//
	// 2019-05-30T03:29:10Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The restoration method that you want to use. Valid values:
	//
	// 	- **BackupSetID**: Data is restored from the backup set. If you use this value, you must also specify the **BackupSetID*	- parameter.
	//
	// 	- **RestoreTime**: Data is restored to a specific point in time. If you use this value, you must also specify the **RestoreTime*	- parameter.
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
