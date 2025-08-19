// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDictsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListDictsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDictsRequest
	GetPageSize() *int32
}

type ListDictsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDictsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDictsRequest) GoString() string {
	return s.String()
}

func (s *ListDictsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDictsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDictsRequest) SetPageNumber(v int32) *ListDictsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDictsRequest) SetPageSize(v int32) *ListDictsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDictsRequest) Validate() error {
	return dara.Validate(s)
}
