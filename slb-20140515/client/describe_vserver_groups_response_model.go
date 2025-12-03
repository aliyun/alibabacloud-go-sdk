// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVServerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVServerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVServerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVServerGroupsResponseBody) *DescribeVServerGroupsResponse
	GetBody() *DescribeVServerGroupsResponseBody
}

type DescribeVServerGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVServerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVServerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVServerGroupsResponse) GetBody() *DescribeVServerGroupsResponseBody {
	return s.Body
}

func (s *DescribeVServerGroupsResponse) SetHeaders(v map[string]*string) *DescribeVServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVServerGroupsResponse) SetStatusCode(v int32) *DescribeVServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVServerGroupsResponse) SetBody(v *DescribeVServerGroupsResponseBody) *DescribeVServerGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeVServerGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
