// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnCertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnCertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnCertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnCertificateListResponseBody) *DescribeDcdnCertificateListResponse
	GetBody() *DescribeDcdnCertificateListResponseBody
}

type DescribeDcdnCertificateListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnCertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnCertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnCertificateListResponse) GetBody() *DescribeDcdnCertificateListResponseBody {
	return s.Body
}

func (s *DescribeDcdnCertificateListResponse) SetHeaders(v map[string]*string) *DescribeDcdnCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnCertificateListResponse) SetStatusCode(v int32) *DescribeDcdnCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnCertificateListResponse) SetBody(v *DescribeDcdnCertificateListResponseBody) *DescribeDcdnCertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnCertificateListResponse) Validate() error {
	return dara.Validate(s)
}
