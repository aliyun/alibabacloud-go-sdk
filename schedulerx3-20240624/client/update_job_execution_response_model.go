// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJobExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJobExecutionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJobExecutionResponseBody) *UpdateJobExecutionResponse
	GetBody() *UpdateJobExecutionResponseBody
}

type UpdateJobExecutionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJobExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJobExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJobExecutionResponse) GetBody() *UpdateJobExecutionResponseBody {
	return s.Body
}

func (s *UpdateJobExecutionResponse) SetHeaders(v map[string]*string) *UpdateJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobExecutionResponse) SetStatusCode(v int32) *UpdateJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobExecutionResponse) SetBody(v *UpdateJobExecutionResponseBody) *UpdateJobExecutionResponse {
	s.Body = v
	return s
}

func (s *UpdateJobExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
