// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWHClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWHClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWHClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *CreateWHClientCertificateResponseBody) *CreateWHClientCertificateResponse
	GetBody() *CreateWHClientCertificateResponseBody
}

type CreateWHClientCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWHClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWHClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWHClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateWHClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWHClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWHClientCertificateResponse) GetBody() *CreateWHClientCertificateResponseBody {
	return s.Body
}

func (s *CreateWHClientCertificateResponse) SetHeaders(v map[string]*string) *CreateWHClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateWHClientCertificateResponse) SetStatusCode(v int32) *CreateWHClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWHClientCertificateResponse) SetBody(v *CreateWHClientCertificateResponseBody) *CreateWHClientCertificateResponse {
	s.Body = v
	return s
}

func (s *CreateWHClientCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
