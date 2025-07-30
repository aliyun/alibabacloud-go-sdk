// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse
	GetBody() *DescribeBackupPolicyResponseBody
}

type DescribeBackupPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupPolicyResponse) GetBody() *DescribeBackupPolicyResponseBody {
	return s.Body
}

func (s *DescribeBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPolicyResponse) SetStatusCode(v int32) *DescribeBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupPolicyResponse) Validate() error {
	return dara.Validate(s)
}
