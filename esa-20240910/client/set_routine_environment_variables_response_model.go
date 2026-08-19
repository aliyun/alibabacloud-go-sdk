// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRoutineEnvironmentVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetRoutineEnvironmentVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetRoutineEnvironmentVariablesResponse
	GetStatusCode() *int32
	SetBody(v *SetRoutineEnvironmentVariablesResponseBody) *SetRoutineEnvironmentVariablesResponse
	GetBody() *SetRoutineEnvironmentVariablesResponseBody
}

type SetRoutineEnvironmentVariablesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRoutineEnvironmentVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRoutineEnvironmentVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetRoutineEnvironmentVariablesResponse) GoString() string {
	return s.String()
}

func (s *SetRoutineEnvironmentVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetRoutineEnvironmentVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetRoutineEnvironmentVariablesResponse) GetBody() *SetRoutineEnvironmentVariablesResponseBody {
	return s.Body
}

func (s *SetRoutineEnvironmentVariablesResponse) SetHeaders(v map[string]*string) *SetRoutineEnvironmentVariablesResponse {
	s.Headers = v
	return s
}

func (s *SetRoutineEnvironmentVariablesResponse) SetStatusCode(v int32) *SetRoutineEnvironmentVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRoutineEnvironmentVariablesResponse) SetBody(v *SetRoutineEnvironmentVariablesResponseBody) *SetRoutineEnvironmentVariablesResponse {
	s.Body = v
	return s
}

func (s *SetRoutineEnvironmentVariablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
