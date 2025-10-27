// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomBlockInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomBlockInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomBlockInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomBlockInstancesResponseBody) *DescribeCustomBlockInstancesResponse
	GetBody() *DescribeCustomBlockInstancesResponseBody
}

type DescribeCustomBlockInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomBlockInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomBlockInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomBlockInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomBlockInstancesResponse) GetBody() *DescribeCustomBlockInstancesResponseBody {
	return s.Body
}

func (s *DescribeCustomBlockInstancesResponse) SetHeaders(v map[string]*string) *DescribeCustomBlockInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomBlockInstancesResponse) SetStatusCode(v int32) *DescribeCustomBlockInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponse) SetBody(v *DescribeCustomBlockInstancesResponseBody) *DescribeCustomBlockInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomBlockInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
