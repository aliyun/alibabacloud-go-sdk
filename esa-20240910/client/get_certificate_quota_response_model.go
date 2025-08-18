// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCertificateQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCertificateQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetCertificateQuotaResponseBody) *GetCertificateQuotaResponse
	GetBody() *GetCertificateQuotaResponseBody
}

type GetCertificateQuotaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCertificateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCertificateQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetCertificateQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCertificateQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCertificateQuotaResponse) GetBody() *GetCertificateQuotaResponseBody {
	return s.Body
}

func (s *GetCertificateQuotaResponse) SetHeaders(v map[string]*string) *GetCertificateQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetCertificateQuotaResponse) SetStatusCode(v int32) *GetCertificateQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCertificateQuotaResponse) SetBody(v *GetCertificateQuotaResponseBody) *GetCertificateQuotaResponse {
	s.Body = v
	return s
}

func (s *GetCertificateQuotaResponse) Validate() error {
	return dara.Validate(s)
}
