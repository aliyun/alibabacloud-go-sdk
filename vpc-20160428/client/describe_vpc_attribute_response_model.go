// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcAttributeResponseBody) *DescribeVpcAttributeResponse
	GetBody() *DescribeVpcAttributeResponseBody
}

type DescribeVpcAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcAttributeResponse) GetBody() *DescribeVpcAttributeResponseBody {
	return s.Body
}

func (s *DescribeVpcAttributeResponse) SetHeaders(v map[string]*string) *DescribeVpcAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcAttributeResponse) SetStatusCode(v int32) *DescribeVpcAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcAttributeResponse) SetBody(v *DescribeVpcAttributeResponseBody) *DescribeVpcAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcAttributeResponse) Validate() error {
	return dara.Validate(s)
}
