// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCertificateRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCertificateRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCertificateRequestResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCertificateRequestResponseBody) *DeleteCertificateRequestResponse
	GetBody() *DeleteCertificateRequestResponseBody
}

type DeleteCertificateRequestResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCertificateRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCertificateRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCertificateRequestResponse) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCertificateRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCertificateRequestResponse) GetBody() *DeleteCertificateRequestResponseBody {
	return s.Body
}

func (s *DeleteCertificateRequestResponse) SetHeaders(v map[string]*string) *DeleteCertificateRequestResponse {
	s.Headers = v
	return s
}

func (s *DeleteCertificateRequestResponse) SetStatusCode(v int32) *DeleteCertificateRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCertificateRequestResponse) SetBody(v *DeleteCertificateRequestResponseBody) *DeleteCertificateRequestResponse {
	s.Body = v
	return s
}

func (s *DeleteCertificateRequestResponse) Validate() error {
	return dara.Validate(s)
}
