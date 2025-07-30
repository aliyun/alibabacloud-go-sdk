// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRevokeCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRevokeCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRevokeCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ListRevokeCertificateResponseBody) *ListRevokeCertificateResponse
	GetBody() *ListRevokeCertificateResponseBody
}

type ListRevokeCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRevokeCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRevokeCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRevokeCertificateResponse) GoString() string {
	return s.String()
}

func (s *ListRevokeCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRevokeCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRevokeCertificateResponse) GetBody() *ListRevokeCertificateResponseBody {
	return s.Body
}

func (s *ListRevokeCertificateResponse) SetHeaders(v map[string]*string) *ListRevokeCertificateResponse {
	s.Headers = v
	return s
}

func (s *ListRevokeCertificateResponse) SetStatusCode(v int32) *ListRevokeCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRevokeCertificateResponse) SetBody(v *ListRevokeCertificateResponseBody) *ListRevokeCertificateResponse {
	s.Body = v
	return s
}

func (s *ListRevokeCertificateResponse) Validate() error {
	return dara.Validate(s)
}
