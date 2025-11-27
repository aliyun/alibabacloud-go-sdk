// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableCrossRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeAvailableCrossRegionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAvailableCrossRegionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAvailableCrossRegionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAvailableCrossRegionRequest
	GetResourceOwnerId() *int64
}

type DescribeAvailableCrossRegionRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent zone list.
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

func (s DescribeAvailableCrossRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableCrossRegionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCrossRegionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAvailableCrossRegionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableCrossRegionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAvailableCrossRegionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableCrossRegionRequest) SetOwnerId(v int64) *DescribeAvailableCrossRegionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableCrossRegionRequest) SetRegionId(v string) *DescribeAvailableCrossRegionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableCrossRegionRequest) SetResourceOwnerAccount(v string) *DescribeAvailableCrossRegionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableCrossRegionRequest) SetResourceOwnerId(v int64) *DescribeAvailableCrossRegionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableCrossRegionRequest) Validate() error {
	return dara.Validate(s)
}
