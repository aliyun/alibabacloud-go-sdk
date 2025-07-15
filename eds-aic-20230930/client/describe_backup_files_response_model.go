// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupFilesResponseBody) *DescribeBackupFilesResponse
	GetBody() *DescribeBackupFilesResponseBody
}

type DescribeBackupFilesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupFilesResponse) GetBody() *DescribeBackupFilesResponseBody {
	return s.Body
}

func (s *DescribeBackupFilesResponse) SetHeaders(v map[string]*string) *DescribeBackupFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupFilesResponse) SetStatusCode(v int32) *DescribeBackupFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupFilesResponse) SetBody(v *DescribeBackupFilesResponseBody) *DescribeBackupFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupFilesResponse) Validate() error {
	return dara.Validate(s)
}
