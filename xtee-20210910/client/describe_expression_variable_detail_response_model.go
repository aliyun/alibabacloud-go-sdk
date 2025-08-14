// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariableDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressionVariableDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressionVariableDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressionVariableDetailResponseBody) *DescribeExpressionVariableDetailResponse
	GetBody() *DescribeExpressionVariableDetailResponseBody
}

type DescribeExpressionVariableDetailResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressionVariableDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressionVariableDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressionVariableDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressionVariableDetailResponse) GetBody() *DescribeExpressionVariableDetailResponseBody {
	return s.Body
}

func (s *DescribeExpressionVariableDetailResponse) SetHeaders(v map[string]*string) *DescribeExpressionVariableDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressionVariableDetailResponse) SetStatusCode(v int32) *DescribeExpressionVariableDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressionVariableDetailResponse) SetBody(v *DescribeExpressionVariableDetailResponseBody) *DescribeExpressionVariableDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressionVariableDetailResponse) Validate() error {
	return dara.Validate(s)
}
