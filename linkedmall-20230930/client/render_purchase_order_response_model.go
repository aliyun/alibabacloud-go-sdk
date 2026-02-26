// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenderPurchaseOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenderPurchaseOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenderPurchaseOrderResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseOrderRenderResult) *RenderPurchaseOrderResponse
	GetBody() *PurchaseOrderRenderResult
}

type RenderPurchaseOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseOrderRenderResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenderPurchaseOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s RenderPurchaseOrderResponse) GoString() string {
	return s.String()
}

func (s *RenderPurchaseOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenderPurchaseOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenderPurchaseOrderResponse) GetBody() *PurchaseOrderRenderResult {
	return s.Body
}

func (s *RenderPurchaseOrderResponse) SetHeaders(v map[string]*string) *RenderPurchaseOrderResponse {
	s.Headers = v
	return s
}

func (s *RenderPurchaseOrderResponse) SetStatusCode(v int32) *RenderPurchaseOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RenderPurchaseOrderResponse) SetBody(v *PurchaseOrderRenderResult) *RenderPurchaseOrderResponse {
	s.Body = v
	return s
}

func (s *RenderPurchaseOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
