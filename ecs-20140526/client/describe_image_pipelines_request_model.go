// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagePipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImagePipelineId(v []*string) *DescribeImagePipelinesRequest
	GetImagePipelineId() []*string
	SetMaxResults(v int32) *DescribeImagePipelinesRequest
	GetMaxResults() *int32
	SetName(v string) *DescribeImagePipelinesRequest
	GetName() *string
	SetNextToken(v string) *DescribeImagePipelinesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeImagePipelinesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeImagePipelinesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeImagePipelinesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeImagePipelinesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeImagePipelinesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeImagePipelinesRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeImagePipelinesRequestTag) *DescribeImagePipelinesRequest
	GetTag() []*DescribeImagePipelinesRequestTag
}

type DescribeImagePipelinesRequest struct {
	// The image template ID. Valid values of N: 1 to 20.
	//
	// example:
	//
	// ip-2ze5tsl5bp6nf2b3****
	ImagePipelineId []*string `json:"ImagePipelineId,omitempty" xml:"ImagePipelineId,omitempty" type:"Repeated"`
	// The maximum number of entries per page for paging queries. Valid values: 1 to 500.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The template name.
	//
	// example:
	//
	// testImagePipeline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token. Set this parameter to the value of `NextToken` returned by the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group. When you use this parameter to filter resources, the resource count cannot exceed 1000.
	//
	// >Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*DescribeImagePipelinesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesRequest) GetImagePipelineId() []*string {
	return s.ImagePipelineId
}

func (s *DescribeImagePipelinesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeImagePipelinesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeImagePipelinesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImagePipelinesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeImagePipelinesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeImagePipelinesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImagePipelinesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeImagePipelinesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeImagePipelinesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeImagePipelinesRequest) GetTag() []*DescribeImagePipelinesRequestTag {
	return s.Tag
}

func (s *DescribeImagePipelinesRequest) SetImagePipelineId(v []*string) *DescribeImagePipelinesRequest {
	s.ImagePipelineId = v
	return s
}

func (s *DescribeImagePipelinesRequest) SetMaxResults(v int32) *DescribeImagePipelinesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImagePipelinesRequest) SetName(v string) *DescribeImagePipelinesRequest {
	s.Name = &v
	return s
}

func (s *DescribeImagePipelinesRequest) SetNextToken(v string) *DescribeImagePipelinesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImagePipelinesRequest) SetOwnerAccount(v string) *DescribeImagePipelinesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeImagePipelinesRequest) SetOwnerId(v int64) *DescribeImagePipelinesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImagePipelinesRequest) SetRegionId(v string) *DescribeImagePipelinesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImagePipelinesRequest) SetResourceGroupId(v string) *DescribeImagePipelinesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImagePipelinesRequest) SetResourceOwnerAccount(v string) *DescribeImagePipelinesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeImagePipelinesRequest) SetResourceOwnerId(v int64) *DescribeImagePipelinesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeImagePipelinesRequest) SetTag(v []*DescribeImagePipelinesRequestTag) *DescribeImagePipelinesRequest {
	s.Tag = v
	return s
}

func (s *DescribeImagePipelinesRequest) Validate() error {
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

type DescribeImagePipelinesRequestTag struct {
	// The tag key. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImagePipelinesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeImagePipelinesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeImagePipelinesRequestTag) SetKey(v string) *DescribeImagePipelinesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeImagePipelinesRequestTag) SetValue(v string) *DescribeImagePipelinesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeImagePipelinesRequestTag) Validate() error {
	return dara.Validate(s)
}
