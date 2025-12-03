// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterModelRequest
	GetClusterId() *string
	SetOwnerId(v int64) *DescribeClusterModelRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeClusterModelRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeClusterModelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClusterModelRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeClusterModelRequest
	GetZoneId() *string
}

type DescribeClusterModelRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterModelRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterModelRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterModelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClusterModelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterModelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClusterModelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterModelRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClusterModelRequest) SetClusterId(v string) *DescribeClusterModelRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterModelRequest) SetOwnerId(v int64) *DescribeClusterModelRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterModelRequest) SetRegionId(v string) *DescribeClusterModelRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterModelRequest) SetResourceOwnerAccount(v string) *DescribeClusterModelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterModelRequest) SetResourceOwnerId(v int64) *DescribeClusterModelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterModelRequest) SetZoneId(v string) *DescribeClusterModelRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterModelRequest) Validate() error {
	return dara.Validate(s)
}
