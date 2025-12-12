// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNodeInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterNodeInfosRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterNodeInfosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterNodeInfosRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBClusterNodeInfosRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBClusterNodeInfosRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBClusterNodeInfosRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterNodeInfosRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterNodeInfosRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterNodeInfosRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterNodeInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNodeInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeInfosRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterNodeInfosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterNodeInfosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterNodeInfosRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClusterNodeInfosRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBClusterNodeInfosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterNodeInfosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterNodeInfosRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterNodeInfosRequest) SetDBClusterId(v string) *DescribeDBClusterNodeInfosRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterNodeInfosRequest) SetOwnerAccount(v string) *DescribeDBClusterNodeInfosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNodeInfosRequest) SetOwnerId(v int64) *DescribeDBClusterNodeInfosRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterNodeInfosRequest) SetPageNumber(v int32) *DescribeDBClusterNodeInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClusterNodeInfosRequest) SetPageSize(v int32) *DescribeDBClusterNodeInfosRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClusterNodeInfosRequest) SetRegionId(v string) *DescribeDBClusterNodeInfosRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterNodeInfosRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterNodeInfosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNodeInfosRequest) SetResourceOwnerId(v int64) *DescribeDBClusterNodeInfosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterNodeInfosRequest) Validate() error {
	return dara.Validate(s)
}
