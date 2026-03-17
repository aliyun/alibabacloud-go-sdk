// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteDistributionStrategiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeRouteDistributionStrategiesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRouteDistributionStrategiesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRouteDistributionStrategiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteDistributionStrategiesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRouteDistributionStrategiesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRouteDistributionStrategiesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRouteDistributionStrategiesRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeRouteDistributionStrategiesRequest
	GetSmartAGId() *string
	SetSourceType(v string) *DescribeRouteDistributionStrategiesRequest
	GetSourceType() *string
}

type DescribeRouteDistributionStrategiesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Smart Access Gateway (SAG) instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-erx3qta5xg5zyq****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The type of route. Valid values:
	//
	// 	- **cloud**: Alibaba Cloud-facing routes. These routes allow SAG instances to access resources deployed on Alibaba Cloud.
	//
	// 	- **local**: on-premises network-facing routes. These routes allow SAG instances to communicate with on-premises terminals.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeRouteDistributionStrategiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteDistributionStrategiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteDistributionStrategiesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRouteDistributionStrategiesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRouteDistributionStrategiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteDistributionStrategiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteDistributionStrategiesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouteDistributionStrategiesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRouteDistributionStrategiesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRouteDistributionStrategiesRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeRouteDistributionStrategiesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeRouteDistributionStrategiesRequest) SetOwnerAccount(v string) *DescribeRouteDistributionStrategiesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesRequest) SetOwnerId(v int64) *DescribeRouteDistributionStrategiesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesRequest) SetPageNumber(v int32) *DescribeRouteDistributionStrategiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesRequest) SetPageSize(v int32) *DescribeRouteDistributionStrategiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesRequest) SetRegionId(v string) *DescribeRouteDistributionStrategiesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesRequest) SetResourceOwnerAccount(v string) *DescribeRouteDistributionStrategiesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesRequest) SetResourceOwnerId(v int64) *DescribeRouteDistributionStrategiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesRequest) SetSmartAGId(v string) *DescribeRouteDistributionStrategiesRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesRequest) SetSourceType(v string) *DescribeRouteDistributionStrategiesRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesRequest) Validate() error {
	return dara.Validate(s)
}
