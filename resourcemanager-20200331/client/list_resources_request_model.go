// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v bool) *ListResourcesRequest
	GetCount() *bool
	SetPageNumber(v int32) *ListResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourcesRequest
	GetPageSize() *int32
	SetRegion(v string) *ListResourcesRequest
	GetRegion() *string
	SetResourceGroupId(v string) *ListResourcesRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ListResourcesRequest
	GetResourceId() *string
	SetResourceType(v string) *ListResourcesRequest
	GetResourceType() *string
	SetResourceTypes(v []*ListResourcesRequestResourceTypes) *ListResourcesRequest
	GetResourceTypes() []*ListResourcesRequestResourceTypes
	SetService(v string) *ListResourcesRequest
	GetService() *string
}

type ListResourcesRequest struct {
	// Specifies whether to query the total number of resources. Default value: true. Valid values:
	//
	// - true: Queries the total number of resources and returns the result in the TotalCount parameter.
	//
	// - false: Does not query the total number of resources. The TotalCount parameter returns -1, and the resources are still paginated based on the PageNumber and PageSize parameters.
	Count *bool `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
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
	// You can obtain the resource type from the **Resource Type*	- column in [Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of resource types. A maximum of 50 resource types are supported.
	//
	// > If you configure the Settings for ResourceTypes, the Service and ResourceType request parameter specified separately do not take effect.
	ResourceTypes []*ListResourcesRequestResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The code of the cloud service.
	//
	// You can obtain the code from the **Cloud Service Code*	- column in [Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) GetCount() *bool {
	return s.Count
}

func (s *ListResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesRequest) GetRegion() *string {
	return s.Region
}

func (s *ListResourcesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourcesRequest) GetResourceTypes() []*ListResourcesRequestResourceTypes {
	return s.ResourceTypes
}

func (s *ListResourcesRequest) GetService() *string {
	return s.Service
}

func (s *ListResourcesRequest) SetCount(v bool) *ListResourcesRequest {
	s.Count = &v
	return s
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetRegion(v string) *ListResourcesRequest {
	s.Region = &v
	return s
}

func (s *ListResourcesRequest) SetResourceGroupId(v string) *ListResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourcesRequest) SetResourceId(v string) *ListResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesRequest) SetResourceType(v string) *ListResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesRequest) SetResourceTypes(v []*ListResourcesRequestResourceTypes) *ListResourcesRequest {
	s.ResourceTypes = v
	return s
}

func (s *ListResourcesRequest) SetService(v string) *ListResourcesRequest {
	s.Service = &v
	return s
}

func (s *ListResourcesRequest) Validate() error {
	if s.ResourceTypes != nil {
		for _, item := range s.ResourceTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourcesRequestResourceTypes struct {
	// The resource type.
	//
	// Valid values of N: 1 to 50.
	//
	// You can obtain the resource type from the **Resource Type*	- column in [Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
	//
	// > The Service and ResourceType parameters in ResourceTypes must be specified at the same time. Otherwise, they do not take effect.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The code of the cloud service.
	//
	// Valid values of N: 1 to 50.
	//
	// You can obtain the code from the **Cloud Service Code*	- column in [Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
	//
	// > The Service and ResourceType parameters in ResourceTypes must be specified at the same time. Otherwise, they do not take effect.
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListResourcesRequestResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesRequestResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourcesRequestResourceTypes) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourcesRequestResourceTypes) GetService() *string {
	return s.Service
}

func (s *ListResourcesRequestResourceTypes) SetResourceType(v string) *ListResourcesRequestResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesRequestResourceTypes) SetService(v string) *ListResourcesRequestResourceTypes {
	s.Service = &v
	return s
}

func (s *ListResourcesRequestResourceTypes) Validate() error {
	return dara.Validate(s)
}
