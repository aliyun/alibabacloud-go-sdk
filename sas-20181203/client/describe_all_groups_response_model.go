// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllGroupsResponseBody) *DescribeAllGroupsResponse
	GetBody() *DescribeAllGroupsResponseBody
}

type DescribeAllGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllGroupsResponse) GetBody() *DescribeAllGroupsResponseBody {
	return s.Body
}

func (s *DescribeAllGroupsResponse) SetHeaders(v map[string]*string) *DescribeAllGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllGroupsResponse) SetStatusCode(v int32) *DescribeAllGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllGroupsResponse) SetBody(v *DescribeAllGroupsResponseBody) *DescribeAllGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeAllGroupsResponse) Validate() error {
	return dara.Validate(s)
}
