// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFundAccountTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v string) *CreateFundAccountTransferRequest
	GetAmount() *string
	SetCurrency(v string) *CreateFundAccountTransferRequest
	GetCurrency() *string
	SetFinanceType(v string) *CreateFundAccountTransferRequest
	GetFinanceType() *string
	SetFromFundAccountId(v int64) *CreateFundAccountTransferRequest
	GetFromFundAccountId() *int64
	SetRemark(v string) *CreateFundAccountTransferRequest
	GetRemark() *string
	SetToFundAccountId(v int64) *CreateFundAccountTransferRequest
	GetToFundAccountId() *int64
	SetTransferType(v string) *CreateFundAccountTransferRequest
	GetTransferType() *string
}

type CreateFundAccountTransferRequest struct {
	// The transfer amount.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency.
	//
	// This parameter is required.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The asset type. Currently, only cash is supported. Valid values:
	//
	// - cash: cash.
	//
	// This parameter is required.
	//
	// example:
	//
	// cash
	FinanceType *string `json:"FinanceType,omitempty" xml:"FinanceType,omitempty"`
	// The ID of the source account (payer account ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123212323
	FromFundAccountId *int64 `json:"FromFundAccountId,omitempty" xml:"FromFundAccountId,omitempty"`
	// The remarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 转账的备注
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the destination account (payee account ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// 11112231
	ToFundAccountId *int64 `json:"ToFundAccountId,omitempty" xml:"ToFundAccountId,omitempty"`
	// The transfer type. Valid values:
	//
	// - grant: transfer.
	//
	// - recycle: revocation.
	//
	// This parameter is required.
	//
	// example:
	//
	// grant
	TransferType *string `json:"TransferType,omitempty" xml:"TransferType,omitempty"`
}

func (s CreateFundAccountTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFundAccountTransferRequest) GoString() string {
	return s.String()
}

func (s *CreateFundAccountTransferRequest) GetAmount() *string {
	return s.Amount
}

func (s *CreateFundAccountTransferRequest) GetCurrency() *string {
	return s.Currency
}

func (s *CreateFundAccountTransferRequest) GetFinanceType() *string {
	return s.FinanceType
}

func (s *CreateFundAccountTransferRequest) GetFromFundAccountId() *int64 {
	return s.FromFundAccountId
}

func (s *CreateFundAccountTransferRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateFundAccountTransferRequest) GetToFundAccountId() *int64 {
	return s.ToFundAccountId
}

func (s *CreateFundAccountTransferRequest) GetTransferType() *string {
	return s.TransferType
}

func (s *CreateFundAccountTransferRequest) SetAmount(v string) *CreateFundAccountTransferRequest {
	s.Amount = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetCurrency(v string) *CreateFundAccountTransferRequest {
	s.Currency = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetFinanceType(v string) *CreateFundAccountTransferRequest {
	s.FinanceType = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetFromFundAccountId(v int64) *CreateFundAccountTransferRequest {
	s.FromFundAccountId = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetRemark(v string) *CreateFundAccountTransferRequest {
	s.Remark = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetToFundAccountId(v int64) *CreateFundAccountTransferRequest {
	s.ToFundAccountId = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetTransferType(v string) *CreateFundAccountTransferRequest {
	s.TransferType = &v
	return s
}

func (s *CreateFundAccountTransferRequest) Validate() error {
	return dara.Validate(s)
}
