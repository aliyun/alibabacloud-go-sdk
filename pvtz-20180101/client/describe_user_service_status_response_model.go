// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserServiceStatusResponseBody) *DescribeUserServiceStatusResponse
	GetBody() *DescribeUserServiceStatusResponseBody
}

type DescribeUserServiceStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserServiceStatusResponse) GetBody() *DescribeUserServiceStatusResponseBody {
	return s.Body
}

func (s *DescribeUserServiceStatusResponse) SetHeaders(v map[string]*string) *DescribeUserServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserServiceStatusResponse) SetStatusCode(v int32) *DescribeUserServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserServiceStatusResponse) SetBody(v *DescribeUserServiceStatusResponseBody) *DescribeUserServiceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserServiceStatusResponse) Validate() error {
	return dara.Validate(s)
}
