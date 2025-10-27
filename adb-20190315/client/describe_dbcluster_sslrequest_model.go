// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterSSLRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterSSLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterSSLRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBClusterSSLRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterSSLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterSSLRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterSSLRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612293.html) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterSSLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterSSLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterSSLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterSSLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterSSLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterSSLRequest) SetDBClusterId(v string) *DescribeDBClusterSSLRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetOwnerAccount(v string) *DescribeDBClusterSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetOwnerId(v int64) *DescribeDBClusterSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetRegionId(v string) *DescribeDBClusterSSLRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetResourceOwnerId(v int64) *DescribeDBClusterSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) Validate() error {
	return dara.Validate(s)
}
