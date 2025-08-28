// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcGrantRulesToEcrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcGrantRulesToEcrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcGrantRulesToEcrResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcGrantRulesToEcrResponseBody) *DescribeVpcGrantRulesToEcrResponse
	GetBody() *DescribeVpcGrantRulesToEcrResponseBody
}

type DescribeVpcGrantRulesToEcrResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcGrantRulesToEcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcGrantRulesToEcrResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcGrantRulesToEcrResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcGrantRulesToEcrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcGrantRulesToEcrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcGrantRulesToEcrResponse) GetBody() *DescribeVpcGrantRulesToEcrResponseBody {
	return s.Body
}

func (s *DescribeVpcGrantRulesToEcrResponse) SetHeaders(v map[string]*string) *DescribeVpcGrantRulesToEcrResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponse) SetStatusCode(v int32) *DescribeVpcGrantRulesToEcrResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponse) SetBody(v *DescribeVpcGrantRulesToEcrResponseBody) *DescribeVpcGrantRulesToEcrResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponse) Validate() error {
	return dara.Validate(s)
}
