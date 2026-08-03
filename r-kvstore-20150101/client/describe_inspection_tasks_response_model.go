// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInspectionTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInspectionTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInspectionTasksResponseBody) *DescribeInspectionTasksResponse
	GetBody() *DescribeInspectionTasksResponseBody
}

type DescribeInspectionTasksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInspectionTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInspectionTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInspectionTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInspectionTasksResponse) GetBody() *DescribeInspectionTasksResponseBody {
	return s.Body
}

func (s *DescribeInspectionTasksResponse) SetHeaders(v map[string]*string) *DescribeInspectionTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeInspectionTasksResponse) SetStatusCode(v int32) *DescribeInspectionTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInspectionTasksResponse) SetBody(v *DescribeInspectionTasksResponseBody) *DescribeInspectionTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeInspectionTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
