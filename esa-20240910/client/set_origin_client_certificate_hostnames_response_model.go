// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetOriginClientCertificateHostnamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetOriginClientCertificateHostnamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetOriginClientCertificateHostnamesResponse
	GetStatusCode() *int32
	SetBody(v *SetOriginClientCertificateHostnamesResponseBody) *SetOriginClientCertificateHostnamesResponse
	GetBody() *SetOriginClientCertificateHostnamesResponseBody
}

type SetOriginClientCertificateHostnamesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetOriginClientCertificateHostnamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetOriginClientCertificateHostnamesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetOriginClientCertificateHostnamesResponse) GoString() string {
	return s.String()
}

func (s *SetOriginClientCertificateHostnamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetOriginClientCertificateHostnamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetOriginClientCertificateHostnamesResponse) GetBody() *SetOriginClientCertificateHostnamesResponseBody {
	return s.Body
}

func (s *SetOriginClientCertificateHostnamesResponse) SetHeaders(v map[string]*string) *SetOriginClientCertificateHostnamesResponse {
	s.Headers = v
	return s
}

func (s *SetOriginClientCertificateHostnamesResponse) SetStatusCode(v int32) *SetOriginClientCertificateHostnamesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetOriginClientCertificateHostnamesResponse) SetBody(v *SetOriginClientCertificateHostnamesResponseBody) *SetOriginClientCertificateHostnamesResponse {
	s.Body = v
	return s
}

func (s *SetOriginClientCertificateHostnamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
