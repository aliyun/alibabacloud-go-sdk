// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOssPolicyResponseBody) *DescribeOssPolicyResponse
	GetBody() *DescribeOssPolicyResponseBody
}

type DescribeOssPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssPolicyResponse) GetBody() *DescribeOssPolicyResponseBody {
	return s.Body
}

func (s *DescribeOssPolicyResponse) SetHeaders(v map[string]*string) *DescribeOssPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssPolicyResponse) SetStatusCode(v int32) *DescribeOssPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssPolicyResponse) SetBody(v *DescribeOssPolicyResponseBody) *DescribeOssPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeOssPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
