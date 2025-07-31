// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuditPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuditPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuditPolicyResponseBody) *DescribeAuditPolicyResponse
	GetBody() *DescribeAuditPolicyResponseBody
}

type DescribeAuditPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuditPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuditPolicyResponse) GetBody() *DescribeAuditPolicyResponseBody {
	return s.Body
}

func (s *DescribeAuditPolicyResponse) SetHeaders(v map[string]*string) *DescribeAuditPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditPolicyResponse) SetStatusCode(v int32) *DescribeAuditPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditPolicyResponse) SetBody(v *DescribeAuditPolicyResponseBody) *DescribeAuditPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeAuditPolicyResponse) Validate() error {
	return dara.Validate(s)
}
