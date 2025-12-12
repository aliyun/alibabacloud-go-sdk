// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterStatusSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeDBClusterStatusSetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterStatusSetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBClusterStatusSetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterStatusSetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterStatusSetRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterStatusSetRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterStatusSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStatusSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusSetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterStatusSetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterStatusSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterStatusSetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterStatusSetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterStatusSetRequest) SetOwnerAccount(v string) *DescribeDBClusterStatusSetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterStatusSetRequest) SetOwnerId(v int64) *DescribeDBClusterStatusSetRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterStatusSetRequest) SetRegionId(v string) *DescribeDBClusterStatusSetRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterStatusSetRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterStatusSetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterStatusSetRequest) SetResourceOwnerId(v int64) *DescribeDBClusterStatusSetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterStatusSetRequest) Validate() error {
	return dara.Validate(s)
}
