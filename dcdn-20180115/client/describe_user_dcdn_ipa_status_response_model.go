// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDcdnIpaStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserDcdnIpaStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserDcdnIpaStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserDcdnIpaStatusResponseBody) *DescribeUserDcdnIpaStatusResponse
	GetBody() *DescribeUserDcdnIpaStatusResponseBody
}

type DescribeUserDcdnIpaStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserDcdnIpaStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserDcdnIpaStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDcdnIpaStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnIpaStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserDcdnIpaStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserDcdnIpaStatusResponse) GetBody() *DescribeUserDcdnIpaStatusResponseBody {
	return s.Body
}

func (s *DescribeUserDcdnIpaStatusResponse) SetHeaders(v map[string]*string) *DescribeUserDcdnIpaStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponse) SetStatusCode(v int32) *DescribeUserDcdnIpaStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponse) SetBody(v *DescribeUserDcdnIpaStatusResponseBody) *DescribeUserDcdnIpaStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponse) Validate() error {
	return dara.Validate(s)
}
