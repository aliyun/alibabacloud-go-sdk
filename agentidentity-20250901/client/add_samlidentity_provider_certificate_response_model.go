// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSAMLIdentityProviderCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSAMLIdentityProviderCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSAMLIdentityProviderCertificateResponse
	GetStatusCode() *int32
	SetBody(v *AddSAMLIdentityProviderCertificateResponseBody) *AddSAMLIdentityProviderCertificateResponse
	GetBody() *AddSAMLIdentityProviderCertificateResponseBody
}

type AddSAMLIdentityProviderCertificateResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSAMLIdentityProviderCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSAMLIdentityProviderCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSAMLIdentityProviderCertificateResponse) GoString() string {
	return s.String()
}

func (s *AddSAMLIdentityProviderCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSAMLIdentityProviderCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSAMLIdentityProviderCertificateResponse) GetBody() *AddSAMLIdentityProviderCertificateResponseBody {
	return s.Body
}

func (s *AddSAMLIdentityProviderCertificateResponse) SetHeaders(v map[string]*string) *AddSAMLIdentityProviderCertificateResponse {
	s.Headers = v
	return s
}

func (s *AddSAMLIdentityProviderCertificateResponse) SetStatusCode(v int32) *AddSAMLIdentityProviderCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSAMLIdentityProviderCertificateResponse) SetBody(v *AddSAMLIdentityProviderCertificateResponseBody) *AddSAMLIdentityProviderCertificateResponse {
	s.Body = v
	return s
}

func (s *AddSAMLIdentityProviderCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
