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
	// The search keyword. Fuzzy match on version names is supported.
	//
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results per page. Default value: 20. Minimum value: 1. Maximum value: 100.
	//
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
