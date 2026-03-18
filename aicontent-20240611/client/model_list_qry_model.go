// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelListQry interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ModelListQry
	GetKeyword() *string
	SetPage(v int32) *ModelListQry
	GetPage() *int32
	SetPageSize(v int32) *ModelListQry
	GetPageSize() *int32
	SetStatus(v int32) *ModelListQry
	GetStatus() *int32
}

type ModelListQry struct {
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

func (s ModelListQry) String() string {
	return dara.Prettify(s)
}

func (s ModelListQry) GoString() string {
	return s.String()
}

func (s *ModelListQry) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelListQry) GetPage() *int32 {
	return s.Page
}

func (s *ModelListQry) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelListQry) GetStatus() *int32 {
	return s.Status
}

func (s *ModelListQry) SetKeyword(v string) *ModelListQry {
	s.Keyword = &v
	return s
}

func (s *ModelListQry) SetPage(v int32) *ModelListQry {
	s.Page = &v
	return s
}

func (s *ModelListQry) SetPageSize(v int32) *ModelListQry {
	s.PageSize = &v
	return s
}

func (s *ModelListQry) SetStatus(v int32) *ModelListQry {
	s.Status = &v
	return s
}

func (s *ModelListQry) Validate() error {
	return dara.Validate(s)
}
