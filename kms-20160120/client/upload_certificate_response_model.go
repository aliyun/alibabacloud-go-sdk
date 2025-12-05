// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadCertificateResponse
	GetStatusCode() *int32
	SetBody(v *UploadCertificateResponseBody) *UploadCertificateResponse
	GetBody() *UploadCertificateResponseBody
}

type UploadCertificateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadCertificateResponse) GetBody() *UploadCertificateResponseBody {
	return s.Body
}

func (s *UploadCertificateResponse) SetHeaders(v map[string]*string) *UploadCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadCertificateResponse) SetStatusCode(v int32) *UploadCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadCertificateResponse) SetBody(v *UploadCertificateResponseBody) *UploadCertificateResponse {
	s.Body = v
	return s
}

func (s *UploadCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
