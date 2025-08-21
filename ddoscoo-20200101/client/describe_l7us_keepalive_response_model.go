// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL7UsKeepaliveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeL7UsKeepaliveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeL7UsKeepaliveResponse
	GetStatusCode() *int32
	SetBody(v *DescribeL7UsKeepaliveResponseBody) *DescribeL7UsKeepaliveResponse
	GetBody() *DescribeL7UsKeepaliveResponseBody
}

type DescribeL7UsKeepaliveResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeL7UsKeepaliveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeL7UsKeepaliveResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7UsKeepaliveResponse) GoString() string {
	return s.String()
}

func (s *DescribeL7UsKeepaliveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeL7UsKeepaliveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeL7UsKeepaliveResponse) GetBody() *DescribeL7UsKeepaliveResponseBody {
	return s.Body
}

func (s *DescribeL7UsKeepaliveResponse) SetHeaders(v map[string]*string) *DescribeL7UsKeepaliveResponse {
	s.Headers = v
	return s
}

func (s *DescribeL7UsKeepaliveResponse) SetStatusCode(v int32) *DescribeL7UsKeepaliveResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeL7UsKeepaliveResponse) SetBody(v *DescribeL7UsKeepaliveResponseBody) *DescribeL7UsKeepaliveResponse {
	s.Body = v
	return s
}

func (s *DescribeL7UsKeepaliveResponse) Validate() error {
	return dara.Validate(s)
}
