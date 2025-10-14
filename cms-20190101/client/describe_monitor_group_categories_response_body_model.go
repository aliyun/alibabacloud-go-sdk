// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeMonitorGroupCategoriesResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeMonitorGroupCategoriesResponseBody
	GetMessage() *string
	SetMonitorGroupCategories(v *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) *DescribeMonitorGroupCategoriesResponseBody
	GetMonitorGroupCategories() *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories
	SetRequestId(v string) *DescribeMonitorGroupCategoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMonitorGroupCategoriesResponseBody
	GetSuccess() *bool
}

type DescribeMonitorGroupCategoriesResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The cloud services to which the resources in the application group belong and the number of resources that belong to the cloud service.
	MonitorGroupCategories *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories `json:"MonitorGroupCategories,omitempty" xml:"MonitorGroupCategories,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9E0347B0-EBC3-4769-A78D-D96F21C6BB52
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitorGroupCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeMonitorGroupCategoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitorGroupCategoriesResponseBody) GetMonitorGroupCategories() *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories {
	return s.MonitorGroupCategories
}

func (s *DescribeMonitorGroupCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorGroupCategoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitorGroupCategoriesResponseBody) SetCode(v int32) *DescribeMonitorGroupCategoriesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBody) SetMessage(v string) *DescribeMonitorGroupCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBody) SetMonitorGroupCategories(v *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) *DescribeMonitorGroupCategoriesResponseBody {
	s.MonitorGroupCategories = v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBody) SetRequestId(v string) *DescribeMonitorGroupCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBody) SetSuccess(v bool) *DescribeMonitorGroupCategoriesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBody) Validate() error {
	if s.MonitorGroupCategories != nil {
		if err := s.MonitorGroupCategories.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories struct {
	// The ID of the application group.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The cloud services to which the resources in the application group belong and the number of resources that belong to the cloud service.
	MonitorGroupCategory *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory `json:"MonitorGroupCategory,omitempty" xml:"MonitorGroupCategory,omitempty" type:"Struct"`
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) GetMonitorGroupCategory() *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory {
	return s.MonitorGroupCategory
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) SetGroupId(v int64) *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) SetMonitorGroupCategory(v *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory) *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories {
	s.MonitorGroupCategory = v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) Validate() error {
	if s.MonitorGroupCategory != nil {
		if err := s.MonitorGroupCategory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory struct {
	CategoryItem []*DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem `json:"CategoryItem,omitempty" xml:"CategoryItem,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory) GetCategoryItem() []*DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem {
	return s.CategoryItem
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory) SetCategoryItem(v []*DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory {
	s.CategoryItem = v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory) Validate() error {
	if s.CategoryItem != nil {
		for _, item := range s.CategoryItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem struct {
	// The abbreviation of the cloud service name.
	//
	// >  For more information about how to obtain the abbreviation of a cloud service name, see `metricCategory` in the response parameter `Labels` of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
	//
	// example:
	//
	// ECS
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The number of resources that belong to the cloud service.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) GetCategory() *string {
	return s.Category
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) GetCount() *int32 {
	return s.Count
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) SetCategory(v string) *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) SetCount(v int32) *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem {
	s.Count = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) Validate() error {
	return dara.Validate(s)
}
