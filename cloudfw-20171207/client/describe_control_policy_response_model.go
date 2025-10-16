// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeControlPolicyResponseBody) *DescribeControlPolicyResponse
	GetBody() *DescribeControlPolicyResponseBody
}

type DescribeControlPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeControlPolicyResponse) GetBody() *DescribeControlPolicyResponseBody {
	return s.Body
}

func (s *DescribeControlPolicyResponse) SetHeaders(v map[string]*string) *DescribeControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeControlPolicyResponse) SetStatusCode(v int32) *DescribeControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeControlPolicyResponse) SetBody(v *DescribeControlPolicyResponseBody) *DescribeControlPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
