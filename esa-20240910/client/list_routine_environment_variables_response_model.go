// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineEnvironmentVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoutineEnvironmentVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoutineEnvironmentVariablesResponse
	GetStatusCode() *int32
	SetBody(v *ListRoutineEnvironmentVariablesResponseBody) *ListRoutineEnvironmentVariablesResponse
	GetBody() *ListRoutineEnvironmentVariablesResponseBody
}

type ListRoutineEnvironmentVariablesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoutineEnvironmentVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoutineEnvironmentVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineEnvironmentVariablesResponse) GoString() string {
	return s.String()
}

func (s *ListRoutineEnvironmentVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoutineEnvironmentVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoutineEnvironmentVariablesResponse) GetBody() *ListRoutineEnvironmentVariablesResponseBody {
	return s.Body
}

func (s *ListRoutineEnvironmentVariablesResponse) SetHeaders(v map[string]*string) *ListRoutineEnvironmentVariablesResponse {
	s.Headers = v
	return s
}

func (s *ListRoutineEnvironmentVariablesResponse) SetStatusCode(v int32) *ListRoutineEnvironmentVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesResponse) SetBody(v *ListRoutineEnvironmentVariablesResponseBody) *ListRoutineEnvironmentVariablesResponse {
	s.Body = v
	return s
}

func (s *ListRoutineEnvironmentVariablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
