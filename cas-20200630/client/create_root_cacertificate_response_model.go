// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRootCACertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRootCACertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRootCACertificateResponse
	GetStatusCode() *int32
	SetBody(v *CreateRootCACertificateResponseBody) *CreateRootCACertificateResponse
	GetBody() *CreateRootCACertificateResponseBody
}

type CreateRootCACertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRootCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRootCACertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRootCACertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateRootCACertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRootCACertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRootCACertificateResponse) GetBody() *CreateRootCACertificateResponseBody {
	return s.Body
}

func (s *CreateRootCACertificateResponse) SetHeaders(v map[string]*string) *CreateRootCACertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateRootCACertificateResponse) SetStatusCode(v int32) *CreateRootCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRootCACertificateResponse) SetBody(v *CreateRootCACertificateResponseBody) *CreateRootCACertificateResponse {
	s.Body = v
	return s
}

func (s *CreateRootCACertificateResponse) Validate() error {
	return dara.Validate(s)
}
