// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlowLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlowLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlowLogsResponseBody) *DescribeSlowLogsResponse
	GetBody() *DescribeSlowLogsResponseBody
}

type DescribeSlowLogsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlowLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlowLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlowLogsResponse) GetBody() *DescribeSlowLogsResponseBody {
	return s.Body
}

func (s *DescribeSlowLogsResponse) SetHeaders(v map[string]*string) *DescribeSlowLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogsResponse) SetStatusCode(v int32) *DescribeSlowLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowLogsResponse) SetBody(v *DescribeSlowLogsResponseBody) *DescribeSlowLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeSlowLogsResponse) Validate() error {
	return dara.Validate(s)
}
