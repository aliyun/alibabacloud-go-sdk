// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDataListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupDataListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupDataListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupDataListResponseBody) *DescribeBackupDataListResponse
	GetBody() *DescribeBackupDataListResponseBody
}

type DescribeBackupDataListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupDataListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupDataListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDataListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupDataListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupDataListResponse) GetBody() *DescribeBackupDataListResponseBody {
	return s.Body
}

func (s *DescribeBackupDataListResponse) SetHeaders(v map[string]*string) *DescribeBackupDataListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupDataListResponse) SetStatusCode(v int32) *DescribeBackupDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupDataListResponse) SetBody(v *DescribeBackupDataListResponseBody) *DescribeBackupDataListResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupDataListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
