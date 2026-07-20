// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetThirdPartId(v string) *InvoiceDeleteRequest
	GetThirdPartId() *string
}

type InvoiceDeleteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 340049
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s InvoiceDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s InvoiceDeleteRequest) GoString() string {
	return s.String()
}

func (s *InvoiceDeleteRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *InvoiceDeleteRequest) SetThirdPartId(v string) *InvoiceDeleteRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceDeleteRequest) Validate() error {
	return dara.Validate(s)
}
