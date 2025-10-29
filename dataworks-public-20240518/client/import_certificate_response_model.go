// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ImportCertificateResponseBody) *ImportCertificateResponse
	GetBody() *ImportCertificateResponseBody
}

type ImportCertificateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportCertificateResponse) GoString() string {
	return s.String()
}

func (s *ImportCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportCertificateResponse) GetBody() *ImportCertificateResponseBody {
	return s.Body
}

func (s *ImportCertificateResponse) SetHeaders(v map[string]*string) *ImportCertificateResponse {
	s.Headers = v
	return s
}

func (s *ImportCertificateResponse) SetStatusCode(v int32) *ImportCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportCertificateResponse) SetBody(v *ImportCertificateResponseBody) *ImportCertificateResponse {
	s.Body = v
	return s
}

func (s *ImportCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
