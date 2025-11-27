// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCInstancesResponseBody) *DescribeRCInstancesResponse
	GetBody() *DescribeRCInstancesResponseBody
}

type DescribeRCInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCInstancesResponse) GetBody() *DescribeRCInstancesResponseBody {
	return s.Body
}

func (s *DescribeRCInstancesResponse) SetHeaders(v map[string]*string) *DescribeRCInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCInstancesResponse) SetStatusCode(v int32) *DescribeRCInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCInstancesResponse) SetBody(v *DescribeRCInstancesResponseBody) *DescribeRCInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeRCInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
