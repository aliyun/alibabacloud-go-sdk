// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortViewSourceIspsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortViewSourceIspsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortViewSourceIspsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortViewSourceIspsResponseBody) *DescribePortViewSourceIspsResponse
	GetBody() *DescribePortViewSourceIspsResponseBody
}

type DescribePortViewSourceIspsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortViewSourceIspsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortViewSourceIspsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceIspsResponse) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceIspsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortViewSourceIspsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortViewSourceIspsResponse) GetBody() *DescribePortViewSourceIspsResponseBody {
	return s.Body
}

func (s *DescribePortViewSourceIspsResponse) SetHeaders(v map[string]*string) *DescribePortViewSourceIspsResponse {
	s.Headers = v
	return s
}

func (s *DescribePortViewSourceIspsResponse) SetStatusCode(v int32) *DescribePortViewSourceIspsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortViewSourceIspsResponse) SetBody(v *DescribePortViewSourceIspsResponseBody) *DescribePortViewSourceIspsResponse {
	s.Body = v
	return s
}

func (s *DescribePortViewSourceIspsResponse) Validate() error {
	return dara.Validate(s)
}
