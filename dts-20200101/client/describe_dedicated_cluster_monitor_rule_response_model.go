// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedClusterMonitorRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDedicatedClusterMonitorRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDedicatedClusterMonitorRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDedicatedClusterMonitorRuleResponseBody) *DescribeDedicatedClusterMonitorRuleResponse
	GetBody() *DescribeDedicatedClusterMonitorRuleResponseBody
}

type DescribeDedicatedClusterMonitorRuleResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedClusterMonitorRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedClusterMonitorRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterMonitorRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterMonitorRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDedicatedClusterMonitorRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDedicatedClusterMonitorRuleResponse) GetBody() *DescribeDedicatedClusterMonitorRuleResponseBody {
	return s.Body
}

func (s *DescribeDedicatedClusterMonitorRuleResponse) SetHeaders(v map[string]*string) *DescribeDedicatedClusterMonitorRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponse) SetStatusCode(v int32) *DescribeDedicatedClusterMonitorRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponse) SetBody(v *DescribeDedicatedClusterMonitorRuleResponseBody) *DescribeDedicatedClusterMonitorRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
