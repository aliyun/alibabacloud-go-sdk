// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersInGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUsersInGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUsersInGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUsersInGroupResponseBody) *DescribeUsersInGroupResponse
	GetBody() *DescribeUsersInGroupResponseBody
}

type DescribeUsersInGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsersInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsersInGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersInGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUsersInGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUsersInGroupResponse) GetBody() *DescribeUsersInGroupResponseBody {
	return s.Body
}

func (s *DescribeUsersInGroupResponse) SetHeaders(v map[string]*string) *DescribeUsersInGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersInGroupResponse) SetStatusCode(v int32) *DescribeUsersInGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsersInGroupResponse) SetBody(v *DescribeUsersInGroupResponseBody) *DescribeUsersInGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeUsersInGroupResponse) Validate() error {
	return dara.Validate(s)
}
