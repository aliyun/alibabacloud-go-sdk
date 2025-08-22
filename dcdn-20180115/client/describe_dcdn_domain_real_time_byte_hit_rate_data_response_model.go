// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeByteHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeByteHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRealTimeByteHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) *DescribeDcdnDomainRealTimeByteHitRateDataResponse
	GetBody() *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody
}

type DescribeDcdnDomainRealTimeByteHitRateDataResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponse) GetBody() *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeByteHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRealTimeByteHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponse) SetBody(v *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) *DescribeDcdnDomainRealTimeByteHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponse) Validate() error {
	return dara.Validate(s)
}
