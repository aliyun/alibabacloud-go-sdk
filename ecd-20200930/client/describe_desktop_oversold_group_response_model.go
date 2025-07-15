// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopOversoldGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopOversoldGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopOversoldGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopOversoldGroupResponseBody) *DescribeDesktopOversoldGroupResponse
	GetBody() *DescribeDesktopOversoldGroupResponseBody
}

type DescribeDesktopOversoldGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopOversoldGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopOversoldGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopOversoldGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopOversoldGroupResponse) GetBody() *DescribeDesktopOversoldGroupResponseBody {
	return s.Body
}

func (s *DescribeDesktopOversoldGroupResponse) SetHeaders(v map[string]*string) *DescribeDesktopOversoldGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopOversoldGroupResponse) SetStatusCode(v int32) *DescribeDesktopOversoldGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponse) SetBody(v *DescribeDesktopOversoldGroupResponseBody) *DescribeDesktopOversoldGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopOversoldGroupResponse) Validate() error {
	return dara.Validate(s)
}
