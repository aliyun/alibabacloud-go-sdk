// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSpecReviewTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *ListSpecReviewTasksRequest
	GetPage() *int32
	SetPageNumber(v int32) *ListSpecReviewTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSpecReviewTasksRequest
	GetPageSize() *int32
	SetSize(v int32) *ListSpecReviewTasksRequest
	GetSize() *int32
	SetStatus(v string) *ListSpecReviewTasksRequest
	GetStatus() *string
	SetType(v string) *ListSpecReviewTasksRequest
	GetType() *string
}

type ListSpecReviewTasksRequest struct {
	// example:
	//
	// 1
	Page       *int32 `json:"page,omitempty" xml:"page,omitempty"`
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 10
	Size   *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// QUOTA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSpecReviewTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSpecReviewTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSpecReviewTasksRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSpecReviewTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSpecReviewTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSpecReviewTasksRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListSpecReviewTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSpecReviewTasksRequest) GetType() *string {
	return s.Type
}

func (s *ListSpecReviewTasksRequest) SetPage(v int32) *ListSpecReviewTasksRequest {
	s.Page = &v
	return s
}

func (s *ListSpecReviewTasksRequest) SetPageNumber(v int32) *ListSpecReviewTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSpecReviewTasksRequest) SetPageSize(v int32) *ListSpecReviewTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSpecReviewTasksRequest) SetSize(v int32) *ListSpecReviewTasksRequest {
	s.Size = &v
	return s
}

func (s *ListSpecReviewTasksRequest) SetStatus(v string) *ListSpecReviewTasksRequest {
	s.Status = &v
	return s
}

func (s *ListSpecReviewTasksRequest) SetType(v string) *ListSpecReviewTasksRequest {
	s.Type = &v
	return s
}

func (s *ListSpecReviewTasksRequest) Validate() error {
	return dara.Validate(s)
}
