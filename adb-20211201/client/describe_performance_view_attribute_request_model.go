// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePerformanceViewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePerformanceViewAttributeRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribePerformanceViewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePerformanceViewAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribePerformanceViewAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePerformanceViewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePerformanceViewAttributeRequest
	GetResourceOwnerId() *int64
	SetViewName(v string) *DescribePerformanceViewAttributeRequest
	GetViewName() *string
}

type DescribePerformanceViewAttributeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
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
	// The name of the view.
	//
	// This parameter is required.
	//
	// example:
	//
	// Basic
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s DescribePerformanceViewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePerformanceViewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribePerformanceViewAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePerformanceViewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePerformanceViewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePerformanceViewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePerformanceViewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePerformanceViewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePerformanceViewAttributeRequest) GetViewName() *string {
	return s.ViewName
}

func (s *DescribePerformanceViewAttributeRequest) SetDBClusterId(v string) *DescribePerformanceViewAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePerformanceViewAttributeRequest) SetOwnerAccount(v string) *DescribePerformanceViewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePerformanceViewAttributeRequest) SetOwnerId(v int64) *DescribePerformanceViewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePerformanceViewAttributeRequest) SetRegionId(v string) *DescribePerformanceViewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePerformanceViewAttributeRequest) SetResourceOwnerAccount(v string) *DescribePerformanceViewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePerformanceViewAttributeRequest) SetResourceOwnerId(v int64) *DescribePerformanceViewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePerformanceViewAttributeRequest) SetViewName(v string) *DescribePerformanceViewAttributeRequest {
	s.ViewName = &v
	return s
}

func (s *DescribePerformanceViewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
