// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthnSourceMappingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserAuthnSourceMappingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserAuthnSourceMappingsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserAuthnSourceMappingsResponseBody) *ListUserAuthnSourceMappingsResponse
	GetBody() *ListUserAuthnSourceMappingsResponseBody
}

type ListUserAuthnSourceMappingsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserAuthnSourceMappingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserAuthnSourceMappingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthnSourceMappingsResponse) GoString() string {
	return s.String()
}

func (s *ListUserAuthnSourceMappingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserAuthnSourceMappingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserAuthnSourceMappingsResponse) GetBody() *ListUserAuthnSourceMappingsResponseBody {
	return s.Body
}

func (s *ListUserAuthnSourceMappingsResponse) SetHeaders(v map[string]*string) *ListUserAuthnSourceMappingsResponse {
	s.Headers = v
	return s
}

func (s *ListUserAuthnSourceMappingsResponse) SetStatusCode(v int32) *ListUserAuthnSourceMappingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserAuthnSourceMappingsResponse) SetBody(v *ListUserAuthnSourceMappingsResponseBody) *ListUserAuthnSourceMappingsResponse {
	s.Body = v
	return s
}

func (s *ListUserAuthnSourceMappingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
