// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSplitPurchaseOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SplitPurchaseOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SplitPurchaseOrderResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseOrderRenderResult) *SplitPurchaseOrderResponse
	GetBody() *PurchaseOrderRenderResult
}

type SplitPurchaseOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseOrderRenderResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SplitPurchaseOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s SplitPurchaseOrderResponse) GoString() string {
	return s.String()
}

func (s *SplitPurchaseOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SplitPurchaseOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SplitPurchaseOrderResponse) GetBody() *PurchaseOrderRenderResult {
	return s.Body
}

func (s *SplitPurchaseOrderResponse) SetHeaders(v map[string]*string) *SplitPurchaseOrderResponse {
	s.Headers = v
	return s
}

func (s *SplitPurchaseOrderResponse) SetStatusCode(v int32) *SplitPurchaseOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SplitPurchaseOrderResponse) SetBody(v *PurchaseOrderRenderResult) *SplitPurchaseOrderResponse {
	s.Body = v
	return s
}

func (s *SplitPurchaseOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
