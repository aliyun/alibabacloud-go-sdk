// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *ResetAccountPasswordRequest
	GetAccountDescription() *string
	SetAccountName(v string) *ResetAccountPasswordRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ResetAccountPasswordRequest
	GetAccountPassword() *string
	SetDBClusterId(v string) *ResetAccountPasswordRequest
	GetDBClusterId() *string
	SetEngine(v string) *ResetAccountPasswordRequest
	GetEngine() *string
	SetResourceGroupName(v string) *ResetAccountPasswordRequest
	GetResourceGroupName() *string
}

type ResetAccountPasswordRequest struct {
	// The description of the account.
	//
	// - The description cannot start with `http://` or `https://`.
	//
	// - The description must be 2 to 256 characters in length.
	//
	// example:
	//
	// AccDesc
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The database account.
	//
	// > You can call the [DescribeAccounts](https://help.aliyun.com/document_detail/612430.html) operation to query the database account information of a specified cluster, including the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - The following special characters are supported: `!@#$%^&*()_+-=`
	//
	// - The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test_accout1
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// <props="china">The ID of the cluster. The cluster can be an Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine. Valid values:
	//
	// - **AnalyticDB*	- (default): the AnalyticDB for MySQL engine.
	//
	// - **Clickhouse**: the wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine            *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *ResetAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ResetAccountPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetAccountPasswordRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ResetAccountPasswordRequest) GetEngine() *string {
	return s.Engine
}

func (s *ResetAccountPasswordRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ResetAccountPasswordRequest) SetAccountDescription(v string) *ResetAccountPasswordRequest {
	s.AccountDescription = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBClusterId(v string) *ResetAccountPasswordRequest {
	s.DBClusterId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetEngine(v string) *ResetAccountPasswordRequest {
	s.Engine = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetResourceGroupName(v string) *ResetAccountPasswordRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ResetAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
