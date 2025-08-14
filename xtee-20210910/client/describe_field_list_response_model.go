// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFieldListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFieldListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFieldListResponseBody) *DescribeFieldListResponse
	GetBody() *DescribeFieldListResponseBody
}

type DescribeFieldListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFieldListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFieldListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFieldListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFieldListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFieldListResponse) GetBody() *DescribeFieldListResponseBody {
	return s.Body
}

func (s *DescribeFieldListResponse) SetHeaders(v map[string]*string) *DescribeFieldListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFieldListResponse) SetStatusCode(v int32) *DescribeFieldListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFieldListResponse) SetBody(v *DescribeFieldListResponseBody) *DescribeFieldListResponse {
	s.Body = v
	return s
}

func (s *DescribeFieldListResponse) Validate() error {
	return dara.Validate(s)
}
