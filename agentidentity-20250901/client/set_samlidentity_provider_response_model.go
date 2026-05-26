// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSAMLIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSAMLIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSAMLIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *SetSAMLIdentityProviderResponseBody) *SetSAMLIdentityProviderResponse
	GetBody() *SetSAMLIdentityProviderResponseBody
}

type SetSAMLIdentityProviderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSAMLIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSAMLIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *SetSAMLIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSAMLIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSAMLIdentityProviderResponse) GetBody() *SetSAMLIdentityProviderResponseBody {
	return s.Body
}

func (s *SetSAMLIdentityProviderResponse) SetHeaders(v map[string]*string) *SetSAMLIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *SetSAMLIdentityProviderResponse) SetStatusCode(v int32) *SetSAMLIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSAMLIdentityProviderResponse) SetBody(v *SetSAMLIdentityProviderResponseBody) *SetSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *SetSAMLIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
