// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupsResponseBody) *DescribeGroupsResponse
	GetBody() *DescribeGroupsResponseBody
}

type DescribeGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupsResponse) GetBody() *DescribeGroupsResponseBody {
	return s.Body
}

func (s *DescribeGroupsResponse) SetHeaders(v map[string]*string) *DescribeGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupsResponse) SetStatusCode(v int32) *DescribeGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupsResponse) SetBody(v *DescribeGroupsResponseBody) *DescribeGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
