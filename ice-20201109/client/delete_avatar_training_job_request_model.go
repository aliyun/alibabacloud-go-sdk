// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAvatarTrainingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DeleteAvatarTrainingJobRequest
	GetJobId() *string
}

type DeleteAvatarTrainingJobRequest struct {
	// The ID of the digital human training job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteAvatarTrainingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAvatarTrainingJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteAvatarTrainingJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteAvatarTrainingJobRequest) SetJobId(v string) *DeleteAvatarTrainingJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteAvatarTrainingJobRequest) Validate() error {
	return dara.Validate(s)
}
