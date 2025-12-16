// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAnalyzersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListUserAnalyzersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserAnalyzersRequest
	GetPageSize() *int32
}

type ListUserAnalyzersRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListUserAnalyzersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserAnalyzersRequest) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserAnalyzersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserAnalyzersRequest) SetPageNumber(v int32) *ListUserAnalyzersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserAnalyzersRequest) SetPageSize(v int32) *ListUserAnalyzersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserAnalyzersRequest) Validate() error {
	return dara.Validate(s)
}
