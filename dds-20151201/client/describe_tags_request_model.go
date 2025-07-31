// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeTagsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeTagsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTagsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeTagsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTagsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeTagsRequest
	GetResourceType() *string
}

type DescribeTagsRequest struct {
	// The token used to start the next query to retrieve more results.
	//
	// >  This parameter is not required in the first query. If not all results are returned in one query, you can pass in the NextToken value returned in the previous query to perform the query again.
	//
	// example:
	//
	// 212db86****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// dds-bp17e7a04960****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Set the value to **INSTANCE**.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTagsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagsRequest) SetNextToken(v string) *DescribeTagsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerAccount(v string) *DescribeTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerId(v int64) *DescribeTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceGroupId(v string) *DescribeTagsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerAccount(v string) *DescribeTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerId(v int64) *DescribeTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceType(v string) *DescribeTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagsRequest) Validate() error {
	return dara.Validate(s)
}
