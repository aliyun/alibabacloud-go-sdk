// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePurchaseOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PurchaseOrderCreateCmd) *CreatePurchaseOrderRequest
	GetBody() *PurchaseOrderCreateCmd
}

type CreatePurchaseOrderRequest struct {
	// This parameter is required.
	Body *PurchaseOrderCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePurchaseOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePurchaseOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePurchaseOrderRequest) GetBody() *PurchaseOrderCreateCmd {
	return s.Body
}

func (s *CreatePurchaseOrderRequest) SetBody(v *PurchaseOrderCreateCmd) *CreatePurchaseOrderRequest {
	s.Body = v
	return s
}

func (s *CreatePurchaseOrderRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
