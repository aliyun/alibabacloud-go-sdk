// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *StartJobRequestBody) *StartJobRequest
	GetBody() *StartJobRequestBody
}

type StartJobRequest struct {
	// The parameter that is used to start the job.
	//
	// This parameter is required.
	Body *StartJobRequestBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartJobRequest) GoString() string {
	return s.String()
}

func (s *StartJobRequest) GetBody() *StartJobRequestBody {
	return s.Body
}

func (s *StartJobRequest) SetBody(v *StartJobRequestBody) *StartJobRequest {
	s.Body = v
	return s
}

func (s *StartJobRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
