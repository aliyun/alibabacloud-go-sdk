// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskCheckItemResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageContentResource(v *DescribeRiskCheckItemResultResponseBodyPageContentResource) *DescribeRiskCheckItemResultResponseBody
	GetPageContentResource() *DescribeRiskCheckItemResultResponseBodyPageContentResource
	SetRequestId(v string) *DescribeRiskCheckItemResultResponseBody
	GetRequestId() *string
}

type DescribeRiskCheckItemResultResponseBody struct {
	// The pagination information.
	PageContentResource *DescribeRiskCheckItemResultResponseBodyPageContentResource `json:"PageContentResource,omitempty" xml:"PageContentResource,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 3BFB4989-A108-46A4-954E-FF7EF02D1078
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRiskCheckItemResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckItemResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckItemResultResponseBody) GetPageContentResource() *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	return s.PageContentResource
}

func (s *DescribeRiskCheckItemResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskCheckItemResultResponseBody) SetPageContentResource(v *DescribeRiskCheckItemResultResponseBodyPageContentResource) *DescribeRiskCheckItemResultResponseBody {
	s.PageContentResource = v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBody) SetRequestId(v string) *DescribeRiskCheckItemResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBody) Validate() error {
	if s.PageContentResource != nil {
		if err := s.PageContentResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRiskCheckItemResultResponseBodyPageContentResource struct {
	// The data of the affected assets on each page in a dynamic table.
	ContentResource map[string]interface{} `json:"ContentResource,omitempty" xml:"ContentResource,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 3
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRiskCheckItemResultResponseBodyPageContentResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckItemResultResponseBodyPageContentResource) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) GetContentResource() map[string]interface{} {
	return s.ContentResource
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) GetCount() *int32 {
	return s.Count
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) GetPageCount() *int32 {
	return s.PageCount
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetContentResource(v map[string]interface{}) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.ContentResource = v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetCount(v int32) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.Count = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetCurrentPage(v int32) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetPageCount(v int32) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.PageCount = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetPageSize(v int32) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetTotalCount(v int32) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.TotalCount = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) Validate() error {
	return dara.Validate(s)
}
