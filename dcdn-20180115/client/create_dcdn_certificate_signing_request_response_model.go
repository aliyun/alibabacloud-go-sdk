// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnCertificateSigningRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDcdnCertificateSigningRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDcdnCertificateSigningRequestResponse
	GetStatusCode() *int32
	SetBody(v *CreateDcdnCertificateSigningRequestResponseBody) *CreateDcdnCertificateSigningRequestResponse
	GetBody() *CreateDcdnCertificateSigningRequestResponseBody
}

type CreateDcdnCertificateSigningRequestResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDcdnCertificateSigningRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDcdnCertificateSigningRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnCertificateSigningRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateDcdnCertificateSigningRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDcdnCertificateSigningRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDcdnCertificateSigningRequestResponse) GetBody() *CreateDcdnCertificateSigningRequestResponseBody {
	return s.Body
}

func (s *CreateDcdnCertificateSigningRequestResponse) SetHeaders(v map[string]*string) *CreateDcdnCertificateSigningRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponse) SetStatusCode(v int32) *CreateDcdnCertificateSigningRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponse) SetBody(v *CreateDcdnCertificateSigningRequestResponseBody) *CreateDcdnCertificateSigningRequestResponse {
	s.Body = v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponse) Validate() error {
	return dara.Validate(s)
}
