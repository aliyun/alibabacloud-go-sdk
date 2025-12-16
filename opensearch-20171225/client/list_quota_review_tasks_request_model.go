// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaReviewTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListQuotaReviewTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListQuotaReviewTasksRequest
	GetPageSize() *int32
}

type ListQuotaReviewTasksRequest struct {
	// The page number. Default value: 1.
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

func (s ListQuotaReviewTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaReviewTasksRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQuotaReviewTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQuotaReviewTasksRequest) SetPageNumber(v int32) *ListQuotaReviewTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQuotaReviewTasksRequest) SetPageSize(v int32) *ListQuotaReviewTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListQuotaReviewTasksRequest) Validate() error {
	return dara.Validate(s)
}
