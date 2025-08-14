// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskLogListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTaskLogListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTaskLogListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTaskLogListResponseBody) *DescribeTaskLogListResponse
	GetBody() *DescribeTaskLogListResponseBody
}

type DescribeTaskLogListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskLogListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskLogListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskLogListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskLogListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTaskLogListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTaskLogListResponse) GetBody() *DescribeTaskLogListResponseBody {
	return s.Body
}

func (s *DescribeTaskLogListResponse) SetHeaders(v map[string]*string) *DescribeTaskLogListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskLogListResponse) SetStatusCode(v int32) *DescribeTaskLogListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskLogListResponse) SetBody(v *DescribeTaskLogListResponseBody) *DescribeTaskLogListResponse {
	s.Body = v
	return s
}

func (s *DescribeTaskLogListResponse) Validate() error {
	return dara.Validate(s)
}
