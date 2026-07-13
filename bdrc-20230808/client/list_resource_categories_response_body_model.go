// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListResourceCategoriesResponseBodyData) *ListResourceCategoriesResponseBody
	GetData() *ListResourceCategoriesResponseBodyData
	SetRequestId(v string) *ListResourceCategoriesResponseBody
	GetRequestId() *string
}

type ListResourceCategoriesResponseBody struct {
	// The returned data.
	Data *ListResourceCategoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// AE43C4CB-8074-5EBD-9806-8CA6D12800B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourceCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceCategoriesResponseBody) GetData() *ListResourceCategoriesResponseBodyData {
	return s.Data
}

func (s *ListResourceCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceCategoriesResponseBody) SetData(v *ListResourceCategoriesResponseBodyData) *ListResourceCategoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListResourceCategoriesResponseBody) SetRequestId(v string) *ListResourceCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceCategoriesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceCategoriesResponseBodyData struct {
	// The collection of records returned in this request.
	Content []*ListResourceCategoriesResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The maximum number of records returned in this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Indicates the position where the current call returns data from. An empty value indicates that all data has been read.
	//
	// example:
	//
	// eKDyCM0zFQ5op7jVMWmNNA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of data entries under the current request conditions. This parameter is optional and can be left unspecified by default.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceCategoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListResourceCategoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListResourceCategoriesResponseBodyData) GetContent() []*ListResourceCategoriesResponseBodyDataContent {
	return s.Content
}

func (s *ListResourceCategoriesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceCategoriesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceCategoriesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListResourceCategoriesResponseBodyData) SetContent(v []*ListResourceCategoriesResponseBodyDataContent) *ListResourceCategoriesResponseBodyData {
	s.Content = v
	return s
}

func (s *ListResourceCategoriesResponseBodyData) SetMaxResults(v int32) *ListResourceCategoriesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyData) SetNextToken(v string) *ListResourceCategoriesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyData) SetTotalCount(v int64) *ListResourceCategoriesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceCategoriesResponseBodyDataContent struct {
	// The applicable product type. If this parameter is empty, all products are matched.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The resource category ID, which is globally unique.
	//
	// example:
	//
	// rc-123***7890
	ResourceCategoryId *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	// The resource name, which is unique within the namespace.
	//
	// example:
	//
	// My***ResourceCategory
	ResourceCategoryName *string `json:"ResourceCategoryName,omitempty" xml:"ResourceCategoryName,omitempty"`
	// The resource category type. Valid values:
	//
	// - DEFAULT: default group created by the system, cannot be deleted.
	//
	// - CUSTOM: custom group, can be deleted.
	//
	// example:
	//
	// CUSTOM
	ResourceCategoryType *string `json:"ResourceCategoryType,omitempty" xml:"ResourceCategoryType,omitempty"`
	// The number of resources of each type.
	ResourceCount []*ListResourceCategoriesResponseBodyDataContentResourceCount `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty" type:"Repeated"`
	// The resource matcher. If this parameter is empty, no resources are matched.
	//
	// example:
	//
	// {\\"type\\":\\"BOOL\\",\\"operator\\":\\"AND\\",\\"values\\":[{\\"type\\":\\"TAG\\",\\"key\\":\\"createdBy\\",\\"operator\\":\\"EQUAL\\",\\"values\\":[\\"me\\"]}]}
	ResourceMatcher *string `json:"ResourceMatcher,omitempty" xml:"ResourceMatcher,omitempty"`
	// The applicable resource type. If this parameter is empty, all resources are matched.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceCategoriesResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListResourceCategoriesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListResourceCategoriesResponseBodyDataContent) GetProductType() *string {
	return s.ProductType
}

func (s *ListResourceCategoriesResponseBodyDataContent) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *ListResourceCategoriesResponseBodyDataContent) GetResourceCategoryName() *string {
	return s.ResourceCategoryName
}

func (s *ListResourceCategoriesResponseBodyDataContent) GetResourceCategoryType() *string {
	return s.ResourceCategoryType
}

func (s *ListResourceCategoriesResponseBodyDataContent) GetResourceCount() []*ListResourceCategoriesResponseBodyDataContentResourceCount {
	return s.ResourceCount
}

func (s *ListResourceCategoriesResponseBodyDataContent) GetResourceMatcher() *string {
	return s.ResourceMatcher
}

func (s *ListResourceCategoriesResponseBodyDataContent) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceCategoriesResponseBodyDataContent) SetProductType(v string) *ListResourceCategoriesResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyDataContent) SetResourceCategoryId(v string) *ListResourceCategoriesResponseBodyDataContent {
	s.ResourceCategoryId = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyDataContent) SetResourceCategoryName(v string) *ListResourceCategoriesResponseBodyDataContent {
	s.ResourceCategoryName = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyDataContent) SetResourceCategoryType(v string) *ListResourceCategoriesResponseBodyDataContent {
	s.ResourceCategoryType = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyDataContent) SetResourceCount(v []*ListResourceCategoriesResponseBodyDataContentResourceCount) *ListResourceCategoriesResponseBodyDataContent {
	s.ResourceCount = v
	return s
}

func (s *ListResourceCategoriesResponseBodyDataContent) SetResourceMatcher(v string) *ListResourceCategoriesResponseBodyDataContent {
	s.ResourceMatcher = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyDataContent) SetResourceType(v string) *ListResourceCategoriesResponseBodyDataContent {
	s.ResourceType = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyDataContent) Validate() error {
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

type ListResourceCategoriesResponseBodyDataContentResourceCount struct {
	// The number of resources of each type.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceCategoriesResponseBodyDataContentResourceCount) String() string {
	return dara.Prettify(s)
}

func (s ListResourceCategoriesResponseBodyDataContentResourceCount) GoString() string {
	return s.String()
}

func (s *ListResourceCategoriesResponseBodyDataContentResourceCount) GetCount() *int64 {
	return s.Count
}

func (s *ListResourceCategoriesResponseBodyDataContentResourceCount) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceCategoriesResponseBodyDataContentResourceCount) SetCount(v int64) *ListResourceCategoriesResponseBodyDataContentResourceCount {
	s.Count = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyDataContentResourceCount) SetResourceType(v string) *ListResourceCategoriesResponseBodyDataContentResourceCount {
	s.ResourceType = &v
	return s
}

func (s *ListResourceCategoriesResponseBodyDataContentResourceCount) Validate() error {
	return dara.Validate(s)
}
