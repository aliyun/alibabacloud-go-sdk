// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIDBClusterTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIDBClusterTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIDBClusterTasksResponseBody) *DescribeAIDBClusterTasksResponse
	GetBody() *DescribeAIDBClusterTasksResponseBody
}

type DescribeAIDBClusterTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIDBClusterTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIDBClusterTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIDBClusterTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIDBClusterTasksResponse) GetBody() *DescribeAIDBClusterTasksResponseBody {
	return s.Body
}

func (s *DescribeAIDBClusterTasksResponse) SetHeaders(v map[string]*string) *DescribeAIDBClusterTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIDBClusterTasksResponse) SetStatusCode(v int32) *DescribeAIDBClusterTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponse) SetBody(v *DescribeAIDBClusterTasksResponseBody) *DescribeAIDBClusterTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeAIDBClusterTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
