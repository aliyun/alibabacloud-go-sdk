// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainByCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnDomainByCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnDomainByCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnDomainByCertificateResponseBody) *DescribeCdnDomainByCertificateResponse
	GetBody() *DescribeCdnDomainByCertificateResponseBody
}

type DescribeCdnDomainByCertificateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnDomainByCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnDomainByCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainByCertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainByCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnDomainByCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnDomainByCertificateResponse) GetBody() *DescribeCdnDomainByCertificateResponseBody {
	return s.Body
}

func (s *DescribeCdnDomainByCertificateResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainByCertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainByCertificateResponse) SetStatusCode(v int32) *DescribeCdnDomainByCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponse) SetBody(v *DescribeCdnDomainByCertificateResponseBody) *DescribeCdnDomainByCertificateResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnDomainByCertificateResponse) Validate() error {
	return dara.Validate(s)
}
