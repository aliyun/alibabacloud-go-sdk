// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeSrcTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealTimeSrcTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealTimeSrcTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealTimeSrcTrafficDataResponseBody) *DescribeDomainRealTimeSrcTrafficDataResponse
	GetBody() *DescribeDomainRealTimeSrcTrafficDataResponseBody
}

type DescribeDomainRealTimeSrcTrafficDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealTimeSrcTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealTimeSrcTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponse) GetBody() *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeSrcTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponse) SetStatusCode(v int32) *DescribeDomainRealTimeSrcTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponse) SetBody(v *DescribeDomainRealTimeSrcTrafficDataResponseBody) *DescribeDomainRealTimeSrcTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
