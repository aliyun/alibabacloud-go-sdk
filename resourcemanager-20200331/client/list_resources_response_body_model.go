// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourcesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourcesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResourcesResponseBody
	GetRequestId() *string
	SetResources(v *ListResourcesResponseBodyResources) *ListResourcesResponseBody
	GetResources() *ListResourcesResponseBodyResources
	SetTotalCount(v int32) *ListResourcesResponseBody
	GetTotalCount() *int32
}

type ListResourcesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resources.
	Resources *ListResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourcesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourcesResponseBody) GetResources() *ListResourcesResponseBodyResources {
	return s.Resources
}

func (s *ListResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourcesResponseBody) SetPageNumber(v int32) *ListResourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponseBody) SetPageSize(v int32) *ListResourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) SetResources(v *ListResourcesResponseBodyResources) *ListResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBody) SetTotalCount(v int32) *ListResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourcesResponseBody) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcesResponseBodyResources struct {
	Resource []*ListResourcesResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s ListResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResources) GetResource() []*ListResourcesResponseBodyResourcesResource {
	return s.Resource
}

func (s *ListResourcesResponseBodyResources) SetResource(v []*ListResourcesResponseBodyResourcesResource) *ListResourcesResponseBodyResources {
	s.Resource = v
	return s
}

func (s *ListResourcesResponseBodyResources) Validate() error {
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

type ListResourcesResponseBodyResourcesResource struct {
	// The time when the resource was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-uPJpP****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// i-23v38****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the Alibaba Cloud service.
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListResourcesResponseBodyResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesResource) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListResourcesResponseBodyResourcesResource) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourcesResponseBodyResourcesResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourcesResponseBodyResourcesResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourcesResponseBodyResourcesResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourcesResponseBodyResourcesResource) GetService() *string {
	return s.Service
}

func (s *ListResourcesResponseBodyResourcesResource) SetCreateDate(v string) *ListResourcesResponseBodyResourcesResource {
	s.CreateDate = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) SetRegionId(v string) *ListResourcesResponseBodyResourcesResource {
	s.RegionId = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) SetResourceGroupId(v string) *ListResourcesResponseBodyResourcesResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) SetResourceId(v string) *ListResourcesResponseBodyResourcesResource {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) SetResourceType(v string) *ListResourcesResponseBodyResourcesResource {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) SetService(v string) *ListResourcesResponseBodyResourcesResource {
	s.Service = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResource) Validate() error {
	return dara.Validate(s)
}
