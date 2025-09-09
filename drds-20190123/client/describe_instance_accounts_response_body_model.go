// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceAccounts(v *DescribeInstanceAccountsResponseBodyInstanceAccounts) *DescribeInstanceAccountsResponseBody
	GetInstanceAccounts() *DescribeInstanceAccountsResponseBodyInstanceAccounts
	SetRequestId(v string) *DescribeInstanceAccountsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceAccountsResponseBody
	GetSuccess() *bool
}

type DescribeInstanceAccountsResponseBody struct {
	// Indicates the information about the instance accounts.
	InstanceAccounts *DescribeInstanceAccountsResponseBodyInstanceAccounts `json:"InstanceAccounts,omitempty" xml:"InstanceAccounts,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E2E4056D-57EB-4353-8355-2E6284******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBody) GetInstanceAccounts() *DescribeInstanceAccountsResponseBodyInstanceAccounts {
	return s.InstanceAccounts
}

func (s *DescribeInstanceAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAccountsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceAccountsResponseBody) SetInstanceAccounts(v *DescribeInstanceAccountsResponseBodyInstanceAccounts) *DescribeInstanceAccountsResponseBody {
	s.InstanceAccounts = v
	return s
}

func (s *DescribeInstanceAccountsResponseBody) SetRequestId(v string) *DescribeInstanceAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBody) SetSuccess(v bool) *DescribeInstanceAccountsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAccountsResponseBodyInstanceAccounts struct {
	InstanceAccount []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount `json:"InstanceAccount,omitempty" xml:"InstanceAccount,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccounts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccounts) GetInstanceAccount() []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	return s.InstanceAccount
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccounts) SetInstanceAccount(v []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) *DescribeInstanceAccountsResponseBodyInstanceAccounts {
	s.InstanceAccount = v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccounts) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount struct {
	// Indicates the username of an instance account.
	//
	// example:
	//
	// test_rds3
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates the type of an instance account. Valid values:
	//
	// 	- **0**: The instance account is a privileged account.
	//
	// 	- **1**: The instance account is a standard account.
	//
	// example:
	//
	// 1
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// Indicates the information about the permissions of an account on a database.
	DbPrivileges *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges `json:"DbPrivileges,omitempty" xml:"DbPrivileges,omitempty" type:"Struct"`
	// Indicates the description of an account. By default, if 0 is the value of the AccountType parameter, **Created by DRDS*	- is returned as the value of the Description parameter. If 1 is the value of the AccountType parameter, an empty string is returned as the value of the Description parameter. You can modify the description of an account on the Accounts page in the PolarDB-X console.
	//
	// example:
	//
	// Created by DRDS
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates an IP address that is allowed to access the database. The value **%*	- indicates that each IP address is allowed to access the database. \\</note>
	//
	// example:
	//
	// %
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) GetAccountType() *int32 {
	return s.AccountType
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) GetDbPrivileges() *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges {
	return s.DbPrivileges
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) GetHost() *string {
	return s.Host
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetAccountName(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetAccountType(v int32) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetDbPrivileges(v *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.DbPrivileges = v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetDescription(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) SetHost(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount {
	s.Host = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccount) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges struct {
	DbPrivilege []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege `json:"DbPrivilege,omitempty" xml:"DbPrivilege,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) GetDbPrivilege() []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege {
	return s.DbPrivilege
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) SetDbPrivilege(v []*DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges {
	s.DbPrivilege = v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivileges) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege struct {
	// Indicates the name of a database.
	//
	// example:
	//
	// test_rds3
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates the permissions that an account is granted on the database. Valid values:
	//
	// 	- **R**: The account is granted the permissions that are required to read the data of the database.
	//
	// 	- **W**: The account is granted the permissions that are required to write data to the database.
	//
	// 	- **DDL**: The account is granted the permissions that are required to perform DDL operations on the database.
	//
	// 	- **DML**: The account is granted the permissions that are required to perform DML operations on the database.
	//
	// example:
	//
	// R
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) GetDbName() *string {
	return s.DbName
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) GetPrivilege() *string {
	return s.Privilege
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) SetDbName(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege {
	s.DbName = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) SetPrivilege(v string) *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege {
	s.Privilege = &v
	return s
}

func (s *DescribeInstanceAccountsResponseBodyInstanceAccountsInstanceAccountDbPrivilegesDbPrivilege) Validate() error {
	return dara.Validate(s)
}
