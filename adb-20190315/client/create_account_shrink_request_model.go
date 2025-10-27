// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateAccountShrinkRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateAccountShrinkRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateAccountShrinkRequest
	GetAccountPassword() *string
	SetAccountType(v string) *CreateAccountShrinkRequest
	GetAccountType() *string
	SetDBClusterId(v string) *CreateAccountShrinkRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CreateAccountShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAccountShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateAccountShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAccountShrinkRequest
	GetResourceOwnerId() *int64
	SetTagShrink(v string) *CreateAccountShrinkRequest
	GetTagShrink() *string
}

type CreateAccountShrinkRequest struct {
	// The description of the database account.
	//
	// 	- The description cannot start with `http://` or `https://`.
	//
	// 	- The description can be up to 256 characters in length.
	//
	// example:
	//
	// Test account
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// 	- The name must start with a lowercase letter and end with a lowercase letter or a digit.
	//
	// 	- The name can contain lowercase letters, digits, and underscores (_).
	//
	// 	- The name must be 2 to 16 characters in length.
	//
	// 	- Reserved account names such as root, admin, and opsadmin cannot be used.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- Special characters include `! @ # $ % ^ & 	- ( ) _ + - =`
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test_accout1
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **Normal**: standard account. Up to 256 standard accounts can be created for a cluster.
	//
	// 	- **Super*	- (default): privileged account. Only a single privileged account can be created for a cluster.
	//
	// >  If a cluster does not have accounts, you can specify this parameter to create a privileged account or standard account. If a cluster has a privileged account, you must set this parameter to Normal to create a standard account. Otherwise, the operation fails. After an account is created, the privileged account has permissions on all databases of the cluster. The standard account does not have permissions and must be granted permissions on specific databases by the privileged account. For more information, see GRANT.
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to view cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagShrink            *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateAccountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountShrinkRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateAccountShrinkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAccountShrinkRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateAccountShrinkRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAccountShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAccountShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAccountShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccountShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccountShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAccountShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateAccountShrinkRequest) SetAccountDescription(v string) *CreateAccountShrinkRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetAccountName(v string) *CreateAccountShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetAccountPassword(v string) *CreateAccountShrinkRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetAccountType(v string) *CreateAccountShrinkRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetDBClusterId(v string) *CreateAccountShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetOwnerAccount(v string) *CreateAccountShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetOwnerId(v int64) *CreateAccountShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetResourceOwnerAccount(v string) *CreateAccountShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetResourceOwnerId(v int64) *CreateAccountShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountShrinkRequest) SetTagShrink(v string) *CreateAccountShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateAccountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
