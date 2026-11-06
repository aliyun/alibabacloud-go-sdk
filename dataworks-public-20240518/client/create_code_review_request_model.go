// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateCodeReviewRequest
	GetDescription() *string
	SetPipelineRunId(v string) *CreateCodeReviewRequest
	GetPipelineRunId() *string
	SetProjectId(v int64) *CreateCodeReviewRequest
	GetProjectId() *int64
	SetReviewers(v string) *CreateCodeReviewRequest
	GetReviewers() *string
}

type CreateCodeReviewRequest struct {
	// The description of the code review. The maximum length is 1,024 characters.
	//
	// example:
	//
	// Code review
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of a single release pipeline.
	//
	// This parameter is required.
	//
	// example:
	//
	// pipeline-id
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23620
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specifies the reviewers. Separate multiple accounts with commas (,). At least one reviewer is required, and the maximum length is 512 characters after deduplication.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123456,2345678901234567
	Reviewers *string `json:"Reviewers,omitempty" xml:"Reviewers,omitempty"`
}

func (s CreateCodeReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeReviewRequest) GoString() string {
	return s.String()
}

func (s *CreateCodeReviewRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCodeReviewRequest) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *CreateCodeReviewRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateCodeReviewRequest) GetReviewers() *string {
	return s.Reviewers
}

func (s *CreateCodeReviewRequest) SetDescription(v string) *CreateCodeReviewRequest {
	s.Description = &v
	return s
}

func (s *CreateCodeReviewRequest) SetPipelineRunId(v string) *CreateCodeReviewRequest {
	s.PipelineRunId = &v
	return s
}

func (s *CreateCodeReviewRequest) SetProjectId(v int64) *CreateCodeReviewRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCodeReviewRequest) SetReviewers(v string) *CreateCodeReviewRequest {
	s.Reviewers = &v
	return s
}

func (s *CreateCodeReviewRequest) Validate() error {
	return dara.Validate(s)
}
