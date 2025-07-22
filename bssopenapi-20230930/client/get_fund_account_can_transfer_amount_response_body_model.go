// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanTransferAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAmount(v string) *GetFundAccountCanTransferAmountResponseBody
	GetAvailableAmount() *string
	SetCashAmount(v string) *GetFundAccountCanTransferAmountResponseBody
	GetCashAmount() *string
	SetCurrency(v string) *GetFundAccountCanTransferAmountResponseBody
	GetCurrency() *string
	SetFundAccountEcid(v string) *GetFundAccountCanTransferAmountResponseBody
	GetFundAccountEcid() *string
	SetFundAccountId(v int64) *GetFundAccountCanTransferAmountResponseBody
	GetFundAccountId() *int64
	SetFundAccountName(v string) *GetFundAccountCanTransferAmountResponseBody
	GetFundAccountName() *string
	SetFundAccountOwnerAccountId(v int64) *GetFundAccountCanTransferAmountResponseBody
	GetFundAccountOwnerAccountId() *int64
	SetMaxTransferableAmount(v string) *GetFundAccountCanTransferAmountResponseBody
	GetMaxTransferableAmount() *string
	SetMetadata(v interface{}) *GetFundAccountCanTransferAmountResponseBody
	GetMetadata() interface{}
	SetNbid(v string) *GetFundAccountCanTransferAmountResponseBody
	GetNbid() *string
	SetRequestId(v string) *GetFundAccountCanTransferAmountResponseBody
	GetRequestId() *string
	SetSite(v string) *GetFundAccountCanTransferAmountResponseBody
	GetSite() *string
	SetTransferAmount(v string) *GetFundAccountCanTransferAmountResponseBody
	GetTransferAmount() *string
}

type GetFundAccountCanTransferAmountResponseBody struct {
	// example:
	//
	// 100
	AvailableAmount *string `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// example:
	//
	// 500
	CashAmount *string `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 2032121324
	FundAccountEcid *string `json:"FundAccountEcid,omitempty" xml:"FundAccountEcid,omitempty"`
	// example:
	//
	// 12332112
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 云某的账户
	FundAccountName *string `json:"FundAccountName,omitempty" xml:"FundAccountName,omitempty"`
	// example:
	//
	// 154738212323
	FundAccountOwnerAccountId *int64 `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// 100
	MaxTransferableAmount *string `json:"MaxTransferableAmount,omitempty" xml:"MaxTransferableAmount,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 2684210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 26842
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// example:
	//
	// 100
	TransferAmount *string `json:"TransferAmount,omitempty" xml:"TransferAmount,omitempty"`
}

func (s GetFundAccountCanTransferAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanTransferAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetAvailableAmount() *string {
	return s.AvailableAmount
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetCashAmount() *string {
	return s.CashAmount
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetCurrency() *string {
	return s.Currency
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetFundAccountEcid() *string {
	return s.FundAccountEcid
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetFundAccountName() *string {
	return s.FundAccountName
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetFundAccountOwnerAccountId() *int64 {
	return s.FundAccountOwnerAccountId
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetMaxTransferableAmount() *string {
	return s.MaxTransferableAmount
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetNbid() *string {
	return s.Nbid
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetSite() *string {
	return s.Site
}

func (s *GetFundAccountCanTransferAmountResponseBody) GetTransferAmount() *string {
	return s.TransferAmount
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetAvailableAmount(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.AvailableAmount = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetCashAmount(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.CashAmount = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetCurrency(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.Currency = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetFundAccountEcid(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.FundAccountEcid = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetFundAccountId(v int64) *GetFundAccountCanTransferAmountResponseBody {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetFundAccountName(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.FundAccountName = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetFundAccountOwnerAccountId(v int64) *GetFundAccountCanTransferAmountResponseBody {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetMaxTransferableAmount(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.MaxTransferableAmount = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetMetadata(v interface{}) *GetFundAccountCanTransferAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetNbid(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.Nbid = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetRequestId(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetSite(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.Site = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetTransferAmount(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.TransferAmount = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) Validate() error {
	return dara.Validate(s)
}
