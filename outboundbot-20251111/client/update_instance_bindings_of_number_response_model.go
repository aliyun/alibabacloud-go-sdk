// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceBindingsOfNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceBindingsOfNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceBindingsOfNumberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceBindingsOfNumberResponseBody) *UpdateInstanceBindingsOfNumberResponse
	GetBody() *UpdateInstanceBindingsOfNumberResponseBody
}

type UpdateInstanceBindingsOfNumberResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceBindingsOfNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceBindingsOfNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceBindingsOfNumberResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceBindingsOfNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceBindingsOfNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceBindingsOfNumberResponse) GetBody() *UpdateInstanceBindingsOfNumberResponseBody {
	return s.Body
}

func (s *UpdateInstanceBindingsOfNumberResponse) SetHeaders(v map[string]*string) *UpdateInstanceBindingsOfNumberResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceBindingsOfNumberResponse) SetStatusCode(v int32) *UpdateInstanceBindingsOfNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberResponse) SetBody(v *UpdateInstanceBindingsOfNumberResponseBody) *UpdateInstanceBindingsOfNumberResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceBindingsOfNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
