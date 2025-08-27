// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceAddRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *InvoiceAddRequest
	GetAddress() *string
	SetBankName(v string) *InvoiceAddRequest
	GetBankName() *string
	SetBankNo(v string) *InvoiceAddRequest
	GetBankNo() *string
	SetTaxNo(v string) *InvoiceAddRequest
	GetTaxNo() *string
	SetTel(v string) *InvoiceAddRequest
	GetTel() *string
	SetThirdPartId(v string) *InvoiceAddRequest
	GetThirdPartId() *string
	SetTitle(v string) *InvoiceAddRequest
	GetTitle() *string
	SetType(v int32) *InvoiceAddRequest
	GetType() *int32
	SetUnitType(v int32) *InvoiceAddRequest
	GetUnitType() *int32
}

type InvoiceAddRequest struct {
	// example:
	//
	// xxx
	Address  *string `json:"address,omitempty" xml:"address,omitempty"`
	BankName *string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// example:
	//
	// 123456
	BankNo *string `json:"bank_no,omitempty" xml:"bank_no,omitempty"`
	// example:
	//
	// 12345
	TaxNo *string `json:"tax_no,omitempty" xml:"tax_no,omitempty"`
	// example:
	//
	// 123
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

func (s InvoiceAddRequest) String() string {
	return dara.Prettify(s)
}

func (s InvoiceAddRequest) GoString() string {
	return s.String()
}

func (s *InvoiceAddRequest) GetAddress() *string {
	return s.Address
}

func (s *InvoiceAddRequest) GetBankName() *string {
	return s.BankName
}

func (s *InvoiceAddRequest) GetBankNo() *string {
	return s.BankNo
}

func (s *InvoiceAddRequest) GetTaxNo() *string {
	return s.TaxNo
}

func (s *InvoiceAddRequest) GetTel() *string {
	return s.Tel
}

func (s *InvoiceAddRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *InvoiceAddRequest) GetTitle() *string {
	return s.Title
}

func (s *InvoiceAddRequest) GetType() *int32 {
	return s.Type
}

func (s *InvoiceAddRequest) GetUnitType() *int32 {
	return s.UnitType
}

func (s *InvoiceAddRequest) SetAddress(v string) *InvoiceAddRequest {
	s.Address = &v
	return s
}

func (s *InvoiceAddRequest) SetBankName(v string) *InvoiceAddRequest {
	s.BankName = &v
	return s
}

func (s *InvoiceAddRequest) SetBankNo(v string) *InvoiceAddRequest {
	s.BankNo = &v
	return s
}

func (s *InvoiceAddRequest) SetTaxNo(v string) *InvoiceAddRequest {
	s.TaxNo = &v
	return s
}

func (s *InvoiceAddRequest) SetTel(v string) *InvoiceAddRequest {
	s.Tel = &v
	return s
}

func (s *InvoiceAddRequest) SetThirdPartId(v string) *InvoiceAddRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceAddRequest) SetTitle(v string) *InvoiceAddRequest {
	s.Title = &v
	return s
}

func (s *InvoiceAddRequest) SetType(v int32) *InvoiceAddRequest {
	s.Type = &v
	return s
}

func (s *InvoiceAddRequest) SetUnitType(v int32) *InvoiceAddRequest {
	s.UnitType = &v
	return s
}

func (s *InvoiceAddRequest) Validate() error {
	return dara.Validate(s)
}
