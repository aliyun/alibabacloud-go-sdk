// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOriginClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOriginClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOriginClientCertificateResponseBody) *DeleteOriginClientCertificateResponse
	GetBody() *DeleteOriginClientCertificateResponseBody
}

type DeleteOriginClientCertificateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOriginClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOriginClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteOriginClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOriginClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOriginClientCertificateResponse) GetBody() *DeleteOriginClientCertificateResponseBody {
	return s.Body
}

func (s *DeleteOriginClientCertificateResponse) SetHeaders(v map[string]*string) *DeleteOriginClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteOriginClientCertificateResponse) SetStatusCode(v int32) *DeleteOriginClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOriginClientCertificateResponse) SetBody(v *DeleteOriginClientCertificateResponseBody) *DeleteOriginClientCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteOriginClientCertificateResponse) Validate() error {
	return dara.Validate(s)
}
