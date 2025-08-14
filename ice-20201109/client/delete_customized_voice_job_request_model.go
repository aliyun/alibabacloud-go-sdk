// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizedVoiceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DeleteCustomizedVoiceJobRequest
	GetJobId() *string
}

type DeleteCustomizedVoiceJobRequest struct {
	// The ID of the human voice cloning job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteCustomizedVoiceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizedVoiceJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedVoiceJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteCustomizedVoiceJobRequest) SetJobId(v string) *DeleteCustomizedVoiceJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteCustomizedVoiceJobRequest) Validate() error {
	return dara.Validate(s)
}
