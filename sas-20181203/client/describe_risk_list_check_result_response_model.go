// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskListCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskListCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskListCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskListCheckResultResponseBody) *DescribeRiskListCheckResultResponse
	GetBody() *DescribeRiskListCheckResultResponseBody
}

type DescribeRiskListCheckResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskListCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskListCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskListCheckResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskListCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskListCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskListCheckResultResponse) GetBody() *DescribeRiskListCheckResultResponseBody {
	return s.Body
}

func (s *DescribeRiskListCheckResultResponse) SetHeaders(v map[string]*string) *DescribeRiskListCheckResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskListCheckResultResponse) SetStatusCode(v int32) *DescribeRiskListCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskListCheckResultResponse) SetBody(v *DescribeRiskListCheckResultResponseBody) *DescribeRiskListCheckResultResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskListCheckResultResponse) Validate() error {
	return dara.Validate(s)
}
