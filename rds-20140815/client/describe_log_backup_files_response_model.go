// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogBackupFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogBackupFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogBackupFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogBackupFilesResponseBody) *DescribeLogBackupFilesResponse
	GetBody() *DescribeLogBackupFilesResponseBody
}

type DescribeLogBackupFilesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogBackupFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogBackupFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogBackupFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogBackupFilesResponse) GetBody() *DescribeLogBackupFilesResponseBody {
	return s.Body
}

func (s *DescribeLogBackupFilesResponse) SetHeaders(v map[string]*string) *DescribeLogBackupFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogBackupFilesResponse) SetStatusCode(v int32) *DescribeLogBackupFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogBackupFilesResponse) SetBody(v *DescribeLogBackupFilesResponseBody) *DescribeLogBackupFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeLogBackupFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
