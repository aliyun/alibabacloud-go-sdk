// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeReqHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeReqHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRealTimeReqHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) *DescribeDcdnDomainRealTimeReqHitRateDataResponse
	GetBody() *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody
}

type DescribeDcdnDomainRealTimeReqHitRateDataResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponse) GetBody() *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeReqHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRealTimeReqHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponse) SetBody(v *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) *DescribeDcdnDomainRealTimeReqHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponse) Validate() error {
	return dara.Validate(s)
}
