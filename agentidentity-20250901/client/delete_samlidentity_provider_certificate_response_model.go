// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSAMLIdentityProviderCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSAMLIdentityProviderCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSAMLIdentityProviderCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSAMLIdentityProviderCertificateResponseBody) *DeleteSAMLIdentityProviderCertificateResponse
	GetBody() *DeleteSAMLIdentityProviderCertificateResponseBody
}

type DeleteSAMLIdentityProviderCertificateResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSAMLIdentityProviderCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSAMLIdentityProviderCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSAMLIdentityProviderCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteSAMLIdentityProviderCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSAMLIdentityProviderCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSAMLIdentityProviderCertificateResponse) GetBody() *DeleteSAMLIdentityProviderCertificateResponseBody {
	return s.Body
}

func (s *DeleteSAMLIdentityProviderCertificateResponse) SetHeaders(v map[string]*string) *DeleteSAMLIdentityProviderCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteSAMLIdentityProviderCertificateResponse) SetStatusCode(v int32) *DeleteSAMLIdentityProviderCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSAMLIdentityProviderCertificateResponse) SetBody(v *DeleteSAMLIdentityProviderCertificateResponseBody) *DeleteSAMLIdentityProviderCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteSAMLIdentityProviderCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
