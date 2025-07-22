// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterAllowedPrefixHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse
	GetBody() *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody
}

type DescribeExpressConnectRouterAllowedPrefixHistoryResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) GetBody() *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	return s.Body
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) SetBody(v *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) Validate() error {
	return dara.Validate(s)
}
