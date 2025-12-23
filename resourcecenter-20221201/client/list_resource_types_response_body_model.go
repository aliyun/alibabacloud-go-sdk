// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListResourceTypesResponseBody
	GetRequestId() *string
	SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody
	GetResourceTypes() []*ListResourceTypesResponseBodyResourceTypes
}

type ListResourceTypesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E5556E4C-479A-5BBB-B325-F07563E7E917
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource types.
	ResourceTypes []*ListResourceTypesResponseBodyResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceTypesResponseBody) GetResourceTypes() []*ListResourceTypesResponseBodyResourceTypes {
	return s.ResourceTypes
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *ListResourceTypesResponseBody) Validate() error {
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

type ListResourceTypesResponseBodyResourceTypes struct {
	// The code mapping of the resource type.
	CodeMapping *ListResourceTypesResponseBodyResourceTypesCodeMapping `json:"CodeMapping,omitempty" xml:"CodeMapping,omitempty" type:"Struct"`
	// The supported filter conditions.
	FilterKeys []*string `json:"FilterKeys,omitempty" xml:"FilterKeys,omitempty" type:"Repeated"`
	// The name of the Alibaba Cloud service.
	//
	// example:
	//
	// Container Service for Kubernetes
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The name of supported related resource types.
	RelatedResourceTypes []*string `json:"RelatedResourceTypes,omitempty" xml:"RelatedResourceTypes,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The name of the resource type.
	//
	// example:
	//
	// Cluster
	ResourceTypeName *string `json:"ResourceTypeName,omitempty" xml:"ResourceTypeName,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetCodeMapping() *ListResourceTypesResponseBodyResourceTypesCodeMapping {
	return s.CodeMapping
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetFilterKeys() []*string {
	return s.FilterKeys
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetProductName() *string {
	return s.ProductName
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetRelatedResourceTypes() []*string {
	return s.RelatedResourceTypes
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceTypesResponseBodyResourceTypes) GetResourceTypeName() *string {
	return s.ResourceTypeName
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetCodeMapping(v *ListResourceTypesResponseBodyResourceTypesCodeMapping) *ListResourceTypesResponseBodyResourceTypes {
	s.CodeMapping = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetFilterKeys(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.FilterKeys = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProductName(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ProductName = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetRelatedResourceTypes(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.RelatedResourceTypes = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceType(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceTypeName(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceTypeName = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) Validate() error {
	if s.CodeMapping != nil {
		if err := s.CodeMapping.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceTypesResponseBodyResourceTypesCodeMapping struct {
	// The resource group.
	//
	// example:
	//
	// cs.cluster
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The tag.
	//
	// example:
	//
	// cs.cluster
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypesCodeMapping) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesCodeMapping) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) GetTag() *string {
	return s.Tag
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) SetResourceGroup(v string) *ListResourceTypesResponseBodyResourceTypesCodeMapping {
	s.ResourceGroup = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) SetTag(v string) *ListResourceTypesResponseBodyResourceTypesCodeMapping {
	s.Tag = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) Validate() error {
	return dara.Validate(s)
}
