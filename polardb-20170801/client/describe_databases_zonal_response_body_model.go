// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabasesZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*DescribeDatabasesZonalResponseBodyDatabases) *DescribeDatabasesZonalResponseBody
	GetDatabases() []*DescribeDatabasesZonalResponseBodyDatabases
	SetMaxResults(v int32) *DescribeDatabasesZonalResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDatabasesZonalResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDatabasesZonalResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDatabasesZonalResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDatabasesZonalResponseBody
	GetRequestId() *string
}

type DescribeDatabasesZonalResponseBody struct {
	// The details of the databases.
	Databases []*DescribeDatabasesZonalResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The maximum number of entries returned for the current request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. This is the NextToken value from the previous API call. If there are no more results, do not specify this parameter.
	//
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 7
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDatabasesZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesZonalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesZonalResponseBody) GetDatabases() []*DescribeDatabasesZonalResponseBodyDatabases {
	return s.Databases
}

func (s *DescribeDatabasesZonalResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDatabasesZonalResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDatabasesZonalResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabasesZonalResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDatabasesZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatabasesZonalResponseBody) SetDatabases(v []*DescribeDatabasesZonalResponseBodyDatabases) *DescribeDatabasesZonalResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDatabasesZonalResponseBody) SetMaxResults(v int32) *DescribeDatabasesZonalResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBody) SetNextToken(v string) *DescribeDatabasesZonalResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBody) SetPageNumber(v int32) *DescribeDatabasesZonalResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBody) SetPageRecordCount(v int32) *DescribeDatabasesZonalResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBody) SetRequestId(v string) *DescribeDatabasesZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBody) Validate() error {
	if s.Databases != nil {
		for _, item := range s.Databases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDatabasesZonalResponseBodyDatabases struct {
	// The details of the database accounts.
	//
	// > If the cluster is a PolarDB for MySQL cluster, privileged accounts are not included.
	Accounts []*DescribeDatabasesZonalResponseBodyDatabasesAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The character set.
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
	// The status of the database. Valid values:
	//
	// - **Creating**
	//
	// - **Running**
	//
	// - **Deleting**
	//
	// example:
	//
	// Running
	DBStatus *string `json:"DBStatus,omitempty" xml:"DBStatus,omitempty"`
	// The database engine type. Valid values:
	//
	// - **MySQL**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The ID of the primary node that corresponds to the database in a Multi-master Cluster (Database/Table) edition cluster.
	//
	// example:
	//
	// 2
	MasterID *string `json:"MasterID,omitempty" xml:"MasterID,omitempty"`
}

func (s DescribeDatabasesZonalResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesZonalResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) GetAccounts() []*DescribeDatabasesZonalResponseBodyDatabasesAccounts {
	return s.Accounts
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) GetCharacterSetName() *string {
	return s.CharacterSetName
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) GetDBDescription() *string {
	return s.DBDescription
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) GetDBStatus() *string {
	return s.DBStatus
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) GetMasterID() *string {
	return s.MasterID
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) SetAccounts(v []*DescribeDatabasesZonalResponseBodyDatabasesAccounts) *DescribeDatabasesZonalResponseBodyDatabases {
	s.Accounts = v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) SetCharacterSetName(v string) *DescribeDatabasesZonalResponseBodyDatabases {
	s.CharacterSetName = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) SetDBDescription(v string) *DescribeDatabasesZonalResponseBodyDatabases {
	s.DBDescription = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) SetDBName(v string) *DescribeDatabasesZonalResponseBodyDatabases {
	s.DBName = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) SetDBStatus(v string) *DescribeDatabasesZonalResponseBodyDatabases {
	s.DBStatus = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) SetEngine(v string) *DescribeDatabasesZonalResponseBodyDatabases {
	s.Engine = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) SetMasterID(v string) *DescribeDatabasesZonalResponseBodyDatabases {
	s.MasterID = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabases) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDatabasesZonalResponseBodyDatabasesAccounts struct {
	// The account name.
	//
	// > If the cluster is a PolarDB for MySQL cluster, privileged accounts are not included.
	//
	// example:
	//
	// test_acc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions of the account. Valid values:
	//
	// - **ReadWrite**
	//
	// - **ReadOnly**
	//
	// - **DMLOnly**
	//
	// - **DDLOnly**
	//
	// - **ReadIndex**
	//
	// example:
	//
	// ReadOnly
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The status of the account. Valid values:
	//
	// - **Creating**
	//
	// - **Available**
	//
	// - **Deleting**
	//
	// example:
	//
	// Available
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The authorization status. Valid values:
	//
	// - **Empowering**: Permissions are being granted.
	//
	// - **Empowered**: Permissions are granted.
	//
	// - **Removing**: Permissions are being revoked.
	//
	// example:
	//
	// Empowered
	PrivilegeStatus *string `json:"PrivilegeStatus,omitempty" xml:"PrivilegeStatus,omitempty"`
}

func (s DescribeDatabasesZonalResponseBodyDatabasesAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesZonalResponseBodyDatabasesAccounts) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesZonalResponseBodyDatabasesAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeDatabasesZonalResponseBodyDatabasesAccounts) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *DescribeDatabasesZonalResponseBodyDatabasesAccounts) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeDatabasesZonalResponseBodyDatabasesAccounts) GetPrivilegeStatus() *string {
	return s.PrivilegeStatus
}

func (s *DescribeDatabasesZonalResponseBodyDatabasesAccounts) SetAccountName(v string) *DescribeDatabasesZonalResponseBodyDatabasesAccounts {
	s.AccountName = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabasesAccounts) SetAccountPrivilege(v string) *DescribeDatabasesZonalResponseBodyDatabasesAccounts {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabasesAccounts) SetAccountStatus(v string) *DescribeDatabasesZonalResponseBodyDatabasesAccounts {
	s.AccountStatus = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabasesAccounts) SetPrivilegeStatus(v string) *DescribeDatabasesZonalResponseBodyDatabasesAccounts {
	s.PrivilegeStatus = &v
	return s
}

func (s *DescribeDatabasesZonalResponseBodyDatabasesAccounts) Validate() error {
	return dara.Validate(s)
}
