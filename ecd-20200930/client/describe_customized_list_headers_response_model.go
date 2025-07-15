// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedListHeadersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomizedListHeadersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomizedListHeadersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomizedListHeadersResponseBody) *DescribeCustomizedListHeadersResponse
	GetBody() *DescribeCustomizedListHeadersResponseBody
}

type DescribeCustomizedListHeadersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizedListHeadersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizedListHeadersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedListHeadersResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedListHeadersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomizedListHeadersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomizedListHeadersResponse) GetBody() *DescribeCustomizedListHeadersResponseBody {
	return s.Body
}

func (s *DescribeCustomizedListHeadersResponse) SetHeaders(v map[string]*string) *DescribeCustomizedListHeadersResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizedListHeadersResponse) SetStatusCode(v int32) *DescribeCustomizedListHeadersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizedListHeadersResponse) SetBody(v *DescribeCustomizedListHeadersResponseBody) *DescribeCustomizedListHeadersResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomizedListHeadersResponse) Validate() error {
	return dara.Validate(s)
}
