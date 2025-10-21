// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *ListVariablesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListVariablesRequest
	GetPageSize() *int32
}

type ListVariablesRequest struct {
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVariablesRequest) GoString() string {
	return s.String()
}

func (s *ListVariablesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListVariablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVariablesRequest) SetPageIndex(v int32) *ListVariablesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListVariablesRequest) SetPageSize(v int32) *ListVariablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVariablesRequest) Validate() error {
	return dara.Validate(s)
}
