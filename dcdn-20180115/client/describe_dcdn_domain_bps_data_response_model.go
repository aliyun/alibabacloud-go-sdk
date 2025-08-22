// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainBpsDataResponseBody) *DescribeDcdnDomainBpsDataResponse
	GetBody() *DescribeDcdnDomainBpsDataResponseBody
}

type DescribeDcdnDomainBpsDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainBpsDataResponse) GetBody() *DescribeDcdnDomainBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponse) SetBody(v *DescribeDcdnDomainBpsDataResponseBody) *DescribeDcdnDomainBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
