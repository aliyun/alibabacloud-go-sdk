// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseOrderStatusResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PurchaseOrderStatusResult
	GetRequestId() *string
	SetStatus(v string) *PurchaseOrderStatusResult
	GetStatus() *string
}

type PurchaseOrderStatusResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PurchaseOrderStatusResult) String() string {
	return dara.Prettify(s)
}

func (s PurchaseOrderStatusResult) GoString() string {
	return s.String()
}

func (s *PurchaseOrderStatusResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PurchaseOrderStatusResult) GetStatus() *string {
	return s.Status
}

func (s *PurchaseOrderStatusResult) SetRequestId(v string) *PurchaseOrderStatusResult {
	s.RequestId = &v
	return s
}

func (s *PurchaseOrderStatusResult) SetStatus(v string) *PurchaseOrderStatusResult {
	s.Status = &v
	return s
}

func (s *PurchaseOrderStatusResult) Validate() error {
	return dara.Validate(s)
}
