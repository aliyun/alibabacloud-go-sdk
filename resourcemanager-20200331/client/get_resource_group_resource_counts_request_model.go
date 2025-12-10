// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupResourceCountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupByKey(v string) *GetResourceGroupResourceCountsRequest
	GetGroupByKey() *string
	SetResourceGroupId(v string) *GetResourceGroupResourceCountsRequest
	GetResourceGroupId() *string
	SetResourceRegionId(v string) *GetResourceGroupResourceCountsRequest
	GetResourceRegionId() *string
	SetResourceTypes(v []*GetResourceGroupResourceCountsRequestResourceTypes) *GetResourceGroupResourceCountsRequest
	GetResourceTypes() []*GetResourceGroupResourceCountsRequestResourceTypes
}

type GetResourceGroupResourceCountsRequest struct {
	// The dimension by which resources are queried.
	//
	// > If you do not specify a dimension, no results are returned.
	//
	// Valid values:
	//
	// 	- ResourceGroupId
	//
	// 	- ResourceType
	//
	// example:
	//
	// ResourceGroupId
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// The resource group ID.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/2716558.html) operation to obtain the ID.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The region ID of the resources.
	//
	// example:
	//
	// cn-shanghai
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource types.
	ResourceTypes []*GetResourceGroupResourceCountsRequestResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s GetResourceGroupResourceCountsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResourceCountsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResourceCountsRequest) GetGroupByKey() *string {
	return s.GroupByKey
}

func (s *GetResourceGroupResourceCountsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetResourceGroupResourceCountsRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *GetResourceGroupResourceCountsRequest) GetResourceTypes() []*GetResourceGroupResourceCountsRequestResourceTypes {
	return s.ResourceTypes
}

func (s *GetResourceGroupResourceCountsRequest) SetGroupByKey(v string) *GetResourceGroupResourceCountsRequest {
	s.GroupByKey = &v
	return s
}

func (s *GetResourceGroupResourceCountsRequest) SetResourceGroupId(v string) *GetResourceGroupResourceCountsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetResourceGroupResourceCountsRequest) SetResourceRegionId(v string) *GetResourceGroupResourceCountsRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GetResourceGroupResourceCountsRequest) SetResourceTypes(v []*GetResourceGroupResourceCountsRequestResourceTypes) *GetResourceGroupResourceCountsRequest {
	s.ResourceTypes = v
	return s
}

func (s *GetResourceGroupResourceCountsRequest) Validate() error {
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

type GetResourceGroupResourceCountsRequestResourceTypes struct {
	// The resource type.
	//
	// You can obtain the resource type from the **Resource type*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// instance
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	// The service code.
	//
	// You can obtain the code from the **Service code*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s GetResourceGroupResourceCountsRequestResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResourceCountsRequestResourceTypes) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResourceCountsRequestResourceTypes) GetResourceTypeCode() *string {
	return s.ResourceTypeCode
}

func (s *GetResourceGroupResourceCountsRequestResourceTypes) GetService() *string {
	return s.Service
}

func (s *GetResourceGroupResourceCountsRequestResourceTypes) SetResourceTypeCode(v string) *GetResourceGroupResourceCountsRequestResourceTypes {
	s.ResourceTypeCode = &v
	return s
}

func (s *GetResourceGroupResourceCountsRequestResourceTypes) SetService(v string) *GetResourceGroupResourceCountsRequestResourceTypes {
	s.Service = &v
	return s
}

func (s *GetResourceGroupResourceCountsRequestResourceTypes) Validate() error {
	return dara.Validate(s)
}
