// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v *DescribeDatabasesResponseBodyDatabases) *DescribeDatabasesResponseBody
	GetDatabases() *DescribeDatabasesResponseBodyDatabases
	SetPageNumber(v int32) *DescribeDatabasesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDatabasesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDatabasesResponseBody
	GetRequestId() *string
}

type DescribeDatabasesResponseBody struct {
	// Details about databases.
	Databases *DescribeDatabasesResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E1DF8CA6-2300-448B-9ABF-760C4B******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBody) GetDatabases() *DescribeDatabasesResponseBodyDatabases {
	return s.Databases
}

func (s *DescribeDatabasesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabasesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatabasesResponseBody) SetDatabases(v *DescribeDatabasesResponseBodyDatabases) *DescribeDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDatabasesResponseBody) SetPageNumber(v int32) *DescribeDatabasesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabasesResponseBody) SetPageRecordCount(v int32) *DescribeDatabasesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDatabasesResponseBody) SetRequestId(v string) *DescribeDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDatabasesResponseBodyDatabases struct {
	Database []*DescribeDatabasesResponseBodyDatabasesDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabases) GetDatabase() []*DescribeDatabasesResponseBodyDatabasesDatabase {
	return s.Database
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDatabase(v []*DescribeDatabasesResponseBodyDatabasesDatabase) *DescribeDatabasesResponseBodyDatabases {
	s.Database = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}

type DescribeDatabasesResponseBodyDatabasesDatabase struct {
	// Details about the accounts.
	//
	// > A PolarDB for MySQL cluster does not support privileged accounts.
	Accounts *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The character set that the database uses. For more information, see [Character set tables](https://help.aliyun.com/document_detail/99716.html).
	//
	// example:
	//
	// utf8mb4
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	// The description of the database.
	//
	// example:
	//
	// test_des
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test_db
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The state of the database. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Running**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Running
	DBStatus *string `json:"DBStatus,omitempty" xml:"DBStatus,omitempty"`
	// The type of the database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **Oracle**
	//
	// 	- **PostgreSQL**
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The ID of the primary node in the cluster of Multi-master Cluster (Database/Table) Edition.
	//
	// example:
	//
	// 2
	MasterID *string `json:"MasterID,omitempty" xml:"MasterID,omitempty"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabase) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabase) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetAccounts() *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts {
	return s.Accounts
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetCharacterSetName() *string {
	return s.CharacterSetName
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetDBDescription() *string {
	return s.DBDescription
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetDBStatus() *string {
	return s.DBStatus
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetMasterID() *string {
	return s.MasterID
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetAccounts(v *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.Accounts = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetCharacterSetName(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.CharacterSetName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetDBDescription(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.DBDescription = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetDBName(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.DBName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetDBStatus(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.DBStatus = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetEngine(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.Engine = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetMasterID(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.MasterID = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) Validate() error {
	return dara.Validate(s)
}

type DescribeDatabasesResponseBodyDatabasesDatabaseAccounts struct {
	Account []*DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) GetAccount() []*DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount {
	return s.Account
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) SetAccount(v []*DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts {
	s.Account = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) Validate() error {
	return dara.Validate(s)
}

type DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount struct {
	// The username of the account.
	//
	// > A PolarDB for MySQL cluster does not support privileged accounts.
	//
	// example:
	//
	// test_acc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions that are granted to the account. Valid values:
	//
	// 	- **ReadWrite**: read and write permissions
	//
	// 	- **ReadOnly**: read-only permissions
	//
	// 	- **DMLOnly**: The account is granted the permissions to execute only DML statements on the database.
	//
	// 	- **DDLOnly**: The account is granted the permissions to execute only DDL statements on the database.
	//
	// 	- **ReadIndex**: The account has the read and index permissions on the database.
	//
	// example:
	//
	// ReadOnly
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The state of the account. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Available**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Available
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The authorization state of the account. Valid values:
	//
	// 	- **Empowering**: The system is granting permissions to the account.
	//
	// 	- **Empowered**: Permissions are granted to the account.
	//
	// 	- **Removing**: The system is revoking permissions from the account.
	//
	// example:
	//
	// Empowered
	PrivilegeStatus *string `json:"PrivilegeStatus,omitempty" xml:"PrivilegeStatus,omitempty"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) GetPrivilegeStatus() *string {
	return s.PrivilegeStatus
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) SetAccountName(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) SetAccountPrivilege(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) SetAccountStatus(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) SetPrivilegeStatus(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount {
	s.PrivilegeStatus = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) Validate() error {
	return dara.Validate(s)
}
