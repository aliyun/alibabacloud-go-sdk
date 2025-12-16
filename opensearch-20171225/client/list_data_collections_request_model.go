// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCollectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListDataCollectionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataCollectionsRequest
	GetPageSize() *int32
}

type ListDataCollectionsRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDataCollectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCollectionsRequest) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataCollectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCollectionsRequest) SetPageNumber(v int32) *ListDataCollectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataCollectionsRequest) SetPageSize(v int32) *ListDataCollectionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCollectionsRequest) Validate() error {
	return dara.Validate(s)
}
