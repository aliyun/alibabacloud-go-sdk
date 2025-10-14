// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadOriginClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadOriginClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadOriginClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *UploadOriginClientCertificateResponseBody) *UploadOriginClientCertificateResponse
	GetBody() *UploadOriginClientCertificateResponseBody
}

type UploadOriginClientCertificateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadOriginClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadOriginClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadOriginClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadOriginClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadOriginClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadOriginClientCertificateResponse) GetBody() *UploadOriginClientCertificateResponseBody {
	return s.Body
}

func (s *UploadOriginClientCertificateResponse) SetHeaders(v map[string]*string) *UploadOriginClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadOriginClientCertificateResponse) SetStatusCode(v int32) *UploadOriginClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadOriginClientCertificateResponse) SetBody(v *UploadOriginClientCertificateResponseBody) *UploadOriginClientCertificateResponse {
	s.Body = v
	return s
}

func (s *UploadOriginClientCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
