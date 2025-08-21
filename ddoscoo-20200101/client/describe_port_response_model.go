// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortResponseBody) *DescribePortResponse
	GetBody() *DescribePortResponseBody
}

type DescribePortResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortResponse) GoString() string {
	return s.String()
}

func (s *DescribePortResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortResponse) GetBody() *DescribePortResponseBody {
	return s.Body
}

func (s *DescribePortResponse) SetHeaders(v map[string]*string) *DescribePortResponse {
	s.Headers = v
	return s
}

func (s *DescribePortResponse) SetStatusCode(v int32) *DescribePortResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortResponse) SetBody(v *DescribePortResponseBody) *DescribePortResponse {
	s.Body = v
	return s
}

func (s *DescribePortResponse) Validate() error {
	return dara.Validate(s)
}
