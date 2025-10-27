// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScaItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyScaItemResponseBodyPageInfo) *DescribePropertyScaItemResponseBody
	GetPageInfo() *DescribePropertyScaItemResponseBodyPageInfo
	SetPropertyItems(v []*DescribePropertyScaItemResponseBodyPropertyItems) *DescribePropertyScaItemResponseBody
	GetPropertyItems() []*DescribePropertyScaItemResponseBodyPropertyItems
	SetRequestId(v string) *DescribePropertyScaItemResponseBody
	GetRequestId() *string
}

type DescribePropertyScaItemResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyScaItemResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the information about middleware fingerprints.
	PropertyItems []*DescribePropertyScaItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3F4236AB-7070-538D-85EB-98EBFE6C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyScaItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaItemResponseBody) GetPageInfo() *DescribePropertyScaItemResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyScaItemResponseBody) GetPropertyItems() []*DescribePropertyScaItemResponseBodyPropertyItems {
	return s.PropertyItems
}

func (s *DescribePropertyScaItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyScaItemResponseBody) SetPageInfo(v *DescribePropertyScaItemResponseBodyPageInfo) *DescribePropertyScaItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyScaItemResponseBody) SetPropertyItems(v []*DescribePropertyScaItemResponseBodyPropertyItems) *DescribePropertyScaItemResponseBody {
	s.PropertyItems = v
	return s
}

func (s *DescribePropertyScaItemResponseBody) SetRequestId(v string) *DescribePropertyScaItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyScaItemResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PropertyItems != nil {
		for _, item := range s.PropertyItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePropertyScaItemResponseBodyPageInfo struct {
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
	// 27
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyScaItemResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaItemResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyScaItemResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyScaItemResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyScaItemResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyScaItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertyScaItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyScaItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyScaItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyScaItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyScaItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyScaItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyScaItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyScaItemResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyScaItemResponseBodyPropertyItems struct {
	// The type of the middleware, database, or web service. Valid values:
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
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The number of servers on which the middleware is run.
	//
	// example:
	//
	// 23
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the middleware.
	//
	// example:
	//
	// kubelet
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The display name of the middleware type.
	//
	// example:
	//
	// Docker Component
	TypeDisplay *string `json:"TypeDisplay,omitempty" xml:"TypeDisplay,omitempty"`
}

func (s DescribePropertyScaItemResponseBodyPropertyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaItemResponseBodyPropertyItems) GetBizType() *string {
	return s.BizType
}

func (s *DescribePropertyScaItemResponseBodyPropertyItems) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyScaItemResponseBodyPropertyItems) GetName() *string {
	return s.Name
}

func (s *DescribePropertyScaItemResponseBodyPropertyItems) GetTypeDisplay() *string {
	return s.TypeDisplay
}

func (s *DescribePropertyScaItemResponseBodyPropertyItems) SetBizType(v string) *DescribePropertyScaItemResponseBodyPropertyItems {
	s.BizType = &v
	return s
}

func (s *DescribePropertyScaItemResponseBodyPropertyItems) SetCount(v int32) *DescribePropertyScaItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

func (s *DescribePropertyScaItemResponseBodyPropertyItems) SetName(v string) *DescribePropertyScaItemResponseBodyPropertyItems {
	s.Name = &v
	return s
}

func (s *DescribePropertyScaItemResponseBodyPropertyItems) SetTypeDisplay(v string) *DescribePropertyScaItemResponseBodyPropertyItems {
	s.TypeDisplay = &v
	return s
}

func (s *DescribePropertyScaItemResponseBodyPropertyItems) Validate() error {
	return dara.Validate(s)
}
