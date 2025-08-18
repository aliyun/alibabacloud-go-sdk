// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadClientCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadClientCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadClientCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *UploadClientCaCertificateResponseBody) *UploadClientCaCertificateResponse
	GetBody() *UploadClientCaCertificateResponseBody
}

type UploadClientCaCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadClientCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadClientCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadClientCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadClientCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadClientCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadClientCaCertificateResponse) GetBody() *UploadClientCaCertificateResponseBody {
	return s.Body
}

func (s *UploadClientCaCertificateResponse) SetHeaders(v map[string]*string) *UploadClientCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadClientCaCertificateResponse) SetStatusCode(v int32) *UploadClientCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadClientCaCertificateResponse) SetBody(v *UploadClientCaCertificateResponseBody) *UploadClientCaCertificateResponse {
	s.Body = v
	return s
}

func (s *UploadClientCaCertificateResponse) Validate() error {
	return dara.Validate(s)
}
