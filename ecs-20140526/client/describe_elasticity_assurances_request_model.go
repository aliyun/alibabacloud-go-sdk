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
	// The billing method of instances. Valid values: PostPaid. Only pay-as-you-go is supported.
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance type. You can use the instance type to query only active elasticity assurance services. Released services can only be queried by using `PrivatePoolOptions.Ids`.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum number of entries per page for a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the elasticity assurance service query. Obtain the value from the result of the previous request.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of the elasticity assurance service. Valid values:
	//
	// - ElasticityAssurance: standard elasticity assurance (used when RecurrenceRules is not specified).
	//
	// - TimeDivisionElasticityAssurance: time-division elasticity assurance (used when RecurrenceRules is specified).
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
	// The ID of the region to which the elasticity assurance service belongs. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. When you use this parameter to filter resources, the resource count cannot exceed 1000.
	//
	// >Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the elasticity assurance service. Valid values:
	//
	// - All: all states.
	//
	// - Deactived: pending activation (this state is in invitational preview).
	//
	// - Preparing: being prepared.
	//
	// - Prepared: pending effectiveness.
	//
	// - Active: active.
	//
	// - Released: released.
	//
	// If you do not specify this parameter, elasticity assurance services in all states except Pending and Released are queried.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tag key-value pairs bound to the elasticity assurance service.
	Tag []*DescribeElasticityAssurancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID in the region to which the elasticity assurance service belongs.
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
	// The list of elasticity assurance service IDs. The value can be a JSON array that consists of up to 100 IDs. Separate multiple IDs with commas (,).
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
	// The tag key. N indicates that you can set multiple tag keys for filtering. Valid values of N: 1 to 20.
	//
	// If you use a single tag to filter resources, the resource count with the specified tag cannot exceed 1000. If you use multiple tags to filter resources, the resource count of resources that have all specified tags attached cannot exceed 1000. If the resource count exceeds 1000, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query resources.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. N indicates that you can set multiple tag values for filtering. Valid values of N: 1 to 20.
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
