// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUniBackupPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUniBackupPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUniBackupPoliciesResponseBody) *DescribeUniBackupPoliciesResponse
	GetBody() *DescribeUniBackupPoliciesResponseBody
}

type DescribeUniBackupPoliciesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUniBackupPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUniBackupPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUniBackupPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUniBackupPoliciesResponse) GetBody() *DescribeUniBackupPoliciesResponseBody {
	return s.Body
}

func (s *DescribeUniBackupPoliciesResponse) SetHeaders(v map[string]*string) *DescribeUniBackupPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUniBackupPoliciesResponse) SetStatusCode(v int32) *DescribeUniBackupPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponse) SetBody(v *DescribeUniBackupPoliciesResponseBody) *DescribeUniBackupPoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeUniBackupPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
