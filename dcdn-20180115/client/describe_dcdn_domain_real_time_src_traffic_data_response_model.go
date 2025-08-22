// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeSrcTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRealTimeSrcTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) *DescribeDcdnDomainRealTimeSrcTrafficDataResponse
	GetBody() *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody
}

type DescribeDcdnDomainRealTimeSrcTrafficDataResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponse) GetBody() *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRealTimeSrcTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponse) SetBody(v *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) *DescribeDcdnDomainRealTimeSrcTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
