// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginClientCertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOriginClientCertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOriginClientCertificatesResponse
	GetStatusCode() *int32
	SetBody(v *ListOriginClientCertificatesResponseBody) *ListOriginClientCertificatesResponse
	GetBody() *ListOriginClientCertificatesResponseBody
}

type ListOriginClientCertificatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOriginClientCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOriginClientCertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOriginClientCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListOriginClientCertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOriginClientCertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOriginClientCertificatesResponse) GetBody() *ListOriginClientCertificatesResponseBody {
	return s.Body
}

func (s *ListOriginClientCertificatesResponse) SetHeaders(v map[string]*string) *ListOriginClientCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListOriginClientCertificatesResponse) SetStatusCode(v int32) *ListOriginClientCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOriginClientCertificatesResponse) SetBody(v *ListOriginClientCertificatesResponseBody) *ListOriginClientCertificatesResponse {
	s.Body = v
	return s
}

func (s *ListOriginClientCertificatesResponse) Validate() error {
	return dara.Validate(s)
}
