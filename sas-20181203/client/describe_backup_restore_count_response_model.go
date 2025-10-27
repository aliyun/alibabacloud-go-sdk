// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupRestoreCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupRestoreCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupRestoreCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupRestoreCountResponseBody) *DescribeBackupRestoreCountResponse
	GetBody() *DescribeBackupRestoreCountResponseBody
}

type DescribeBackupRestoreCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupRestoreCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupRestoreCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupRestoreCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupRestoreCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupRestoreCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupRestoreCountResponse) GetBody() *DescribeBackupRestoreCountResponseBody {
	return s.Body
}

func (s *DescribeBackupRestoreCountResponse) SetHeaders(v map[string]*string) *DescribeBackupRestoreCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupRestoreCountResponse) SetStatusCode(v int32) *DescribeBackupRestoreCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupRestoreCountResponse) SetBody(v *DescribeBackupRestoreCountResponseBody) *DescribeBackupRestoreCountResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupRestoreCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
