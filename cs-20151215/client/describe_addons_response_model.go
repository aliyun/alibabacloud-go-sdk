// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAddonsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAddonsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAddonsResponseBody) *DescribeAddonsResponse
	GetBody() *DescribeAddonsResponseBody
}

type DescribeAddonsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddonsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAddonsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAddonsResponse) GetBody() *DescribeAddonsResponseBody {
	return s.Body
}

func (s *DescribeAddonsResponse) SetHeaders(v map[string]*string) *DescribeAddonsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddonsResponse) SetStatusCode(v int32) *DescribeAddonsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddonsResponse) SetBody(v *DescribeAddonsResponseBody) *DescribeAddonsResponse {
	s.Body = v
	return s
}

func (s *DescribeAddonsResponse) Validate() error {
	return dara.Validate(s)
}
