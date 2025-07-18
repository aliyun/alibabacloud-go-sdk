// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetAccountRequest
	GetAccountName() *string
	SetDBInstanceId(v string) *GetAccountRequest
	GetDBInstanceId() *string
}

type GetAccountRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuser
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s GetAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountRequest) GoString() string {
	return s.String()
}

func (s *GetAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetAccountRequest) SetAccountName(v string) *GetAccountRequest {
	s.AccountName = &v
	return s
}

func (s *GetAccountRequest) SetDBInstanceId(v string) *GetAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GetAccountRequest) Validate() error {
	return dara.Validate(s)
}
