// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeAccountsZonalRequest
	GetAccountName() *string
	SetDBClusterId(v string) *DescribeAccountsZonalRequest
	GetDBClusterId() *string
	SetMaxResults(v int32) *DescribeAccountsZonalRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAccountsZonalRequest
	GetNextToken() *string
	SetNodeType(v string) *DescribeAccountsZonalRequest
	GetNodeType() *string
	SetOwnerAccount(v string) *DescribeAccountsZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAccountsZonalRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAccountsZonalRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccountsZonalRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeAccountsZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccountsZonalRequest
	GetResourceOwnerId() *int64
}

type DescribeAccountsZonalRequest struct {
	// example:
	//
	// test_acc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// Search
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAccountsZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsZonalRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsZonalRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountsZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAccountsZonalRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAccountsZonalRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAccountsZonalRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeAccountsZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAccountsZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccountsZonalRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccountsZonalRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccountsZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccountsZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccountsZonalRequest) SetAccountName(v string) *DescribeAccountsZonalRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsZonalRequest) SetDBClusterId(v string) *DescribeAccountsZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountsZonalRequest) SetMaxResults(v int32) *DescribeAccountsZonalRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAccountsZonalRequest) SetNextToken(v string) *DescribeAccountsZonalRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAccountsZonalRequest) SetNodeType(v string) *DescribeAccountsZonalRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeAccountsZonalRequest) SetOwnerAccount(v string) *DescribeAccountsZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccountsZonalRequest) SetOwnerId(v int64) *DescribeAccountsZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountsZonalRequest) SetPageNumber(v int32) *DescribeAccountsZonalRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsZonalRequest) SetPageSize(v int32) *DescribeAccountsZonalRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountsZonalRequest) SetResourceOwnerAccount(v string) *DescribeAccountsZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountsZonalRequest) SetResourceOwnerId(v int64) *DescribeAccountsZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccountsZonalRequest) Validate() error {
	return dara.Validate(s)
}
