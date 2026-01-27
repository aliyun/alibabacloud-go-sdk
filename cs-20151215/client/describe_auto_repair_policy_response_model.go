// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoRepairPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoRepairPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoRepairPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoRepairPolicyResponseBody) *DescribeAutoRepairPolicyResponse
	GetBody() *DescribeAutoRepairPolicyResponseBody
}

type DescribeAutoRepairPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoRepairPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoRepairPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoRepairPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoRepairPolicyResponse) GetBody() *DescribeAutoRepairPolicyResponseBody {
	return s.Body
}

func (s *DescribeAutoRepairPolicyResponse) SetHeaders(v map[string]*string) *DescribeAutoRepairPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoRepairPolicyResponse) SetStatusCode(v int32) *DescribeAutoRepairPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponse) SetBody(v *DescribeAutoRepairPolicyResponseBody) *DescribeAutoRepairPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoRepairPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
