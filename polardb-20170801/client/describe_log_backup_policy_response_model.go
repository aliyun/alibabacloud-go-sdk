// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogBackupPolicyResponseBody) *DescribeLogBackupPolicyResponse
	GetBody() *DescribeLogBackupPolicyResponseBody
}

type DescribeLogBackupPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogBackupPolicyResponse) GetBody() *DescribeLogBackupPolicyResponseBody {
	return s.Body
}

func (s *DescribeLogBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeLogBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogBackupPolicyResponse) SetStatusCode(v int32) *DescribeLogBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogBackupPolicyResponse) SetBody(v *DescribeLogBackupPolicyResponseBody) *DescribeLogBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeLogBackupPolicyResponse) Validate() error {
	return dara.Validate(s)
}
