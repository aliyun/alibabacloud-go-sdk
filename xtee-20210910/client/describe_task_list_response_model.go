// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTaskListResponseBody) *DescribeTaskListResponse
	GetBody() *DescribeTaskListResponseBody
}

type DescribeTaskListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTaskListResponse) GetBody() *DescribeTaskListResponseBody {
	return s.Body
}

func (s *DescribeTaskListResponse) SetHeaders(v map[string]*string) *DescribeTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskListResponse) SetStatusCode(v int32) *DescribeTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskListResponse) SetBody(v *DescribeTaskListResponseBody) *DescribeTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeTaskListResponse) Validate() error {
	return dara.Validate(s)
}
