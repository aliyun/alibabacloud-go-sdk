// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortConnsCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortConnsCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortConnsCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortConnsCountResponseBody) *DescribePortConnsCountResponse
	GetBody() *DescribePortConnsCountResponseBody
}

type DescribePortConnsCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortConnsCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortConnsCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortConnsCountResponse) GoString() string {
	return s.String()
}

func (s *DescribePortConnsCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortConnsCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortConnsCountResponse) GetBody() *DescribePortConnsCountResponseBody {
	return s.Body
}

func (s *DescribePortConnsCountResponse) SetHeaders(v map[string]*string) *DescribePortConnsCountResponse {
	s.Headers = v
	return s
}

func (s *DescribePortConnsCountResponse) SetStatusCode(v int32) *DescribePortConnsCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortConnsCountResponse) SetBody(v *DescribePortConnsCountResponseBody) *DescribePortConnsCountResponse {
	s.Body = v
	return s
}

func (s *DescribePortConnsCountResponse) Validate() error {
	return dara.Validate(s)
}
