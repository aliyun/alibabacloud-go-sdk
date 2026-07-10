// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmPreBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *ConfirmPreBillRequest
	GetBillBatch() *string
}

type ConfirmPreBillRequest struct {
	// The bill batch date in the format of yyyy-MM-dd, such as 2026-06-21.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-06-21
	BillBatch *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
}

func (s ConfirmPreBillRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPreBillRequest) GoString() string {
	return s.String()
}

func (s *ConfirmPreBillRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *ConfirmPreBillRequest) SetBillBatch(v string) *ConfirmPreBillRequest {
	s.BillBatch = &v
	return s
}

func (s *ConfirmPreBillRequest) Validate() error {
	return dara.Validate(s)
}
