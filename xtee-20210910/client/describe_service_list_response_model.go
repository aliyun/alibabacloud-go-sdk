// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceListResponseBody) *DescribeServiceListResponse
	GetBody() *DescribeServiceListResponseBody
}

type DescribeServiceListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceListResponse) GetBody() *DescribeServiceListResponseBody {
	return s.Body
}

func (s *DescribeServiceListResponse) SetHeaders(v map[string]*string) *DescribeServiceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceListResponse) SetStatusCode(v int32) *DescribeServiceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceListResponse) SetBody(v *DescribeServiceListResponseBody) *DescribeServiceListResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceListResponse) Validate() error {
	return dara.Validate(s)
}
