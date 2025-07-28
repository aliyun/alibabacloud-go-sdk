// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClusterRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudClusterRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudClusterRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudClusterRulesResponseBody) *DescribeHybridCloudClusterRulesResponse
	GetBody() *DescribeHybridCloudClusterRulesResponseBody
}

type DescribeHybridCloudClusterRulesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudClusterRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudClusterRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudClusterRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudClusterRulesResponse) GetBody() *DescribeHybridCloudClusterRulesResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudClusterRulesResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudClusterRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponse) SetStatusCode(v int32) *DescribeHybridCloudClusterRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponse) SetBody(v *DescribeHybridCloudClusterRulesResponseBody) *DescribeHybridCloudClusterRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponse) Validate() error {
	return dara.Validate(s)
}
