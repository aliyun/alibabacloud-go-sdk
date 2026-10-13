// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateOidcAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitiateOidcAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitiateOidcAuthResponse
	GetStatusCode() *int32
	SetBody(v *InitiateOidcAuthResponseBody) *InitiateOidcAuthResponse
	GetBody() *InitiateOidcAuthResponseBody
}

type InitiateOidcAuthResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitiateOidcAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitiateOidcAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s InitiateOidcAuthResponse) GoString() string {
	return s.String()
}

func (s *InitiateOidcAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitiateOidcAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitiateOidcAuthResponse) GetBody() *InitiateOidcAuthResponseBody {
	return s.Body
}

func (s *InitiateOidcAuthResponse) SetHeaders(v map[string]*string) *InitiateOidcAuthResponse {
	s.Headers = v
	return s
}

func (s *InitiateOidcAuthResponse) SetStatusCode(v int32) *InitiateOidcAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *InitiateOidcAuthResponse) SetBody(v *InitiateOidcAuthResponseBody) *InitiateOidcAuthResponse {
	s.Body = v
	return s
}

func (s *InitiateOidcAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
