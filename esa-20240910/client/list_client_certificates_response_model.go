// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientCertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientCertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientCertificatesResponse
	GetStatusCode() *int32
	SetBody(v *ListClientCertificatesResponseBody) *ListClientCertificatesResponse
	GetBody() *ListClientCertificatesResponseBody
}

type ListClientCertificatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientCertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListClientCertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientCertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientCertificatesResponse) GetBody() *ListClientCertificatesResponseBody {
	return s.Body
}

func (s *ListClientCertificatesResponse) SetHeaders(v map[string]*string) *ListClientCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListClientCertificatesResponse) SetStatusCode(v int32) *ListClientCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientCertificatesResponse) SetBody(v *ListClientCertificatesResponseBody) *ListClientCertificatesResponse {
	s.Body = v
	return s
}

func (s *ListClientCertificatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
