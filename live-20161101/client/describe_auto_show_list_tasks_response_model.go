// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoShowListTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoShowListTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoShowListTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoShowListTasksResponseBody) *DescribeAutoShowListTasksResponse
	GetBody() *DescribeAutoShowListTasksResponseBody
}

type DescribeAutoShowListTasksResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoShowListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoShowListTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoShowListTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoShowListTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoShowListTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoShowListTasksResponse) GetBody() *DescribeAutoShowListTasksResponseBody {
	return s.Body
}

func (s *DescribeAutoShowListTasksResponse) SetHeaders(v map[string]*string) *DescribeAutoShowListTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoShowListTasksResponse) SetStatusCode(v int32) *DescribeAutoShowListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoShowListTasksResponse) SetBody(v *DescribeAutoShowListTasksResponseBody) *DescribeAutoShowListTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoShowListTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
