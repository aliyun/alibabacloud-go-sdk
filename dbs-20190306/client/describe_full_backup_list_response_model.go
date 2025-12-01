// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFullBackupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFullBackupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFullBackupListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFullBackupListResponseBody) *DescribeFullBackupListResponse
	GetBody() *DescribeFullBackupListResponseBody
}

type DescribeFullBackupListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFullBackupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFullBackupListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFullBackupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFullBackupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFullBackupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFullBackupListResponse) GetBody() *DescribeFullBackupListResponseBody {
	return s.Body
}

func (s *DescribeFullBackupListResponse) SetHeaders(v map[string]*string) *DescribeFullBackupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFullBackupListResponse) SetStatusCode(v int32) *DescribeFullBackupListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFullBackupListResponse) SetBody(v *DescribeFullBackupListResponseBody) *DescribeFullBackupListResponse {
	s.Body = v
	return s
}

func (s *DescribeFullBackupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
