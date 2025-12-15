// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterBackupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterBackupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterBackupListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterBackupListResponseBody) *DescribeClusterBackupListResponse
	GetBody() *DescribeClusterBackupListResponseBody
}

type DescribeClusterBackupListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterBackupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterBackupListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterBackupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterBackupListResponse) GetBody() *DescribeClusterBackupListResponseBody {
	return s.Body
}

func (s *DescribeClusterBackupListResponse) SetHeaders(v map[string]*string) *DescribeClusterBackupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterBackupListResponse) SetStatusCode(v int32) *DescribeClusterBackupListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterBackupListResponse) SetBody(v *DescribeClusterBackupListResponseBody) *DescribeClusterBackupListResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterBackupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
