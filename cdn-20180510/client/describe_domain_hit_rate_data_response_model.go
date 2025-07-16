// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainHitRateDataResponseBody) *DescribeDomainHitRateDataResponse
	GetBody() *DescribeDomainHitRateDataResponseBody
}

type DescribeDomainHitRateDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainHitRateDataResponse) GetBody() *DescribeDomainHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeDomainHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDomainHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainHitRateDataResponse) SetStatusCode(v int32) *DescribeDomainHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainHitRateDataResponse) SetBody(v *DescribeDomainHitRateDataResponseBody) *DescribeDomainHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainHitRateDataResponse) Validate() error {
	return dara.Validate(s)
}
