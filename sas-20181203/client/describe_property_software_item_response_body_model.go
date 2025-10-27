// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertySoftwareItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertySoftwareItemResponseBodyPageInfo) *DescribePropertySoftwareItemResponseBody
	GetPageInfo() *DescribePropertySoftwareItemResponseBodyPageInfo
	SetPropertyItems(v []*DescribePropertySoftwareItemResponseBodyPropertyItems) *DescribePropertySoftwareItemResponseBody
	GetPropertyItems() []*DescribePropertySoftwareItemResponseBodyPropertyItems
	SetRequestId(v string) *DescribePropertySoftwareItemResponseBody
	GetRequestId() *string
}

type DescribePropertySoftwareItemResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertySoftwareItemResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the software assets returned.
	PropertyItems []*DescribePropertySoftwareItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 3A85CFCF-05C8-451A-9E41-C0D5E96BA407
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertySoftwareItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertySoftwareItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareItemResponseBody) GetPageInfo() *DescribePropertySoftwareItemResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertySoftwareItemResponseBody) GetPropertyItems() []*DescribePropertySoftwareItemResponseBodyPropertyItems {
	return s.PropertyItems
}

func (s *DescribePropertySoftwareItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertySoftwareItemResponseBody) SetPageInfo(v *DescribePropertySoftwareItemResponseBodyPageInfo) *DescribePropertySoftwareItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertySoftwareItemResponseBody) SetPropertyItems(v []*DescribePropertySoftwareItemResponseBodyPropertyItems) *DescribePropertySoftwareItemResponseBody {
	s.PropertyItems = v
	return s
}

func (s *DescribePropertySoftwareItemResponseBody) SetRequestId(v string) *DescribePropertySoftwareItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBody) Validate() error {
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

type DescribePropertySoftwareItemResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
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
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5037
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertySoftwareItemResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertySoftwareItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertySoftwareItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertySoftwareItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertySoftwareItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertySoftwareItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertySoftwareItemResponseBodyPropertyItems struct {
	// The number of servers on which the software is installed.
	//
	// example:
	//
	// 23
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the software.
	//
	// example:
	//
	// aaa_base
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePropertySoftwareItemResponseBodyPropertyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertySoftwareItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareItemResponseBodyPropertyItems) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertySoftwareItemResponseBodyPropertyItems) GetName() *string {
	return s.Name
}

func (s *DescribePropertySoftwareItemResponseBodyPropertyItems) SetCount(v int32) *DescribePropertySoftwareItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBodyPropertyItems) SetName(v string) *DescribePropertySoftwareItemResponseBodyPropertyItems {
	s.Name = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBodyPropertyItems) Validate() error {
	return dara.Validate(s)
}
