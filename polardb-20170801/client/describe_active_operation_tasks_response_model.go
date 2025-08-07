// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActiveOperationTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActiveOperationTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActiveOperationTasksResponseBody) *DescribeActiveOperationTasksResponse
	GetBody() *DescribeActiveOperationTasksResponseBody
}

type DescribeActiveOperationTasksResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActiveOperationTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActiveOperationTasksResponse) GetBody() *DescribeActiveOperationTasksResponseBody {
	return s.Body
}

func (s *DescribeActiveOperationTasksResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTasksResponse) SetStatusCode(v int32) *DescribeActiveOperationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationTasksResponse) SetBody(v *DescribeActiveOperationTasksResponseBody) *DescribeActiveOperationTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeActiveOperationTasksResponse) Validate() error {
	return dara.Validate(s)
}
