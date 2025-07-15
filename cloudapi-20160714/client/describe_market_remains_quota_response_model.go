// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarketRemainsQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMarketRemainsQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMarketRemainsQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMarketRemainsQuotaResponseBody) *DescribeMarketRemainsQuotaResponse
	GetBody() *DescribeMarketRemainsQuotaResponseBody
}

type DescribeMarketRemainsQuotaResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMarketRemainsQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMarketRemainsQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketRemainsQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeMarketRemainsQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMarketRemainsQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMarketRemainsQuotaResponse) GetBody() *DescribeMarketRemainsQuotaResponseBody {
	return s.Body
}

func (s *DescribeMarketRemainsQuotaResponse) SetHeaders(v map[string]*string) *DescribeMarketRemainsQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeMarketRemainsQuotaResponse) SetStatusCode(v int32) *DescribeMarketRemainsQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMarketRemainsQuotaResponse) SetBody(v *DescribeMarketRemainsQuotaResponseBody) *DescribeMarketRemainsQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeMarketRemainsQuotaResponse) Validate() error {
	return dara.Validate(s)
}
