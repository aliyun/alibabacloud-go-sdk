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
	SetDBClusterId(v string) *DescribeAccountsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeAccountsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAccountsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeAccountsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccountsRequest
	GetResourceOwnerId() *int64
	SetTags(v []*DescribeAccountsRequestTags) *DescribeAccountsRequest
	GetTags() []*DescribeAccountsRequestTags
}

type DescribeAccountsRequest struct {
	// The name of the database account.
	//
	// >  If you do not specify this parameter, the information about all database accounts is returned.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of the database account. If you do not specify this parameter, the information about all account types is returned. Valid values:
	//
	// 	- **Normal**: standard account.
	//
	// 	- **Super**: privileged account.
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBClusterId          *string                        `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 []*DescribeAccountsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *DescribeAccountsRequest) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *DescribeAccountsRequest) GetTags() []*DescribeAccountsRequestTags {
	return s.Tags
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountType(v string) *DescribeAccountsRequest {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBClusterId(v string) *DescribeAccountsRequest {
	s.DBClusterId = &v
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

func (s *DescribeAccountsRequest) SetTags(v []*DescribeAccountsRequestTags) *DescribeAccountsRequest {
	s.Tags = v
	return s
}

func (s *DescribeAccountsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountsRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAccountsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAccountsRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAccountsRequestTags) SetKey(v string) *DescribeAccountsRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeAccountsRequestTags) SetValue(v string) *DescribeAccountsRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeAccountsRequestTags) Validate() error {
	return dara.Validate(s)
}
