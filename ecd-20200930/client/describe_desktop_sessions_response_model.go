// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopSessionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopSessionsResponseBody) *DescribeDesktopSessionsResponse
	GetBody() *DescribeDesktopSessionsResponseBody
}

type DescribeDesktopSessionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopSessionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopSessionsResponse) GetBody() *DescribeDesktopSessionsResponseBody {
	return s.Body
}

func (s *DescribeDesktopSessionsResponse) SetHeaders(v map[string]*string) *DescribeDesktopSessionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopSessionsResponse) SetStatusCode(v int32) *DescribeDesktopSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopSessionsResponse) SetBody(v *DescribeDesktopSessionsResponseBody) *DescribeDesktopSessionsResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopSessionsResponse) Validate() error {
	return dara.Validate(s)
}
