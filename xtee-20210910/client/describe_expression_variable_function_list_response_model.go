// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariableFunctionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressionVariableFunctionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressionVariableFunctionListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressionVariableFunctionListResponseBody) *DescribeExpressionVariableFunctionListResponse
	GetBody() *DescribeExpressionVariableFunctionListResponseBody
}

type DescribeExpressionVariableFunctionListResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressionVariableFunctionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressionVariableFunctionListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableFunctionListResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableFunctionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressionVariableFunctionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressionVariableFunctionListResponse) GetBody() *DescribeExpressionVariableFunctionListResponseBody {
	return s.Body
}

func (s *DescribeExpressionVariableFunctionListResponse) SetHeaders(v map[string]*string) *DescribeExpressionVariableFunctionListResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponse) SetStatusCode(v int32) *DescribeExpressionVariableFunctionListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponse) SetBody(v *DescribeExpressionVariableFunctionListResponseBody) *DescribeExpressionVariableFunctionListResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponse) Validate() error {
	return dara.Validate(s)
}
