// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterAttributeRequest
	GetClusterId() *string
	SetOwnerId(v int64) *DescribeClusterAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeClusterAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeClusterAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClusterAttributeRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeClusterAttributeRequest
	GetZoneId() *string
}

type DescribeClusterAttributeRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClusterAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClusterAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterAttributeRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClusterAttributeRequest) SetClusterId(v string) *DescribeClusterAttributeRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterAttributeRequest) SetOwnerId(v int64) *DescribeClusterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterAttributeRequest) SetRegionId(v string) *DescribeClusterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterAttributeRequest) SetResourceOwnerAccount(v string) *DescribeClusterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterAttributeRequest) SetResourceOwnerId(v int64) *DescribeClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterAttributeRequest) SetZoneId(v string) *DescribeClusterAttributeRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
