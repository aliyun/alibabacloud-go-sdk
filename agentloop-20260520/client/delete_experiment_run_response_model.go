// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExperimentRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExperimentRunResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExperimentRunResponseBody) *DeleteExperimentRunResponse
	GetBody() *DeleteExperimentRunResponseBody
}

type DeleteExperimentRunResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentRunResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentRunResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExperimentRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExperimentRunResponse) GetBody() *DeleteExperimentRunResponseBody {
	return s.Body
}

func (s *DeleteExperimentRunResponse) SetHeaders(v map[string]*string) *DeleteExperimentRunResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentRunResponse) SetStatusCode(v int32) *DeleteExperimentRunResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentRunResponse) SetBody(v *DeleteExperimentRunResponseBody) *DeleteExperimentRunResponse {
	s.Body = v
	return s
}

func (s *DeleteExperimentRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
