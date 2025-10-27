// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyProcItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyProcItemResponseBodyPageInfo) *DescribePropertyProcItemResponseBody
	GetPageInfo() *DescribePropertyProcItemResponseBodyPageInfo
	SetPropertyItems(v []*DescribePropertyProcItemResponseBodyPropertyItems) *DescribePropertyProcItemResponseBody
	GetPropertyItems() []*DescribePropertyProcItemResponseBodyPropertyItems
	SetRequestId(v string) *DescribePropertyProcItemResponseBody
	GetRequestId() *string
}

type DescribePropertyProcItemResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyProcItemResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the processes returned.
	PropertyItems []*DescribePropertyProcItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// BD8AD4C6-A169-4FA3-BA1F-ED40ED52973B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyProcItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyProcItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcItemResponseBody) GetPageInfo() *DescribePropertyProcItemResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyProcItemResponseBody) GetPropertyItems() []*DescribePropertyProcItemResponseBodyPropertyItems {
	return s.PropertyItems
}

func (s *DescribePropertyProcItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyProcItemResponseBody) SetPageInfo(v *DescribePropertyProcItemResponseBodyPageInfo) *DescribePropertyProcItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyProcItemResponseBody) SetPropertyItems(v []*DescribePropertyProcItemResponseBodyPropertyItems) *DescribePropertyProcItemResponseBody {
	s.PropertyItems = v
	return s
}

func (s *DescribePropertyProcItemResponseBody) SetRequestId(v string) *DescribePropertyProcItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyProcItemResponseBody) Validate() error {
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

type DescribePropertyProcItemResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 5
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
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 372
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyProcItemResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyProcItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertyProcItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyProcItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyProcItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyProcItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyProcItemResponseBodyPropertyItems struct {
	// The number of servers on which the process runs.
	//
	// example:
	//
	// 8888
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// .ss
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePropertyProcItemResponseBodyPropertyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyProcItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcItemResponseBodyPropertyItems) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyProcItemResponseBodyPropertyItems) GetName() *string {
	return s.Name
}

func (s *DescribePropertyProcItemResponseBodyPropertyItems) SetCount(v int32) *DescribePropertyProcItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

func (s *DescribePropertyProcItemResponseBodyPropertyItems) SetName(v string) *DescribePropertyProcItemResponseBodyPropertyItems {
	s.Name = &v
	return s
}

func (s *DescribePropertyProcItemResponseBodyPropertyItems) Validate() error {
	return dara.Validate(s)
}
