// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshTaskByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRefreshTaskByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRefreshTaskByIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRefreshTaskByIdResponseBody) *DescribeRefreshTaskByIdResponse
	GetBody() *DescribeRefreshTaskByIdResponseBody
}

type DescribeRefreshTaskByIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRefreshTaskByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRefreshTaskByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTaskByIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTaskByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRefreshTaskByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRefreshTaskByIdResponse) GetBody() *DescribeRefreshTaskByIdResponseBody {
	return s.Body
}

func (s *DescribeRefreshTaskByIdResponse) SetHeaders(v map[string]*string) *DescribeRefreshTaskByIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeRefreshTaskByIdResponse) SetStatusCode(v int32) *DescribeRefreshTaskByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponse) SetBody(v *DescribeRefreshTaskByIdResponseBody) *DescribeRefreshTaskByIdResponse {
	s.Body = v
	return s
}

func (s *DescribeRefreshTaskByIdResponse) Validate() error {
	return dara.Validate(s)
}
