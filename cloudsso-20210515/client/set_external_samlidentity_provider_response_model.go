// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetExternalSAMLIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetExternalSAMLIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetExternalSAMLIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *SetExternalSAMLIdentityProviderResponseBody) *SetExternalSAMLIdentityProviderResponse
	GetBody() *SetExternalSAMLIdentityProviderResponseBody
}

type SetExternalSAMLIdentityProviderResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetExternalSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetExternalSAMLIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s SetExternalSAMLIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *SetExternalSAMLIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetExternalSAMLIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetExternalSAMLIdentityProviderResponse) GetBody() *SetExternalSAMLIdentityProviderResponseBody {
	return s.Body
}

func (s *SetExternalSAMLIdentityProviderResponse) SetHeaders(v map[string]*string) *SetExternalSAMLIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponse) SetStatusCode(v int32) *SetExternalSAMLIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponse) SetBody(v *SetExternalSAMLIdentityProviderResponseBody) *SetExternalSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
