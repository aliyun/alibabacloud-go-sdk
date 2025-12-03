// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterWhiteListRequest
	GetClusterId() *string
	SetOwnerId(v int64) *DescribeClusterWhiteListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeClusterWhiteListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeClusterWhiteListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClusterWhiteListRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeClusterWhiteListRequest
	GetZoneId() *string
}

type DescribeClusterWhiteListRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterWhiteListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterWhiteListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClusterWhiteListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterWhiteListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClusterWhiteListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterWhiteListRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClusterWhiteListRequest) SetClusterId(v string) *DescribeClusterWhiteListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) SetOwnerId(v int64) *DescribeClusterWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) SetRegionId(v string) *DescribeClusterWhiteListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) SetResourceOwnerAccount(v string) *DescribeClusterWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) SetResourceOwnerId(v int64) *DescribeClusterWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) SetZoneId(v string) *DescribeClusterWhiteListRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
