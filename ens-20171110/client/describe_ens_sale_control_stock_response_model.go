// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsSaleControlStockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsSaleControlStockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsSaleControlStockResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsSaleControlStockResponseBody) *DescribeEnsSaleControlStockResponse
	GetBody() *DescribeEnsSaleControlStockResponseBody
}

type DescribeEnsSaleControlStockResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsSaleControlStockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsSaleControlStockResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsSaleControlStockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsSaleControlStockResponse) GetBody() *DescribeEnsSaleControlStockResponseBody {
	return s.Body
}

func (s *DescribeEnsSaleControlStockResponse) SetHeaders(v map[string]*string) *DescribeEnsSaleControlStockResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsSaleControlStockResponse) SetStatusCode(v int32) *DescribeEnsSaleControlStockResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponse) SetBody(v *DescribeEnsSaleControlStockResponseBody) *DescribeEnsSaleControlStockResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsSaleControlStockResponse) Validate() error {
	return dara.Validate(s)
}
