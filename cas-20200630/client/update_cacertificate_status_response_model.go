// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCACertificateStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCACertificateStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCACertificateStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCACertificateStatusResponseBody) *UpdateCACertificateStatusResponse
	GetBody() *UpdateCACertificateStatusResponseBody
}

type UpdateCACertificateStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCACertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCACertificateStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCACertificateStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCACertificateStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCACertificateStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCACertificateStatusResponse) GetBody() *UpdateCACertificateStatusResponseBody {
	return s.Body
}

func (s *UpdateCACertificateStatusResponse) SetHeaders(v map[string]*string) *UpdateCACertificateStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCACertificateStatusResponse) SetStatusCode(v int32) *UpdateCACertificateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCACertificateStatusResponse) SetBody(v *UpdateCACertificateStatusResponseBody) *UpdateCACertificateStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateCACertificateStatusResponse) Validate() error {
	return dara.Validate(s)
}
