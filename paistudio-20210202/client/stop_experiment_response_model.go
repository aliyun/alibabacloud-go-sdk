// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopExperimentResponse
	GetStatusCode() *int32
	SetBody(v *StopExperimentResponseBody) *StopExperimentResponse
	GetBody() *StopExperimentResponseBody
}

type StopExperimentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s StopExperimentResponse) GoString() string {
	return s.String()
}

func (s *StopExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopExperimentResponse) GetBody() *StopExperimentResponseBody {
	return s.Body
}

func (s *StopExperimentResponse) SetHeaders(v map[string]*string) *StopExperimentResponse {
	s.Headers = v
	return s
}

func (s *StopExperimentResponse) SetStatusCode(v int32) *StopExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *StopExperimentResponse) SetBody(v *StopExperimentResponseBody) *StopExperimentResponse {
	s.Body = v
	return s
}

func (s *StopExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
