// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationEnvironmentVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationEnvironmentVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationEnvironmentVariablesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationEnvironmentVariablesResponseBody) *UpdateApplicationEnvironmentVariablesResponse
	GetBody() *UpdateApplicationEnvironmentVariablesResponseBody
}

type UpdateApplicationEnvironmentVariablesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationEnvironmentVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationEnvironmentVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationEnvironmentVariablesResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationEnvironmentVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationEnvironmentVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationEnvironmentVariablesResponse) GetBody() *UpdateApplicationEnvironmentVariablesResponseBody {
	return s.Body
}

func (s *UpdateApplicationEnvironmentVariablesResponse) SetHeaders(v map[string]*string) *UpdateApplicationEnvironmentVariablesResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponse) SetStatusCode(v int32) *UpdateApplicationEnvironmentVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponse) SetBody(v *UpdateApplicationEnvironmentVariablesResponseBody) *UpdateApplicationEnvironmentVariablesResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
