// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSiteOriginClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadSiteOriginClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadSiteOriginClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *UploadSiteOriginClientCertificateResponseBody) *UploadSiteOriginClientCertificateResponse
	GetBody() *UploadSiteOriginClientCertificateResponseBody
}

type UploadSiteOriginClientCertificateResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadSiteOriginClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadSiteOriginClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadSiteOriginClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadSiteOriginClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadSiteOriginClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadSiteOriginClientCertificateResponse) GetBody() *UploadSiteOriginClientCertificateResponseBody {
	return s.Body
}

func (s *UploadSiteOriginClientCertificateResponse) SetHeaders(v map[string]*string) *UploadSiteOriginClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadSiteOriginClientCertificateResponse) SetStatusCode(v int32) *UploadSiteOriginClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponse) SetBody(v *UploadSiteOriginClientCertificateResponseBody) *UploadSiteOriginClientCertificateResponse {
	s.Body = v
	return s
}

func (s *UploadSiteOriginClientCertificateResponse) Validate() error {
	return dara.Validate(s)
}
