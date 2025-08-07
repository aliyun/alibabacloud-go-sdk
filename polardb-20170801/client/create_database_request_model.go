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
	SetCollate(v string) *CreateDatabaseRequest
	GetCollate() *string
	SetCtype(v string) *CreateDatabaseRequest
	GetCtype() *string
	SetDBClusterId(v string) *CreateDatabaseRequest
	GetDBClusterId() *string
	SetDBDescription(v string) *CreateDatabaseRequest
	GetDBDescription() *string
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
	// The name of the account that is authorized to access the database. You can call the [DescribeAccounts](https://help.aliyun.com/document_detail/98107.html) operation to query account information.
	//
	// >- You can specify only a standard account. By default, privileged accounts have all permissions on all databases. You do not need to grant privileged accounts the permissions to access the database.
	//
	// >- This parameter is required for PolarDB for PostgreSQL (Compatible with Oracle) clusters or PolarDB for PostgreSQL clusters. This parameter is optional for PolarDB for MySQL clusters.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions that are granted to the account. Valid values:
	//
	// 	- **ReadWrite**: read and write permissions.
	//
	// 	- **ReadOnly**: read-only permissions.
	//
	// 	- **DMLOnly**: permissions only to execute DML statements on the database.
	//
	// 	- **DDLOnly**: permissions only to execute DDL statements on the database.
	//
	// 	- **ReadIndex**: read-only and index permissions.
	//
	// The default value is **ReadWrite**.
	//
	// >
	//
	// 	- This parameter is valid only when the **AccountName*	- parameter is specified.
	//
	// 	- For a PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL cluster, this parameter is optional. If **AccountName*	- is specified, it is the account of the database owner.
	//
	// 	- For a PolarDB for MySQL cluster, this parameter is optional.
	//
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The character set that is used by the cluster. For more information, see [Character set tables](https://help.aliyun.com/document_detail/99716.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// utf8
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	// The language that defines the collation rules in the database.
	//
	// >
	//
	// 	- The language must be compatible with the character set that is specified by **CharacterSetName**.
	//
	// 	- This parameter is required for a PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL cluster. This parameter is optional for a PolarDB for MySQL cluster. To view the valid values of this parameter, perform the following steps: Log on to the PolarDB console and click the ID of the cluster. In the left-side navigation pane, choose **Settings and Management*	- > **Databases**. Then, click **Create Database**.
	//
	// example:
	//
	// C
	Collate *string `json:"Collate,omitempty" xml:"Collate,omitempty"`
	// The language that indicates the character type of the database.
	//
	// >- The language must be compatible with the character set that is specified by **CharacterSetName**.
	//
	// >- The value that you specify must be the same as the value of **Collate**.
	//
	// >- This parameter is required for PolarDB for PostgreSQL (Compatible with Oracle) clusters or PolarDB for PostgreSQL clusters. This parameter is optional for PolarDB for MySQL clusters.
	//
	// To view the valid values for this parameter, perform the following steps: Log on to the PolarDB console and click the ID of a cluster. In the left-side navigation pane, choose **Settings and Management*	- > **Databases**. Then, click **Create Database**.
	//
	// example:
	//
	// C
	Ctype *string `json:"Ctype,omitempty" xml:"Ctype,omitempty"`
	// The ID of cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the database. The description must meet the following requirements:
	//
	// 	- It cannot start with `http://` or `https://`.
	//
	// 	- It must be 2 to 256 characters in length.
	//
	// > This parameter is required for a PolarDB for Oracle or PolarDB for PostgreSQL cluster. This parameter is optional for a PolarDB for MySQL cluster.
	//
	// example:
	//
	// testdesc
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// The name of the database. The name must meet the following requirements:
	//
	// 	- The name can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must start with a lowercase letter and end with a lowercase letter or a digit. The name must be 1 to 64 characters in length.
	//
	// > Do not use reserved words as database names, such as `test` or `mysql`.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDB
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

func (s *CreateDatabaseRequest) GetCollate() *string {
	return s.Collate
}

func (s *CreateDatabaseRequest) GetCtype() *string {
	return s.Ctype
}

func (s *CreateDatabaseRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDatabaseRequest) GetDBDescription() *string {
	return s.DBDescription
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

func (s *CreateDatabaseRequest) SetCollate(v string) *CreateDatabaseRequest {
	s.Collate = &v
	return s
}

func (s *CreateDatabaseRequest) SetCtype(v string) *CreateDatabaseRequest {
	s.Ctype = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBClusterId(v string) *CreateDatabaseRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBDescription(v string) *CreateDatabaseRequest {
	s.DBDescription = &v
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
