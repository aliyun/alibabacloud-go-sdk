// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariableVersionDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressionVariableVersionDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressionVariableVersionDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressionVariableVersionDetailResponseBody) *DescribeExpressionVariableVersionDetailResponse
	GetBody() *DescribeExpressionVariableVersionDetailResponseBody
}

type DescribeExpressionVariableVersionDetailResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressionVariableVersionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressionVariableVersionDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableVersionDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableVersionDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressionVariableVersionDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressionVariableVersionDetailResponse) GetBody() *DescribeExpressionVariableVersionDetailResponseBody {
	return s.Body
}

func (s *DescribeExpressionVariableVersionDetailResponse) SetHeaders(v map[string]*string) *DescribeExpressionVariableVersionDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponse) SetStatusCode(v int32) *DescribeExpressionVariableVersionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponse) SetBody(v *DescribeExpressionVariableVersionDetailResponseBody) *DescribeExpressionVariableVersionDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
