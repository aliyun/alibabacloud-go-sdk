// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProcessTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProcessTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProcessTasksResponseBody) *DescribeProcessTasksResponse
	GetBody() *DescribeProcessTasksResponseBody
}

type DescribeProcessTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProcessTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProcessTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProcessTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProcessTasksResponse) GetBody() *DescribeProcessTasksResponseBody {
	return s.Body
}

func (s *DescribeProcessTasksResponse) SetHeaders(v map[string]*string) *DescribeProcessTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessTasksResponse) SetStatusCode(v int32) *DescribeProcessTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProcessTasksResponse) SetBody(v *DescribeProcessTasksResponseBody) *DescribeProcessTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeProcessTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
