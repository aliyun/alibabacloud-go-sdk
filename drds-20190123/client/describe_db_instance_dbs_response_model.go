// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbInstanceDbsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDbInstanceDbsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDbInstanceDbsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDbInstanceDbsResponseBody) *DescribeDbInstanceDbsResponse
	GetBody() *DescribeDbInstanceDbsResponseBody
}

type DescribeDbInstanceDbsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDbInstanceDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDbInstanceDbsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstanceDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDbInstanceDbsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDbInstanceDbsResponse) GetBody() *DescribeDbInstanceDbsResponseBody {
	return s.Body
}

func (s *DescribeDbInstanceDbsResponse) SetHeaders(v map[string]*string) *DescribeDbInstanceDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbInstanceDbsResponse) SetStatusCode(v int32) *DescribeDbInstanceDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbInstanceDbsResponse) SetBody(v *DescribeDbInstanceDbsResponseBody) *DescribeDbInstanceDbsResponse {
	s.Body = v
	return s
}

func (s *DescribeDbInstanceDbsResponse) Validate() error {
	return dara.Validate(s)
}
