// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainH2FingerprintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainH2FingerprintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainH2FingerprintResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainH2FingerprintResponseBody) *DescribeDomainH2FingerprintResponse
	GetBody() *DescribeDomainH2FingerprintResponseBody
}

type DescribeDomainH2FingerprintResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainH2FingerprintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainH2FingerprintResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainH2FingerprintResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainH2FingerprintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainH2FingerprintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainH2FingerprintResponse) GetBody() *DescribeDomainH2FingerprintResponseBody {
	return s.Body
}

func (s *DescribeDomainH2FingerprintResponse) SetHeaders(v map[string]*string) *DescribeDomainH2FingerprintResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainH2FingerprintResponse) SetStatusCode(v int32) *DescribeDomainH2FingerprintResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainH2FingerprintResponse) SetBody(v *DescribeDomainH2FingerprintResponseBody) *DescribeDomainH2FingerprintResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainH2FingerprintResponse) Validate() error {
	return dara.Validate(s)
}
