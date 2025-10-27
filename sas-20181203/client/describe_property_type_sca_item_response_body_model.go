// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyTypeScaItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyTypeScaItemResponseBodyPageInfo) *DescribePropertyTypeScaItemResponseBody
	GetPageInfo() *DescribePropertyTypeScaItemResponseBodyPageInfo
	SetPropertyTypeItems(v []*DescribePropertyTypeScaItemResponseBodyPropertyTypeItems) *DescribePropertyTypeScaItemResponseBody
	GetPropertyTypeItems() []*DescribePropertyTypeScaItemResponseBodyPropertyTypeItems
	SetRequestId(v string) *DescribePropertyTypeScaItemResponseBody
	GetRequestId() *string
}

type DescribePropertyTypeScaItemResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyTypeScaItemResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the middleware types.
	PropertyTypeItems []*DescribePropertyTypeScaItemResponseBodyPropertyTypeItems `json:"PropertyTypeItems,omitempty" xml:"PropertyTypeItems,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// B7A839E8-70AE-591D-8D9E-C5419A22****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyTypeScaItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyTypeScaItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyTypeScaItemResponseBody) GetPageInfo() *DescribePropertyTypeScaItemResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyTypeScaItemResponseBody) GetPropertyTypeItems() []*DescribePropertyTypeScaItemResponseBodyPropertyTypeItems {
	return s.PropertyTypeItems
}

func (s *DescribePropertyTypeScaItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyTypeScaItemResponseBody) SetPageInfo(v *DescribePropertyTypeScaItemResponseBodyPageInfo) *DescribePropertyTypeScaItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyTypeScaItemResponseBody) SetPropertyTypeItems(v []*DescribePropertyTypeScaItemResponseBodyPropertyTypeItems) *DescribePropertyTypeScaItemResponseBody {
	s.PropertyTypeItems = v
	return s
}

func (s *DescribePropertyTypeScaItemResponseBody) SetRequestId(v string) *DescribePropertyTypeScaItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyTypeScaItemResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PropertyTypeItems != nil {
		for _, item := range s.PropertyTypeItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePropertyTypeScaItemResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyTypeScaItemResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyTypeScaItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyTypeScaItemResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyTypeScaItemResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyTypeScaItemResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyTypeScaItemResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyTypeScaItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertyTypeScaItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyTypeScaItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyTypeScaItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyTypeScaItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyTypeScaItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyTypeScaItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyTypeScaItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyTypeScaItemResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyTypeScaItemResponseBodyPropertyTypeItems struct {
	// The name of the middleware type.
	//
	// example:
	//
	// Docker Component
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the middleware. Valid values:
	//
	// 	- **system_service**: system service
	//
	// 	- **software_library**: software library
	//
	// 	- **docker_component**: container component
	//
	// 	- **database**: database
	//
	// 	- **web_container**: web container
	//
	// 	- **jar**: JAR package
	//
	// 	- **web_framework**: web framework
	//
	// example:
	//
	// docker_component
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePropertyTypeScaItemResponseBodyPropertyTypeItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyTypeScaItemResponseBodyPropertyTypeItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyTypeScaItemResponseBodyPropertyTypeItems) GetName() *string {
	return s.Name
}

func (s *DescribePropertyTypeScaItemResponseBodyPropertyTypeItems) GetType() *string {
	return s.Type
}

func (s *DescribePropertyTypeScaItemResponseBodyPropertyTypeItems) SetName(v string) *DescribePropertyTypeScaItemResponseBodyPropertyTypeItems {
	s.Name = &v
	return s
}

func (s *DescribePropertyTypeScaItemResponseBodyPropertyTypeItems) SetType(v string) *DescribePropertyTypeScaItemResponseBodyPropertyTypeItems {
	s.Type = &v
	return s
}

func (s *DescribePropertyTypeScaItemResponseBodyPropertyTypeItems) Validate() error {
	return dara.Validate(s)
}
