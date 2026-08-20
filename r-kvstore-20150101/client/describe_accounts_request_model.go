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
	SetInstanceId(v string) *DescribeAccountsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeAccountsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAccountsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAccountsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccountsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeAccountsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccountsRequest
	GetResourceOwnerId() *int64
	SetSearchAccountName(v string) *DescribeAccountsRequest
	GetSearchAccountName() *string
	SetSecurityToken(v string) *DescribeAccountsRequest
	GetSecurityToken() *string
}

type DescribeAccountsRequest struct {
	// The name of the account to query.
	//
	// example:
	//
	// demoaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The fuzzy match string for the account name.
	//
	// example:
	//
	// test
	SearchAccountName *string `json:"SearchAccountName,omitempty" xml:"SearchAccountName,omitempty"`
	SecurityToken     *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *DescribeAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAccountsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAccountsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccountsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccountsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccountsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccountsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccountsRequest) GetSearchAccountName() *string {
	return s.SearchAccountName
}

func (s *DescribeAccountsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetInstanceId(v string) *DescribeAccountsRequest {
	s.InstanceId = &v
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

func (s *DescribeAccountsRequest) SetPageNumber(v int32) *DescribeAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsRequest) SetPageSize(v int32) *DescribeAccountsRequest {
	s.PageSize = &v
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

func (s *DescribeAccountsRequest) SetSearchAccountName(v string) *DescribeAccountsRequest {
	s.SearchAccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetSecurityToken(v string) *DescribeAccountsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAccountsRequest) Validate() error {
	return dara.Validate(s)
}
