// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSpecificIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSpecificIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSpecificIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *SetSpecificIdentityProviderResponseBody) *SetSpecificIdentityProviderResponse
	GetBody() *SetSpecificIdentityProviderResponseBody
}

type SetSpecificIdentityProviderResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSpecificIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSpecificIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSpecificIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *SetSpecificIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSpecificIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSpecificIdentityProviderResponse) GetBody() *SetSpecificIdentityProviderResponseBody {
	return s.Body
}

func (s *SetSpecificIdentityProviderResponse) SetHeaders(v map[string]*string) *SetSpecificIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *SetSpecificIdentityProviderResponse) SetStatusCode(v int32) *SetSpecificIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSpecificIdentityProviderResponse) SetBody(v *SetSpecificIdentityProviderResponseBody) *SetSpecificIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *SetSpecificIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
