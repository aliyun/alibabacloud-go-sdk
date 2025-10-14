// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupDynamicRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorGroupDynamicRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorGroupDynamicRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorGroupDynamicRulesResponseBody) *DescribeMonitorGroupDynamicRulesResponse
	GetBody() *DescribeMonitorGroupDynamicRulesResponseBody
}

type DescribeMonitorGroupDynamicRulesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorGroupDynamicRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorGroupDynamicRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorGroupDynamicRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorGroupDynamicRulesResponse) GetBody() *DescribeMonitorGroupDynamicRulesResponseBody {
	return s.Body
}

func (s *DescribeMonitorGroupDynamicRulesResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupDynamicRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponse) SetStatusCode(v int32) *DescribeMonitorGroupDynamicRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponse) SetBody(v *DescribeMonitorGroupDynamicRulesResponseBody) *DescribeMonitorGroupDynamicRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
