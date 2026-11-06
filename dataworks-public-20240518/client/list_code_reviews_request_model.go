// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeReviewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *ListCodeReviewsRequest
	GetCreator() *string
	SetPageNumber(v int32) *ListCodeReviewsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCodeReviewsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListCodeReviewsRequest
	GetProjectId() *int64
	SetReviewer(v string) *ListCodeReviewsRequest
	GetReviewer() *string
	SetStatus(v string) *ListCodeReviewsRequest
	GetStatus() *string
}

type ListCodeReviewsRequest struct {
	// The account ID of the submitter.
	//
	// example:
	//
	// 209000000000000002
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23620
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The account ID of the specified reviewer. Exact match is used.
	//
	// example:
	//
	// 209000000000000001
	Reviewer *string `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	// The review status.
	//
	// example:
	//
	// Reviewing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCodeReviewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCodeReviewsRequest) GoString() string {
	return s.String()
}

func (s *ListCodeReviewsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListCodeReviewsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCodeReviewsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCodeReviewsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCodeReviewsRequest) GetReviewer() *string {
	return s.Reviewer
}

func (s *ListCodeReviewsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCodeReviewsRequest) SetCreator(v string) *ListCodeReviewsRequest {
	s.Creator = &v
	return s
}

func (s *ListCodeReviewsRequest) SetPageNumber(v int32) *ListCodeReviewsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCodeReviewsRequest) SetPageSize(v int32) *ListCodeReviewsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCodeReviewsRequest) SetProjectId(v int64) *ListCodeReviewsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCodeReviewsRequest) SetReviewer(v string) *ListCodeReviewsRequest {
	s.Reviewer = &v
	return s
}

func (s *ListCodeReviewsRequest) SetStatus(v string) *ListCodeReviewsRequest {
	s.Status = &v
	return s
}

func (s *ListCodeReviewsRequest) Validate() error {
	return dara.Validate(s)
}
