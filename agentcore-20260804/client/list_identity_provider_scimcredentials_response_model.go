// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProviderSCIMCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIdentityProviderSCIMCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIdentityProviderSCIMCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *ListIdentityProviderSCIMCredentialsResponseBody) *ListIdentityProviderSCIMCredentialsResponse
	GetBody() *ListIdentityProviderSCIMCredentialsResponseBody
}

type ListIdentityProviderSCIMCredentialsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdentityProviderSCIMCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdentityProviderSCIMCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderSCIMCredentialsResponse) GoString() string {
	return s.String()
}

func (s *ListIdentityProviderSCIMCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIdentityProviderSCIMCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIdentityProviderSCIMCredentialsResponse) GetBody() *ListIdentityProviderSCIMCredentialsResponseBody {
	return s.Body
}

func (s *ListIdentityProviderSCIMCredentialsResponse) SetHeaders(v map[string]*string) *ListIdentityProviderSCIMCredentialsResponse {
	s.Headers = v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponse) SetStatusCode(v int32) *ListIdentityProviderSCIMCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponse) SetBody(v *ListIdentityProviderSCIMCredentialsResponseBody) *ListIdentityProviderSCIMCredentialsResponse {
	s.Body = v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
