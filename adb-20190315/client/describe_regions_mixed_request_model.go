// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsMixedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeRegionsMixedRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRegionsMixedRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeRegionsMixedRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRegionsMixedRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRegionsMixedRequest
	GetResourceOwnerId() *int64
}

type DescribeRegionsMixedRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsMixedRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsMixedRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsMixedRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRegionsMixedRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRegionsMixedRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsMixedRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRegionsMixedRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRegionsMixedRequest) SetOwnerAccount(v string) *DescribeRegionsMixedRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsMixedRequest) SetOwnerId(v int64) *DescribeRegionsMixedRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsMixedRequest) SetRegionId(v string) *DescribeRegionsMixedRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsMixedRequest) SetResourceOwnerAccount(v string) *DescribeRegionsMixedRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsMixedRequest) SetResourceOwnerId(v int64) *DescribeRegionsMixedRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionsMixedRequest) Validate() error {
	return dara.Validate(s)
}
