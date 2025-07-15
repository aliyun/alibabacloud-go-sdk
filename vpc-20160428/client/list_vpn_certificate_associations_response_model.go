// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpnCertificateAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpnCertificateAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpnCertificateAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *ListVpnCertificateAssociationsResponseBody) *ListVpnCertificateAssociationsResponse
	GetBody() *ListVpnCertificateAssociationsResponseBody
}

type ListVpnCertificateAssociationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpnCertificateAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpnCertificateAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpnCertificateAssociationsResponse) GoString() string {
	return s.String()
}

func (s *ListVpnCertificateAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpnCertificateAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpnCertificateAssociationsResponse) GetBody() *ListVpnCertificateAssociationsResponseBody {
	return s.Body
}

func (s *ListVpnCertificateAssociationsResponse) SetHeaders(v map[string]*string) *ListVpnCertificateAssociationsResponse {
	s.Headers = v
	return s
}

func (s *ListVpnCertificateAssociationsResponse) SetStatusCode(v int32) *ListVpnCertificateAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpnCertificateAssociationsResponse) SetBody(v *ListVpnCertificateAssociationsResponseBody) *ListVpnCertificateAssociationsResponse {
	s.Body = v
	return s
}

func (s *ListVpnCertificateAssociationsResponse) Validate() error {
	return dara.Validate(s)
}
