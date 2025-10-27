// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyCronItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyCronItemResponseBodyPageInfo) *DescribePropertyCronItemResponseBody
	GetPageInfo() *DescribePropertyCronItemResponseBodyPageInfo
	SetPropertyItems(v []*DescribePropertyCronItemResponseBodyPropertyItems) *DescribePropertyCronItemResponseBody
	GetPropertyItems() []*DescribePropertyCronItemResponseBodyPropertyItems
	SetRequestId(v string) *DescribePropertyCronItemResponseBody
	GetRequestId() *string
}

type DescribePropertyCronItemResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyCronItemResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of scheduled tasks.
	PropertyItems []*DescribePropertyCronItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 291B49F9-1685-4005-9D34-606B6F78****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyCronItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCronItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronItemResponseBody) GetPageInfo() *DescribePropertyCronItemResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyCronItemResponseBody) GetPropertyItems() []*DescribePropertyCronItemResponseBodyPropertyItems {
	return s.PropertyItems
}

func (s *DescribePropertyCronItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyCronItemResponseBody) SetPageInfo(v *DescribePropertyCronItemResponseBodyPageInfo) *DescribePropertyCronItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyCronItemResponseBody) SetPropertyItems(v []*DescribePropertyCronItemResponseBodyPropertyItems) *DescribePropertyCronItemResponseBody {
	s.PropertyItems = v
	return s
}

func (s *DescribePropertyCronItemResponseBody) SetRequestId(v string) *DescribePropertyCronItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyCronItemResponseBody) Validate() error {
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

type DescribePropertyCronItemResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 11
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
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyCronItemResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCronItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronItemResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyCronItemResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyCronItemResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyCronItemResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyCronItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertyCronItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyCronItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyCronItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyCronItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyCronItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyCronItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyCronItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyCronItemResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyCronItemResponseBodyPropertyItems struct {
	// The number of servers on which the scheduled task is run.
	//
	// example:
	//
	// 23
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The path to the scheduled task.
	//
	// example:
	//
	// /data
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribePropertyCronItemResponseBodyPropertyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCronItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronItemResponseBodyPropertyItems) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyCronItemResponseBodyPropertyItems) GetSource() *string {
	return s.Source
}

func (s *DescribePropertyCronItemResponseBodyPropertyItems) SetCount(v int32) *DescribePropertyCronItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

func (s *DescribePropertyCronItemResponseBodyPropertyItems) SetSource(v string) *DescribePropertyCronItemResponseBodyPropertyItems {
	s.Source = &v
	return s
}

func (s *DescribePropertyCronItemResponseBodyPropertyItems) Validate() error {
	return dara.Validate(s)
}
