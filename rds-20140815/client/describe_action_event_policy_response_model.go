// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActionEventPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActionEventPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActionEventPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActionEventPolicyResponseBody) *DescribeActionEventPolicyResponse
	GetBody() *DescribeActionEventPolicyResponseBody
}

type DescribeActionEventPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActionEventPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActionEventPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActionEventPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeActionEventPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActionEventPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActionEventPolicyResponse) GetBody() *DescribeActionEventPolicyResponseBody {
	return s.Body
}

func (s *DescribeActionEventPolicyResponse) SetHeaders(v map[string]*string) *DescribeActionEventPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeActionEventPolicyResponse) SetStatusCode(v int32) *DescribeActionEventPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActionEventPolicyResponse) SetBody(v *DescribeActionEventPolicyResponseBody) *DescribeActionEventPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeActionEventPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
