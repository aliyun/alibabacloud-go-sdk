// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAccessStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmAccessStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmAccessStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmAccessStrategiesResponseBody) *DescribeDnsGtmAccessStrategiesResponse
	GetBody() *DescribeDnsGtmAccessStrategiesResponseBody
}

type DescribeDnsGtmAccessStrategiesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmAccessStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmAccessStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmAccessStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmAccessStrategiesResponse) GetBody() *DescribeDnsGtmAccessStrategiesResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmAccessStrategiesResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAccessStrategiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponse) SetStatusCode(v int32) *DescribeDnsGtmAccessStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponse) SetBody(v *DescribeDnsGtmAccessStrategiesResponseBody) *DescribeDnsGtmAccessStrategiesResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponse) Validate() error {
	return dara.Validate(s)
}
