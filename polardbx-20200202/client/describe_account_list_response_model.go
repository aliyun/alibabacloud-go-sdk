// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountListResponseBody) *DescribeAccountListResponse
	GetBody() *DescribeAccountListResponseBody
}

type DescribeAccountListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountListResponse) GetBody() *DescribeAccountListResponseBody {
	return s.Body
}

func (s *DescribeAccountListResponse) SetHeaders(v map[string]*string) *DescribeAccountListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountListResponse) SetStatusCode(v int32) *DescribeAccountListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountListResponse) SetBody(v *DescribeAccountListResponseBody) *DescribeAccountListResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountListResponse) Validate() error {
	return dara.Validate(s)
}
