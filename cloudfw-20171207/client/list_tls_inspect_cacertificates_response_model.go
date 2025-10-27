// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTlsInspectCACertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTlsInspectCACertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTlsInspectCACertificatesResponse
	GetStatusCode() *int32
	SetBody(v *ListTlsInspectCACertificatesResponseBody) *ListTlsInspectCACertificatesResponse
	GetBody() *ListTlsInspectCACertificatesResponseBody
}

type ListTlsInspectCACertificatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTlsInspectCACertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTlsInspectCACertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTlsInspectCACertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListTlsInspectCACertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTlsInspectCACertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTlsInspectCACertificatesResponse) GetBody() *ListTlsInspectCACertificatesResponseBody {
	return s.Body
}

func (s *ListTlsInspectCACertificatesResponse) SetHeaders(v map[string]*string) *ListTlsInspectCACertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListTlsInspectCACertificatesResponse) SetStatusCode(v int32) *ListTlsInspectCACertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponse) SetBody(v *ListTlsInspectCACertificatesResponseBody) *ListTlsInspectCACertificatesResponse {
	s.Body = v
	return s
}

func (s *ListTlsInspectCACertificatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
