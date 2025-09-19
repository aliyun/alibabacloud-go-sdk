// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAntiBruteForceRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAntiBruteForceRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAntiBruteForceRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAntiBruteForceRulesResponseBody) *DescribeAntiBruteForceRulesResponse
	GetBody() *DescribeAntiBruteForceRulesResponseBody
}

type DescribeAntiBruteForceRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntiBruteForceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntiBruteForceRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAntiBruteForceRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAntiBruteForceRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAntiBruteForceRulesResponse) GetBody() *DescribeAntiBruteForceRulesResponseBody {
	return s.Body
}

func (s *DescribeAntiBruteForceRulesResponse) SetHeaders(v map[string]*string) *DescribeAntiBruteForceRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntiBruteForceRulesResponse) SetStatusCode(v int32) *DescribeAntiBruteForceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponse) SetBody(v *DescribeAntiBruteForceRulesResponseBody) *DescribeAntiBruteForceRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeAntiBruteForceRulesResponse) Validate() error {
	return dara.Validate(s)
}
