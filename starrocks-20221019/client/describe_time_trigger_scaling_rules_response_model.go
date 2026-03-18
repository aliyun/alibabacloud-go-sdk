// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTimeTriggerScalingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTimeTriggerScalingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTimeTriggerScalingRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTimeTriggerScalingRulesResponseBody) *DescribeTimeTriggerScalingRulesResponse
	GetBody() *DescribeTimeTriggerScalingRulesResponseBody
}

type DescribeTimeTriggerScalingRulesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTimeTriggerScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTimeTriggerScalingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimeTriggerScalingRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTimeTriggerScalingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTimeTriggerScalingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTimeTriggerScalingRulesResponse) GetBody() *DescribeTimeTriggerScalingRulesResponseBody {
	return s.Body
}

func (s *DescribeTimeTriggerScalingRulesResponse) SetHeaders(v map[string]*string) *DescribeTimeTriggerScalingRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponse) SetStatusCode(v int32) *DescribeTimeTriggerScalingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponse) SetBody(v *DescribeTimeTriggerScalingRulesResponseBody) *DescribeTimeTriggerScalingRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeTimeTriggerScalingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
