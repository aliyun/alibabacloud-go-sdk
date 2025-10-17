// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadInstanceCACertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadInstanceCACertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadInstanceCACertificateResponse
	GetStatusCode() *int32
	SetBody(v *DownloadInstanceCACertificateResponseBody) *DownloadInstanceCACertificateResponse
	GetBody() *DownloadInstanceCACertificateResponseBody
}

type DownloadInstanceCACertificateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadInstanceCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadInstanceCACertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadInstanceCACertificateResponse) GoString() string {
	return s.String()
}

func (s *DownloadInstanceCACertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadInstanceCACertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadInstanceCACertificateResponse) GetBody() *DownloadInstanceCACertificateResponseBody {
	return s.Body
}

func (s *DownloadInstanceCACertificateResponse) SetHeaders(v map[string]*string) *DownloadInstanceCACertificateResponse {
	s.Headers = v
	return s
}

func (s *DownloadInstanceCACertificateResponse) SetStatusCode(v int32) *DownloadInstanceCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadInstanceCACertificateResponse) SetBody(v *DownloadInstanceCACertificateResponseBody) *DownloadInstanceCACertificateResponse {
	s.Body = v
	return s
}

func (s *DownloadInstanceCACertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
