// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DeleteAccountRequest
	GetAccountName() *string
	SetDBClusterId(v string) *DeleteAccountRequest
	GetDBClusterId() *string
	SetEngine(v string) *DeleteAccountRequest
	GetEngine() *string
}

type DeleteAccountRequest struct {
	// The name of the database account.
	//
	// >  You can call the [DescribeAccounts](https://help.aliyun.com/document_detail/612430.html) operation to query the information about database accounts for a cluster, including the account name.
	//
	// This parameter is required.
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
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DeleteAccountRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAccountRequest) GetEngine() *string {
	return s.Engine
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetDBClusterId(v string) *DeleteAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAccountRequest) SetEngine(v string) *DeleteAccountRequest {
	s.Engine = &v
	return s
}

func (s *DeleteAccountRequest) Validate() error {
	return dara.Validate(s)
}
