// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOrderResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOrderResponseBody) *DescribeOrderResponse
	GetBody() *DescribeOrderResponseBody
}

type DescribeOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOrderResponse) GetBody() *DescribeOrderResponseBody {
	return s.Body
}

func (s *DescribeOrderResponse) SetHeaders(v map[string]*string) *DescribeOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeOrderResponse) SetStatusCode(v int32) *DescribeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOrderResponse) SetBody(v *DescribeOrderResponseBody) *DescribeOrderResponse {
	s.Body = v
	return s
}

func (s *DescribeOrderResponse) Validate() error {
	return dara.Validate(s)
}
