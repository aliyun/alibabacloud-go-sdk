// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainCertificateInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainCertificateInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainCertificateInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainCertificateInfoResponseBody) *DescribeDcdnDomainCertificateInfoResponse
	GetBody() *DescribeDcdnDomainCertificateInfoResponseBody
}

type DescribeDcdnDomainCertificateInfoResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainCertificateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainCertificateInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCertificateInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainCertificateInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainCertificateInfoResponse) GetBody() *DescribeDcdnDomainCertificateInfoResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainCertificateInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainCertificateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponse) SetStatusCode(v int32) *DescribeDcdnDomainCertificateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponse) SetBody(v *DescribeDcdnDomainCertificateInfoResponseBody) *DescribeDcdnDomainCertificateInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponse) Validate() error {
	return dara.Validate(s)
}
