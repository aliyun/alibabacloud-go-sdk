// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExperimentRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExperimentRunResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExperimentRunResponseBody) *UpdateExperimentRunResponse
	GetBody() *UpdateExperimentRunResponseBody
}

type UpdateExperimentRunResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentRunResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentRunResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExperimentRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExperimentRunResponse) GetBody() *UpdateExperimentRunResponseBody {
	return s.Body
}

func (s *UpdateExperimentRunResponse) SetHeaders(v map[string]*string) *UpdateExperimentRunResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentRunResponse) SetStatusCode(v int32) *UpdateExperimentRunResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentRunResponse) SetBody(v *UpdateExperimentRunResponseBody) *UpdateExperimentRunResponse {
	s.Body = v
	return s
}

func (s *UpdateExperimentRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
