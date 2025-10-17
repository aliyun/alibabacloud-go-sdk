// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupLogsResponseBody) *DescribeBackupLogsResponse
	GetBody() *DescribeBackupLogsResponseBody
}

type DescribeBackupLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupLogsResponse) GetBody() *DescribeBackupLogsResponseBody {
	return s.Body
}

func (s *DescribeBackupLogsResponse) SetHeaders(v map[string]*string) *DescribeBackupLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupLogsResponse) SetStatusCode(v int32) *DescribeBackupLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupLogsResponse) SetBody(v *DescribeBackupLogsResponseBody) *DescribeBackupLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
