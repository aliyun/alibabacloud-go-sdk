// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowsersInput interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserName(v string) *ListBrowsersInput
	GetBrowserName() *string
	SetPageNumber(v int) *ListBrowsersInput
	GetPageNumber() *int
	SetPageSize(v int) *ListBrowsersInput
	GetPageSize() *int
}

type ListBrowsersInput struct {
	// 按浏览器名称过滤
	BrowserName *string `json:"browserName,omitempty" xml:"browserName,omitempty"`
	PageNumber  *int    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListBrowsersInput) String() string {
	return dara.Prettify(s)
}

func (s ListBrowsersInput) GoString() string {
	return s.String()
}

func (s *ListBrowsersInput) GetBrowserName() *string {
	return s.BrowserName
}

func (s *ListBrowsersInput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListBrowsersInput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListBrowsersInput) SetBrowserName(v string) *ListBrowsersInput {
	s.BrowserName = &v
	return s
}

func (s *ListBrowsersInput) SetPageNumber(v int) *ListBrowsersInput {
	s.PageNumber = &v
	return s
}

func (s *ListBrowsersInput) SetPageSize(v int) *ListBrowsersInput {
	s.PageSize = &v
	return s
}

func (s *ListBrowsersInput) Validate() error {
	return dara.Validate(s)
}
