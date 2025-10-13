// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRootCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRootCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRootCertificateResponse
	GetStatusCode() *int32
	SetBody(v *GetRootCertificateResponseBody) *GetRootCertificateResponse
	GetBody() *GetRootCertificateResponseBody
}

type GetRootCertificateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRootCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRootCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRootCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetRootCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRootCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRootCertificateResponse) GetBody() *GetRootCertificateResponseBody {
	return s.Body
}

func (s *GetRootCertificateResponse) SetHeaders(v map[string]*string) *GetRootCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetRootCertificateResponse) SetStatusCode(v int32) *GetRootCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRootCertificateResponse) SetBody(v *GetRootCertificateResponseBody) *GetRootCertificateResponse {
	s.Body = v
	return s
}

func (s *GetRootCertificateResponse) Validate() error {
	return dara.Validate(s)
}
