// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertificateRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCertificateRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCertificateRequestResponse
	GetStatusCode() *int32
	SetBody(v *CreateCertificateRequestResponseBody) *CreateCertificateRequestResponse
	GetBody() *CreateCertificateRequestResponseBody
}

type CreateCertificateRequestResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertificateRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCertificateRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCertificateRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCertificateRequestResponse) GetBody() *CreateCertificateRequestResponseBody {
	return s.Body
}

func (s *CreateCertificateRequestResponse) SetHeaders(v map[string]*string) *CreateCertificateRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateRequestResponse) SetStatusCode(v int32) *CreateCertificateRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertificateRequestResponse) SetBody(v *CreateCertificateRequestResponseBody) *CreateCertificateRequestResponse {
	s.Body = v
	return s
}

func (s *CreateCertificateRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
