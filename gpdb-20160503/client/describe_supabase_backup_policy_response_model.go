// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupabaseBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupabaseBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupabaseBackupPolicyResponseBody) *DescribeSupabaseBackupPolicyResponse
	GetBody() *DescribeSupabaseBackupPolicyResponseBody
}

type DescribeSupabaseBackupPolicyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupabaseBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupabaseBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupabaseBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupabaseBackupPolicyResponse) GetBody() *DescribeSupabaseBackupPolicyResponseBody {
	return s.Body
}

func (s *DescribeSupabaseBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeSupabaseBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupabaseBackupPolicyResponse) SetStatusCode(v int32) *DescribeSupabaseBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupabaseBackupPolicyResponse) SetBody(v *DescribeSupabaseBackupPolicyResponseBody) *DescribeSupabaseBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeSupabaseBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
