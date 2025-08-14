// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvatarTrainingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetAvatarTrainingJobRequest
	GetJobId() *string
}

type GetAvatarTrainingJobRequest struct {
	// The ID of the digital human training job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAvatarTrainingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAvatarTrainingJobRequest) GoString() string {
	return s.String()
}

func (s *GetAvatarTrainingJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetAvatarTrainingJobRequest) SetJobId(v string) *GetAvatarTrainingJobRequest {
	s.JobId = &v
	return s
}

func (s *GetAvatarTrainingJobRequest) Validate() error {
	return dara.Validate(s)
}
