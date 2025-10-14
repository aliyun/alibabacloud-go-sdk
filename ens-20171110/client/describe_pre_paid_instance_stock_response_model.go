// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrePaidInstanceStockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrePaidInstanceStockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrePaidInstanceStockResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrePaidInstanceStockResponseBody) *DescribePrePaidInstanceStockResponse
	GetBody() *DescribePrePaidInstanceStockResponseBody
}

type DescribePrePaidInstanceStockResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrePaidInstanceStockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrePaidInstanceStockResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrePaidInstanceStockResponse) GoString() string {
	return s.String()
}

func (s *DescribePrePaidInstanceStockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrePaidInstanceStockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrePaidInstanceStockResponse) GetBody() *DescribePrePaidInstanceStockResponseBody {
	return s.Body
}

func (s *DescribePrePaidInstanceStockResponse) SetHeaders(v map[string]*string) *DescribePrePaidInstanceStockResponse {
	s.Headers = v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) SetStatusCode(v int32) *DescribePrePaidInstanceStockResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) SetBody(v *DescribePrePaidInstanceStockResponseBody) *DescribePrePaidInstanceStockResponse {
	s.Body = v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
