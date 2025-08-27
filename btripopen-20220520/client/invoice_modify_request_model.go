// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceModifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *InvoiceModifyRequest
	GetAddress() *string
	SetBankName(v string) *InvoiceModifyRequest
	GetBankName() *string
	SetBankNo(v string) *InvoiceModifyRequest
	GetBankNo() *string
	SetTaxNo(v string) *InvoiceModifyRequest
	GetTaxNo() *string
	SetTel(v string) *InvoiceModifyRequest
	GetTel() *string
	SetThirdPartId(v string) *InvoiceModifyRequest
	GetThirdPartId() *string
	SetTitle(v string) *InvoiceModifyRequest
	GetTitle() *string
	SetType(v int32) *InvoiceModifyRequest
	GetType() *int32
	SetUnitType(v int32) *InvoiceModifyRequest
	GetUnitType() *int32
}

type InvoiceModifyRequest struct {
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 12345678
	BankName *string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// example:
	//
	// 2222
	BankNo *string `json:"bank_no,omitempty" xml:"bank_no,omitempty"`
	// example:
	//
	// 122333121
	TaxNo *string `json:"tax_no,omitempty" xml:"tax_no,omitempty"`
	// example:
	//
	// 12345678
	Tel *string `json:"tel,omitempty" xml:"tel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 340049
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type     *int32 `json:"type,omitempty" xml:"type,omitempty"`
	UnitType *int32 `json:"unit_type,omitempty" xml:"unit_type,omitempty"`
}

func (s InvoiceModifyRequest) String() string {
	return dara.Prettify(s)
}

func (s InvoiceModifyRequest) GoString() string {
	return s.String()
}

func (s *InvoiceModifyRequest) GetAddress() *string {
	return s.Address
}

func (s *InvoiceModifyRequest) GetBankName() *string {
	return s.BankName
}

func (s *InvoiceModifyRequest) GetBankNo() *string {
	return s.BankNo
}

func (s *InvoiceModifyRequest) GetTaxNo() *string {
	return s.TaxNo
}

func (s *InvoiceModifyRequest) GetTel() *string {
	return s.Tel
}

func (s *InvoiceModifyRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *InvoiceModifyRequest) GetTitle() *string {
	return s.Title
}

func (s *InvoiceModifyRequest) GetType() *int32 {
	return s.Type
}

func (s *InvoiceModifyRequest) GetUnitType() *int32 {
	return s.UnitType
}

func (s *InvoiceModifyRequest) SetAddress(v string) *InvoiceModifyRequest {
	s.Address = &v
	return s
}

func (s *InvoiceModifyRequest) SetBankName(v string) *InvoiceModifyRequest {
	s.BankName = &v
	return s
}

func (s *InvoiceModifyRequest) SetBankNo(v string) *InvoiceModifyRequest {
	s.BankNo = &v
	return s
}

func (s *InvoiceModifyRequest) SetTaxNo(v string) *InvoiceModifyRequest {
	s.TaxNo = &v
	return s
}

func (s *InvoiceModifyRequest) SetTel(v string) *InvoiceModifyRequest {
	s.Tel = &v
	return s
}

func (s *InvoiceModifyRequest) SetThirdPartId(v string) *InvoiceModifyRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceModifyRequest) SetTitle(v string) *InvoiceModifyRequest {
	s.Title = &v
	return s
}

func (s *InvoiceModifyRequest) SetType(v int32) *InvoiceModifyRequest {
	s.Type = &v
	return s
}

func (s *InvoiceModifyRequest) SetUnitType(v int32) *InvoiceModifyRequest {
	s.UnitType = &v
	return s
}

func (s *InvoiceModifyRequest) Validate() error {
	return dara.Validate(s)
}
