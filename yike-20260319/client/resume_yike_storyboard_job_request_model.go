// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeYikeStoryboardJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ResumeYikeStoryboardJobRequest
	GetJobId() *string
}

type ResumeYikeStoryboardJobRequest struct {
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s ResumeYikeStoryboardJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeYikeStoryboardJobRequest) GoString() string {
	return s.String()
}

func (s *ResumeYikeStoryboardJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *ResumeYikeStoryboardJobRequest) SetJobId(v string) *ResumeYikeStoryboardJobRequest {
	s.JobId = &v
	return s
}

func (s *ResumeYikeStoryboardJobRequest) Validate() error {
	return dara.Validate(s)
}
