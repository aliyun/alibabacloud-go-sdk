// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyPortItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyPortItemResponseBodyPageInfo) *DescribePropertyPortItemResponseBody
	GetPageInfo() *DescribePropertyPortItemResponseBodyPageInfo
	SetPropertyItems(v []*DescribePropertyPortItemResponseBodyPropertyItems) *DescribePropertyPortItemResponseBody
	GetPropertyItems() []*DescribePropertyPortItemResponseBodyPropertyItems
	SetRequestId(v string) *DescribePropertyPortItemResponseBody
	GetRequestId() *string
}

type DescribePropertyPortItemResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyPortItemResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the ports returned.
	PropertyItems []*DescribePropertyPortItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyPortItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyPortItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortItemResponseBody) GetPageInfo() *DescribePropertyPortItemResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyPortItemResponseBody) GetPropertyItems() []*DescribePropertyPortItemResponseBodyPropertyItems {
	return s.PropertyItems
}

func (s *DescribePropertyPortItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyPortItemResponseBody) SetPageInfo(v *DescribePropertyPortItemResponseBodyPageInfo) *DescribePropertyPortItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyPortItemResponseBody) SetPropertyItems(v []*DescribePropertyPortItemResponseBodyPropertyItems) *DescribePropertyPortItemResponseBody {
	s.PropertyItems = v
	return s
}

func (s *DescribePropertyPortItemResponseBody) SetRequestId(v string) *DescribePropertyPortItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyPortItemResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyPortItemResponseBodyPageInfo struct {
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
	// 163
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyPortItemResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyPortItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertyPortItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyPortItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyPortItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyPortItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyPortItemResponseBodyPropertyItems struct {
	// The number of servers that use the port.
	//
	// example:
	//
	// 495
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The number of the listening port.
	//
	// example:
	//
	// 22
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The network protocol of the port.
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
}

func (s DescribePropertyPortItemResponseBodyPropertyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyPortItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortItemResponseBodyPropertyItems) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyPortItemResponseBodyPropertyItems) GetPort() *string {
	return s.Port
}

func (s *DescribePropertyPortItemResponseBodyPropertyItems) GetProto() *string {
	return s.Proto
}

func (s *DescribePropertyPortItemResponseBodyPropertyItems) SetCount(v int32) *DescribePropertyPortItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPropertyItems) SetPort(v string) *DescribePropertyPortItemResponseBodyPropertyItems {
	s.Port = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPropertyItems) SetProto(v string) *DescribePropertyPortItemResponseBodyPropertyItems {
	s.Proto = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPropertyItems) Validate() error {
	return dara.Validate(s)
}
