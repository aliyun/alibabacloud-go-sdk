// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListServiceVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListServiceVersionsRequest
	GetPageSize() *int32
}

type ListServiceVersionsRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListServiceVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServiceVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServiceVersionsRequest) SetPageNumber(v int32) *ListServiceVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServiceVersionsRequest) SetPageSize(v int32) *ListServiceVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListServiceVersionsRequest) Validate() error {
	return dara.Validate(s)
}
