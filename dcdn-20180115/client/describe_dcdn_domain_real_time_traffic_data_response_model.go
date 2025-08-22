// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRealTimeTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRealTimeTrafficDataResponseBody) *DescribeDcdnDomainRealTimeTrafficDataResponse
	GetBody() *DescribeDcdnDomainRealTimeTrafficDataResponseBody
}

type DescribeDcdnDomainRealTimeTrafficDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRealTimeTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponse) GetBody() *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRealTimeTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponse) SetBody(v *DescribeDcdnDomainRealTimeTrafficDataResponseBody) *DescribeDcdnDomainRealTimeTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
