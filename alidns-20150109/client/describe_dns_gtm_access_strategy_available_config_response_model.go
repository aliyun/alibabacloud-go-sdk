// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAccessStrategyAvailableConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmAccessStrategyAvailableConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmAccessStrategyAvailableConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) *DescribeDnsGtmAccessStrategyAvailableConfigResponse
	GetBody() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponse) GetBody() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAccessStrategyAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponse) SetStatusCode(v int32) *DescribeDnsGtmAccessStrategyAvailableConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponse) SetBody(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) *DescribeDnsGtmAccessStrategyAvailableConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponse) Validate() error {
	return dara.Validate(s)
}
