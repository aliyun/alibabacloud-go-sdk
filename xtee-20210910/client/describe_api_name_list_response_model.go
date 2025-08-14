// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiNameListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiNameListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiNameListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiNameListResponseBody) *DescribeApiNameListResponse
	GetBody() *DescribeApiNameListResponseBody
}

type DescribeApiNameListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiNameListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiNameListResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiNameListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiNameListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiNameListResponse) GetBody() *DescribeApiNameListResponseBody {
	return s.Body
}

func (s *DescribeApiNameListResponse) SetHeaders(v map[string]*string) *DescribeApiNameListResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiNameListResponse) SetStatusCode(v int32) *DescribeApiNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiNameListResponse) SetBody(v *DescribeApiNameListResponseBody) *DescribeApiNameListResponse {
	s.Body = v
	return s
}

func (s *DescribeApiNameListResponse) Validate() error {
	return dara.Validate(s)
}
