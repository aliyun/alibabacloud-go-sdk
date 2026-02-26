// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenderPurchaseOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PurchaseOrderRenderQuery) *RenderPurchaseOrderRequest
	GetBody() *PurchaseOrderRenderQuery
}

type RenderPurchaseOrderRequest struct {
	// This parameter is required.
	Body *PurchaseOrderRenderQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenderPurchaseOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s RenderPurchaseOrderRequest) GoString() string {
	return s.String()
}

func (s *RenderPurchaseOrderRequest) GetBody() *PurchaseOrderRenderQuery {
	return s.Body
}

func (s *RenderPurchaseOrderRequest) SetBody(v *PurchaseOrderRenderQuery) *RenderPurchaseOrderRequest {
	s.Body = v
	return s
}

func (s *RenderPurchaseOrderRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
