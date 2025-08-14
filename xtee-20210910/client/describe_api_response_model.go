// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiResponseBody) *DescribeApiResponse
	GetBody() *DescribeApiResponseBody
}

type DescribeApiResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiResponse) GetBody() *DescribeApiResponseBody {
	return s.Body
}

func (s *DescribeApiResponse) SetHeaders(v map[string]*string) *DescribeApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiResponse) SetStatusCode(v int32) *DescribeApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiResponse) SetBody(v *DescribeApiResponseBody) *DescribeApiResponse {
	s.Body = v
	return s
}

func (s *DescribeApiResponse) Validate() error {
	return dara.Validate(s)
}
