// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvancedBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvancedBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvancedBackupPolicyResponseBody) *DescribeAdvancedBackupPolicyResponse
	GetBody() *DescribeAdvancedBackupPolicyResponseBody
}

type DescribeAdvancedBackupPolicyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvancedBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvancedBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvancedBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvancedBackupPolicyResponse) GetBody() *DescribeAdvancedBackupPolicyResponseBody {
	return s.Body
}

func (s *DescribeAdvancedBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeAdvancedBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponse) SetStatusCode(v int32) *DescribeAdvancedBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponse) SetBody(v *DescribeAdvancedBackupPolicyResponseBody) *DescribeAdvancedBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvancedBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
