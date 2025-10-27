// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNetInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterNetInfoRequest
	GetDBClusterId() *string
	SetEngine(v string) *DescribeDBClusterNetInfoRequest
	GetEngine() *string
	SetOwnerAccount(v string) *DescribeDBClusterNetInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterNetInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBClusterNetInfoRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterNetInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterNetInfoRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterNetInfoRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Engine       *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterNetInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterNetInfoRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClusterNetInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterNetInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterNetInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterNetInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterNetInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterNetInfoRequest) SetDBClusterId(v string) *DescribeDBClusterNetInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetEngine(v string) *DescribeDBClusterNetInfoRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetOwnerAccount(v string) *DescribeDBClusterNetInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetOwnerId(v int64) *DescribeDBClusterNetInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetRegionId(v string) *DescribeDBClusterNetInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterNetInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetResourceOwnerId(v int64) *DescribeDBClusterNetInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) Validate() error {
	return dara.Validate(s)
}
