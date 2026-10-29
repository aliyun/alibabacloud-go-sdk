// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateDatabaseRequest
	GetAccountName() *string
	SetAccountPrivilege(v string) *CreateDatabaseRequest
	GetAccountPrivilege() *string
	SetCharacterSetName(v string) *CreateDatabaseRequest
	GetCharacterSetName() *string
	SetCollationName(v string) *CreateDatabaseRequest
	GetCollationName() *string
	SetDBDescription(v string) *CreateDatabaseRequest
	GetDBDescription() *string
	SetDBInstanceId(v string) *CreateDatabaseRequest
	GetDBInstanceId() *string
	SetDBName(v string) *CreateDatabaseRequest
	GetDBName() *string
	SetOwnerAccount(v string) *CreateDatabaseRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDatabaseRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateDatabaseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDatabaseRequest
	GetResourceOwnerId() *int64
}

type CreateDatabaseRequest struct {
	// The name of the account. You can call the [DescribeAccounts](https://help.aliyun.com/document_detail/610454.html) operation to query the name.
	//
	// example:
	//
	// test1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of permission granted to the account. If you specify multiple database names, you must specify the same number of permission types in the same order and separate them with commas (,).
	//
	// Different database engines support different permission types. Valid values:
	//
	// > For more information about account permissions, see [Permission list for ApsaraDB RDS for MySQL and ApsaraDB RDS for MariaDB](https://help.aliyun.com/document_detail/146395.html), [Permission list for ApsaraDB RDS for SQL Server](https://help.aliyun.com/document_detail/95692.html), and [Permission list for ApsaraDB RDS for PostgreSQL](https://help.aliyun.com/document_detail/257684.html).
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for MySQL and ApsaraDB RDS for MariaDB</summary>
	//
	// - **ReadWrite**: read and write
	//
	// - **ReadOnly**: read-only
	//
	// - **DDLOnly**: DDL-only
	//
	// - **DMLOnly**: DML-only
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
	// **DBOwner**: database owner
	//
	// > For more information about fine-grained permission management for accounts, see [Best practices for permission management in ApsaraDB RDS for PostgreSQL](https://help.aliyun.com/document_detail/352149.html).
	//
	// </details>
	//
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The character set. Valid values:
	//
	// 	- ApsaraDB RDS for MySQL and ApsaraDB RDS for MariaDB: **utf8**, **gbk**, **latin1**, and **utf8mb4**.
	//
	// 	- ApsaraDB RDS for SQL Server: **Chinese_PRC_CI_AS**, **Chinese_PRC_CS_AS**, **SQL_Latin1_General_CP1_CI_AS**, **SQL_Latin1_General_CP1_CS_AS**, and **Chinese_PRC_BIN**.
	//
	// 	- ApsaraDB RDS for PostgreSQL: You must configure the character set, collation, and ctype in the `character set,<collation>,<ctype>` format. Example: `UTF8,C,en_US.utf8`.
	//
	//     - Valid values for the character set: **KOI8U, UTF8, WIN866, WIN874, WIN1250, WIN1251, WIN1252, WIN1253, WIN1254, WIN1255, WIN1256, WIN1257, WIN1258, EUC_CN, EUC_KR, EUC_TW, EUC_JP, EUC_JIS_2004, KOI8R, MULE_INTERNAL, LATIN1, LATIN2, LATIN3, LATIN4, LATIN5, LATIN6, LATIN7, LATIN8, LATIN9, LATIN10, ISO_8859_5, ISO_8859_6, ISO_8859_7, ISO_8859_8, and SQL_ASCII**.
	//
	//     - Valid values for **collation**: You can execute the `SELECT DISTINCT collname FROM pg_collation;` statement to query the values. If you do not configure this parameter, the default value is **C**.
	//
	//     - Valid values for **ctype**: You can execute the `SELECT DISTINCT collctype FROM pg_collation;` statement to query the values. If you do not configure this parameter, the default value is **en_US.utf8**.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbk
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	// The collation. This parameter is only supported for ApsaraDB RDS for MySQL instances. Specify a collation that matches the character set. For example, if the character set is utf8mb4, the collation must be utf8mb4_bin or utf8mb4_general_ci.
	//
	// example:
	//
	// gbk_chinese_ci
	CollationName *string `json:"CollationName,omitempty" xml:"CollationName,omitempty"`
	// The description of the database. The description must be 2 to 256 characters in length. It must start with a Chinese or English letter and can contain digits, Chinese characters, English letters, underscores (_), and hyphens (-).
	//
	// > It cannot start with http:// or https://.
	//
	// example:
	//
	// testdb
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// Instance ID of the instance. You can call the DescribeDBInstances operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// > 	- The name must be 2 to 64 characters in length.
	//
	// > 	- It must start with a letter and end with a letter or digit.
	//
	// > 	- It can contain lowercase letters, digits, underscores (_), and hyphens (-).
	//
	// > 	- The database name must be unique within the instance.
	//
	// > 	- For other invalid characters, see [Reserved keywords](https://help.aliyun.com/document_detail/26317.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_mysql
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateDatabaseRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *CreateDatabaseRequest) GetCharacterSetName() *string {
	return s.CharacterSetName
}

func (s *CreateDatabaseRequest) GetCollationName() *string {
	return s.CollationName
}

func (s *CreateDatabaseRequest) GetDBDescription() *string {
	return s.DBDescription
}

func (s *CreateDatabaseRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDatabaseRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateDatabaseRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDatabaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDatabaseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDatabaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDatabaseRequest) SetAccountName(v string) *CreateDatabaseRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDatabaseRequest) SetAccountPrivilege(v string) *CreateDatabaseRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateDatabaseRequest) SetCharacterSetName(v string) *CreateDatabaseRequest {
	s.CharacterSetName = &v
	return s
}

func (s *CreateDatabaseRequest) SetCollationName(v string) *CreateDatabaseRequest {
	s.CollationName = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBDescription(v string) *CreateDatabaseRequest {
	s.DBDescription = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBInstanceId(v string) *CreateDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBName(v string) *CreateDatabaseRequest {
	s.DBName = &v
	return s
}

func (s *CreateDatabaseRequest) SetOwnerAccount(v string) *CreateDatabaseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDatabaseRequest) SetOwnerId(v int64) *CreateDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDatabaseRequest) SetResourceOwnerAccount(v string) *CreateDatabaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDatabaseRequest) SetResourceOwnerId(v int64) *CreateDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
