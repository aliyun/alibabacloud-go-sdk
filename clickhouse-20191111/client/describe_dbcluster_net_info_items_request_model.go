// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNetInfoItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterNetInfoItemsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterNetInfoItemsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterNetInfoItemsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterNetInfoItemsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterNetInfoItemsRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterNetInfoItemsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterNetInfoItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoItemsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterNetInfoItemsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterNetInfoItemsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterNetInfoItemsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterNetInfoItemsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterNetInfoItemsRequest) SetDBClusterId(v string) *DescribeDBClusterNetInfoItemsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsRequest) SetOwnerAccount(v string) *DescribeDBClusterNetInfoItemsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsRequest) SetOwnerId(v int64) *DescribeDBClusterNetInfoItemsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterNetInfoItemsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsRequest) SetResourceOwnerId(v int64) *DescribeDBClusterNetInfoItemsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsRequest) Validate() error {
	return dara.Validate(s)
}
