// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDecisionResultTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDecisionResultTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDecisionResultTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDecisionResultTrendResponseBody) *DescribeDecisionResultTrendResponse
	GetBody() *DescribeDecisionResultTrendResponseBody
}

type DescribeDecisionResultTrendResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDecisionResultTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDecisionResultTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDecisionResultTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDecisionResultTrendResponse) GetBody() *DescribeDecisionResultTrendResponseBody {
	return s.Body
}

func (s *DescribeDecisionResultTrendResponse) SetHeaders(v map[string]*string) *DescribeDecisionResultTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeDecisionResultTrendResponse) SetStatusCode(v int32) *DescribeDecisionResultTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDecisionResultTrendResponse) SetBody(v *DescribeDecisionResultTrendResponseBody) *DescribeDecisionResultTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeDecisionResultTrendResponse) Validate() error {
	return dara.Validate(s)
}
