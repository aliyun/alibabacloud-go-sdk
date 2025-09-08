// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAggregateFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAggregateFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAggregateFunctionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAggregateFunctionResponseBody) *DescribeAggregateFunctionResponse
	GetBody() *DescribeAggregateFunctionResponseBody
}

type DescribeAggregateFunctionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAggregateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAggregateFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAggregateFunctionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAggregateFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAggregateFunctionResponse) GetBody() *DescribeAggregateFunctionResponseBody {
	return s.Body
}

func (s *DescribeAggregateFunctionResponse) SetHeaders(v map[string]*string) *DescribeAggregateFunctionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAggregateFunctionResponse) SetStatusCode(v int32) *DescribeAggregateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAggregateFunctionResponse) SetBody(v *DescribeAggregateFunctionResponseBody) *DescribeAggregateFunctionResponse {
	s.Body = v
	return s
}

func (s *DescribeAggregateFunctionResponse) Validate() error {
	return dara.Validate(s)
}
