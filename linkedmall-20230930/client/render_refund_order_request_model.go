// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenderRefundOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *RefundRenderCmd) *RenderRefundOrderRequest
	GetBody() *RefundRenderCmd
}

type RenderRefundOrderRequest struct {
	// This parameter is required.
	Body *RefundRenderCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenderRefundOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s RenderRefundOrderRequest) GoString() string {
	return s.String()
}

func (s *RenderRefundOrderRequest) GetBody() *RefundRenderCmd {
	return s.Body
}

func (s *RenderRefundOrderRequest) SetBody(v *RefundRenderCmd) *RenderRefundOrderRequest {
	s.Body = v
	return s
}

func (s *RenderRefundOrderRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
