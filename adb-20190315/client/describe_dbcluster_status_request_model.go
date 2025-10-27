// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeDBClusterStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBClusterStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterStatusRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterStatusRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterStatusRequest) SetOwnerAccount(v string) *DescribeDBClusterStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterStatusRequest) SetOwnerId(v int64) *DescribeDBClusterStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterStatusRequest) SetRegionId(v string) *DescribeDBClusterStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterStatusRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterStatusRequest) SetResourceOwnerId(v int64) *DescribeDBClusterStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterStatusRequest) Validate() error {
	return dara.Validate(s)
}
