// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortMaxConnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortMaxConnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortMaxConnsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortMaxConnsResponseBody) *DescribePortMaxConnsResponse
	GetBody() *DescribePortMaxConnsResponseBody
}

type DescribePortMaxConnsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortMaxConnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortMaxConnsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortMaxConnsResponse) GoString() string {
	return s.String()
}

func (s *DescribePortMaxConnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortMaxConnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortMaxConnsResponse) GetBody() *DescribePortMaxConnsResponseBody {
	return s.Body
}

func (s *DescribePortMaxConnsResponse) SetHeaders(v map[string]*string) *DescribePortMaxConnsResponse {
	s.Headers = v
	return s
}

func (s *DescribePortMaxConnsResponse) SetStatusCode(v int32) *DescribePortMaxConnsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortMaxConnsResponse) SetBody(v *DescribePortMaxConnsResponseBody) *DescribePortMaxConnsResponse {
	s.Body = v
	return s
}

func (s *DescribePortMaxConnsResponse) Validate() error {
	return dara.Validate(s)
}
