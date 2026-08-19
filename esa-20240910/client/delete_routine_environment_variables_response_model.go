// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineEnvironmentVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRoutineEnvironmentVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRoutineEnvironmentVariablesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRoutineEnvironmentVariablesResponseBody) *DeleteRoutineEnvironmentVariablesResponse
	GetBody() *DeleteRoutineEnvironmentVariablesResponseBody
}

type DeleteRoutineEnvironmentVariablesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRoutineEnvironmentVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRoutineEnvironmentVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineEnvironmentVariablesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineEnvironmentVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRoutineEnvironmentVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRoutineEnvironmentVariablesResponse) GetBody() *DeleteRoutineEnvironmentVariablesResponseBody {
	return s.Body
}

func (s *DeleteRoutineEnvironmentVariablesResponse) SetHeaders(v map[string]*string) *DeleteRoutineEnvironmentVariablesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesResponse) SetStatusCode(v int32) *DeleteRoutineEnvironmentVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesResponse) SetBody(v *DeleteRoutineEnvironmentVariablesResponseBody) *DeleteRoutineEnvironmentVariablesResponse {
	s.Body = v
	return s
}

func (s *DeleteRoutineEnvironmentVariablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
