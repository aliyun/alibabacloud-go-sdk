// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQueryConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQueryConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQueryConfigsResponseBody) *DescribeQueryConfigsResponse
	GetBody() *DescribeQueryConfigsResponseBody
}

type DescribeQueryConfigsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQueryConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQueryConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeQueryConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQueryConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQueryConfigsResponse) GetBody() *DescribeQueryConfigsResponseBody {
	return s.Body
}

func (s *DescribeQueryConfigsResponse) SetHeaders(v map[string]*string) *DescribeQueryConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeQueryConfigsResponse) SetStatusCode(v int32) *DescribeQueryConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQueryConfigsResponse) SetBody(v *DescribeQueryConfigsResponseBody) *DescribeQueryConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeQueryConfigsResponse) Validate() error {
	return dara.Validate(s)
}
