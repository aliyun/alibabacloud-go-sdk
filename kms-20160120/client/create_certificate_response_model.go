// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCertificateResponse
	GetStatusCode() *int32
	SetBody(v *CreateCertificateResponseBody) *CreateCertificateResponse
	GetBody() *CreateCertificateResponseBody
}

type CreateCertificateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCertificateResponse) GetBody() *CreateCertificateResponseBody {
	return s.Body
}

func (s *CreateCertificateResponse) SetHeaders(v map[string]*string) *CreateCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateResponse) SetStatusCode(v int32) *CreateCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertificateResponse) SetBody(v *CreateCertificateResponseBody) *CreateCertificateResponse {
	s.Body = v
	return s
}

func (s *CreateCertificateResponse) Validate() error {
	return dara.Validate(s)
}
