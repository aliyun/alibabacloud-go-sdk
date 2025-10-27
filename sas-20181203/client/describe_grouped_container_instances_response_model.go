// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedContainerInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupedContainerInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupedContainerInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupedContainerInstancesResponseBody) *DescribeGroupedContainerInstancesResponse
	GetBody() *DescribeGroupedContainerInstancesResponseBody
}

type DescribeGroupedContainerInstancesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupedContainerInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupedContainerInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedContainerInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedContainerInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupedContainerInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupedContainerInstancesResponse) GetBody() *DescribeGroupedContainerInstancesResponseBody {
	return s.Body
}

func (s *DescribeGroupedContainerInstancesResponse) SetHeaders(v map[string]*string) *DescribeGroupedContainerInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedContainerInstancesResponse) SetStatusCode(v int32) *DescribeGroupedContainerInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponse) SetBody(v *DescribeGroupedContainerInstancesResponseBody) *DescribeGroupedContainerInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupedContainerInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
