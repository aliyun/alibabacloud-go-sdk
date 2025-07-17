// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundPayAsYouGoOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefundPayAsYouGoOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefundPayAsYouGoOrderResponse
	GetStatusCode() *int32
	SetBody(v *RefundPayAsYouGoOrderResponseBody) *RefundPayAsYouGoOrderResponse
	GetBody() *RefundPayAsYouGoOrderResponseBody
}

type RefundPayAsYouGoOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundPayAsYouGoOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundPayAsYouGoOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s RefundPayAsYouGoOrderResponse) GoString() string {
	return s.String()
}

func (s *RefundPayAsYouGoOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefundPayAsYouGoOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefundPayAsYouGoOrderResponse) GetBody() *RefundPayAsYouGoOrderResponseBody {
	return s.Body
}

func (s *RefundPayAsYouGoOrderResponse) SetHeaders(v map[string]*string) *RefundPayAsYouGoOrderResponse {
	s.Headers = v
	return s
}

func (s *RefundPayAsYouGoOrderResponse) SetStatusCode(v int32) *RefundPayAsYouGoOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundPayAsYouGoOrderResponse) SetBody(v *RefundPayAsYouGoOrderResponseBody) *RefundPayAsYouGoOrderResponse {
	s.Body = v
	return s
}

func (s *RefundPayAsYouGoOrderResponse) Validate() error {
	return dara.Validate(s)
}
