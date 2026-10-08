// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateAccountRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateAccountRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateAccountRequest
	GetAccountPassword() *string
	SetAccountType(v string) *CreateAccountRequest
	GetAccountType() *string
	SetCheckPolicy(v bool) *CreateAccountRequest
	GetCheckPolicy() *bool
	SetDBInstanceId(v string) *CreateAccountRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *CreateAccountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAccountRequest
	GetResourceOwnerId() *int64
}

type CreateAccountRequest struct {
	// The description of the account. The description must be 2 to 256 characters in length. It must start with a letter or a Chinese character and can contain digits, Chinese characters, letters, underscores (_), and hyphens (-).
	//
	// >The description cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testuser
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// > The name must be unique and can contain uppercase letters (supported only by MySQL), lowercase letters, digits, or underscores. For specific naming conventions, refer to the tutorials for each engine: [Create a MySQL account](https://help.aliyun.com/document_detail/96089.html), [Create a PostgreSQL account](https://help.aliyun.com/document_detail/96753.html), [Create a SQL Server account](https://help.aliyun.com/document_detail/95810.html), [Create a MariaDB account](https://help.aliyun.com/document_detail/97132.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// > 	- The password must be 8 to 32 characters in length.
	//
	// > 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters (`!@#$%^&*()_+-=`).
	//
	// This parameter is required.
	//
	// example:
	//
	// Test123456
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The type of the account. Valid values:
	//
	// - **Normal*	- (default): standard account.
	//
	// - **Super**: privileged account. You can create at most one privileged account per instance.
	//
	// - **Sysadmin*	- (SQL Server instances only): database account with SA permissions. Before you create this account, check whether the instance meets the [prerequisites](https://help.aliyun.com/document_detail/170736.html).
	//
	// - **GlobalRO*	- (SQL Server instances only): global read-only account. You can create at most two global read-only accounts per instance. The database engine version of the instance must be SQL Server 2016 or later, and the instance type must be dedicated or general-purpose.
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The [account password policy](https://help.aliyun.com/document_detail/2845728.html) for the SQL Server instance. Valid values:
	//
	// - **true**: The policy is applied.
	//
	// - **false**: The policy is not applied.
	//
	// > - If you set this parameter to true, you must first [configure the SQL Server account password policy](https://help.aliyun.com/document_detail/2848317.html).
	//
	// > - This parameter does not support SQL Server instances of the [shared instance type](https://help.aliyun.com/document_detail/57184.html), [2008 R2 edition](https://help.aliyun.com/document_detail/145468.html), or [serverless type](https://help.aliyun.com/document_detail/603466.html).
	//
	// example:
	//
	// true
	CheckPolicy *bool `json:"CheckPolicy,omitempty" xml:"CheckPolicy,omitempty"`
	// The instance ID. You can call [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateAccountRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAccountRequest) GetCheckPolicy() *bool {
	return s.CheckPolicy
}

func (s *CreateAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateAccountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetCheckPolicy(v bool) *CreateAccountRequest {
	s.CheckPolicy = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceId(v string) *CreateAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerAccount(v string) *CreateAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerAccount(v string) *CreateAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerId(v int64) *CreateAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	return dara.Validate(s)
}
