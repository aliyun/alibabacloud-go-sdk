// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRangeListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribePortRangeListsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePortRangeListsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribePortRangeListsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePortRangeListsRequest
	GetOwnerId() *int64
	SetPortRangeListId(v []*string) *DescribePortRangeListsRequest
	GetPortRangeListId() []*string
	SetPortRangeListName(v string) *DescribePortRangeListsRequest
	GetPortRangeListName() *string
	SetRegionId(v string) *DescribePortRangeListsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribePortRangeListsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribePortRangeListsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePortRangeListsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribePortRangeListsRequestTag) *DescribePortRangeListsRequest
	GetTag() []*DescribePortRangeListsRequestTag
}

type DescribePortRangeListsRequest struct {
	// The number of entries per page for a paged query.
	//
	// - Maximum value: 100.
	//
	// - Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// 727d41872117f2816343eeb432fbc5bfd21dc824589d2a4be0b5e8707e68181f
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the port range list. Valid values of N: 0 to 100.
	PortRangeListId []*string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty" type:"Repeated"`
	// The name of the port range list. The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. The name cannot start with http://, https://, com.aliyun, or com.alibabacloud. The name can contain letters, Chinese characters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// PortRangeListNameSample
	PortRangeListName *string `json:"PortRangeListName,omitempty" xml:"PortRangeListName,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/2679950.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. When you use this parameter to filter resources, the resource count cannot exceed 1000. You can invoke [ListResourceGroups](https://help.aliyun.com/document_detail/2716558.html) to query the list of resource groups.
	//
	// >Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags bound to the port range list.
	Tag []*DescribePortRangeListsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribePortRangeListsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListsRequest) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePortRangeListsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePortRangeListsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePortRangeListsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePortRangeListsRequest) GetPortRangeListId() []*string {
	return s.PortRangeListId
}

func (s *DescribePortRangeListsRequest) GetPortRangeListName() *string {
	return s.PortRangeListName
}

func (s *DescribePortRangeListsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePortRangeListsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortRangeListsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePortRangeListsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePortRangeListsRequest) GetTag() []*DescribePortRangeListsRequestTag {
	return s.Tag
}

func (s *DescribePortRangeListsRequest) SetMaxResults(v int32) *DescribePortRangeListsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetNextToken(v string) *DescribePortRangeListsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetOwnerAccount(v string) *DescribePortRangeListsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetOwnerId(v int64) *DescribePortRangeListsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetPortRangeListId(v []*string) *DescribePortRangeListsRequest {
	s.PortRangeListId = v
	return s
}

func (s *DescribePortRangeListsRequest) SetPortRangeListName(v string) *DescribePortRangeListsRequest {
	s.PortRangeListName = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetRegionId(v string) *DescribePortRangeListsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetResourceGroupId(v string) *DescribePortRangeListsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetResourceOwnerAccount(v string) *DescribePortRangeListsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetResourceOwnerId(v int64) *DescribePortRangeListsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetTag(v []*DescribePortRangeListsRequestTag) *DescribePortRangeListsRequest {
	s.Tag = v
	return s
}

func (s *DescribePortRangeListsRequest) Validate() error {
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

type DescribePortRangeListsRequestTag struct {
	// The tag key of the instance. Valid values of N: 1 to 20.
	//
	// If you use a single tag to filter resources, the resource count with this tag cannot exceed 1000. If you use multiple tags to filter resources, the resource count of resources that have all specified tags attached cannot exceed 1000. If the resource count exceeds 1000, call [ListTagResources](https://help.aliyun.com/document_detail/110425.html) to query resources.
	//
	// example:
	//
	// key for PortRangeList
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value for PortRangeList
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePortRangeListsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribePortRangeListsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribePortRangeListsRequestTag) SetKey(v string) *DescribePortRangeListsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribePortRangeListsRequestTag) SetValue(v string) *DescribePortRangeListsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribePortRangeListsRequestTag) Validate() error {
	return dara.Validate(s)
}
