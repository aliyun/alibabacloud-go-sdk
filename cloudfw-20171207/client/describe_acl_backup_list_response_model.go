// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclBackupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAclBackupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAclBackupListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAclBackupListResponseBody) *DescribeAclBackupListResponse
	GetBody() *DescribeAclBackupListResponseBody
}

type DescribeAclBackupListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclBackupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclBackupListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclBackupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclBackupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAclBackupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAclBackupListResponse) GetBody() *DescribeAclBackupListResponseBody {
	return s.Body
}

func (s *DescribeAclBackupListResponse) SetHeaders(v map[string]*string) *DescribeAclBackupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclBackupListResponse) SetStatusCode(v int32) *DescribeAclBackupListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclBackupListResponse) SetBody(v *DescribeAclBackupListResponseBody) *DescribeAclBackupListResponse {
	s.Body = v
	return s
}

func (s *DescribeAclBackupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
