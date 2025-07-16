// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeReqHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealTimeReqHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealTimeReqHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealTimeReqHitRateDataResponseBody) *DescribeDomainRealTimeReqHitRateDataResponse
	GetBody() *DescribeDomainRealTimeReqHitRateDataResponseBody
}

type DescribeDomainRealTimeReqHitRateDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealTimeReqHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealTimeReqHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeReqHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeReqHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealTimeReqHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealTimeReqHitRateDataResponse) GetBody() *DescribeDomainRealTimeReqHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRealTimeReqHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeReqHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponse) SetStatusCode(v int32) *DescribeDomainRealTimeReqHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponse) SetBody(v *DescribeDomainRealTimeReqHitRateDataResponseBody) *DescribeDomainRealTimeReqHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponse) Validate() error {
	return dara.Validate(s)
}
