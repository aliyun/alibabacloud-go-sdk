// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListContactRequest
	GetCurrentPage() *int32
	SetKeyword(v string) *ListContactRequest
	GetKeyword() *string
	SetShowSize(v int32) *ListContactRequest
	GetShowSize() *int32
}

type ListContactRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword used in the query. For example, you can specify a keyword in names, email addresses, and mobile phone numbers.
	//
	// example:
	//
	// 186
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of contacts per page.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListContactRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContactRequest) GoString() string {
	return s.String()
}

func (s *ListContactRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListContactRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListContactRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListContactRequest) SetCurrentPage(v int32) *ListContactRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListContactRequest) SetKeyword(v string) *ListContactRequest {
	s.Keyword = &v
	return s
}

func (s *ListContactRequest) SetShowSize(v int32) *ListContactRequest {
	s.ShowSize = &v
	return s
}

func (s *ListContactRequest) Validate() error {
	return dara.Validate(s)
}
