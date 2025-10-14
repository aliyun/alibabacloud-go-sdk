// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCertificateHostnamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientCertificateHostnamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientCertificateHostnamesResponse
	GetStatusCode() *int32
	SetBody(v *GetClientCertificateHostnamesResponseBody) *GetClientCertificateHostnamesResponse
	GetBody() *GetClientCertificateHostnamesResponseBody
}

type GetClientCertificateHostnamesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientCertificateHostnamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientCertificateHostnamesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientCertificateHostnamesResponse) GoString() string {
	return s.String()
}

func (s *GetClientCertificateHostnamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientCertificateHostnamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientCertificateHostnamesResponse) GetBody() *GetClientCertificateHostnamesResponseBody {
	return s.Body
}

func (s *GetClientCertificateHostnamesResponse) SetHeaders(v map[string]*string) *GetClientCertificateHostnamesResponse {
	s.Headers = v
	return s
}

func (s *GetClientCertificateHostnamesResponse) SetStatusCode(v int32) *GetClientCertificateHostnamesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientCertificateHostnamesResponse) SetBody(v *GetClientCertificateHostnamesResponseBody) *GetClientCertificateHostnamesResponse {
	s.Body = v
	return s
}

func (s *GetClientCertificateHostnamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
