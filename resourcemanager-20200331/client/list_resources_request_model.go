// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
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
	// For more information about the supported resource types, see the **Resource type*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource types. A maximum of 50 resource types are supported.
	//
	// >  If you configure `ResourceTypes`, you must configure both `Service` and `ResourceType`. Otherwise, the configured Service or ResourceType does not take effect.
	ResourceTypes []*ListResourcesRequestResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud service.
	//
	// You can obtain the ID from the **Service code*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
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
	// For more information about the supported resource types, see the **Resource type*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// >  You must configure both `Service` and `ResourceType` in `ResourceTypes`. Otherwise, the two parameters do not take effect.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the Alibaba Cloud service.
	//
	// Valid values of N: 1 to 50.
	//
	// You can obtain the ID from the **Service code*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// >  You must configure both `Service` and `ResourceType` in `ResourceTypes`. Otherwise, the two parameters do not take effect.
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
