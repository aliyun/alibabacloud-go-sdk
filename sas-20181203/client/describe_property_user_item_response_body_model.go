// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyUserItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyUserItemResponseBodyPageInfo) *DescribePropertyUserItemResponseBody
	GetPageInfo() *DescribePropertyUserItemResponseBodyPageInfo
	SetPropertyItems(v []*DescribePropertyUserItemResponseBodyPropertyItems) *DescribePropertyUserItemResponseBody
	GetPropertyItems() []*DescribePropertyUserItemResponseBodyPropertyItems
	SetRequestId(v string) *DescribePropertyUserItemResponseBody
	GetRequestId() *string
}

type DescribePropertyUserItemResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyUserItemResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the account information returned.
	PropertyItems []*DescribePropertyUserItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 76068BE2-F9C4-4EDD-967B-F503B8CCDD3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyUserItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUserItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserItemResponseBody) GetPageInfo() *DescribePropertyUserItemResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyUserItemResponseBody) GetPropertyItems() []*DescribePropertyUserItemResponseBodyPropertyItems {
	return s.PropertyItems
}

func (s *DescribePropertyUserItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyUserItemResponseBody) SetPageInfo(v *DescribePropertyUserItemResponseBodyPageInfo) *DescribePropertyUserItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyUserItemResponseBody) SetPropertyItems(v []*DescribePropertyUserItemResponseBodyPropertyItems) *DescribePropertyUserItemResponseBody {
	s.PropertyItems = v
	return s
}

func (s *DescribePropertyUserItemResponseBody) SetRequestId(v string) *DescribePropertyUserItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyUserItemResponseBody) Validate() error {
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

type DescribePropertyUserItemResponseBodyPageInfo struct {
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
	// 114
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyUserItemResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUserItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertyUserItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyUserItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyUserItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyUserItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyUserItemResponseBodyPropertyItems struct {
	// The number of servers that belong to the account.
	//
	// example:
	//
	// 384
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// adm
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribePropertyUserItemResponseBodyPropertyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUserItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserItemResponseBodyPropertyItems) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyUserItemResponseBodyPropertyItems) GetUser() *string {
	return s.User
}

func (s *DescribePropertyUserItemResponseBodyPropertyItems) SetCount(v int32) *DescribePropertyUserItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

func (s *DescribePropertyUserItemResponseBodyPropertyItems) SetUser(v string) *DescribePropertyUserItemResponseBodyPropertyItems {
	s.User = &v
	return s
}

func (s *DescribePropertyUserItemResponseBodyPropertyItems) Validate() error {
	return dara.Validate(s)
}
