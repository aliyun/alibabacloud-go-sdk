// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlowLogRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlowLogRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlowLogRecordsResponseBody) *DescribeSlowLogRecordsResponse
	GetBody() *DescribeSlowLogRecordsResponseBody
}

type DescribeSlowLogRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlowLogRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlowLogRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlowLogRecordsResponse) GetBody() *DescribeSlowLogRecordsResponseBody {
	return s.Body
}

func (s *DescribeSlowLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeSlowLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetStatusCode(v int32) *DescribeSlowLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetBody(v *DescribeSlowLogRecordsResponseBody) *DescribeSlowLogRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeSlowLogRecordsResponse) Validate() error {
	return dara.Validate(s)
}
