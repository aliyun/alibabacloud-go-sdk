// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupTablesResponseBody) *DescribeBackupTablesResponse
	GetBody() *DescribeBackupTablesResponseBody
}

type DescribeBackupTablesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupTablesResponse) GetBody() *DescribeBackupTablesResponseBody {
	return s.Body
}

func (s *DescribeBackupTablesResponse) SetHeaders(v map[string]*string) *DescribeBackupTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupTablesResponse) SetStatusCode(v int32) *DescribeBackupTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupTablesResponse) SetBody(v *DescribeBackupTablesResponseBody) *DescribeBackupTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
