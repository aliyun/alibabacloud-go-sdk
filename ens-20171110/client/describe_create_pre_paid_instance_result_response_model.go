// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreatePrePaidInstanceResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCreatePrePaidInstanceResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCreatePrePaidInstanceResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCreatePrePaidInstanceResultResponseBody) *DescribeCreatePrePaidInstanceResultResponse
	GetBody() *DescribeCreatePrePaidInstanceResultResponseBody
}

type DescribeCreatePrePaidInstanceResultResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCreatePrePaidInstanceResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCreatePrePaidInstanceResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCreatePrePaidInstanceResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCreatePrePaidInstanceResultResponse) GetBody() *DescribeCreatePrePaidInstanceResultResponseBody {
	return s.Body
}

func (s *DescribeCreatePrePaidInstanceResultResponse) SetHeaders(v map[string]*string) *DescribeCreatePrePaidInstanceResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponse) SetStatusCode(v int32) *DescribeCreatePrePaidInstanceResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponse) SetBody(v *DescribeCreatePrePaidInstanceResultResponseBody) *DescribeCreatePrePaidInstanceResultResponse {
	s.Body = v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponse) Validate() error {
	return dara.Validate(s)
}
