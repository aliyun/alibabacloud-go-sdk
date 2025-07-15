// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopGroupsResponseBody) *DescribeDesktopGroupsResponse
	GetBody() *DescribeDesktopGroupsResponseBody
}

type DescribeDesktopGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopGroupsResponse) GetBody() *DescribeDesktopGroupsResponseBody {
	return s.Body
}

func (s *DescribeDesktopGroupsResponse) SetHeaders(v map[string]*string) *DescribeDesktopGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopGroupsResponse) SetStatusCode(v int32) *DescribeDesktopGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopGroupsResponse) SetBody(v *DescribeDesktopGroupsResponseBody) *DescribeDesktopGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopGroupsResponse) Validate() error {
	return dara.Validate(s)
}
