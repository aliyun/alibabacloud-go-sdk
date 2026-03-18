// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestLogListQry interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *RequestLogListQry
	GetKeyword() *string
	SetPage(v int32) *RequestLogListQry
	GetPage() *int32
	SetPageSize(v int32) *RequestLogListQry
	GetPageSize() *int32
	SetStatus(v int32) *RequestLogListQry
	GetStatus() *int32
}

type RequestLogListQry struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s RequestLogListQry) String() string {
	return dara.Prettify(s)
}

func (s RequestLogListQry) GoString() string {
	return s.String()
}

func (s *RequestLogListQry) GetKeyword() *string {
	return s.Keyword
}

func (s *RequestLogListQry) GetPage() *int32 {
	return s.Page
}

func (s *RequestLogListQry) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RequestLogListQry) GetStatus() *int32 {
	return s.Status
}

func (s *RequestLogListQry) SetKeyword(v string) *RequestLogListQry {
	s.Keyword = &v
	return s
}

func (s *RequestLogListQry) SetPage(v int32) *RequestLogListQry {
	s.Page = &v
	return s
}

func (s *RequestLogListQry) SetPageSize(v int32) *RequestLogListQry {
	s.PageSize = &v
	return s
}

func (s *RequestLogListQry) SetStatus(v int32) *RequestLogListQry {
	s.Status = &v
	return s
}

func (s *RequestLogListQry) Validate() error {
	return dara.Validate(s)
}
