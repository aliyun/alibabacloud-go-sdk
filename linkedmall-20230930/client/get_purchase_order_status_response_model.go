// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPurchaseOrderStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPurchaseOrderStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPurchaseOrderStatusResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseOrderStatusResult) *GetPurchaseOrderStatusResponse
	GetBody() *PurchaseOrderStatusResult
}

type GetPurchaseOrderStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseOrderStatusResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPurchaseOrderStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPurchaseOrderStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPurchaseOrderStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPurchaseOrderStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPurchaseOrderStatusResponse) GetBody() *PurchaseOrderStatusResult {
	return s.Body
}

func (s *GetPurchaseOrderStatusResponse) SetHeaders(v map[string]*string) *GetPurchaseOrderStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPurchaseOrderStatusResponse) SetStatusCode(v int32) *GetPurchaseOrderStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPurchaseOrderStatusResponse) SetBody(v *PurchaseOrderStatusResult) *GetPurchaseOrderStatusResponse {
	s.Body = v
	return s
}

func (s *GetPurchaseOrderStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
