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
	SetDBInstanceId(v string) *DescribeAccountsRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeAccountsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAccountsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeAccountsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccountsRequest
	GetResourceOwnerId() *int64
}

type DescribeAccountsRequest struct {
	// The name of the account. Set the value to **root**.
	//
	// example:
	//
	// root
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1fd530f271****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribeAccountsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAccountsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAccountsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccountsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccountsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBInstanceId(v string) *DescribeAccountsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerAccount(v string) *DescribeAccountsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerId(v int64) *DescribeAccountsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountsRequest) SetResourceOwnerAccount(v string) *DescribeAccountsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountsRequest) SetResourceOwnerId(v int64) *DescribeAccountsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccountsRequest) Validate() error {
	return dara.Validate(s)
}
