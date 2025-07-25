// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAccessStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmAccessStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmAccessStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmAccessStrategyResponseBody) *DescribeDnsGtmAccessStrategyResponse
	GetBody() *DescribeDnsGtmAccessStrategyResponseBody
}

type DescribeDnsGtmAccessStrategyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmAccessStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmAccessStrategyResponse) GetBody() *DescribeDnsGtmAccessStrategyResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponse) SetStatusCode(v int32) *DescribeDnsGtmAccessStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponse) SetBody(v *DescribeDnsGtmAccessStrategyResponseBody) *DescribeDnsGtmAccessStrategyResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponse) Validate() error {
	return dara.Validate(s)
}
