// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetIds(v string) *DescribeDeploymentSetsRequest
	GetDeploymentSetIds() *string
	SetDeploymentSetName(v string) *DescribeDeploymentSetsRequest
	GetDeploymentSetName() *string
	SetDomain(v string) *DescribeDeploymentSetsRequest
	GetDomain() *string
	SetGranularity(v string) *DescribeDeploymentSetsRequest
	GetGranularity() *string
	SetNetworkType(v string) *DescribeDeploymentSetsRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeDeploymentSetsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDeploymentSetsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDeploymentSetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDeploymentSetsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDeploymentSetsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDeploymentSetsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDeploymentSetsRequest
	GetResourceOwnerId() *int64
	SetStrategy(v string) *DescribeDeploymentSetsRequest
	GetStrategy() *string
	SetType(v string) *DescribeDeploymentSetsRequest
	GetType() *string
}

type DescribeDeploymentSetsRequest struct {
	// The IDs of the deployment sets. The value can be a JSON array that consists of up to 100 deployment set IDs. Sample format: `["ds-xxxxxxxxx", "ds-yyyyyyyyy", … "ds-zzzzzzzzz"]`.
	//
	// example:
	//
	// ["ds-bp67acfmxazb4ph****", "ds-bp67acfmxazb4pi****", … "ds-bp67acfmxazb4pj****"]
	DeploymentSetIds *string `json:"DeploymentSetIds,omitempty" xml:"DeploymentSetIds,omitempty"`
	// The name of the deployment set. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// testDeploymentSetName
	DeploymentSetName *string `json:"DeploymentSetName,omitempty" xml:"DeploymentSetName,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Starts at 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the deployment set is located. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The deployment strategy. Valid values:
	//
	// - Availability: high availability strategy.
	//
	// - AvailabilityGroup: high availability group strategy.
	//
	// - LowLatency: low-latency strategy.
	//
	// example:
	//
	// Availability
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The deployment type. Valid values:
	//
	// - host: Ensures that the instances in the deployment set are deployed on different hosts.
	//
	// - sw: Ensures that the instances in the deployment set are deployed on different switches.
	//
	// - rack: Ensures that the instances in the deployment set are deployed on different racks.
	//
	// Default value: host.
	//
	// example:
	//
	// host
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDeploymentSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetsRequest) GetDeploymentSetIds() *string {
	return s.DeploymentSetIds
}

func (s *DescribeDeploymentSetsRequest) GetDeploymentSetName() *string {
	return s.DeploymentSetName
}

func (s *DescribeDeploymentSetsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDeploymentSetsRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *DescribeDeploymentSetsRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDeploymentSetsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDeploymentSetsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDeploymentSetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDeploymentSetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeploymentSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeploymentSetsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDeploymentSetsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDeploymentSetsRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeDeploymentSetsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDeploymentSetsRequest) SetDeploymentSetIds(v string) *DescribeDeploymentSetsRequest {
	s.DeploymentSetIds = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetDeploymentSetName(v string) *DescribeDeploymentSetsRequest {
	s.DeploymentSetName = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetDomain(v string) *DescribeDeploymentSetsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetGranularity(v string) *DescribeDeploymentSetsRequest {
	s.Granularity = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetNetworkType(v string) *DescribeDeploymentSetsRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetOwnerAccount(v string) *DescribeDeploymentSetsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetOwnerId(v int64) *DescribeDeploymentSetsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetPageNumber(v int32) *DescribeDeploymentSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetPageSize(v int32) *DescribeDeploymentSetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetRegionId(v string) *DescribeDeploymentSetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetResourceOwnerAccount(v string) *DescribeDeploymentSetsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetResourceOwnerId(v int64) *DescribeDeploymentSetsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetStrategy(v string) *DescribeDeploymentSetsRequest {
	s.Strategy = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetType(v string) *DescribeDeploymentSetsRequest {
	s.Type = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) Validate() error {
	return dara.Validate(s)
}
