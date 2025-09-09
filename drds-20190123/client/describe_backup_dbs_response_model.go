// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDbsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupDbsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupDbsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupDbsResponseBody) *DescribeBackupDbsResponse
	GetBody() *DescribeBackupDbsResponseBody
}

type DescribeBackupDbsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupDbsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupDbsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupDbsResponse) GetBody() *DescribeBackupDbsResponseBody {
	return s.Body
}

func (s *DescribeBackupDbsResponse) SetHeaders(v map[string]*string) *DescribeBackupDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupDbsResponse) SetStatusCode(v int32) *DescribeBackupDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupDbsResponse) SetBody(v *DescribeBackupDbsResponseBody) *DescribeBackupDbsResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupDbsResponse) Validate() error {
	return dara.Validate(s)
}
