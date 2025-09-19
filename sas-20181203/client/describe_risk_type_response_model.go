// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskTypeResponseBody) *DescribeRiskTypeResponse
	GetBody() *DescribeRiskTypeResponseBody
}

type DescribeRiskTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskTypeResponse) GetBody() *DescribeRiskTypeResponseBody {
	return s.Body
}

func (s *DescribeRiskTypeResponse) SetHeaders(v map[string]*string) *DescribeRiskTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskTypeResponse) SetStatusCode(v int32) *DescribeRiskTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskTypeResponse) SetBody(v *DescribeRiskTypeResponseBody) *DescribeRiskTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskTypeResponse) Validate() error {
	return dara.Validate(s)
}
