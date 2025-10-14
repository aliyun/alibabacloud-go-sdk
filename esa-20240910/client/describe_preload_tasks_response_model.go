// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreloadTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePreloadTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePreloadTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribePreloadTasksResponseBody) *DescribePreloadTasksResponse
	GetBody() *DescribePreloadTasksResponseBody
}

type DescribePreloadTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePreloadTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePreloadTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePreloadTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribePreloadTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePreloadTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePreloadTasksResponse) GetBody() *DescribePreloadTasksResponseBody {
	return s.Body
}

func (s *DescribePreloadTasksResponse) SetHeaders(v map[string]*string) *DescribePreloadTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribePreloadTasksResponse) SetStatusCode(v int32) *DescribePreloadTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePreloadTasksResponse) SetBody(v *DescribePreloadTasksResponseBody) *DescribePreloadTasksResponse {
	s.Body = v
	return s
}

func (s *DescribePreloadTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
