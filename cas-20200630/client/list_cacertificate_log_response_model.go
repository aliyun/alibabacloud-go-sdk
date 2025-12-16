// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCACertificateLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCACertificateLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCACertificateLogResponse
	GetStatusCode() *int32
	SetBody(v *ListCACertificateLogResponseBody) *ListCACertificateLogResponse
	GetBody() *ListCACertificateLogResponseBody
}

type ListCACertificateLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCACertificateLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCACertificateLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCACertificateLogResponse) GoString() string {
	return s.String()
}

func (s *ListCACertificateLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCACertificateLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCACertificateLogResponse) GetBody() *ListCACertificateLogResponseBody {
	return s.Body
}

func (s *ListCACertificateLogResponse) SetHeaders(v map[string]*string) *ListCACertificateLogResponse {
	s.Headers = v
	return s
}

func (s *ListCACertificateLogResponse) SetStatusCode(v int32) *ListCACertificateLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCACertificateLogResponse) SetBody(v *ListCACertificateLogResponseBody) *ListCACertificateLogResponse {
	s.Body = v
	return s
}

func (s *ListCACertificateLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
