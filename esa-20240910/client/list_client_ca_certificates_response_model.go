// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientCaCertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientCaCertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientCaCertificatesResponse
	GetStatusCode() *int32
	SetBody(v *ListClientCaCertificatesResponseBody) *ListClientCaCertificatesResponse
	GetBody() *ListClientCaCertificatesResponseBody
}

type ListClientCaCertificatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientCaCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientCaCertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientCaCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListClientCaCertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientCaCertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientCaCertificatesResponse) GetBody() *ListClientCaCertificatesResponseBody {
	return s.Body
}

func (s *ListClientCaCertificatesResponse) SetHeaders(v map[string]*string) *ListClientCaCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListClientCaCertificatesResponse) SetStatusCode(v int32) *ListClientCaCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientCaCertificatesResponse) SetBody(v *ListClientCaCertificatesResponseBody) *ListClientCaCertificatesResponse {
	s.Body = v
	return s
}

func (s *ListClientCaCertificatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
