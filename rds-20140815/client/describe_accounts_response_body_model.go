// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody
	GetAccounts() *DescribeAccountsResponseBodyAccounts
	SetPageNumber(v int32) *DescribeAccountsResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeAccountsResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeAccountsResponseBody
	GetResourceGroupId() *string
	SetSystemAdminAccountFirstActivationTime(v string) *DescribeAccountsResponseBody
	GetSystemAdminAccountFirstActivationTime() *string
	SetSystemAdminAccountStatus(v string) *DescribeAccountsResponseBody
	GetSystemAdminAccountStatus() *string
	SetTotalRecordCount(v int32) *DescribeAccountsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeAccountsResponseBody struct {
	Accounts                              *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	PageNumber                            *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId                             *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId                       *string                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SystemAdminAccountFirstActivationTime *string                               `json:"SystemAdminAccountFirstActivationTime,omitempty" xml:"SystemAdminAccountFirstActivationTime,omitempty"`
	SystemAdminAccountStatus              *string                               `json:"SystemAdminAccountStatus,omitempty" xml:"SystemAdminAccountStatus,omitempty"`
	TotalRecordCount                      *int32                                `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) GetAccounts() *DescribeAccountsResponseBodyAccounts {
	return s.Accounts
}

func (s *DescribeAccountsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountsResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAccountsResponseBody) GetSystemAdminAccountFirstActivationTime() *string {
	return s.SystemAdminAccountFirstActivationTime
}

func (s *DescribeAccountsResponseBody) GetSystemAdminAccountStatus() *string {
	return s.SystemAdminAccountStatus
}

func (s *DescribeAccountsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeAccountsResponseBody) SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeAccountsResponseBody) SetPageNumber(v int32) *DescribeAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetResourceGroupId(v string) *DescribeAccountsResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetSystemAdminAccountFirstActivationTime(v string) *DescribeAccountsResponseBody {
	s.SystemAdminAccountFirstActivationTime = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetSystemAdminAccountStatus(v string) *DescribeAccountsResponseBody {
	s.SystemAdminAccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetTotalRecordCount(v int32) *DescribeAccountsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAccountsResponseBody) Validate() error {
	if s.Accounts != nil {
		if err := s.Accounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccounts struct {
	DBInstanceAccount []*DescribeAccountsResponseBodyAccountsDBInstanceAccount `json:"DBInstanceAccount,omitempty" xml:"DBInstanceAccount,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccounts) GetDBInstanceAccount() []*DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	return s.DBInstanceAccount
}

func (s *DescribeAccountsResponseBodyAccounts) SetDBInstanceAccount(v []*DescribeAccountsResponseBodyAccountsDBInstanceAccount) *DescribeAccountsResponseBodyAccounts {
	s.DBInstanceAccount = v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) Validate() error {
	if s.DBInstanceAccount != nil {
		for _, item := range s.DBInstanceAccount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccountsDBInstanceAccount struct {
	AccountDescription *string                                                                  `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string                                                                  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountStatus      *string                                                                  `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountType        *string                                                                  `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	BypassRLS          *string                                                                  `json:"BypassRLS,omitempty" xml:"BypassRLS,omitempty"`
	CheckPolicy        *bool                                                                    `json:"CheckPolicy,omitempty" xml:"CheckPolicy,omitempty"`
	CreateDB           *string                                                                  `json:"CreateDB,omitempty" xml:"CreateDB,omitempty"`
	CreateRole         *string                                                                  `json:"CreateRole,omitempty" xml:"CreateRole,omitempty"`
	DBInstanceId       *string                                                                  `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DatabasePrivileges *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivileges `json:"DatabasePrivileges,omitempty" xml:"DatabasePrivileges,omitempty" type:"Struct"`
	PasswordExpireTime *string                                                                  `json:"PasswordExpireTime,omitempty" xml:"PasswordExpireTime,omitempty"`
	PrivExceeded       *string                                                                  `json:"PrivExceeded,omitempty" xml:"PrivExceeded,omitempty"`
	Replication        *string                                                                  `json:"Replication,omitempty" xml:"Replication,omitempty"`
	ValidUntil         *string                                                                  `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccount) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetBypassRLS() *string {
	return s.BypassRLS
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetCheckPolicy() *bool {
	return s.CheckPolicy
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetCreateDB() *string {
	return s.CreateDB
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetCreateRole() *string {
	return s.CreateRole
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetDatabasePrivileges() *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivileges {
	return s.DatabasePrivileges
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetPasswordExpireTime() *string {
	return s.PasswordExpireTime
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetPrivExceeded() *string {
	return s.PrivExceeded
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetReplication() *string {
	return s.Replication
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) GetValidUntil() *string {
	return s.ValidUntil
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetAccountType(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetBypassRLS(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.BypassRLS = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetCheckPolicy(v bool) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.CheckPolicy = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetCreateDB(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.CreateDB = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetCreateRole(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.CreateRole = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetDBInstanceId(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetDatabasePrivileges(v *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivileges) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.DatabasePrivileges = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetPasswordExpireTime(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.PasswordExpireTime = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetPrivExceeded(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.PrivExceeded = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetReplication(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.Replication = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) SetValidUntil(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccount {
	s.ValidUntil = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccount) Validate() error {
	if s.DatabasePrivileges != nil {
		if err := s.DatabasePrivileges.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivileges struct {
	DatabasePrivilege []*DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege `json:"DatabasePrivilege,omitempty" xml:"DatabasePrivilege,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivileges) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivileges) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivileges) GetDatabasePrivilege() []*DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege {
	return s.DatabasePrivilege
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivileges) SetDatabasePrivilege(v []*DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege) *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivileges {
	s.DatabasePrivilege = v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivileges) Validate() error {
	if s.DatabasePrivilege != nil {
		for _, item := range s.DatabasePrivilege {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege struct {
	AccountPrivilege       *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	AccountPrivilegeDetail *string `json:"AccountPrivilegeDetail,omitempty" xml:"AccountPrivilegeDetail,omitempty"`
	DBName                 *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege) GetAccountPrivilegeDetail() *string {
	return s.AccountPrivilegeDetail
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege) SetAccountPrivilege(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege) SetAccountPrivilegeDetail(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege {
	s.AccountPrivilegeDetail = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege) SetDBName(v string) *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege {
	s.DBName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDBInstanceAccountDatabasePrivilegesDatabasePrivilege) Validate() error {
	return dara.Validate(s)
}
