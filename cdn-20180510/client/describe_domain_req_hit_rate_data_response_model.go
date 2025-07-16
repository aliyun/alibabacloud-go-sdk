// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainReqHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainReqHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainReqHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainReqHitRateDataResponseBody) *DescribeDomainReqHitRateDataResponse
	GetBody() *DescribeDomainReqHitRateDataResponseBody
}

type DescribeDomainReqHitRateDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainReqHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainReqHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainReqHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainReqHitRateDataResponse) GetBody() *DescribeDomainReqHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeDomainReqHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDomainReqHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainReqHitRateDataResponse) SetStatusCode(v int32) *DescribeDomainReqHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponse) SetBody(v *DescribeDomainReqHitRateDataResponseBody) *DescribeDomainReqHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainReqHitRateDataResponse) Validate() error {
	return dara.Validate(s)
}
