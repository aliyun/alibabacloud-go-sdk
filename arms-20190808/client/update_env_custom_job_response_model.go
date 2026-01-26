// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvCustomJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnvCustomJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnvCustomJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEnvCustomJobResponseBody) *UpdateEnvCustomJobResponse
	GetBody() *UpdateEnvCustomJobResponseBody
}

type UpdateEnvCustomJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnvCustomJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnvCustomJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvCustomJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvCustomJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnvCustomJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnvCustomJobResponse) GetBody() *UpdateEnvCustomJobResponseBody {
	return s.Body
}

func (s *UpdateEnvCustomJobResponse) SetHeaders(v map[string]*string) *UpdateEnvCustomJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvCustomJobResponse) SetStatusCode(v int32) *UpdateEnvCustomJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnvCustomJobResponse) SetBody(v *UpdateEnvCustomJobResponseBody) *UpdateEnvCustomJobResponse {
	s.Body = v
	return s
}

func (s *UpdateEnvCustomJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
