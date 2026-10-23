// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupLogDownloadLinksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupLogDownloadLinksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupLogDownloadLinksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupLogDownloadLinksResponseBody) *DescribeBackupLogDownloadLinksResponse
	GetBody() *DescribeBackupLogDownloadLinksResponseBody
}

type DescribeBackupLogDownloadLinksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupLogDownloadLinksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupLogDownloadLinksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLogDownloadLinksResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogDownloadLinksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupLogDownloadLinksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupLogDownloadLinksResponse) GetBody() *DescribeBackupLogDownloadLinksResponseBody {
	return s.Body
}

func (s *DescribeBackupLogDownloadLinksResponse) SetHeaders(v map[string]*string) *DescribeBackupLogDownloadLinksResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponse) SetStatusCode(v int32) *DescribeBackupLogDownloadLinksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponse) SetBody(v *DescribeBackupLogDownloadLinksResponseBody) *DescribeBackupLogDownloadLinksResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
