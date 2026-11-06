// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetCodeReviewRequest
	GetId() *string
	SetProjectId(v int64) *GetCodeReviewRequest
	GetProjectId() *int64
}

type GetCodeReviewRequest struct {
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
}

func (s GetCodeReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCodeReviewRequest) GoString() string {
	return s.String()
}

func (s *GetCodeReviewRequest) GetId() *string {
	return s.Id
}

func (s *GetCodeReviewRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetCodeReviewRequest) SetId(v string) *GetCodeReviewRequest {
	s.Id = &v
	return s
}

func (s *GetCodeReviewRequest) SetProjectId(v int64) *GetCodeReviewRequest {
	s.ProjectId = &v
	return s
}

func (s *GetCodeReviewRequest) Validate() error {
	return dara.Validate(s)
}
