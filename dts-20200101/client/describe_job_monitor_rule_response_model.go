// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMonitorRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobMonitorRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobMonitorRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobMonitorRuleResponseBody) *DescribeJobMonitorRuleResponse
	GetBody() *DescribeJobMonitorRuleResponseBody
}

type DescribeJobMonitorRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobMonitorRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobMonitorRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMonitorRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobMonitorRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobMonitorRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobMonitorRuleResponse) GetBody() *DescribeJobMonitorRuleResponseBody {
	return s.Body
}

func (s *DescribeJobMonitorRuleResponse) SetHeaders(v map[string]*string) *DescribeJobMonitorRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobMonitorRuleResponse) SetStatusCode(v int32) *DescribeJobMonitorRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobMonitorRuleResponse) SetBody(v *DescribeJobMonitorRuleResponseBody) *DescribeJobMonitorRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeJobMonitorRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
