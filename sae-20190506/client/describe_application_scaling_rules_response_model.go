// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationScalingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationScalingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationScalingRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationScalingRulesResponseBody) *DescribeApplicationScalingRulesResponse
	GetBody() *DescribeApplicationScalingRulesResponseBody
}

type DescribeApplicationScalingRulesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationScalingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationScalingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationScalingRulesResponse) GetBody() *DescribeApplicationScalingRulesResponseBody {
	return s.Body
}

func (s *DescribeApplicationScalingRulesResponse) SetHeaders(v map[string]*string) *DescribeApplicationScalingRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationScalingRulesResponse) SetStatusCode(v int32) *DescribeApplicationScalingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponse) SetBody(v *DescribeApplicationScalingRulesResponseBody) *DescribeApplicationScalingRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationScalingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
