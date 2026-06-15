// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceByTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeResourceByTagsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeResourceByTagsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeResourceByTagsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeResourceByTagsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeResourceByTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeResourceByTagsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeResourceByTagsRequest
	GetResourceType() *string
	SetTag(v []*DescribeResourceByTagsRequestTag) *DescribeResourceByTagsRequest
	GetTag() []*DescribeResourceByTagsRequestTag
}

type DescribeResourceByTagsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number to return. The minimum value is 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the resource is located. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// - `instance`: ECS instance.
	//
	// - `disk`: disk.
	//
	// - `snapshot`: snapshot.
	//
	// - `image`: image.
	//
	// - `securitygroup`: security group.
	//
	// - `volume`: volume.
	//
	// - `eni`: elastic network interface.
	//
	// - `ddh`: dedicated host.
	//
	// - `keypair`: key pair.
	//
	// - `launchtemplate`: launch template.
	//
	// All values must be in lowercase.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag list.
	Tag []*DescribeResourceByTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeResourceByTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceByTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeResourceByTagsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeResourceByTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeResourceByTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceByTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeResourceByTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeResourceByTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourceByTagsRequest) GetTag() []*DescribeResourceByTagsRequestTag {
	return s.Tag
}

func (s *DescribeResourceByTagsRequest) SetOwnerId(v int64) *DescribeResourceByTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeResourceByTagsRequest) SetPageNumber(v int32) *DescribeResourceByTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceByTagsRequest) SetPageSize(v int32) *DescribeResourceByTagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceByTagsRequest) SetRegionId(v string) *DescribeResourceByTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceByTagsRequest) SetResourceOwnerAccount(v string) *DescribeResourceByTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeResourceByTagsRequest) SetResourceOwnerId(v int64) *DescribeResourceByTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeResourceByTagsRequest) SetResourceType(v string) *DescribeResourceByTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourceByTagsRequest) SetTag(v []*DescribeResourceByTagsRequestTag) *DescribeResourceByTagsRequest {
	s.Tag = v
	return s
}

func (s *DescribeResourceByTagsRequest) Validate() error {
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

type DescribeResourceByTagsRequestTag struct {
	// The tag key.
	//
	// > For best compatibility, we recommend that you use the `Tag.N.Key` parameter.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. An empty string is allowed. The value can be up to 128 characters in length.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeResourceByTagsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByTagsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeResourceByTagsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeResourceByTagsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeResourceByTagsRequestTag) SetKey(v string) *DescribeResourceByTagsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeResourceByTagsRequestTag) SetValue(v string) *DescribeResourceByTagsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeResourceByTagsRequestTag) Validate() error {
	return dara.Validate(s)
}
