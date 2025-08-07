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
	SetNodeType(v string) *DescribeAccountsRequest
	GetNodeType() *string
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
}

type DescribeAccountsRequest struct {
	// The name of the account.
	//
	// example:
	//
	// test_acc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node type. This parameter is required when you query the PolarDB search node account.
	//
	// 	- Search
	//
	// example:
	//
	// Search
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the page to return. Set this parameter to an integer that is larger than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *DescribeAccountsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAccountsRequest) GetNodeType() *string {
	return s.NodeType
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

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBClusterId(v string) *DescribeAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountsRequest) SetNodeType(v string) *DescribeAccountsRequest {
	s.NodeType = &v
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

func (s *DescribeAccountsRequest) Validate() error {
	return dara.Validate(s)
}
