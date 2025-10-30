// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataShareInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataShareInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataShareInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataShareInstancesResponseBody) *DescribeDataShareInstancesResponse
	GetBody() *DescribeDataShareInstancesResponseBody
}

type DescribeDataShareInstancesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataShareInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataShareInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataShareInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataShareInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataShareInstancesResponse) GetBody() *DescribeDataShareInstancesResponseBody {
	return s.Body
}

func (s *DescribeDataShareInstancesResponse) SetHeaders(v map[string]*string) *DescribeDataShareInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataShareInstancesResponse) SetStatusCode(v int32) *DescribeDataShareInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataShareInstancesResponse) SetBody(v *DescribeDataShareInstancesResponseBody) *DescribeDataShareInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeDataShareInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
