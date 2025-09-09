// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTimesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupTimesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupTimesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupTimesResponseBody) *DescribeBackupTimesResponse
	GetBody() *DescribeBackupTimesResponseBody
}

type DescribeBackupTimesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupTimesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupTimesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTimesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupTimesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupTimesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupTimesResponse) GetBody() *DescribeBackupTimesResponseBody {
	return s.Body
}

func (s *DescribeBackupTimesResponse) SetHeaders(v map[string]*string) *DescribeBackupTimesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupTimesResponse) SetStatusCode(v int32) *DescribeBackupTimesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupTimesResponse) SetBody(v *DescribeBackupTimesResponseBody) *DescribeBackupTimesResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupTimesResponse) Validate() error {
	return dara.Validate(s)
}
