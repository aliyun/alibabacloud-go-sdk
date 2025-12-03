// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCACertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadCACertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadCACertificateResponse
	GetStatusCode() *int32
	SetBody(v *UploadCACertificateResponseBody) *UploadCACertificateResponse
	GetBody() *UploadCACertificateResponseBody
}

type UploadCACertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadCACertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadCACertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadCACertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadCACertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadCACertificateResponse) GetBody() *UploadCACertificateResponseBody {
	return s.Body
}

func (s *UploadCACertificateResponse) SetHeaders(v map[string]*string) *UploadCACertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadCACertificateResponse) SetStatusCode(v int32) *UploadCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadCACertificateResponse) SetBody(v *UploadCACertificateResponseBody) *UploadCACertificateResponse {
	s.Body = v
	return s
}

func (s *UploadCACertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
