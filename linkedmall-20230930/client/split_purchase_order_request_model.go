// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSplitPurchaseOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PurchaseOrderRenderQuery) *SplitPurchaseOrderRequest
	GetBody() *PurchaseOrderRenderQuery
}

type SplitPurchaseOrderRequest struct {
	Body *PurchaseOrderRenderQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SplitPurchaseOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s SplitPurchaseOrderRequest) GoString() string {
	return s.String()
}

func (s *SplitPurchaseOrderRequest) GetBody() *PurchaseOrderRenderQuery {
	return s.Body
}

func (s *SplitPurchaseOrderRequest) SetBody(v *PurchaseOrderRenderQuery) *SplitPurchaseOrderRequest {
	s.Body = v
	return s
}

func (s *SplitPurchaseOrderRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
