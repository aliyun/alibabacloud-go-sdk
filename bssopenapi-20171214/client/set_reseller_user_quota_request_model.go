// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResellerUserQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v string) *SetResellerUserQuotaRequest
	GetAmount() *string
	SetCurrency(v string) *SetResellerUserQuotaRequest
	GetCurrency() *string
	SetOutBizId(v string) *SetResellerUserQuotaRequest
	GetOutBizId() *string
	SetOwnerId(v int64) *SetResellerUserQuotaRequest
	GetOwnerId() *int64
}

type SetResellerUserQuotaRequest struct {
	// The quota of a quota ledger. Unit: CNY.
	//
	// This parameter is required.
	//
	// example:
	//
	// 750
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// You do not need to set the parameter.
	//
	// example:
	//
	// N/A
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The ID of the business.
	//
	// example:
	//
	// OD2022040818295234777795624764689
	OutBizId *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	// This parameter is required.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SetResellerUserQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s SetResellerUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetResellerUserQuotaRequest) GetAmount() *string {
	return s.Amount
}

func (s *SetResellerUserQuotaRequest) GetCurrency() *string {
	return s.Currency
}

func (s *SetResellerUserQuotaRequest) GetOutBizId() *string {
	return s.OutBizId
}

func (s *SetResellerUserQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetResellerUserQuotaRequest) SetAmount(v string) *SetResellerUserQuotaRequest {
	s.Amount = &v
	return s
}

func (s *SetResellerUserQuotaRequest) SetCurrency(v string) *SetResellerUserQuotaRequest {
	s.Currency = &v
	return s
}

func (s *SetResellerUserQuotaRequest) SetOutBizId(v string) *SetResellerUserQuotaRequest {
	s.OutBizId = &v
	return s
}

func (s *SetResellerUserQuotaRequest) SetOwnerId(v int64) *SetResellerUserQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *SetResellerUserQuotaRequest) Validate() error {
	return dara.Validate(s)
}
