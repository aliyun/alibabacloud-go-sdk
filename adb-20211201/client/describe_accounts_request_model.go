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
	SetDBClusterId(v string) *DescribeAccountsRequest
	GetDBClusterId() *string
	SetEngine(v string) *DescribeAccountsRequest
	GetEngine() *string
	SetOwnerId(v string) *DescribeAccountsRequest
	GetOwnerId() *string
}

type DescribeAccountsRequest struct {
	// The name of the database account.
	//
	// > If you do not specify this parameter, the information about all database accounts in the cluster is returned.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine of the cluster. Valid values:
	//
	// 	- **AnalyticDB*	- (default): the AnalyticDB for MySQL engine.
	//
	// 	- **Clickhouse**: the wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine  *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeAccountsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAccountsRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAccountsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBClusterId(v string) *DescribeAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountsRequest) SetEngine(v string) *DescribeAccountsRequest {
	s.Engine = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerId(v string) *DescribeAccountsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountsRequest) Validate() error {
	return dara.Validate(s)
}
