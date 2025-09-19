// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListStrategyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhiteListStrategyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhiteListStrategyListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhiteListStrategyListResponseBody) *DescribeWhiteListStrategyListResponse
	GetBody() *DescribeWhiteListStrategyListResponseBody
}

type DescribeWhiteListStrategyListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhiteListStrategyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhiteListStrategyListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhiteListStrategyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhiteListStrategyListResponse) GetBody() *DescribeWhiteListStrategyListResponseBody {
	return s.Body
}

func (s *DescribeWhiteListStrategyListResponse) SetHeaders(v map[string]*string) *DescribeWhiteListStrategyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteListStrategyListResponse) SetStatusCode(v int32) *DescribeWhiteListStrategyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhiteListStrategyListResponse) SetBody(v *DescribeWhiteListStrategyListResponseBody) *DescribeWhiteListStrategyListResponse {
	s.Body = v
	return s
}

func (s *DescribeWhiteListStrategyListResponse) Validate() error {
	return dara.Validate(s)
}
