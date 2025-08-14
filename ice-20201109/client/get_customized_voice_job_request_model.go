// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomizedVoiceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetCustomizedVoiceJobRequest
	GetJobId() *string
}

type GetCustomizedVoiceJobRequest struct {
	// The ID of the human voice cloning job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetCustomizedVoiceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizedVoiceJobRequest) GoString() string {
	return s.String()
}

func (s *GetCustomizedVoiceJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetCustomizedVoiceJobRequest) SetJobId(v string) *GetCustomizedVoiceJobRequest {
	s.JobId = &v
	return s
}

func (s *GetCustomizedVoiceJobRequest) Validate() error {
	return dara.Validate(s)
}
