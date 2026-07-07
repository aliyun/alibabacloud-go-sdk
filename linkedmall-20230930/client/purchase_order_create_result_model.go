// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseOrderCreateResult interface {
	dara.Model
	String() string
	GoString() string
	SetPurchaseOrderId(v string) *PurchaseOrderCreateResult
	GetPurchaseOrderId() *string
	SetRequestId(v string) *PurchaseOrderCreateResult
	GetRequestId() *string
}

type PurchaseOrderCreateResult struct {
	// Purchase Order ID
	//
	// example:
	//
	// 6692****5696
	PurchaseOrderId *string `json:"purchaseOrderId,omitempty" xml:"purchaseOrderId,omitempty"`
	// API Request requestId
	//
	// example:
	//
	// 841471F6-5D61-1331-8C38-2****B55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PurchaseOrderCreateResult) String() string {
	return dara.Prettify(s)
}

func (s PurchaseOrderCreateResult) GoString() string {
	return s.String()
}

func (s *PurchaseOrderCreateResult) GetPurchaseOrderId() *string {
	return s.PurchaseOrderId
}

func (s *PurchaseOrderCreateResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PurchaseOrderCreateResult) SetPurchaseOrderId(v string) *PurchaseOrderCreateResult {
	s.PurchaseOrderId = &v
	return s
}

func (s *PurchaseOrderCreateResult) SetRequestId(v string) *PurchaseOrderCreateResult {
	s.RequestId = &v
	return s
}

func (s *PurchaseOrderCreateResult) Validate() error {
	return dara.Validate(s)
}
