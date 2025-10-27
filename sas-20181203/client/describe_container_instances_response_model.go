// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerInstancesResponseBody) *DescribeContainerInstancesResponse
	GetBody() *DescribeContainerInstancesResponseBody
}

type DescribeContainerInstancesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerInstancesResponse) GetBody() *DescribeContainerInstancesResponseBody {
	return s.Body
}

func (s *DescribeContainerInstancesResponse) SetHeaders(v map[string]*string) *DescribeContainerInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerInstancesResponse) SetStatusCode(v int32) *DescribeContainerInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerInstancesResponse) SetBody(v *DescribeContainerInstancesResponseBody) *DescribeContainerInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
