// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskTimeVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskTimeVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskTimeVariablesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskTimeVariablesResponseBody) *UpdateTaskTimeVariablesResponse
	GetBody() *UpdateTaskTimeVariablesResponseBody
}

type UpdateTaskTimeVariablesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskTimeVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskTimeVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskTimeVariablesResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskTimeVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskTimeVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskTimeVariablesResponse) GetBody() *UpdateTaskTimeVariablesResponseBody {
	return s.Body
}

func (s *UpdateTaskTimeVariablesResponse) SetHeaders(v map[string]*string) *UpdateTaskTimeVariablesResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskTimeVariablesResponse) SetStatusCode(v int32) *UpdateTaskTimeVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskTimeVariablesResponse) SetBody(v *UpdateTaskTimeVariablesResponseBody) *UpdateTaskTimeVariablesResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskTimeVariablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
