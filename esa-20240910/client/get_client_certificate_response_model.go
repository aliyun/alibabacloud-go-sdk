// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *GetClientCertificateResponseBody) *GetClientCertificateResponse
	GetBody() *GetClientCertificateResponseBody
}

type GetClientCertificateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientCertificateResponse) GetBody() *GetClientCertificateResponseBody {
	return s.Body
}

func (s *GetClientCertificateResponse) SetHeaders(v map[string]*string) *GetClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetClientCertificateResponse) SetStatusCode(v int32) *GetClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientCertificateResponse) SetBody(v *GetClientCertificateResponseBody) *GetClientCertificateResponse {
	s.Body = v
	return s
}

func (s *GetClientCertificateResponse) Validate() error {
	return dara.Validate(s)
}
