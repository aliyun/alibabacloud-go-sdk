// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPagesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPagesResponseBody
	GetPageSize() *int32
	SetPages(v []*ListPagesResponseBodyPages) *ListPagesResponseBody
	GetPages() []*ListPagesResponseBodyPages
	SetRequestId(v string) *ListPagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPagesResponseBody
	GetTotalCount() *int32
	SetUsage(v int64) *ListPagesResponseBody
	GetUsage() *int64
}

type ListPagesResponseBody struct {
	// The page number returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The custom error pages. Each element in the array contains error page-specific information.
	Pages []*ListPagesResponseBodyPages `json:"Pages,omitempty" xml:"Pages,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of custom error pages after filtering.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of custom error pages that you created.
	//
	// example:
	//
	// 10
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListPagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPagesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPagesResponseBody) GetPages() []*ListPagesResponseBodyPages {
	return s.Pages
}

func (s *ListPagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPagesResponseBody) GetUsage() *int64 {
	return s.Usage
}

func (s *ListPagesResponseBody) SetPageNumber(v int32) *ListPagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPagesResponseBody) SetPageSize(v int32) *ListPagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPagesResponseBody) SetPages(v []*ListPagesResponseBodyPages) *ListPagesResponseBody {
	s.Pages = v
	return s
}

func (s *ListPagesResponseBody) SetRequestId(v string) *ListPagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPagesResponseBody) SetTotalCount(v int32) *ListPagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPagesResponseBody) SetUsage(v int64) *ListPagesResponseBody {
	s.Usage = &v
	return s
}

func (s *ListPagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPagesResponseBodyPages struct {
	// The Base64-encoded content of the error page. The content type is specified by the Content-Type field.
	//
	// This parameter is required.
	//
	// example:
	//
	// PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The Content-Type field in the HTTP header.
	//
	// This parameter is required.
	//
	// example:
	//
	// text/html
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The description of the custom error page.
	//
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom error page.[](~~2850223~~)
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the custom error page.
	//
	// example:
	//
	// custom
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The name of the custom error page.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the custom error page was last modified.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPagesResponseBodyPages) String() string {
	return dara.Prettify(s)
}

func (s ListPagesResponseBodyPages) GoString() string {
	return s.String()
}

func (s *ListPagesResponseBodyPages) GetContent() *string {
	return s.Content
}

func (s *ListPagesResponseBodyPages) GetContentType() *string {
	return s.ContentType
}

func (s *ListPagesResponseBodyPages) GetDescription() *string {
	return s.Description
}

func (s *ListPagesResponseBodyPages) GetId() *int64 {
	return s.Id
}

func (s *ListPagesResponseBodyPages) GetKind() *string {
	return s.Kind
}

func (s *ListPagesResponseBodyPages) GetName() *string {
	return s.Name
}

func (s *ListPagesResponseBodyPages) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPagesResponseBodyPages) SetContent(v string) *ListPagesResponseBodyPages {
	s.Content = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetContentType(v string) *ListPagesResponseBodyPages {
	s.ContentType = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetDescription(v string) *ListPagesResponseBodyPages {
	s.Description = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetId(v int64) *ListPagesResponseBodyPages {
	s.Id = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetKind(v string) *ListPagesResponseBodyPages {
	s.Kind = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetName(v string) *ListPagesResponseBodyPages {
	s.Name = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetUpdateTime(v string) *ListPagesResponseBodyPages {
	s.UpdateTime = &v
	return s
}

func (s *ListPagesResponseBodyPages) Validate() error {
	return dara.Validate(s)
}
