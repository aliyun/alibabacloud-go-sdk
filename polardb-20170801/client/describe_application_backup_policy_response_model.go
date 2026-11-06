// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationBackupPolicyResponseBody) *DescribeApplicationBackupPolicyResponse
	GetBody() *DescribeApplicationBackupPolicyResponseBody
}

type DescribeApplicationBackupPolicyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationBackupPolicyResponse) GetBody() *DescribeApplicationBackupPolicyResponseBody {
	return s.Body
}

func (s *DescribeApplicationBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeApplicationBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationBackupPolicyResponse) SetStatusCode(v int32) *DescribeApplicationBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationBackupPolicyResponse) SetBody(v *DescribeApplicationBackupPolicyResponseBody) *DescribeApplicationBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
