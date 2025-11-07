// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeProductCodeResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeProductCodeResponseBodyItems) *DescribeProductCodeResponseBody
	GetItems() []*DescribeProductCodeResponseBodyItems
	SetPageSize(v int32) *DescribeProductCodeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeProductCodeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeProductCodeResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *DescribeProductCodeResponseBody
	GetTotalPage() *int32
}

type DescribeProductCodeResponseBody struct {
	// Current query page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of product code information.
	Items []*DescribeProductCodeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Number of products per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ID of this request.
	//
	// example:
	//
	// 7FBBADA3-9A66-5759-8AF8-2F99F5BE13F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of returned results.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 3
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeProductCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductCodeResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeProductCodeResponseBody) GetItems() []*DescribeProductCodeResponseBodyItems {
	return s.Items
}

func (s *DescribeProductCodeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeProductCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProductCodeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeProductCodeResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeProductCodeResponseBody) SetCurrentPage(v int32) *DescribeProductCodeResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeProductCodeResponseBody) SetItems(v []*DescribeProductCodeResponseBodyItems) *DescribeProductCodeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeProductCodeResponseBody) SetPageSize(v int32) *DescribeProductCodeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeProductCodeResponseBody) SetRequestId(v string) *DescribeProductCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductCodeResponseBody) SetTotalCount(v int32) *DescribeProductCodeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeProductCodeResponseBody) SetTotalPage(v int32) *DescribeProductCodeResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeProductCodeResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProductCodeResponseBodyItems struct {
	// Product code.
	//
	// example:
	//
	// ID_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Name corresponding to the product code.
	//
	// example:
	//
	// APP认证
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s DescribeProductCodeResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductCodeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeProductCodeResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeProductCodeResponseBodyItems) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeProductCodeResponseBodyItems) SetProductCode(v string) *DescribeProductCodeResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeProductCodeResponseBodyItems) SetProductName(v string) *DescribeProductCodeResponseBodyItems {
	s.ProductName = &v
	return s
}

func (s *DescribeProductCodeResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
