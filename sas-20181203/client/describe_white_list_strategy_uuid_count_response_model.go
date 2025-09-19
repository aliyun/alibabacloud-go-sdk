// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListStrategyUuidCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhiteListStrategyUuidCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhiteListStrategyUuidCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhiteListStrategyUuidCountResponseBody) *DescribeWhiteListStrategyUuidCountResponse
	GetBody() *DescribeWhiteListStrategyUuidCountResponseBody
}

type DescribeWhiteListStrategyUuidCountResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhiteListStrategyUuidCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhiteListStrategyUuidCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyUuidCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyUuidCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhiteListStrategyUuidCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhiteListStrategyUuidCountResponse) GetBody() *DescribeWhiteListStrategyUuidCountResponseBody {
	return s.Body
}

func (s *DescribeWhiteListStrategyUuidCountResponse) SetHeaders(v map[string]*string) *DescribeWhiteListStrategyUuidCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteListStrategyUuidCountResponse) SetStatusCode(v int32) *DescribeWhiteListStrategyUuidCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhiteListStrategyUuidCountResponse) SetBody(v *DescribeWhiteListStrategyUuidCountResponseBody) *DescribeWhiteListStrategyUuidCountResponse {
	s.Body = v
	return s
}

func (s *DescribeWhiteListStrategyUuidCountResponse) Validate() error {
	return dara.Validate(s)
}
