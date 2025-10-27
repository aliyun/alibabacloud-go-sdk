// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAntiBruteForceRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceAntiBruteForceRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceAntiBruteForceRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceAntiBruteForceRulesResponseBody) *DescribeInstanceAntiBruteForceRulesResponse
	GetBody() *DescribeInstanceAntiBruteForceRulesResponseBody
}

type DescribeInstanceAntiBruteForceRulesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAntiBruteForceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAntiBruteForceRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAntiBruteForceRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAntiBruteForceRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceAntiBruteForceRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceAntiBruteForceRulesResponse) GetBody() *DescribeInstanceAntiBruteForceRulesResponseBody {
	return s.Body
}

func (s *DescribeInstanceAntiBruteForceRulesResponse) SetHeaders(v map[string]*string) *DescribeInstanceAntiBruteForceRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponse) SetStatusCode(v int32) *DescribeInstanceAntiBruteForceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponse) SetBody(v *DescribeInstanceAntiBruteForceRulesResponseBody) *DescribeInstanceAntiBruteForceRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
