// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadServerCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadServerCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadServerCertificateResponse
	GetStatusCode() *int32
	SetBody(v *UploadServerCertificateResponseBody) *UploadServerCertificateResponse
	GetBody() *UploadServerCertificateResponseBody
}

type UploadServerCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadServerCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadServerCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadServerCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadServerCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadServerCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadServerCertificateResponse) GetBody() *UploadServerCertificateResponseBody {
	return s.Body
}

func (s *UploadServerCertificateResponse) SetHeaders(v map[string]*string) *UploadServerCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadServerCertificateResponse) SetStatusCode(v int32) *UploadServerCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadServerCertificateResponse) SetBody(v *UploadServerCertificateResponseBody) *UploadServerCertificateResponse {
	s.Body = v
	return s
}

func (s *UploadServerCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
