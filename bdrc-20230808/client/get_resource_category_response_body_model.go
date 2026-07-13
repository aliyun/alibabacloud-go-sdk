// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetResourceCategoryResponseBodyData) *GetResourceCategoryResponseBody
	GetData() *GetResourceCategoryResponseBodyData
	SetRequestId(v string) *GetResourceCategoryResponseBody
	GetRequestId() *string
}

type GetResourceCategoryResponseBody struct {
	// The returned data.
	Data *GetResourceCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// AE43C4CB-8074-5EBD-9806-8CA6D12800B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceCategoryResponseBody) GetData() *GetResourceCategoryResponseBodyData {
	return s.Data
}

func (s *GetResourceCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceCategoryResponseBody) SetData(v *GetResourceCategoryResponseBodyData) *GetResourceCategoryResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceCategoryResponseBody) SetRequestId(v string) *GetResourceCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceCategoryResponseBodyData struct {
	// Applicable product type. If empty, matches all products.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Resource category ID, globally unique.
	//
	// example:
	//
	// rc-123***7890
	ResourceCategoryId *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	// Resource name, unique within the namespace.
	//
	// example:
	//
	// My***ResourceCategory
	ResourceCategoryName *string `json:"ResourceCategoryName,omitempty" xml:"ResourceCategoryName,omitempty"`
	// Resource category type. Valid values:
	//
	// - DEFAULT: default group, created by the system, cannot be deleted.
	//
	// - CUSTOM: custom group, can be deleted.
	//
	// example:
	//
	// CUSTOM
	ResourceCategoryType *string `json:"ResourceCategoryType,omitempty" xml:"ResourceCategoryType,omitempty"`
	// Number of resources by type.
	ResourceCount []*GetResourceCategoryResponseBodyDataResourceCount `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty" type:"Repeated"`
	// Resource matcher. If empty, no resources are matched.
	//
	// example:
	//
	// {\\"type\\":\\"BOOL\\",\\"operator\\":\\"AND\\",\\"values\\":[{\\"type\\":\\"TAG\\",\\"key\\":\\"createdBy\\",\\"operator\\":\\"EQUAL\\",\\"values\\":[\\"me\\"]}]}
	ResourceMatcher *string `json:"ResourceMatcher,omitempty" xml:"ResourceMatcher,omitempty"`
	// Applicable resource type. If empty, matches all resources.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceCategoryResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *GetResourceCategoryResponseBodyData) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *GetResourceCategoryResponseBodyData) GetResourceCategoryName() *string {
	return s.ResourceCategoryName
}

func (s *GetResourceCategoryResponseBodyData) GetResourceCategoryType() *string {
	return s.ResourceCategoryType
}

func (s *GetResourceCategoryResponseBodyData) GetResourceCount() []*GetResourceCategoryResponseBodyDataResourceCount {
	return s.ResourceCount
}

func (s *GetResourceCategoryResponseBodyData) GetResourceMatcher() *string {
	return s.ResourceMatcher
}

func (s *GetResourceCategoryResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceCategoryResponseBodyData) SetProductType(v string) *GetResourceCategoryResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetResourceCategoryResponseBodyData) SetResourceCategoryId(v string) *GetResourceCategoryResponseBodyData {
	s.ResourceCategoryId = &v
	return s
}

func (s *GetResourceCategoryResponseBodyData) SetResourceCategoryName(v string) *GetResourceCategoryResponseBodyData {
	s.ResourceCategoryName = &v
	return s
}

func (s *GetResourceCategoryResponseBodyData) SetResourceCategoryType(v string) *GetResourceCategoryResponseBodyData {
	s.ResourceCategoryType = &v
	return s
}

func (s *GetResourceCategoryResponseBodyData) SetResourceCount(v []*GetResourceCategoryResponseBodyDataResourceCount) *GetResourceCategoryResponseBodyData {
	s.ResourceCount = v
	return s
}

func (s *GetResourceCategoryResponseBodyData) SetResourceMatcher(v string) *GetResourceCategoryResponseBodyData {
	s.ResourceMatcher = &v
	return s
}

func (s *GetResourceCategoryResponseBodyData) SetResourceType(v string) *GetResourceCategoryResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetResourceCategoryResponseBodyData) Validate() error {
	if s.ResourceCount != nil {
		for _, item := range s.ResourceCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceCategoryResponseBodyDataResourceCount struct {
	// Number of resources by type.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceCategoryResponseBodyDataResourceCount) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCategoryResponseBodyDataResourceCount) GoString() string {
	return s.String()
}

func (s *GetResourceCategoryResponseBodyDataResourceCount) GetCount() *int64 {
	return s.Count
}

func (s *GetResourceCategoryResponseBodyDataResourceCount) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceCategoryResponseBodyDataResourceCount) SetCount(v int64) *GetResourceCategoryResponseBodyDataResourceCount {
	s.Count = &v
	return s
}

func (s *GetResourceCategoryResponseBodyDataResourceCount) SetResourceType(v string) *GetResourceCategoryResponseBodyDataResourceCount {
	s.ResourceType = &v
	return s
}

func (s *GetResourceCategoryResponseBodyDataResourceCount) Validate() error {
	return dara.Validate(s)
}
