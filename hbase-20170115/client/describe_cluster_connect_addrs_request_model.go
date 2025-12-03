// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterConnectAddrsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterConnectAddrsRequest
	GetClusterId() *string
	SetOwnerId(v int64) *DescribeClusterConnectAddrsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeClusterConnectAddrsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeClusterConnectAddrsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClusterConnectAddrsRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeClusterConnectAddrsRequest
	GetZoneId() *string
}

type DescribeClusterConnectAddrsRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterConnectAddrsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterConnectAddrsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClusterConnectAddrsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterConnectAddrsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClusterConnectAddrsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterConnectAddrsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClusterConnectAddrsRequest) SetClusterId(v string) *DescribeClusterConnectAddrsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) SetOwnerId(v int64) *DescribeClusterConnectAddrsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) SetRegionId(v string) *DescribeClusterConnectAddrsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) SetResourceOwnerAccount(v string) *DescribeClusterConnectAddrsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) SetResourceOwnerId(v int64) *DescribeClusterConnectAddrsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) SetZoneId(v string) *DescribeClusterConnectAddrsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterConnectAddrsRequest) Validate() error {
	return dara.Validate(s)
}
