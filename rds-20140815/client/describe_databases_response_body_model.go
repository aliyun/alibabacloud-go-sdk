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
	SetRequestId(v string) *DescribeDatabasesResponseBody
	GetRequestId() *string
}

type DescribeDatabasesResponseBody struct {
	// The information about the databases.
	Databases *DescribeDatabasesResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2603CA96-B17D-4903-BC04-61A2C829CD94
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

func (s *DescribeDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatabasesResponseBody) SetDatabases(v *DescribeDatabasesResponseBodyDatabases) *DescribeDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDatabasesResponseBody) SetRequestId(v string) *DescribeDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabasesResponseBody) Validate() error {
	if s.Databases != nil {
		if err := s.Databases.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.Database != nil {
		for _, item := range s.Database {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDatabasesResponseBodyDatabasesDatabase struct {
	// The information about the account. Each account has specific permissions on the database.
	Accounts *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The advanced information about the database.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	AdvancedInfo *DescribeDatabasesResponseBodyDatabasesDatabaseAdvancedInfo `json:"AdvancedInfo,omitempty" xml:"AdvancedInfo,omitempty" type:"Struct"`
	// The basic information about the database.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	BasicInfo *DescribeDatabasesResponseBodyDatabasesDatabaseBasicInfo `json:"BasicInfo,omitempty" xml:"BasicInfo,omitempty" type:"Struct"`
	// The name of the character set.
	//
	// example:
	//
	// utf8
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	// The collation of the character set. The example value C indicates localization.
	//
	// >  This parameter is returned only for instances that run PostgreSQL.
	//
	// example:
	//
	// C
	Collate *string `json:"Collate,omitempty" xml:"Collate,omitempty"`
	// The limit on the number of concurrent requests. The value -1 indicates that the number of concurrent requests is unlimited.
	//
	// >  This parameter is returned only for instances that run PostgreSQL.
	//
	// example:
	//
	// -1
	ConnLimit *string `json:"ConnLimit,omitempty" xml:"ConnLimit,omitempty"`
	// The type of the character set.
	//
	// >  This parameter is returned only for instances that run PostgreSQL.
	//
	// example:
	//
	// en_US.utf8
	Ctype *string `json:"Ctype,omitempty" xml:"Ctype,omitempty"`
	// The description of the database.
	//
	// example:
	//
	// testdb
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// The ID of the instance to which the database belongs.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name.
	//
	// example:
	//
	// testDB01
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The database status. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Running**
	//
	// 	- **Deleting**
	//
	// 	- **Cold**
	//
	// example:
	//
	// Creating
	DBStatus      *string `json:"DBStatus,omitempty" xml:"DBStatus,omitempty"`
	DuckDBEnabled *bool   `json:"DuckDBEnabled,omitempty" xml:"DuckDBEnabled,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The runtime information about the database.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	RuntimeInfo *DescribeDatabasesResponseBodyDatabasesDatabaseRuntimeInfo `json:"RuntimeInfo,omitempty" xml:"RuntimeInfo,omitempty" type:"Struct"`
	// The database tablespace.
	//
	// >  This parameter is returned only for instances that run PostgreSQL.
	//
	// example:
	//
	// pg_default
	Tablespace *string `json:"Tablespace,omitempty" xml:"Tablespace,omitempty"`
	// The total number of entries returned.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetAdvancedInfo() *DescribeDatabasesResponseBodyDatabasesDatabaseAdvancedInfo {
	return s.AdvancedInfo
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetBasicInfo() *DescribeDatabasesResponseBodyDatabasesDatabaseBasicInfo {
	return s.BasicInfo
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetCharacterSetName() *string {
	return s.CharacterSetName
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetCollate() *string {
	return s.Collate
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetConnLimit() *string {
	return s.ConnLimit
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetCtype() *string {
	return s.Ctype
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetDBDescription() *string {
	return s.DBDescription
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetDBStatus() *string {
	return s.DBStatus
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetDuckDBEnabled() *bool {
	return s.DuckDBEnabled
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetRuntimeInfo() *DescribeDatabasesResponseBodyDatabasesDatabaseRuntimeInfo {
	return s.RuntimeInfo
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetTablespace() *string {
	return s.Tablespace
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetAccounts(v *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.Accounts = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetAdvancedInfo(v *DescribeDatabasesResponseBodyDatabasesDatabaseAdvancedInfo) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.AdvancedInfo = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetBasicInfo(v *DescribeDatabasesResponseBodyDatabasesDatabaseBasicInfo) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.BasicInfo = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetCharacterSetName(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.CharacterSetName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetCollate(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.Collate = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetConnLimit(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.ConnLimit = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetCtype(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.Ctype = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetDBDescription(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.DBDescription = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetDBInstanceId(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.DBInstanceId = &v
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

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetDuckDBEnabled(v bool) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.DuckDBEnabled = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetEngine(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.Engine = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetPageNumber(v int32) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetPageSize(v int32) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetRuntimeInfo(v *DescribeDatabasesResponseBodyDatabasesDatabaseRuntimeInfo) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.RuntimeInfo = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetTablespace(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.Tablespace = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetTotalCount(v int32) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.TotalCount = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) Validate() error {
	if s.Accounts != nil {
		if err := s.Accounts.Validate(); err != nil {
			return err
		}
	}
	if s.AdvancedInfo != nil {
		if err := s.AdvancedInfo.Validate(); err != nil {
			return err
		}
	}
	if s.BasicInfo != nil {
		if err := s.BasicInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeInfo != nil {
		if err := s.RuntimeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDatabasesResponseBodyDatabasesDatabaseAccounts struct {
	AccountPrivilegeInfo []*DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo `json:"AccountPrivilegeInfo,omitempty" xml:"AccountPrivilegeInfo,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) GetAccountPrivilegeInfo() []*DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo {
	return s.AccountPrivilegeInfo
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) SetAccountPrivilegeInfo(v []*DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo) *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts {
	s.AccountPrivilegeInfo = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) Validate() error {
	if s.AccountPrivilegeInfo != nil {
		for _, item := range s.AccountPrivilegeInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo struct {
	// The account username.
	//
	// example:
	//
	// test
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The permissions that the account has on the database. Valid values:
	//
	// 	- **ReadWrite**: read and write permissions
	//
	// 	- **ReadOnly**: read-only permissions
	//
	// 	- **DMLOnly**: DML-only permissions
	//
	// 	- **DDLOnly**: DDL-only permissions
	//
	// example:
	//
	// DMLOnly
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The permission that the account has on the database.
	//
	// example:
	//
	// SELECT
	AccountPrivilegeDetail *string `json:"AccountPrivilegeDetail,omitempty" xml:"AccountPrivilegeDetail,omitempty"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo) GetAccount() *string {
	return s.Account
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo) GetAccountPrivilegeDetail() *string {
	return s.AccountPrivilegeDetail
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo) SetAccount(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo {
	s.Account = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo) SetAccountPrivilege(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo) SetAccountPrivilegeDetail(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo {
	s.AccountPrivilegeDetail = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccountPrivilegeInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDatabasesResponseBodyDatabasesDatabaseAdvancedInfo struct {
	AdvancedDbProperty []map[string]interface{} `json:"AdvancedDbProperty,omitempty" xml:"AdvancedDbProperty,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAdvancedInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAdvancedInfo) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAdvancedInfo) GetAdvancedDbProperty() []map[string]interface{} {
	return s.AdvancedDbProperty
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAdvancedInfo) SetAdvancedDbProperty(v []map[string]interface{}) *DescribeDatabasesResponseBodyDatabasesDatabaseAdvancedInfo {
	s.AdvancedDbProperty = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAdvancedInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDatabasesResponseBodyDatabasesDatabaseBasicInfo struct {
	BasicDbProperty []map[string]interface{} `json:"BasicDbProperty,omitempty" xml:"BasicDbProperty,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseBasicInfo) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseBasicInfo) GetBasicDbProperty() []map[string]interface{} {
	return s.BasicDbProperty
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseBasicInfo) SetBasicDbProperty(v []map[string]interface{}) *DescribeDatabasesResponseBodyDatabasesDatabaseBasicInfo {
	s.BasicDbProperty = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseBasicInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDatabasesResponseBodyDatabasesDatabaseRuntimeInfo struct {
	RuntimeDbProperty []map[string]interface{} `json:"RuntimeDbProperty,omitempty" xml:"RuntimeDbProperty,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseRuntimeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseRuntimeInfo) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseRuntimeInfo) GetRuntimeDbProperty() []map[string]interface{} {
	return s.RuntimeDbProperty
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseRuntimeInfo) SetRuntimeDbProperty(v []map[string]interface{}) *DescribeDatabasesResponseBodyDatabasesDatabaseRuntimeInfo {
	s.RuntimeDbProperty = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseRuntimeInfo) Validate() error {
	return dara.Validate(s)
}
