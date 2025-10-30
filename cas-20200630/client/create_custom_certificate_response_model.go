// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomCertificateResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomCertificateResponseBody) *CreateCustomCertificateResponse
	GetBody() *CreateCustomCertificateResponseBody
}

type CreateCustomCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomCertificateResponse) GetBody() *CreateCustomCertificateResponseBody {
	return s.Body
}

func (s *CreateCustomCertificateResponse) SetHeaders(v map[string]*string) *CreateCustomCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomCertificateResponse) SetStatusCode(v int32) *CreateCustomCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomCertificateResponse) SetBody(v *CreateCustomCertificateResponseBody) *CreateCustomCertificateResponse {
	s.Body = v
	return s
}

func (s *CreateCustomCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
