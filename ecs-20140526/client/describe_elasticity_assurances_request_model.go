// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticityAssurancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *DescribeElasticityAssurancesRequestPrivatePoolOptions) *DescribeElasticityAssurancesRequest
	GetPrivatePoolOptions() *DescribeElasticityAssurancesRequestPrivatePoolOptions
	SetInstanceChargeType(v string) *DescribeElasticityAssurancesRequest
	GetInstanceChargeType() *string
	SetInstanceType(v string) *DescribeElasticityAssurancesRequest
	GetInstanceType() *string
	SetMaxResults(v int32) *DescribeElasticityAssurancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeElasticityAssurancesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeElasticityAssurancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeElasticityAssurancesRequest
	GetOwnerId() *int64
	SetPackageType(v string) *DescribeElasticityAssurancesRequest
	GetPackageType() *string
	SetPlatform(v string) *DescribeElasticityAssurancesRequest
	GetPlatform() *string
	SetRegionId(v string) *DescribeElasticityAssurancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeElasticityAssurancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeElasticityAssurancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeElasticityAssurancesRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeElasticityAssurancesRequest
	GetStatus() *string
	SetTag(v []*DescribeElasticityAssurancesRequestTag) *DescribeElasticityAssurancesRequest
	GetTag() []*DescribeElasticityAssurancesRequestTag
	SetZoneId(v string) *DescribeElasticityAssurancesRequest
	GetZoneId() *string
}

type DescribeElasticityAssurancesRequest struct {
	PrivatePoolOptions *DescribeElasticityAssurancesRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The billing method of the instances. Only `PostPaid` (pay-as-you-go) is supported.
	//
	// Default value: `PostPaid`.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance type. You can use this parameter to query only active elasticity assurances. To query released elasticity assurances, you must use `PrivatePoolOptions.Ids`.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of the Elasticity Assurance. Valid values:
	//
	// - `ElasticityAssurance`: a standard elasticity assurance. This type of elasticity assurance is created when you do not specify `RecurrenceRules`.
	//
	// - `TimeDivisionElasticityAssurance`: a time-division elasticity assurance. This type of elasticity assurance is created when you specify `RecurrenceRules`.
	//
	// example:
	//
	// ElasticityAssurance
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The ID of the region where the Elasticity Assurance is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. When you use this parameter to filter resources, the number of matching resources cannot exceed 1,000.
	//
	// > Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the Elasticity Assurance. Valid values:
	//
	// - `All`: all statuses.
	//
	// - `Deactivated`: The Elasticity Assurance is pending activation. This status is available only for invitational preview.
	//
	// - `Preparing`: The Elasticity Assurance is being prepared.
	//
	// - `Prepared`: The Elasticity Assurance is ready to take effect.
	//
	// - `Active`: The Elasticity Assurance is active.
	//
	// - `Released`: The Elasticity Assurance is released.
	//
	// If you do not specify this parameter, elasticity assurances in all states are returned, except for those in the `Pending` and `Released` states.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags used to filter Elasticity Assurances.
	Tag []*DescribeElasticityAssurancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the zone where the Elasticity Assurance is located.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeElasticityAssurancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesRequest) GetPrivatePoolOptions() *DescribeElasticityAssurancesRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *DescribeElasticityAssurancesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeElasticityAssurancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeElasticityAssurancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeElasticityAssurancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeElasticityAssurancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeElasticityAssurancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeElasticityAssurancesRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeElasticityAssurancesRequest) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeElasticityAssurancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeElasticityAssurancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeElasticityAssurancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeElasticityAssurancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeElasticityAssurancesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeElasticityAssurancesRequest) GetTag() []*DescribeElasticityAssurancesRequestTag {
	return s.Tag
}

func (s *DescribeElasticityAssurancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeElasticityAssurancesRequest) SetPrivatePoolOptions(v *DescribeElasticityAssurancesRequestPrivatePoolOptions) *DescribeElasticityAssurancesRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetInstanceChargeType(v string) *DescribeElasticityAssurancesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetInstanceType(v string) *DescribeElasticityAssurancesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetMaxResults(v int32) *DescribeElasticityAssurancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetNextToken(v string) *DescribeElasticityAssurancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetOwnerAccount(v string) *DescribeElasticityAssurancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetOwnerId(v int64) *DescribeElasticityAssurancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetPackageType(v string) *DescribeElasticityAssurancesRequest {
	s.PackageType = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetPlatform(v string) *DescribeElasticityAssurancesRequest {
	s.Platform = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetRegionId(v string) *DescribeElasticityAssurancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetResourceGroupId(v string) *DescribeElasticityAssurancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetResourceOwnerAccount(v string) *DescribeElasticityAssurancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetResourceOwnerId(v int64) *DescribeElasticityAssurancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetStatus(v string) *DescribeElasticityAssurancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetTag(v []*DescribeElasticityAssurancesRequestTag) *DescribeElasticityAssurancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetZoneId(v string) *DescribeElasticityAssurancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticityAssurancesRequestPrivatePoolOptions struct {
	// The IDs of the elasticity assurances. You can specify a JSON array of up to 100 elasticity assurance IDs.
	//
	// example:
	//
	// ["eap-bp67acfmxazb4****", "eap-bp67acfmxazb5****"]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DescribeElasticityAssurancesRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesRequestPrivatePoolOptions) GetIds() *string {
	return s.Ids
}

func (s *DescribeElasticityAssurancesRequestPrivatePoolOptions) SetIds(v string) *DescribeElasticityAssurancesRequestPrivatePoolOptions {
	s.Ids = &v
	return s
}

func (s *DescribeElasticityAssurancesRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticityAssurancesRequestTag struct {
	// The tag key. You can specify up to 20 tag keys to filter resources.
	//
	// The query returns a maximum of 1,000 resources that match the specified tags. If more than 1,000 resources match the tags, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query all the resources.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 20 tag values.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeElasticityAssurancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeElasticityAssurancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeElasticityAssurancesRequestTag) SetKey(v string) *DescribeElasticityAssurancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeElasticityAssurancesRequestTag) SetValue(v string) *DescribeElasticityAssurancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeElasticityAssurancesRequestTag) Validate() error {
	return dara.Validate(s)
}
