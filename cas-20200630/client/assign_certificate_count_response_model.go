// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignCertificateCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignCertificateCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignCertificateCountResponse
	GetStatusCode() *int32
	SetBody(v *AssignCertificateCountResponseBody) *AssignCertificateCountResponse
	GetBody() *AssignCertificateCountResponseBody
}

type AssignCertificateCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignCertificateCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignCertificateCountResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignCertificateCountResponse) GoString() string {
	return s.String()
}

func (s *AssignCertificateCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignCertificateCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignCertificateCountResponse) GetBody() *AssignCertificateCountResponseBody {
	return s.Body
}

func (s *AssignCertificateCountResponse) SetHeaders(v map[string]*string) *AssignCertificateCountResponse {
	s.Headers = v
	return s
}

func (s *AssignCertificateCountResponse) SetStatusCode(v int32) *AssignCertificateCountResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignCertificateCountResponse) SetBody(v *AssignCertificateCountResponseBody) *AssignCertificateCountResponse {
	s.Body = v
	return s
}

func (s *AssignCertificateCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
