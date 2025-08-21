// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopFingerprintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainTopFingerprintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainTopFingerprintResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainTopFingerprintResponseBody) *DescribeDomainTopFingerprintResponse
	GetBody() *DescribeDomainTopFingerprintResponseBody
}

type DescribeDomainTopFingerprintResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainTopFingerprintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainTopFingerprintResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopFingerprintResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopFingerprintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainTopFingerprintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainTopFingerprintResponse) GetBody() *DescribeDomainTopFingerprintResponseBody {
	return s.Body
}

func (s *DescribeDomainTopFingerprintResponse) SetHeaders(v map[string]*string) *DescribeDomainTopFingerprintResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopFingerprintResponse) SetStatusCode(v int32) *DescribeDomainTopFingerprintResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainTopFingerprintResponse) SetBody(v *DescribeDomainTopFingerprintResponseBody) *DescribeDomainTopFingerprintResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainTopFingerprintResponse) Validate() error {
	return dara.Validate(s)
}
