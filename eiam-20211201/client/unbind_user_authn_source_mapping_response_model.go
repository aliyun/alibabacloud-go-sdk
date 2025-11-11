// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindUserAuthnSourceMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindUserAuthnSourceMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindUserAuthnSourceMappingResponse
	GetStatusCode() *int32
	SetBody(v *UnbindUserAuthnSourceMappingResponseBody) *UnbindUserAuthnSourceMappingResponse
	GetBody() *UnbindUserAuthnSourceMappingResponseBody
}

type UnbindUserAuthnSourceMappingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindUserAuthnSourceMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindUserAuthnSourceMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindUserAuthnSourceMappingResponse) GoString() string {
	return s.String()
}

func (s *UnbindUserAuthnSourceMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindUserAuthnSourceMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindUserAuthnSourceMappingResponse) GetBody() *UnbindUserAuthnSourceMappingResponseBody {
	return s.Body
}

func (s *UnbindUserAuthnSourceMappingResponse) SetHeaders(v map[string]*string) *UnbindUserAuthnSourceMappingResponse {
	s.Headers = v
	return s
}

func (s *UnbindUserAuthnSourceMappingResponse) SetStatusCode(v int32) *UnbindUserAuthnSourceMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindUserAuthnSourceMappingResponse) SetBody(v *UnbindUserAuthnSourceMappingResponseBody) *UnbindUserAuthnSourceMappingResponse {
	s.Body = v
	return s
}

func (s *UnbindUserAuthnSourceMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
