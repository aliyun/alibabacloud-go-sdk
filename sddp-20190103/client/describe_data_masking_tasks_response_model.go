// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataMaskingTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataMaskingTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataMaskingTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataMaskingTasksResponseBody) *DescribeDataMaskingTasksResponse
	GetBody() *DescribeDataMaskingTasksResponseBody
}

type DescribeDataMaskingTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataMaskingTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataMaskingTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataMaskingTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataMaskingTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataMaskingTasksResponse) GetBody() *DescribeDataMaskingTasksResponseBody {
	return s.Body
}

func (s *DescribeDataMaskingTasksResponse) SetHeaders(v map[string]*string) *DescribeDataMaskingTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataMaskingTasksResponse) SetStatusCode(v int32) *DescribeDataMaskingTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataMaskingTasksResponse) SetBody(v *DescribeDataMaskingTasksResponseBody) *DescribeDataMaskingTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeDataMaskingTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
