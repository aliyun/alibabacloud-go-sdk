// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHistoryTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHistoryTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHistoryTasksResponseBody) *DescribeHistoryTasksResponse
	GetBody() *DescribeHistoryTasksResponseBody
}

type DescribeHistoryTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHistoryTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHistoryTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHistoryTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHistoryTasksResponse) GetBody() *DescribeHistoryTasksResponseBody {
	return s.Body
}

func (s *DescribeHistoryTasksResponse) SetHeaders(v map[string]*string) *DescribeHistoryTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeHistoryTasksResponse) SetStatusCode(v int32) *DescribeHistoryTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHistoryTasksResponse) SetBody(v *DescribeHistoryTasksResponseBody) *DescribeHistoryTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeHistoryTasksResponse) Validate() error {
	return dara.Validate(s)
}
