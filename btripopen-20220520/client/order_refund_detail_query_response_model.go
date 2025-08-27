// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderRefundDetailQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OrderRefundDetailQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OrderRefundDetailQueryResponse
	GetStatusCode() *int32
	SetBody(v *OrderRefundDetailQueryResponseBody) *OrderRefundDetailQueryResponse
	GetBody() *OrderRefundDetailQueryResponseBody
}

type OrderRefundDetailQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OrderRefundDetailQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OrderRefundDetailQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s OrderRefundDetailQueryResponse) GoString() string {
	return s.String()
}

func (s *OrderRefundDetailQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OrderRefundDetailQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OrderRefundDetailQueryResponse) GetBody() *OrderRefundDetailQueryResponseBody {
	return s.Body
}

func (s *OrderRefundDetailQueryResponse) SetHeaders(v map[string]*string) *OrderRefundDetailQueryResponse {
	s.Headers = v
	return s
}

func (s *OrderRefundDetailQueryResponse) SetStatusCode(v int32) *OrderRefundDetailQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *OrderRefundDetailQueryResponse) SetBody(v *OrderRefundDetailQueryResponseBody) *OrderRefundDetailQueryResponse {
	s.Body = v
	return s
}

func (s *OrderRefundDetailQueryResponse) Validate() error {
	return dara.Validate(s)
}
