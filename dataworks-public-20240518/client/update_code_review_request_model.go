// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCodeReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateCodeReviewRequest
	GetId() *string
	SetProjectId(v int64) *UpdateCodeReviewRequest
	GetProjectId() *int64
	SetStatus(v string) *UpdateCodeReviewRequest
	GetStatus() *string
}

type UpdateCodeReviewRequest struct {
	// The unique identifier of the review.
	//
	// This parameter is required.
	//
	// example:
	//
	// review-id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23620
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The final decision. Valid values: Accepted, Rejected.
	//
	// This parameter is required.
	//
	// example:
	//
	// Accepted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateCodeReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCodeReviewRequest) GoString() string {
	return s.String()
}

func (s *UpdateCodeReviewRequest) GetId() *string {
	return s.Id
}

func (s *UpdateCodeReviewRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateCodeReviewRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateCodeReviewRequest) SetId(v string) *UpdateCodeReviewRequest {
	s.Id = &v
	return s
}

func (s *UpdateCodeReviewRequest) SetProjectId(v int64) *UpdateCodeReviewRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateCodeReviewRequest) SetStatus(v string) *UpdateCodeReviewRequest {
	s.Status = &v
	return s
}

func (s *UpdateCodeReviewRequest) Validate() error {
	return dara.Validate(s)
}
