// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaBackupPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHanaBackupPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHanaBackupPlansResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHanaBackupPlansResponseBody) *DescribeHanaBackupPlansResponse
	GetBody() *DescribeHanaBackupPlansResponseBody
}

type DescribeHanaBackupPlansResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHanaBackupPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHanaBackupPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHanaBackupPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHanaBackupPlansResponse) GetBody() *DescribeHanaBackupPlansResponseBody {
	return s.Body
}

func (s *DescribeHanaBackupPlansResponse) SetHeaders(v map[string]*string) *DescribeHanaBackupPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaBackupPlansResponse) SetStatusCode(v int32) *DescribeHanaBackupPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaBackupPlansResponse) SetBody(v *DescribeHanaBackupPlansResponseBody) *DescribeHanaBackupPlansResponse {
	s.Body = v
	return s
}

func (s *DescribeHanaBackupPlansResponse) Validate() error {
	return dara.Validate(s)
}
