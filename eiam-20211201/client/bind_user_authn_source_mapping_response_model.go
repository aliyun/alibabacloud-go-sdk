// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindUserAuthnSourceMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindUserAuthnSourceMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindUserAuthnSourceMappingResponse
	GetStatusCode() *int32
	SetBody(v *BindUserAuthnSourceMappingResponseBody) *BindUserAuthnSourceMappingResponse
	GetBody() *BindUserAuthnSourceMappingResponseBody
}

type BindUserAuthnSourceMappingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindUserAuthnSourceMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindUserAuthnSourceMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s BindUserAuthnSourceMappingResponse) GoString() string {
	return s.String()
}

func (s *BindUserAuthnSourceMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindUserAuthnSourceMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindUserAuthnSourceMappingResponse) GetBody() *BindUserAuthnSourceMappingResponseBody {
	return s.Body
}

func (s *BindUserAuthnSourceMappingResponse) SetHeaders(v map[string]*string) *BindUserAuthnSourceMappingResponse {
	s.Headers = v
	return s
}

func (s *BindUserAuthnSourceMappingResponse) SetStatusCode(v int32) *BindUserAuthnSourceMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *BindUserAuthnSourceMappingResponse) SetBody(v *BindUserAuthnSourceMappingResponseBody) *BindUserAuthnSourceMappingResponse {
	s.Body = v
	return s
}

func (s *BindUserAuthnSourceMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
