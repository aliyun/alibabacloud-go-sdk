// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenderRefundOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenderRefundOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenderRefundOrderResponse
	GetStatusCode() *int32
	SetBody(v *RefundRenderResult) *RenderRefundOrderResponse
	GetBody() *RefundRenderResult
}

type RenderRefundOrderResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundRenderResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenderRefundOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s RenderRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *RenderRefundOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenderRefundOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenderRefundOrderResponse) GetBody() *RefundRenderResult {
	return s.Body
}

func (s *RenderRefundOrderResponse) SetHeaders(v map[string]*string) *RenderRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *RenderRefundOrderResponse) SetStatusCode(v int32) *RenderRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RenderRefundOrderResponse) SetBody(v *RefundRenderResult) *RenderRefundOrderResponse {
	s.Body = v
	return s
}

func (s *RenderRefundOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
