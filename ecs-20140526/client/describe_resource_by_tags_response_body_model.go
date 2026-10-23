// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceByTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeResourceByTagsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeResourceByTagsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeResourceByTagsResponseBody
	GetRequestId() *string
	SetResources(v *DescribeResourceByTagsResponseBodyResources) *DescribeResourceByTagsResponseBody
	GetResources() *DescribeResourceByTagsResponseBodyResources
	SetTotalCount(v int32) *DescribeResourceByTagsResponseBody
	GetTotalCount() *int32
}

type DescribeResourceByTagsResponseBody struct {
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  *DescribeResourceByTagsResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourceByTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceByTagsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeResourceByTagsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeResourceByTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceByTagsResponseBody) GetResources() *DescribeResourceByTagsResponseBodyResources {
	return s.Resources
}

func (s *DescribeResourceByTagsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeResourceByTagsResponseBody) SetPageNumber(v int32) *DescribeResourceByTagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceByTagsResponseBody) SetPageSize(v int32) *DescribeResourceByTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceByTagsResponseBody) SetRequestId(v string) *DescribeResourceByTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceByTagsResponseBody) SetResources(v *DescribeResourceByTagsResponseBodyResources) *DescribeResourceByTagsResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeResourceByTagsResponseBody) SetTotalCount(v int32) *DescribeResourceByTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeResourceByTagsResponseBody) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourceByTagsResponseBodyResources struct {
	Resource []*DescribeResourceByTagsResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeResourceByTagsResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByTagsResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeResourceByTagsResponseBodyResources) GetResource() []*DescribeResourceByTagsResponseBodyResourcesResource {
	return s.Resource
}

func (s *DescribeResourceByTagsResponseBodyResources) SetResource(v []*DescribeResourceByTagsResponseBodyResourcesResource) *DescribeResourceByTagsResponseBodyResources {
	s.Resource = v
	return s
}

func (s *DescribeResourceByTagsResponseBodyResources) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourceByTagsResponseBodyResourcesResource struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeResourceByTagsResponseBodyResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByTagsResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeResourceByTagsResponseBodyResourcesResource) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceByTagsResponseBodyResourcesResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeResourceByTagsResponseBodyResourcesResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourceByTagsResponseBodyResourcesResource) SetRegionId(v string) *DescribeResourceByTagsResponseBodyResourcesResource {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceByTagsResponseBodyResourcesResource) SetResourceId(v string) *DescribeResourceByTagsResponseBodyResourcesResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourceByTagsResponseBodyResourcesResource) SetResourceType(v string) *DescribeResourceByTagsResponseBodyResourcesResource {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourceByTagsResponseBodyResourcesResource) Validate() error {
	return dara.Validate(s)
}
