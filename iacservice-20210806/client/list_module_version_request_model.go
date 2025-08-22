// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModuleVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListModuleVersionRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListModuleVersionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModuleVersionRequest
	GetPageSize() *int32
}

type ListModuleVersionRequest struct {
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListModuleVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModuleVersionRequest) GoString() string {
	return s.String()
}

func (s *ListModuleVersionRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListModuleVersionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModuleVersionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModuleVersionRequest) SetKeyword(v string) *ListModuleVersionRequest {
	s.Keyword = &v
	return s
}

func (s *ListModuleVersionRequest) SetPageNumber(v int32) *ListModuleVersionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModuleVersionRequest) SetPageSize(v int32) *ListModuleVersionRequest {
	s.PageSize = &v
	return s
}

func (s *ListModuleVersionRequest) Validate() error {
	return dara.Validate(s)
}
