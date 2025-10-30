// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariablePageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressionVariablePageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressionVariablePageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressionVariablePageResponseBody) *DescribeExpressionVariablePageResponse
	GetBody() *DescribeExpressionVariablePageResponseBody
}

type DescribeExpressionVariablePageResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressionVariablePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressionVariablePageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariablePageResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariablePageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressionVariablePageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressionVariablePageResponse) GetBody() *DescribeExpressionVariablePageResponseBody {
	return s.Body
}

func (s *DescribeExpressionVariablePageResponse) SetHeaders(v map[string]*string) *DescribeExpressionVariablePageResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressionVariablePageResponse) SetStatusCode(v int32) *DescribeExpressionVariablePageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressionVariablePageResponse) SetBody(v *DescribeExpressionVariablePageResponseBody) *DescribeExpressionVariablePageResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressionVariablePageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
