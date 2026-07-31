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
	// The database account.
	//
	// > If you do not specify this parameter, information about all database accounts is returned.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
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
	// - **AnalyticDB*	- (default): the AnalyticDB for MySQL engine
	//
	// - **Clickhouse**: the LindormTable engine
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
