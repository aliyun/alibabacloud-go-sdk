// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterSpaceSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterSpaceSummaryRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterSpaceSummaryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterSpaceSummaryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBClusterSpaceSummaryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterSpaceSummaryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterSpaceSummaryRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterSpaceSummaryRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9v5sa7mm79z4l2
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

func (s DescribeDBClusterSpaceSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterSpaceSummaryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterSpaceSummaryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterSpaceSummaryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterSpaceSummaryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterSpaceSummaryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetDBClusterId(v string) *DescribeDBClusterSpaceSummaryRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetOwnerAccount(v string) *DescribeDBClusterSpaceSummaryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetOwnerId(v int64) *DescribeDBClusterSpaceSummaryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetRegionId(v string) *DescribeDBClusterSpaceSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterSpaceSummaryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetResourceOwnerId(v int64) *DescribeDBClusterSpaceSummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) Validate() error {
	return dara.Validate(s)
}
