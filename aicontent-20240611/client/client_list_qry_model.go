// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientListQry interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ClientListQry
	GetKeyword() *string
	SetPage(v int32) *ClientListQry
	GetPage() *int32
	SetPageSize(v int32) *ClientListQry
	GetPageSize() *int32
	SetStatus(v int32) *ClientListQry
	GetStatus() *int32
}

type ClientListQry struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ClientListQry) String() string {
	return dara.Prettify(s)
}

func (s ClientListQry) GoString() string {
	return s.String()
}

func (s *ClientListQry) GetKeyword() *string {
	return s.Keyword
}

func (s *ClientListQry) GetPage() *int32 {
	return s.Page
}

func (s *ClientListQry) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ClientListQry) GetStatus() *int32 {
	return s.Status
}

func (s *ClientListQry) SetKeyword(v string) *ClientListQry {
	s.Keyword = &v
	return s
}

func (s *ClientListQry) SetPage(v int32) *ClientListQry {
	s.Page = &v
	return s
}

func (s *ClientListQry) SetPageSize(v int32) *ClientListQry {
	s.PageSize = &v
	return s
}

func (s *ClientListQry) SetStatus(v int32) *ClientListQry {
	s.Status = &v
	return s
}

func (s *ClientListQry) Validate() error {
	return dara.Validate(s)
}
