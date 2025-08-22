// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSMCertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnSMCertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnSMCertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnSMCertificateListResponseBody) *DescribeDcdnSMCertificateListResponse
	GetBody() *DescribeDcdnSMCertificateListResponseBody
}

type DescribeDcdnSMCertificateListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnSMCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnSMCertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSMCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnSMCertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnSMCertificateListResponse) GetBody() *DescribeDcdnSMCertificateListResponseBody {
	return s.Body
}

func (s *DescribeDcdnSMCertificateListResponse) SetHeaders(v map[string]*string) *DescribeDcdnSMCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSMCertificateListResponse) SetStatusCode(v int32) *DescribeDcdnSMCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnSMCertificateListResponse) SetBody(v *DescribeDcdnSMCertificateListResponseBody) *DescribeDcdnSMCertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnSMCertificateListResponse) Validate() error {
	return dara.Validate(s)
}
