// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClusterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudClusterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudClusterRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudClusterRuleResponseBody) *DescribeHybridCloudClusterRuleResponse
	GetBody() *DescribeHybridCloudClusterRuleResponseBody
}

type DescribeHybridCloudClusterRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudClusterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudClusterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudClusterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudClusterRuleResponse) GetBody() *DescribeHybridCloudClusterRuleResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudClusterRuleResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudClusterRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudClusterRuleResponse) SetStatusCode(v int32) *DescribeHybridCloudClusterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleResponse) SetBody(v *DescribeHybridCloudClusterRuleResponseBody) *DescribeHybridCloudClusterRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudClusterRuleResponse) Validate() error {
	return dara.Validate(s)
}
