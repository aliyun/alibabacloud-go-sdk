// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExperimentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExperimentResponseBody) *UpdateExperimentResponse
	GetBody() *UpdateExperimentResponseBody
}

type UpdateExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExperimentResponse) GetBody() *UpdateExperimentResponseBody {
	return s.Body
}

func (s *UpdateExperimentResponse) SetHeaders(v map[string]*string) *UpdateExperimentResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentResponse) SetStatusCode(v int32) *UpdateExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentResponse) SetBody(v *UpdateExperimentResponseBody) *UpdateExperimentResponse {
	s.Body = v
	return s
}

func (s *UpdateExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
