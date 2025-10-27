// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupPolicyDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUniBackupPolicyDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUniBackupPolicyDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUniBackupPolicyDetailResponseBody) *DescribeUniBackupPolicyDetailResponse
	GetBody() *DescribeUniBackupPolicyDetailResponseBody
}

type DescribeUniBackupPolicyDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUniBackupPolicyDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUniBackupPolicyDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPolicyDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPolicyDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUniBackupPolicyDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUniBackupPolicyDetailResponse) GetBody() *DescribeUniBackupPolicyDetailResponseBody {
	return s.Body
}

func (s *DescribeUniBackupPolicyDetailResponse) SetHeaders(v map[string]*string) *DescribeUniBackupPolicyDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponse) SetStatusCode(v int32) *DescribeUniBackupPolicyDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponse) SetBody(v *DescribeUniBackupPolicyDetailResponseBody) *DescribeUniBackupPolicyDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
