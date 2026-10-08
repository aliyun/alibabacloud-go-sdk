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
	SetClientToken(v string) *RestoreTableRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *RestoreTableRequest
	GetDBInstanceId() *string
	SetInstantRecovery(v bool) *RestoreTableRequest
	GetInstantRecovery() *bool
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
	SetTableMeta(v string) *RestoreTableRequest
	GetTableMeta() *string
}

type RestoreTableRequest struct {
	// The backup set ID. You can call the DescribeBackups operation to query the backup set list.
	//
	// > You must specify at least one of **BackupId*	- and **RestoreTime**.
	//
	// example:
	//
	// 902****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable fast restoration for individual databases and tables. Valid values:
	//
	// 	- **true**: Enabled.
	//
	// 	- **false**: Disabled.
	//
	// > For more information about fast restoration for individual databases and tables, see [Restore individual databases and tables](https://help.aliyun.com/document_detail/103175.html).
	//
	// example:
	//
	// true
	InstantRecovery      *bool   `json:"InstantRecovery,omitempty" xml:"InstantRecovery,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Any point in time within the backup retention period. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > 	- You must specify at least one of **BackupId*	- and **RestoreTime**.
	//
	// > 	- [Log backup](https://help.aliyun.com/document_detail/98818.html) must be enabled for the instance.
	//
	// example:
	//
	// 2011-06-11T16:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The databases and tables to restore.
	//
	// > ApsaraDB RDS for PostgreSQL supports only the restoration of specific databases, not specific tables.
	//
	// - ApsaraDB RDS for MySQL format: ```[{"type":"db","name":"<Database 1 name>","newname":"<New database 1 name>","tables":[{"type":"table","name":"<Table 1 name in database 1>","newname":"<New table 1 name>"},{"type":"table","name":"<Table 2 name in database 1>","newname":"<New table 2 name>"}]},{"type":"db","name":"<Database 2 name>","newname":"<New database 2 name>","tables":[{"type":"table","name":"<Table 3 name in database 2>","newname":"<New table 3 name>"},{"type":"table","name":"<Table 4 name in database 2>","newname":"<New table 4 name>"}]}]```
	//
	// - ApsaraDB RDS for PostgreSQL format: ```[{"type":"db","name":"<Database 1 name>","newname":"<New database 1 name>"}]```
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"type":"db","name":"testdb1","newname":"testdb1_new","tables":[{"type":"table","name":"testdb1table1","newname":"testdb1table1_new"}]}]
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

func (s *RestoreTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestoreTableRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestoreTableRequest) GetInstantRecovery() *bool {
	return s.InstantRecovery
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

func (s *RestoreTableRequest) GetTableMeta() *string {
	return s.TableMeta
}

func (s *RestoreTableRequest) SetBackupId(v string) *RestoreTableRequest {
	s.BackupId = &v
	return s
}

func (s *RestoreTableRequest) SetClientToken(v string) *RestoreTableRequest {
	s.ClientToken = &v
	return s
}

func (s *RestoreTableRequest) SetDBInstanceId(v string) *RestoreTableRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestoreTableRequest) SetInstantRecovery(v bool) *RestoreTableRequest {
	s.InstantRecovery = &v
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

func (s *RestoreTableRequest) SetTableMeta(v string) *RestoreTableRequest {
	s.TableMeta = &v
	return s
}

func (s *RestoreTableRequest) Validate() error {
	return dara.Validate(s)
}
