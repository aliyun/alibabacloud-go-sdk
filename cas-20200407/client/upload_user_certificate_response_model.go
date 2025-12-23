// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadUserCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadUserCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadUserCertificateResponse
	GetStatusCode() *int32
	SetBody(v *UploadUserCertificateResponseBody) *UploadUserCertificateResponse
	GetBody() *UploadUserCertificateResponseBody
}

type UploadUserCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadUserCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadUserCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadUserCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadUserCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadUserCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadUserCertificateResponse) GetBody() *UploadUserCertificateResponseBody {
	return s.Body
}

func (s *UploadUserCertificateResponse) SetHeaders(v map[string]*string) *UploadUserCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadUserCertificateResponse) SetStatusCode(v int32) *UploadUserCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadUserCertificateResponse) SetBody(v *UploadUserCertificateResponseBody) *UploadUserCertificateResponse {
	s.Body = v
	return s
}

func (s *UploadUserCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
