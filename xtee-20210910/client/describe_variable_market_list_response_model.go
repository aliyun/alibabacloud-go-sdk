// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableMarketListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVariableMarketListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVariableMarketListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVariableMarketListResponseBody) *DescribeVariableMarketListResponse
	GetBody() *DescribeVariableMarketListResponseBody
}

type DescribeVariableMarketListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVariableMarketListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVariableMarketListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableMarketListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVariableMarketListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVariableMarketListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVariableMarketListResponse) GetBody() *DescribeVariableMarketListResponseBody {
	return s.Body
}

func (s *DescribeVariableMarketListResponse) SetHeaders(v map[string]*string) *DescribeVariableMarketListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVariableMarketListResponse) SetStatusCode(v int32) *DescribeVariableMarketListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVariableMarketListResponse) SetBody(v *DescribeVariableMarketListResponseBody) *DescribeVariableMarketListResponse {
	s.Body = v
	return s
}

func (s *DescribeVariableMarketListResponse) Validate() error {
	return dara.Validate(s)
}
