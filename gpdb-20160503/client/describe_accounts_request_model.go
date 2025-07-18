// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeAccountsRequest
	GetAccountName() *string
	SetAccountType(v string) *DescribeAccountsRequest
	GetAccountType() *string
	SetDBInstanceId(v string) *DescribeAccountsRequest
	GetDBInstanceId() *string
}

type DescribeAccountsRequest struct {
	// The name of the database account.
	//
	// example:
	//
	// testuser
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of the database account that you want to query. Valid values: Super, Normal, and All. Default value: Super. If you set this parameter to Super, the information about the privileged account is returned. If you set this parameter to Normal, the information about all standard accounts is returned. If you set this parameter to All, the information about all database accounts is returned.
	//
	// example:
	//
	// Super
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountsRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeAccountsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountType(v string) *DescribeAccountsRequest {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBInstanceId(v string) *DescribeAccountsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsRequest) Validate() error {
	return dara.Validate(s)
}
