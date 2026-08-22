// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenSearchInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenSearchInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenSearchInstancesResponseBody) *DescribeOpenSearchInstancesResponse
	GetBody() *DescribeOpenSearchInstancesResponseBody
}

type DescribeOpenSearchInstancesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenSearchInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenSearchInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenSearchInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenSearchInstancesResponse) GetBody() *DescribeOpenSearchInstancesResponseBody {
	return s.Body
}

func (s *DescribeOpenSearchInstancesResponse) SetHeaders(v map[string]*string) *DescribeOpenSearchInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenSearchInstancesResponse) SetStatusCode(v int32) *DescribeOpenSearchInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponse) SetBody(v *DescribeOpenSearchInstancesResponseBody) *DescribeOpenSearchInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenSearchInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
