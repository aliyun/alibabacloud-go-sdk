// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppLayoutsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppLayoutsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppLayoutsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppLayoutsResponseBody) *DescribeAppLayoutsResponse
	GetBody() *DescribeAppLayoutsResponseBody
}

type DescribeAppLayoutsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppLayoutsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppLayoutsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppLayoutsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppLayoutsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppLayoutsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppLayoutsResponse) GetBody() *DescribeAppLayoutsResponseBody {
	return s.Body
}

func (s *DescribeAppLayoutsResponse) SetHeaders(v map[string]*string) *DescribeAppLayoutsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppLayoutsResponse) SetStatusCode(v int32) *DescribeAppLayoutsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppLayoutsResponse) SetBody(v *DescribeAppLayoutsResponseBody) *DescribeAppLayoutsResponse {
	s.Body = v
	return s
}

func (s *DescribeAppLayoutsResponse) Validate() error {
	return dara.Validate(s)
}
