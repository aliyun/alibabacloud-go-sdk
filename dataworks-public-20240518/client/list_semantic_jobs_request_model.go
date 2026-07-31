// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSemanticJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSemanticJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSemanticJobsRequest
	GetPageSize() *int32
}

type ListSemanticJobsRequest struct {
	// The page number, starting from 1. If this parameter is not specified or set to a value less than or equal to 0, the first page is returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of task definitions per page. If this parameter is not specified or set to a value less than or equal to 0, the default value 50 is used. Maximum value: 200.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSemanticJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSemanticJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSemanticJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSemanticJobsRequest) SetPageNumber(v int32) *ListSemanticJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSemanticJobsRequest) SetPageSize(v int32) *ListSemanticJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSemanticJobsRequest) Validate() error {
	return dara.Validate(s)
}
