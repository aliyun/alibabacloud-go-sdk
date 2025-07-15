// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopGroupSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopGroupSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopGroupSessionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopGroupSessionsResponseBody) *DescribeDesktopGroupSessionsResponse
	GetBody() *DescribeDesktopGroupSessionsResponseBody
}

type DescribeDesktopGroupSessionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopGroupSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopGroupSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupSessionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopGroupSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopGroupSessionsResponse) GetBody() *DescribeDesktopGroupSessionsResponseBody {
	return s.Body
}

func (s *DescribeDesktopGroupSessionsResponse) SetHeaders(v map[string]*string) *DescribeDesktopGroupSessionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopGroupSessionsResponse) SetStatusCode(v int32) *DescribeDesktopGroupSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopGroupSessionsResponse) SetBody(v *DescribeDesktopGroupSessionsResponseBody) *DescribeDesktopGroupSessionsResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopGroupSessionsResponse) Validate() error {
	return dara.Validate(s)
}
