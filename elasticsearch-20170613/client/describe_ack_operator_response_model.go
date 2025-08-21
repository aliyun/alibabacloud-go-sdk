// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckOperatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAckOperatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAckOperatorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAckOperatorResponseBody) *DescribeAckOperatorResponse
	GetBody() *DescribeAckOperatorResponseBody
}

type DescribeAckOperatorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAckOperatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAckOperatorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckOperatorResponse) GoString() string {
	return s.String()
}

func (s *DescribeAckOperatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAckOperatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAckOperatorResponse) GetBody() *DescribeAckOperatorResponseBody {
	return s.Body
}

func (s *DescribeAckOperatorResponse) SetHeaders(v map[string]*string) *DescribeAckOperatorResponse {
	s.Headers = v
	return s
}

func (s *DescribeAckOperatorResponse) SetStatusCode(v int32) *DescribeAckOperatorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAckOperatorResponse) SetBody(v *DescribeAckOperatorResponseBody) *DescribeAckOperatorResponse {
	s.Body = v
	return s
}

func (s *DescribeAckOperatorResponse) Validate() error {
	return dara.Validate(s)
}
