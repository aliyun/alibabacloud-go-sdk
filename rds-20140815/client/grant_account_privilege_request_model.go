// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAccountPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GrantAccountPrivilegeRequest
	GetAccountName() *string
	SetAccountPrivilege(v string) *GrantAccountPrivilegeRequest
	GetAccountPrivilege() *string
	SetDBInstanceId(v string) *GrantAccountPrivilegeRequest
	GetDBInstanceId() *string
	SetDBName(v string) *GrantAccountPrivilegeRequest
	GetDBName() *string
	SetResourceOwnerId(v int64) *GrantAccountPrivilegeRequest
	GetResourceOwnerId() *int64
}

type GrantAccountPrivilegeRequest struct {
	// The account name. You can call [DescribeAccounts](https://help.aliyun.com/document_detail/610454.html) to query the account name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of account permission. If you specify multiple values for DBName, you must specify the same number of permission types in the same order, separated by commas (,).
	//
	// The supported permission types vary by database engine. Valid values:
	//
	// > For more information about account permissions, see [MySQL/MariaDB permission list](https://help.aliyun.com/document_detail/146395.html), [SQL Server permission list](https://help.aliyun.com/document_detail/95692.html), and [PostgreSQL permission list](https://help.aliyun.com/document_detail/257684.html).
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for MySQL/ApsaraDB RDS for MariaDB</summary>
	//
	// - **ReadWrite**: read and write.
	//
	// - **ReadOnly**: read-only.
	//
	// - **DDLOnly**: DDL only.
	//
	// - **DMLOnly**: DML only.
	//
	// </details>
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for SQL Server</summary>
	//
	// - **ReadWrite**: read and write. This permission corresponds to the `db_datawriter` and `db_datareader` database roles in SQL Server.
	//
	// - **ReadOnly**: read-only. This permission corresponds to the `db_datareader` database role in SQL Server.
	//
	// - **DBOwner**: database owner. This permission corresponds to the `db_owner` database role in SQL Server.
	//
	// > For more information about database-level roles, see [Microsoft official documentation](https://learn.microsoft.com/en-us/sql/relational-databases/security/authentication-access/database-level-roles?view=sql-server-ver16).
	//
	// </details>
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for PostgreSQL</summary>
	//
	// **DBOwner**: database owner.
	//
	// > For fine-grained permission management, see [Best practices for PostgreSQL permission management](https://help.aliyun.com/document_detail/352149.html).
	//
	// </details>
	//
	// This parameter is required.
	//
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The instance ID. You can call [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database to which you want to grant access permissions. To grant permissions on multiple databases at a time, separate the database names with commas (,), such as `db1,db2,db3`.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDB1
	DBName          *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GrantAccountPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GrantAccountPrivilegeRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *GrantAccountPrivilegeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GrantAccountPrivilegeRequest) GetDBName() *string {
	return s.DBName
}

func (s *GrantAccountPrivilegeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GrantAccountPrivilegeRequest) SetAccountName(v string) *GrantAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetAccountPrivilege(v string) *GrantAccountPrivilegeRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetDBInstanceId(v string) *GrantAccountPrivilegeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetDBName(v string) *GrantAccountPrivilegeRequest {
	s.DBName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
