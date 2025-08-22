// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserCertificateExpireCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserCertificateExpireCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserCertificateExpireCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserCertificateExpireCountResponseBody) *DescribeDcdnUserCertificateExpireCountResponse
	GetBody() *DescribeDcdnUserCertificateExpireCountResponseBody
}

type DescribeDcdnUserCertificateExpireCountResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserCertificateExpireCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserCertificateExpireCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserCertificateExpireCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserCertificateExpireCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserCertificateExpireCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserCertificateExpireCountResponse) GetBody() *DescribeDcdnUserCertificateExpireCountResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserCertificateExpireCountResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserCertificateExpireCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserCertificateExpireCountResponse) SetStatusCode(v int32) *DescribeDcdnUserCertificateExpireCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserCertificateExpireCountResponse) SetBody(v *DescribeDcdnUserCertificateExpireCountResponseBody) *DescribeDcdnUserCertificateExpireCountResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserCertificateExpireCountResponse) Validate() error {
	return dara.Validate(s)
}
