// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarmPoolInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWarmPoolInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWarmPoolInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWarmPoolInstancesResponseBody) *DescribeWarmPoolInstancesResponse
	GetBody() *DescribeWarmPoolInstancesResponseBody
}

type DescribeWarmPoolInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWarmPoolInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWarmPoolInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarmPoolInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeWarmPoolInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWarmPoolInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWarmPoolInstancesResponse) GetBody() *DescribeWarmPoolInstancesResponseBody {
	return s.Body
}

func (s *DescribeWarmPoolInstancesResponse) SetHeaders(v map[string]*string) *DescribeWarmPoolInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeWarmPoolInstancesResponse) SetStatusCode(v int32) *DescribeWarmPoolInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponse) SetBody(v *DescribeWarmPoolInstancesResponseBody) *DescribeWarmPoolInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeWarmPoolInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
