// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupType(v string) *DescribeScalingGroupsRequest
	GetGroupType() *string
	SetOwnerAccount(v string) *DescribeScalingGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeScalingGroupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeScalingGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScalingGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeScalingGroupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeScalingGroupsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeScalingGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeScalingGroupsRequest
	GetResourceOwnerId() *int64
	SetScalingGroupIds(v []*string) *DescribeScalingGroupsRequest
	GetScalingGroupIds() []*string
	SetScalingGroupName(v string) *DescribeScalingGroupsRequest
	GetScalingGroupName() *string
	SetScalingGroupNames(v []*string) *DescribeScalingGroupsRequest
	GetScalingGroupNames() []*string
	SetTags(v []*DescribeScalingGroupsRequestTags) *DescribeScalingGroupsRequest
	GetTags() []*DescribeScalingGroupsRequestTags
}

type DescribeScalingGroupsRequest struct {
	// The type of instances that are managed by the scaling group. Valid values:
	//
	// 	- ECS: Elastic Compute Service (ECS) instances.
	//
	// 	- ECI: elastic container instances.
	//
	// example:
	//
	// ECS
	GroupType    *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Page starts from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the scaling group that you want to query belongs.
	//
	// >  If no scaling group belongs to the specified resource group, the query result is empty and no error is reported.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the scaling groups that you want to query.
	//
	// The IDs of inactive scaling groups are not included in the query results, and no error is returned.
	ScalingGroupIds []*string `json:"ScalingGroupIds,omitempty" xml:"ScalingGroupIds,omitempty" type:"Repeated"`
	// The name of the scaling group.
	//
	// example:
	//
	// scalinggroup****
	ScalingGroupName *string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	// The names of the scaling groups that you want to query.
	//
	// The names of inactive scaling groups are not displayed in the query results, and no error is reported.
	ScalingGroupNames []*string `json:"ScalingGroupNames,omitempty" xml:"ScalingGroupNames,omitempty" type:"Repeated"`
	// The tags of the scaling group.
	Tags []*DescribeScalingGroupsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeScalingGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeScalingGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScalingGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScalingGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScalingGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScalingGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeScalingGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeScalingGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeScalingGroupsRequest) GetScalingGroupIds() []*string {
	return s.ScalingGroupIds
}

func (s *DescribeScalingGroupsRequest) GetScalingGroupName() *string {
	return s.ScalingGroupName
}

func (s *DescribeScalingGroupsRequest) GetScalingGroupNames() []*string {
	return s.ScalingGroupNames
}

func (s *DescribeScalingGroupsRequest) GetTags() []*DescribeScalingGroupsRequestTags {
	return s.Tags
}

func (s *DescribeScalingGroupsRequest) SetGroupType(v string) *DescribeScalingGroupsRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetOwnerAccount(v string) *DescribeScalingGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetOwnerId(v int64) *DescribeScalingGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetPageNumber(v int32) *DescribeScalingGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetPageSize(v int32) *DescribeScalingGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetRegionId(v string) *DescribeScalingGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetResourceGroupId(v string) *DescribeScalingGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetResourceOwnerAccount(v string) *DescribeScalingGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetResourceOwnerId(v int64) *DescribeScalingGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetScalingGroupIds(v []*string) *DescribeScalingGroupsRequest {
	s.ScalingGroupIds = v
	return s
}

func (s *DescribeScalingGroupsRequest) SetScalingGroupName(v string) *DescribeScalingGroupsRequest {
	s.ScalingGroupName = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetScalingGroupNames(v []*string) *DescribeScalingGroupsRequest {
	s.ScalingGroupNames = v
	return s
}

func (s *DescribeScalingGroupsRequest) SetTags(v []*DescribeScalingGroupsRequestTags) *DescribeScalingGroupsRequest {
	s.Tags = v
	return s
}

func (s *DescribeScalingGroupsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScalingGroupsRequestTags struct {
	// The tag key of the scaling group.
	//
	// example:
	//
	// Department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the scaling group.
	//
	// example:
	//
	// Finance
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScalingGroupsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeScalingGroupsRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeScalingGroupsRequestTags) SetKey(v string) *DescribeScalingGroupsRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeScalingGroupsRequestTags) SetValue(v string) *DescribeScalingGroupsRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeScalingGroupsRequestTags) Validate() error {
	return dara.Validate(s)
}
