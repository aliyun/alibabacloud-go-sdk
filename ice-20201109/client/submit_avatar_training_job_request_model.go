// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAvatarTrainingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitAvatarTrainingJobRequest
	GetJobId() *string
}

type SubmitAvatarTrainingJobRequest struct {
	// The ID of the digital human training job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitAvatarTrainingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAvatarTrainingJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAvatarTrainingJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAvatarTrainingJobRequest) SetJobId(v string) *SubmitAvatarTrainingJobRequest {
	s.JobId = &v
	return s
}

func (s *SubmitAvatarTrainingJobRequest) Validate() error {
	return dara.Validate(s)
}
