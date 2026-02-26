// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePurchaseOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePurchaseOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePurchaseOrderResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseOrderCreateResult) *CreatePurchaseOrderResponse
	GetBody() *PurchaseOrderCreateResult
}

type CreatePurchaseOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseOrderCreateResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePurchaseOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePurchaseOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePurchaseOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePurchaseOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePurchaseOrderResponse) GetBody() *PurchaseOrderCreateResult {
	return s.Body
}

func (s *CreatePurchaseOrderResponse) SetHeaders(v map[string]*string) *CreatePurchaseOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePurchaseOrderResponse) SetStatusCode(v int32) *CreatePurchaseOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePurchaseOrderResponse) SetBody(v *PurchaseOrderCreateResult) *CreatePurchaseOrderResponse {
	s.Body = v
	return s
}

func (s *CreatePurchaseOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
