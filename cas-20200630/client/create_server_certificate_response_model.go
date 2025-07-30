// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServerCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServerCertificateResponse
	GetStatusCode() *int32
	SetBody(v *CreateServerCertificateResponseBody) *CreateServerCertificateResponse
	GetBody() *CreateServerCertificateResponseBody
}

type CreateServerCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServerCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServerCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServerCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServerCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServerCertificateResponse) GetBody() *CreateServerCertificateResponseBody {
	return s.Body
}

func (s *CreateServerCertificateResponse) SetHeaders(v map[string]*string) *CreateServerCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateServerCertificateResponse) SetStatusCode(v int32) *CreateServerCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServerCertificateResponse) SetBody(v *CreateServerCertificateResponseBody) *CreateServerCertificateResponse {
	s.Body = v
	return s
}

func (s *CreateServerCertificateResponse) Validate() error {
	return dara.Validate(s)
}
