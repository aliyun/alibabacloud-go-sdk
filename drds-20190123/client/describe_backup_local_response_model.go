// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupLocalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupLocalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupLocalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupLocalResponseBody) *DescribeBackupLocalResponse
	GetBody() *DescribeBackupLocalResponseBody
}

type DescribeBackupLocalResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupLocalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupLocalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLocalResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupLocalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupLocalResponse) GetBody() *DescribeBackupLocalResponseBody {
	return s.Body
}

func (s *DescribeBackupLocalResponse) SetHeaders(v map[string]*string) *DescribeBackupLocalResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupLocalResponse) SetStatusCode(v int32) *DescribeBackupLocalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupLocalResponse) SetBody(v *DescribeBackupLocalResponseBody) *DescribeBackupLocalResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupLocalResponse) Validate() error {
	return dara.Validate(s)
}
