// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainByCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainByCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainByCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainByCertificateResponseBody) *DescribeDcdnDomainByCertificateResponse
	GetBody() *DescribeDcdnDomainByCertificateResponseBody
}

type DescribeDcdnDomainByCertificateResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainByCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainByCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainByCertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainByCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainByCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainByCertificateResponse) GetBody() *DescribeDcdnDomainByCertificateResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainByCertificateResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainByCertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponse) SetStatusCode(v int32) *DescribeDcdnDomainByCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponse) SetBody(v *DescribeDcdnDomainByCertificateResponseBody) *DescribeDcdnDomainByCertificateResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponse) Validate() error {
	return dara.Validate(s)
}
