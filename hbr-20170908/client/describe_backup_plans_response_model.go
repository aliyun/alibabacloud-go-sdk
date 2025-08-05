// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupPlansResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupPlansResponseBody) *DescribeBackupPlansResponse
	GetBody() *DescribeBackupPlansResponseBody
}

type DescribeBackupPlansResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupPlansResponse) GetBody() *DescribeBackupPlansResponseBody {
	return s.Body
}

func (s *DescribeBackupPlansResponse) SetHeaders(v map[string]*string) *DescribeBackupPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlansResponse) SetStatusCode(v int32) *DescribeBackupPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPlansResponse) SetBody(v *DescribeBackupPlansResponseBody) *DescribeBackupPlansResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupPlansResponse) Validate() error {
	return dara.Validate(s)
}
