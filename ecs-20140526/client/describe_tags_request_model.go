// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeTagsRequest
	GetCategory() *string
	SetOwnerId(v int64) *DescribeTagsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTagsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTagsRequest
	GetRegionId() *string
	SetResourceId(v string) *DescribeTagsRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *DescribeTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTagsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeTagsRequest
	GetResourceType() *string
	SetTag(v []*DescribeTagsRequestTag) *DescribeTagsRequest
	GetTag() []*DescribeTagsRequestTag
}

type DescribeTagsRequest struct {
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           *string                   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*DescribeTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTagsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagsRequest) GetResourceId() *string {
	return s.ResourceId
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

func (s *DescribeTagsRequest) GetTag() []*DescribeTagsRequestTag {
	return s.Tag
}

func (s *DescribeTagsRequest) SetCategory(v string) *DescribeTagsRequest {
	s.Category = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerId(v int64) *DescribeTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetPageNumber(v int32) *DescribeTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsRequest) SetPageSize(v int32) *DescribeTagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceId(v string) *DescribeTagsRequest {
	s.ResourceId = &v
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

func (s *DescribeTagsRequest) SetTag(v []*DescribeTagsRequestTag) *DescribeTagsRequest {
	s.Tag = v
	return s
}

func (s *DescribeTagsRequest) Validate() error {
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

type DescribeTagsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTagsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTagsRequestTag) SetKey(v string) *DescribeTagsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeTagsRequestTag) SetValue(v string) *DescribeTagsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeTagsRequestTag) Validate() error {
	return dara.Validate(s)
}
