// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExperimentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExperimentResponseBody) *DeleteExperimentResponse
	GetBody() *DeleteExperimentResponseBody
}

type DeleteExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExperimentResponse) GetBody() *DeleteExperimentResponseBody {
	return s.Body
}

func (s *DeleteExperimentResponse) SetHeaders(v map[string]*string) *DeleteExperimentResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentResponse) SetStatusCode(v int32) *DeleteExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentResponse) SetBody(v *DeleteExperimentResponseBody) *DeleteExperimentResponse {
	s.Body = v
	return s
}

func (s *DeleteExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
