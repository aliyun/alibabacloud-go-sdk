// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAllAccountsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeAllAccountsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAllAccountsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeAllAccountsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAllAccountsRequest
	GetResourceOwnerId() *int64
}

type DescribeAllAccountsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAllAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAllAccountsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAllAccountsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAllAccountsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAllAccountsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAllAccountsRequest) SetDBClusterId(v string) *DescribeAllAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetOwnerAccount(v string) *DescribeAllAccountsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetOwnerId(v int64) *DescribeAllAccountsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetResourceOwnerAccount(v string) *DescribeAllAccountsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetResourceOwnerId(v int64) *DescribeAllAccountsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAllAccountsRequest) Validate() error {
	return dara.Validate(s)
}
