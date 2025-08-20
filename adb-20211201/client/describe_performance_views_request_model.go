// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePerformanceViewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePerformanceViewsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribePerformanceViewsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePerformanceViewsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribePerformanceViewsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePerformanceViewsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePerformanceViewsRequest
	GetResourceOwnerId() *int64
}

type DescribePerformanceViewsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1ub9grke1****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
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

func (s DescribePerformanceViewsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePerformanceViewsRequest) GoString() string {
	return s.String()
}

func (s *DescribePerformanceViewsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePerformanceViewsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePerformanceViewsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePerformanceViewsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePerformanceViewsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePerformanceViewsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePerformanceViewsRequest) SetDBClusterId(v string) *DescribePerformanceViewsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePerformanceViewsRequest) SetOwnerAccount(v string) *DescribePerformanceViewsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePerformanceViewsRequest) SetOwnerId(v int64) *DescribePerformanceViewsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePerformanceViewsRequest) SetRegionId(v string) *DescribePerformanceViewsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePerformanceViewsRequest) SetResourceOwnerAccount(v string) *DescribePerformanceViewsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePerformanceViewsRequest) SetResourceOwnerId(v int64) *DescribePerformanceViewsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePerformanceViewsRequest) Validate() error {
	return dara.Validate(s)
}
