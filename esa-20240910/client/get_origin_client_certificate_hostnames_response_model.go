// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginClientCertificateHostnamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOriginClientCertificateHostnamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOriginClientCertificateHostnamesResponse
	GetStatusCode() *int32
	SetBody(v *GetOriginClientCertificateHostnamesResponseBody) *GetOriginClientCertificateHostnamesResponse
	GetBody() *GetOriginClientCertificateHostnamesResponseBody
}

type GetOriginClientCertificateHostnamesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOriginClientCertificateHostnamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOriginClientCertificateHostnamesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOriginClientCertificateHostnamesResponse) GoString() string {
	return s.String()
}

func (s *GetOriginClientCertificateHostnamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOriginClientCertificateHostnamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOriginClientCertificateHostnamesResponse) GetBody() *GetOriginClientCertificateHostnamesResponseBody {
	return s.Body
}

func (s *GetOriginClientCertificateHostnamesResponse) SetHeaders(v map[string]*string) *GetOriginClientCertificateHostnamesResponse {
	s.Headers = v
	return s
}

func (s *GetOriginClientCertificateHostnamesResponse) SetStatusCode(v int32) *GetOriginClientCertificateHostnamesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOriginClientCertificateHostnamesResponse) SetBody(v *GetOriginClientCertificateHostnamesResponseBody) *GetOriginClientCertificateHostnamesResponse {
	s.Body = v
	return s
}

func (s *GetOriginClientCertificateHostnamesResponse) Validate() error {
	return dara.Validate(s)
}
