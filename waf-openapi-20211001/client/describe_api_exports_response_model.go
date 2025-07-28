// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiExportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiExportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiExportsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiExportsResponseBody) *DescribeApiExportsResponse
	GetBody() *DescribeApiExportsResponseBody
}

type DescribeApiExportsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiExportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiExportsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiExportsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiExportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiExportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiExportsResponse) GetBody() *DescribeApiExportsResponseBody {
	return s.Body
}

func (s *DescribeApiExportsResponse) SetHeaders(v map[string]*string) *DescribeApiExportsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiExportsResponse) SetStatusCode(v int32) *DescribeApiExportsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiExportsResponse) SetBody(v *DescribeApiExportsResponseBody) *DescribeApiExportsResponse {
	s.Body = v
	return s
}

func (s *DescribeApiExportsResponse) Validate() error {
	return dara.Validate(s)
}
