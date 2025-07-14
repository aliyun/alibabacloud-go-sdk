// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChangeOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChangeOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChangeOrderResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChangeOrderResponseBody) *DescribeChangeOrderResponse
	GetBody() *DescribeChangeOrderResponseBody
}

type DescribeChangeOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChangeOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeChangeOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChangeOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChangeOrderResponse) GetBody() *DescribeChangeOrderResponseBody {
	return s.Body
}

func (s *DescribeChangeOrderResponse) SetHeaders(v map[string]*string) *DescribeChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeChangeOrderResponse) SetStatusCode(v int32) *DescribeChangeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChangeOrderResponse) SetBody(v *DescribeChangeOrderResponseBody) *DescribeChangeOrderResponse {
	s.Body = v
	return s
}

func (s *DescribeChangeOrderResponse) Validate() error {
	return dara.Validate(s)
}
