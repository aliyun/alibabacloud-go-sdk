// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlgorithmVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListAlgorithmVersionsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlgorithmVersionsRequest
	GetPageSize() *int64
}

type ListAlgorithmVersionsRequest struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAlgorithmVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlgorithmVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAlgorithmVersionsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlgorithmVersionsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlgorithmVersionsRequest) SetPageNumber(v int64) *ListAlgorithmVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlgorithmVersionsRequest) SetPageSize(v int64) *ListAlgorithmVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlgorithmVersionsRequest) Validate() error {
	return dara.Validate(s)
}
