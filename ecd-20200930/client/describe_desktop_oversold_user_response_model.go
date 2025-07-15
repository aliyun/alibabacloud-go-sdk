// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopOversoldUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopOversoldUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopOversoldUserResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopOversoldUserResponseBody) *DescribeDesktopOversoldUserResponse
	GetBody() *DescribeDesktopOversoldUserResponseBody
}

type DescribeDesktopOversoldUserResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopOversoldUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopOversoldUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldUserResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopOversoldUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopOversoldUserResponse) GetBody() *DescribeDesktopOversoldUserResponseBody {
	return s.Body
}

func (s *DescribeDesktopOversoldUserResponse) SetHeaders(v map[string]*string) *DescribeDesktopOversoldUserResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopOversoldUserResponse) SetStatusCode(v int32) *DescribeDesktopOversoldUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopOversoldUserResponse) SetBody(v *DescribeDesktopOversoldUserResponseBody) *DescribeDesktopOversoldUserResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopOversoldUserResponse) Validate() error {
	return dara.Validate(s)
}
