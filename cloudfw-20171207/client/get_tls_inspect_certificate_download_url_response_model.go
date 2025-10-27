// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlsInspectCertificateDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTlsInspectCertificateDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTlsInspectCertificateDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetTlsInspectCertificateDownloadUrlResponseBody) *GetTlsInspectCertificateDownloadUrlResponse
	GetBody() *GetTlsInspectCertificateDownloadUrlResponseBody
}

type GetTlsInspectCertificateDownloadUrlResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTlsInspectCertificateDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTlsInspectCertificateDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTlsInspectCertificateDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetTlsInspectCertificateDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTlsInspectCertificateDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTlsInspectCertificateDownloadUrlResponse) GetBody() *GetTlsInspectCertificateDownloadUrlResponseBody {
	return s.Body
}

func (s *GetTlsInspectCertificateDownloadUrlResponse) SetHeaders(v map[string]*string) *GetTlsInspectCertificateDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetTlsInspectCertificateDownloadUrlResponse) SetStatusCode(v int32) *GetTlsInspectCertificateDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTlsInspectCertificateDownloadUrlResponse) SetBody(v *GetTlsInspectCertificateDownloadUrlResponseBody) *GetTlsInspectCertificateDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetTlsInspectCertificateDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
