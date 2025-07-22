// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCallbacksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCallbacksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCallbacksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCallbacksResponseBody) *DescribeCallbacksResponse
	GetBody() *DescribeCallbacksResponseBody
}

type DescribeCallbacksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCallbacksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCallbacksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallbacksResponse) GoString() string {
	return s.String()
}

func (s *DescribeCallbacksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCallbacksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCallbacksResponse) GetBody() *DescribeCallbacksResponseBody {
	return s.Body
}

func (s *DescribeCallbacksResponse) SetHeaders(v map[string]*string) *DescribeCallbacksResponse {
	s.Headers = v
	return s
}

func (s *DescribeCallbacksResponse) SetStatusCode(v int32) *DescribeCallbacksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCallbacksResponse) SetBody(v *DescribeCallbacksResponseBody) *DescribeCallbacksResponse {
	s.Body = v
	return s
}

func (s *DescribeCallbacksResponse) Validate() error {
	return dara.Validate(s)
}
