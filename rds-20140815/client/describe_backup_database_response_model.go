// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupDatabaseResponseBody) *DescribeBackupDatabaseResponse
	GetBody() *DescribeBackupDatabaseResponseBody
}

type DescribeBackupDatabaseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupDatabaseResponse) GetBody() *DescribeBackupDatabaseResponseBody {
	return s.Body
}

func (s *DescribeBackupDatabaseResponse) SetHeaders(v map[string]*string) *DescribeBackupDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupDatabaseResponse) SetStatusCode(v int32) *DescribeBackupDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupDatabaseResponse) SetBody(v *DescribeBackupDatabaseResponseBody) *DescribeBackupDatabaseResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
