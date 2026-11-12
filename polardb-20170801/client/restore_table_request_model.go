// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *RestoreTableRequest
	GetBackupId() *string
	SetDBClusterId(v string) *RestoreTableRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *RestoreTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestoreTableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestoreTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestoreTableRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *RestoreTableRequest
	GetRestoreTime() *string
	SetSecurityToken(v string) *RestoreTableRequest
	GetSecurityToken() *string
	SetTableMeta(v string) *RestoreTableRequest
	GetTableMeta() *string
}

type RestoreTableRequest struct {
	// The backup set ID.
	//
	// > This parameter is required if you want to restore the destination databases and tables from a backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/98102.html) operation to query the backup set ID.
	//
	// > To obtain the backup set ID, call the DescribeBackups operation and specify the StartTime and EndTime parameters. Then, call the DescribeMetaList operation to obtain the metadata of the restorable databases and tables. Operation description: Dependency chain description: To obtain the BackupId, you must call the DescribeBackups operation (the StartTime and EndTime parameters are required). To obtain the database and table names required for TableMeta, you must call the DescribeMetaList operation.
	//
	// example:
	//
	// 111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters within your account.
	//
	// > The destination cluster must be in the Running state. You can call the DescribeDBClusterAttribute operation to query the current state. Operation description: Status prerequisite: The cluster must be in the Running state. You can call the DescribeDBClusterAttribute operation to query the current state.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp***************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore the data. Specify the time in the YYYY-MM-DDThh:mmZ format. The time must be in UTC.
	//
	// >	- This parameter is required if you want to restore data to a specific point in time.
	//
	// >	- You can restore data to any point in time within the last seven days.
	//
	// example:
	//
	// 2020-10-04T01:40:00Z
	RestoreTime   *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The JSON string that contains the information about the destination databases and tables to be restored. All values in the database and table information are of the string type.
	//
	// Example: `[
	//
	//    {
	//
	//        "tables":[
	//
	//            {
	//
	//                "name":"testtb",
	//
	//                "type":"table",
	//
	//                "newname":"testtb_restore"
	//
	//            }
	//
	//        ],
	//
	//        "name":"testdb",
	//
	//        "type":"db",
	//
	//        "newname":"testdb_restore"
	//
	//    }
	//
	// ]`.
	//
	// > You can call the [DescribeMetaList](https://help.aliyun.com/document_detail/194770.html) operation to view the details of the databases and tables that can be restored, and fill in the restorable database and table information in the corresponding positions in the preceding example.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ { "tables":[ { "name":"testtb", "type":"table", "newname":"testtb_restore" } ], "name":"testdb", "type":"db", "newname":"testdb_restore" } ]
	TableMeta *string `json:"TableMeta,omitempty" xml:"TableMeta,omitempty"`
}

func (s RestoreTableRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreTableRequest) GoString() string {
	return s.String()
}

func (s *RestoreTableRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *RestoreTableRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestoreTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestoreTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestoreTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestoreTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestoreTableRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *RestoreTableRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RestoreTableRequest) GetTableMeta() *string {
	return s.TableMeta
}

func (s *RestoreTableRequest) SetBackupId(v string) *RestoreTableRequest {
	s.BackupId = &v
	return s
}

func (s *RestoreTableRequest) SetDBClusterId(v string) *RestoreTableRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestoreTableRequest) SetOwnerAccount(v string) *RestoreTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestoreTableRequest) SetOwnerId(v int64) *RestoreTableRequest {
	s.OwnerId = &v
	return s
}

func (s *RestoreTableRequest) SetResourceOwnerAccount(v string) *RestoreTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestoreTableRequest) SetResourceOwnerId(v int64) *RestoreTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestoreTableRequest) SetRestoreTime(v string) *RestoreTableRequest {
	s.RestoreTime = &v
	return s
}

func (s *RestoreTableRequest) SetSecurityToken(v string) *RestoreTableRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestoreTableRequest) SetTableMeta(v string) *RestoreTableRequest {
	s.TableMeta = &v
	return s
}

func (s *RestoreTableRequest) Validate() error {
	return dara.Validate(s)
}
