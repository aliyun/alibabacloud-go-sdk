// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobWithParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *JobStartParameters) *StartJobWithParamsRequest
	GetBody() *JobStartParameters
}

type StartJobWithParamsRequest struct {
	// The parameter that is used to start the job.
	Body *JobStartParameters `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartJobWithParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s StartJobWithParamsRequest) GoString() string {
	return s.String()
}

func (s *StartJobWithParamsRequest) GetBody() *JobStartParameters {
	return s.Body
}

func (s *StartJobWithParamsRequest) SetBody(v *JobStartParameters) *StartJobWithParamsRequest {
	s.Body = v
	return s
}

func (s *StartJobWithParamsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
