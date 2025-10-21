// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *StopJobRequestBody) *StopJobRequest
	GetBody() *StopJobRequestBody
}

type StopJobRequest struct {
	// The parameter that is used to stop the job.
	//
	// This parameter is required.
	Body *StopJobRequestBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopJobRequest) GoString() string {
	return s.String()
}

func (s *StopJobRequest) GetBody() *StopJobRequestBody {
	return s.Body
}

func (s *StopJobRequest) SetBody(v *StopJobRequestBody) *StopJobRequest {
	s.Body = v
	return s
}

func (s *StopJobRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
