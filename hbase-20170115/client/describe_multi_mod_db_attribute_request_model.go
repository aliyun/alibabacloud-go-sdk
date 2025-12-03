// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiModDbAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeMultiModDbAttributeRequest
	GetClusterId() *string
	SetOwnerId(v int64) *DescribeMultiModDbAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeMultiModDbAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeMultiModDbAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMultiModDbAttributeRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeMultiModDbAttributeRequest
	GetZoneId() *string
}

type DescribeMultiModDbAttributeRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeMultiModDbAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiModDbAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiModDbAttributeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeMultiModDbAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMultiModDbAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMultiModDbAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMultiModDbAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMultiModDbAttributeRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeMultiModDbAttributeRequest) SetClusterId(v string) *DescribeMultiModDbAttributeRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) SetOwnerId(v int64) *DescribeMultiModDbAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) SetRegionId(v string) *DescribeMultiModDbAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) SetResourceOwnerAccount(v string) *DescribeMultiModDbAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) SetResourceOwnerId(v int64) *DescribeMultiModDbAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) SetZoneId(v string) *DescribeMultiModDbAttributeRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeMultiModDbAttributeRequest) Validate() error {
	return dara.Validate(s)
}
