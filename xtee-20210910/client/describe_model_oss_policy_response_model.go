// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOssPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelOssPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelOssPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelOssPolicyResponseBody) *DescribeModelOssPolicyResponse
	GetBody() *DescribeModelOssPolicyResponseBody
}

type DescribeModelOssPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelOssPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelOssPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOssPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelOssPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelOssPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelOssPolicyResponse) GetBody() *DescribeModelOssPolicyResponseBody {
	return s.Body
}

func (s *DescribeModelOssPolicyResponse) SetHeaders(v map[string]*string) *DescribeModelOssPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelOssPolicyResponse) SetStatusCode(v int32) *DescribeModelOssPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelOssPolicyResponse) SetBody(v *DescribeModelOssPolicyResponseBody) *DescribeModelOssPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeModelOssPolicyResponse) Validate() error {
	return dara.Validate(s)
}
