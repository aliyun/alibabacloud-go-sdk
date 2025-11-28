// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeBaaSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeBaaSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeBaaSResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeBaaSResponseBody) *AuthorizeBaaSResponse
	GetBody() *AuthorizeBaaSResponseBody
}

type AuthorizeBaaSResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeBaaSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeBaaSResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeBaaSResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeBaaSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeBaaSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeBaaSResponse) GetBody() *AuthorizeBaaSResponseBody {
	return s.Body
}

func (s *AuthorizeBaaSResponse) SetHeaders(v map[string]*string) *AuthorizeBaaSResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeBaaSResponse) SetStatusCode(v int32) *AuthorizeBaaSResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeBaaSResponse) SetBody(v *AuthorizeBaaSResponseBody) *AuthorizeBaaSResponse {
	s.Body = v
	return s
}

func (s *AuthorizeBaaSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
