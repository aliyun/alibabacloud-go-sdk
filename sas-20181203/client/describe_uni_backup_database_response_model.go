// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUniBackupDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUniBackupDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUniBackupDatabaseResponseBody) *DescribeUniBackupDatabaseResponse
	GetBody() *DescribeUniBackupDatabaseResponseBody
}

type DescribeUniBackupDatabaseResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUniBackupDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUniBackupDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUniBackupDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUniBackupDatabaseResponse) GetBody() *DescribeUniBackupDatabaseResponseBody {
	return s.Body
}

func (s *DescribeUniBackupDatabaseResponse) SetHeaders(v map[string]*string) *DescribeUniBackupDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DescribeUniBackupDatabaseResponse) SetStatusCode(v int32) *DescribeUniBackupDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponse) SetBody(v *DescribeUniBackupDatabaseResponseBody) *DescribeUniBackupDatabaseResponse {
	s.Body = v
	return s
}

func (s *DescribeUniBackupDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
