// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteOriginClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSiteOriginClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSiteOriginClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *GetSiteOriginClientCertificateResponseBody) *GetSiteOriginClientCertificateResponse
	GetBody() *GetSiteOriginClientCertificateResponseBody
}

type GetSiteOriginClientCertificateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteOriginClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteOriginClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSiteOriginClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetSiteOriginClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSiteOriginClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSiteOriginClientCertificateResponse) GetBody() *GetSiteOriginClientCertificateResponseBody {
	return s.Body
}

func (s *GetSiteOriginClientCertificateResponse) SetHeaders(v map[string]*string) *GetSiteOriginClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetSiteOriginClientCertificateResponse) SetStatusCode(v int32) *GetSiteOriginClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponse) SetBody(v *GetSiteOriginClientCertificateResponseBody) *GetSiteOriginClientCertificateResponse {
	s.Body = v
	return s
}

func (s *GetSiteOriginClientCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
