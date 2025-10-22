// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodeGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodeGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodeGroupsResponseBody) *DescribeNodeGroupsResponse
	GetBody() *DescribeNodeGroupsResponseBody
}

type DescribeNodeGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodeGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodeGroupsResponse) GetBody() *DescribeNodeGroupsResponseBody {
	return s.Body
}

func (s *DescribeNodeGroupsResponse) SetHeaders(v map[string]*string) *DescribeNodeGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeGroupsResponse) SetStatusCode(v int32) *DescribeNodeGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeGroupsResponse) SetBody(v *DescribeNodeGroupsResponseBody) *DescribeNodeGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeNodeGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
