// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExperimentRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExperimentRunResponse
	GetStatusCode() *int32
	SetBody(v *CreateExperimentRunResponseBody) *CreateExperimentRunResponse
	GetBody() *CreateExperimentRunResponseBody
}

type CreateExperimentRunResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentRunResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExperimentRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExperimentRunResponse) GetBody() *CreateExperimentRunResponseBody {
	return s.Body
}

func (s *CreateExperimentRunResponse) SetHeaders(v map[string]*string) *CreateExperimentRunResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentRunResponse) SetStatusCode(v int32) *CreateExperimentRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentRunResponse) SetBody(v *CreateExperimentRunResponseBody) *CreateExperimentRunResponse {
	s.Body = v
	return s
}

func (s *CreateExperimentRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
