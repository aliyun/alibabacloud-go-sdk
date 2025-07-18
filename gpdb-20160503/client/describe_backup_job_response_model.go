// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupJobResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupJobResponseBody) *DescribeBackupJobResponse
	GetBody() *DescribeBackupJobResponseBody
}

type DescribeBackupJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupJobResponse) GetBody() *DescribeBackupJobResponseBody {
	return s.Body
}

func (s *DescribeBackupJobResponse) SetHeaders(v map[string]*string) *DescribeBackupJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupJobResponse) SetStatusCode(v int32) *DescribeBackupJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupJobResponse) SetBody(v *DescribeBackupJobResponseBody) *DescribeBackupJobResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupJobResponse) Validate() error {
	return dara.Validate(s)
}
