// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectTrafficQosRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressConnectTrafficQosRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressConnectTrafficQosRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressConnectTrafficQosRuleResponseBody) *DescribeExpressConnectTrafficQosRuleResponse
	GetBody() *DescribeExpressConnectTrafficQosRuleResponseBody
}

type DescribeExpressConnectTrafficQosRuleResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectTrafficQosRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectTrafficQosRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressConnectTrafficQosRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressConnectTrafficQosRuleResponse) GetBody() *DescribeExpressConnectTrafficQosRuleResponseBody {
	return s.Body
}

func (s *DescribeExpressConnectTrafficQosRuleResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectTrafficQosRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponse) SetStatusCode(v int32) *DescribeExpressConnectTrafficQosRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponse) SetBody(v *DescribeExpressConnectTrafficQosRuleResponseBody) *DescribeExpressConnectTrafficQosRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
