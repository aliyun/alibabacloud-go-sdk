// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealTimeTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealTimeTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealTimeTrafficDataResponseBody) *DescribeDomainRealTimeTrafficDataResponse
	GetBody() *DescribeDomainRealTimeTrafficDataResponseBody
}

type DescribeDomainRealTimeTrafficDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealTimeTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealTimeTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealTimeTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealTimeTrafficDataResponse) GetBody() *DescribeDomainRealTimeTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRealTimeTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponse) SetStatusCode(v int32) *DescribeDomainRealTimeTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponse) SetBody(v *DescribeDomainRealTimeTrafficDataResponseBody) *DescribeDomainRealTimeTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
