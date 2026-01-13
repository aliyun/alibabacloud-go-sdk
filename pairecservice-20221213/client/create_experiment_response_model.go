// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExperimentResponse
	GetStatusCode() *int32
	SetBody(v *CreateExperimentResponseBody) *CreateExperimentResponse
	GetBody() *CreateExperimentResponseBody
}

type CreateExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExperimentResponse) GetBody() *CreateExperimentResponseBody {
	return s.Body
}

func (s *CreateExperimentResponse) SetHeaders(v map[string]*string) *CreateExperimentResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentResponse) SetStatusCode(v int32) *CreateExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentResponse) SetBody(v *CreateExperimentResponseBody) *CreateExperimentResponse {
	s.Body = v
	return s
}

func (s *CreateExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
