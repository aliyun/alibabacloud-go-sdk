// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComfyTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComfyTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComfyTasksResponseBody) *DescribeComfyTasksResponse
	GetBody() *DescribeComfyTasksResponseBody
}

type DescribeComfyTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComfyTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComfyTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeComfyTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComfyTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComfyTasksResponse) GetBody() *DescribeComfyTasksResponseBody {
	return s.Body
}

func (s *DescribeComfyTasksResponse) SetHeaders(v map[string]*string) *DescribeComfyTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeComfyTasksResponse) SetStatusCode(v int32) *DescribeComfyTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComfyTasksResponse) SetBody(v *DescribeComfyTasksResponseBody) *DescribeComfyTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeComfyTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
