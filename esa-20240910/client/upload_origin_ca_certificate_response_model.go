// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadOriginCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadOriginCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadOriginCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *UploadOriginCaCertificateResponseBody) *UploadOriginCaCertificateResponse
	GetBody() *UploadOriginCaCertificateResponseBody
}

type UploadOriginCaCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadOriginCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadOriginCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadOriginCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadOriginCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadOriginCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadOriginCaCertificateResponse) GetBody() *UploadOriginCaCertificateResponseBody {
	return s.Body
}

func (s *UploadOriginCaCertificateResponse) SetHeaders(v map[string]*string) *UploadOriginCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadOriginCaCertificateResponse) SetStatusCode(v int32) *UploadOriginCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadOriginCaCertificateResponse) SetBody(v *UploadOriginCaCertificateResponseBody) *UploadOriginCaCertificateResponse {
	s.Body = v
	return s
}

func (s *UploadOriginCaCertificateResponse) Validate() error {
	return dara.Validate(s)
}
