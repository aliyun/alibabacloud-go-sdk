// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainHitRateDataResponseBody) *DescribeDcdnDomainHitRateDataResponse
	GetBody() *DescribeDcdnDomainHitRateDataResponseBody
}

type DescribeDcdnDomainHitRateDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainHitRateDataResponse) GetBody() *DescribeDcdnDomainHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponse) SetBody(v *DescribeDcdnDomainHitRateDataResponseBody) *DescribeDcdnDomainHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponse) Validate() error {
	return dara.Validate(s)
}
