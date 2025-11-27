// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceCrossBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceCrossBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceCrossBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceCrossBackupPolicyResponseBody) *DescribeInstanceCrossBackupPolicyResponse
	GetBody() *DescribeInstanceCrossBackupPolicyResponseBody
}

type DescribeInstanceCrossBackupPolicyResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceCrossBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceCrossBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceCrossBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCrossBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceCrossBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceCrossBackupPolicyResponse) GetBody() *DescribeInstanceCrossBackupPolicyResponseBody {
	return s.Body
}

func (s *DescribeInstanceCrossBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeInstanceCrossBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponse) SetStatusCode(v int32) *DescribeInstanceCrossBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponse) SetBody(v *DescribeInstanceCrossBackupPolicyResponseBody) *DescribeInstanceCrossBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
