// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetDownloadTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupSetDownloadTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupSetDownloadTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupSetDownloadTaskListResponseBody) *DescribeBackupSetDownloadTaskListResponse
	GetBody() *DescribeBackupSetDownloadTaskListResponseBody
}

type DescribeBackupSetDownloadTaskListResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupSetDownloadTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupSetDownloadTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetDownloadTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetDownloadTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupSetDownloadTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupSetDownloadTaskListResponse) GetBody() *DescribeBackupSetDownloadTaskListResponseBody {
	return s.Body
}

func (s *DescribeBackupSetDownloadTaskListResponse) SetHeaders(v map[string]*string) *DescribeBackupSetDownloadTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponse) SetStatusCode(v int32) *DescribeBackupSetDownloadTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponse) SetBody(v *DescribeBackupSetDownloadTaskListResponseBody) *DescribeBackupSetDownloadTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
