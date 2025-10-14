// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteOriginClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSiteOriginClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSiteOriginClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSiteOriginClientCertificateResponseBody) *DeleteSiteOriginClientCertificateResponse
	GetBody() *DeleteSiteOriginClientCertificateResponseBody
}

type DeleteSiteOriginClientCertificateResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSiteOriginClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSiteOriginClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteOriginClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteSiteOriginClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSiteOriginClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSiteOriginClientCertificateResponse) GetBody() *DeleteSiteOriginClientCertificateResponseBody {
	return s.Body
}

func (s *DeleteSiteOriginClientCertificateResponse) SetHeaders(v map[string]*string) *DeleteSiteOriginClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteSiteOriginClientCertificateResponse) SetStatusCode(v int32) *DeleteSiteOriginClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSiteOriginClientCertificateResponse) SetBody(v *DeleteSiteOriginClientCertificateResponseBody) *DeleteSiteOriginClientCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteSiteOriginClientCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
