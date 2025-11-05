// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeByteHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealTimeByteHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealTimeByteHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealTimeByteHitRateDataResponseBody) *DescribeDomainRealTimeByteHitRateDataResponse
	GetBody() *DescribeDomainRealTimeByteHitRateDataResponseBody
}

type DescribeDomainRealTimeByteHitRateDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealTimeByteHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealTimeByteHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeByteHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeByteHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealTimeByteHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealTimeByteHitRateDataResponse) GetBody() *DescribeDomainRealTimeByteHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRealTimeByteHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeByteHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponse) SetStatusCode(v int32) *DescribeDomainRealTimeByteHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponse) SetBody(v *DescribeDomainRealTimeByteHitRateDataResponseBody) *DescribeDomainRealTimeByteHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
