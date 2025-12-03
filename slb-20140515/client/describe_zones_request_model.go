// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeZonesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeZonesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeZonesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeZonesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeZonesRequest
	GetResourceOwnerId() *int64
}

type DescribeZonesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Server Load Balancer (SLB) instance.
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

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeZonesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeZonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeZonesRequest) SetOwnerAccount(v string) *DescribeZonesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeZonesRequest) SetOwnerId(v int64) *DescribeZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerAccount(v string) *DescribeZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerId(v int64) *DescribeZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
