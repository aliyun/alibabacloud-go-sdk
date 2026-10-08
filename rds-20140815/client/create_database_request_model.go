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
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The character set. Valid values:
	//
	// 	- MySQL/MariaDB: **utf8, gbk, latin1, utf8mb4**
	//
	// 	- SQL Server: **Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN**
	//
	// 	- PostgreSQL: You must specify the character set, Collate, and Ctype in the format of `Character set,<Collate>,<Ctype>`. Example: `UTF8,C,en_US.utf8`.
	//
	//     - Valid values for the character set: **KOI8U, UTF8, WIN866, WIN874, WIN1250, WIN1251, WIN1252, WIN1253, WIN1254, WIN1255, WIN1256, WIN1257, WIN1258, EUC_CN, EUC_KR, EUC_TW, EUC_JP, EUC_JIS_2004, KOI8R, MULE_INTERNAL, LATIN1, LATIN2, LATIN3, LATIN4, LATIN5, LATIN6, LATIN7, LATIN8, LATIN9, LATIN10, ISO_8859_5, ISO_8859_6, ISO_8859_7, ISO_8859_8, SQL_ASCII**.
	//
	//     - Valid values for **Collate**: You can run the `SELECT DISTINCT collname FROM pg_collation;` command to query the valid values. If this parameter is not specified, the default value **C*	- is used.
	//
	//     - Valid values for **Ctype**: You can run the `SELECT DISTINCT collctype FROM pg_collation;` command to query the valid values. If this parameter is not specified, the default value **en_US.utf8*	- is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbk
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	// The collation. This parameter is supported only for ApsaraDB RDS for MySQL instances. Specify a collation that matches the character set. For example, if the character set is utf8mb4, the collation must be utf8mb4_bin or utf8mb4_general_ci.
	//
	// example:
	//
	// gbk_chinese_ci
	CollationName *string `json:"CollationName,omitempty" xml:"CollationName,omitempty"`
	// The database description. The description must be 2 to 256 characters in length and can contain letters, digits, Chinese characters, underscores (_), and hyphens (-). The description must start with a Chinese character or a letter.
	//
	// >The description cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testdb
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name.
	//
	// > 	- The name must be 2 to 64 characters in length.
	//
	// > 	- The name must start with a letter and end with a letter or digit.
	//
	// > 	- The name can contain lowercase letters, digits, underscores (_), and hyphens (-).
	//
	// > 	- The database name must be unique within the instance.
	//
	// > 	- For more information about invalid characters, see [Reserved words](https://help.aliyun.com/document_detail/26317.html).
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
