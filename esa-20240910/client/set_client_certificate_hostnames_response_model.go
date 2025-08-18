// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClientCertificateHostnamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetClientCertificateHostnamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetClientCertificateHostnamesResponse
	GetStatusCode() *int32
	SetBody(v *SetClientCertificateHostnamesResponseBody) *SetClientCertificateHostnamesResponse
	GetBody() *SetClientCertificateHostnamesResponseBody
}

type SetClientCertificateHostnamesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetClientCertificateHostnamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetClientCertificateHostnamesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetClientCertificateHostnamesResponse) GoString() string {
	return s.String()
}

func (s *SetClientCertificateHostnamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetClientCertificateHostnamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetClientCertificateHostnamesResponse) GetBody() *SetClientCertificateHostnamesResponseBody {
	return s.Body
}

func (s *SetClientCertificateHostnamesResponse) SetHeaders(v map[string]*string) *SetClientCertificateHostnamesResponse {
	s.Headers = v
	return s
}

func (s *SetClientCertificateHostnamesResponse) SetStatusCode(v int32) *SetClientCertificateHostnamesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetClientCertificateHostnamesResponse) SetBody(v *SetClientCertificateHostnamesResponseBody) *SetClientCertificateHostnamesResponse {
	s.Body = v
	return s
}

func (s *SetClientCertificateHostnamesResponse) Validate() error {
	return dara.Validate(s)
}
