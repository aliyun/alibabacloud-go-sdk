// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopOversoldUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopOversoldUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopOversoldUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopOversoldUserGroupResponseBody) *DescribeDesktopOversoldUserGroupResponse
	GetBody() *DescribeDesktopOversoldUserGroupResponseBody
}

type DescribeDesktopOversoldUserGroupResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopOversoldUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopOversoldUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopOversoldUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopOversoldUserGroupResponse) GetBody() *DescribeDesktopOversoldUserGroupResponseBody {
	return s.Body
}

func (s *DescribeDesktopOversoldUserGroupResponse) SetHeaders(v map[string]*string) *DescribeDesktopOversoldUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponse) SetStatusCode(v int32) *DescribeDesktopOversoldUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponse) SetBody(v *DescribeDesktopOversoldUserGroupResponseBody) *DescribeDesktopOversoldUserGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopOversoldUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
