// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRefreshTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodRefreshTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodRefreshTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodRefreshTasksResponseBody) *DescribeVodRefreshTasksResponse
	GetBody() *DescribeVodRefreshTasksResponseBody
}

type DescribeVodRefreshTasksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodRefreshTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodRefreshTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRefreshTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodRefreshTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodRefreshTasksResponse) GetBody() *DescribeVodRefreshTasksResponseBody {
	return s.Body
}

func (s *DescribeVodRefreshTasksResponse) SetHeaders(v map[string]*string) *DescribeVodRefreshTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodRefreshTasksResponse) SetStatusCode(v int32) *DescribeVodRefreshTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodRefreshTasksResponse) SetBody(v *DescribeVodRefreshTasksResponseBody) *DescribeVodRefreshTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeVodRefreshTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
