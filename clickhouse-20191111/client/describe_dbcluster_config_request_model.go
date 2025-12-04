// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterConfigRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBClusterConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterConfigRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterConfigRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-wz988vja2mor4****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterConfigRequest) SetDBClusterId(v string) *DescribeDBClusterConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetOwnerAccount(v string) *DescribeDBClusterConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetOwnerId(v int64) *DescribeDBClusterConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetRegionId(v string) *DescribeDBClusterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetResourceOwnerId(v int64) *DescribeDBClusterConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) Validate() error {
	return dara.Validate(s)
}
