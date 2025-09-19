// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupPoliciesResponseBody) *DescribeBackupPoliciesResponse
	GetBody() *DescribeBackupPoliciesResponseBody
}

type DescribeBackupPoliciesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupPoliciesResponse) GetBody() *DescribeBackupPoliciesResponseBody {
	return s.Body
}

func (s *DescribeBackupPoliciesResponse) SetHeaders(v map[string]*string) *DescribeBackupPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPoliciesResponse) SetStatusCode(v int32) *DescribeBackupPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPoliciesResponse) SetBody(v *DescribeBackupPoliciesResponseBody) *DescribeBackupPoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupPoliciesResponse) Validate() error {
	return dara.Validate(s)
}
