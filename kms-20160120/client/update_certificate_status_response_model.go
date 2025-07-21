// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCertificateStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCertificateStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCertificateStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCertificateStatusResponseBody) *UpdateCertificateStatusResponse
	GetBody() *UpdateCertificateStatusResponseBody
}

type UpdateCertificateStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCertificateStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCertificateStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCertificateStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCertificateStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCertificateStatusResponse) GetBody() *UpdateCertificateStatusResponseBody {
	return s.Body
}

func (s *UpdateCertificateStatusResponse) SetHeaders(v map[string]*string) *UpdateCertificateStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCertificateStatusResponse) SetStatusCode(v int32) *UpdateCertificateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCertificateStatusResponse) SetBody(v *UpdateCertificateStatusResponseBody) *UpdateCertificateStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateCertificateStatusResponse) Validate() error {
	return dara.Validate(s)
}
