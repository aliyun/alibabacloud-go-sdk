// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteOriginClientCertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSiteOriginClientCertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSiteOriginClientCertificatesResponse
	GetStatusCode() *int32
	SetBody(v *ListSiteOriginClientCertificatesResponseBody) *ListSiteOriginClientCertificatesResponse
	GetBody() *ListSiteOriginClientCertificatesResponseBody
}

type ListSiteOriginClientCertificatesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSiteOriginClientCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSiteOriginClientCertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSiteOriginClientCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListSiteOriginClientCertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSiteOriginClientCertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSiteOriginClientCertificatesResponse) GetBody() *ListSiteOriginClientCertificatesResponseBody {
	return s.Body
}

func (s *ListSiteOriginClientCertificatesResponse) SetHeaders(v map[string]*string) *ListSiteOriginClientCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListSiteOriginClientCertificatesResponse) SetStatusCode(v int32) *ListSiteOriginClientCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponse) SetBody(v *ListSiteOriginClientCertificatesResponseBody) *ListSiteOriginClientCertificatesResponse {
	s.Body = v
	return s
}

func (s *ListSiteOriginClientCertificatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
