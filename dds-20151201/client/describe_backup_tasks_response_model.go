// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupTasksResponseBody) *DescribeBackupTasksResponse
	GetBody() *DescribeBackupTasksResponseBody
}

type DescribeBackupTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupTasksResponse) GetBody() *DescribeBackupTasksResponseBody {
	return s.Body
}

func (s *DescribeBackupTasksResponse) SetHeaders(v map[string]*string) *DescribeBackupTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupTasksResponse) SetStatusCode(v int32) *DescribeBackupTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupTasksResponse) SetBody(v *DescribeBackupTasksResponseBody) *DescribeBackupTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupTasksResponse) Validate() error {
	return dara.Validate(s)
}
