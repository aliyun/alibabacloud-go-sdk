// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResellerConsumeAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdjustType(v string) *ChangeResellerConsumeAmountRequest
	GetAdjustType() *string
	SetAmount(v string) *ChangeResellerConsumeAmountRequest
	GetAmount() *string
	SetBusinessType(v string) *ChangeResellerConsumeAmountRequest
	GetBusinessType() *string
	SetCurrency(v string) *ChangeResellerConsumeAmountRequest
	GetCurrency() *string
	SetExtendMap(v string) *ChangeResellerConsumeAmountRequest
	GetExtendMap() *string
	SetOutBizId(v string) *ChangeResellerConsumeAmountRequest
	GetOutBizId() *string
	SetOwnerId(v int64) *ChangeResellerConsumeAmountRequest
	GetOwnerId() *int64
	SetSource(v string) *ChangeResellerConsumeAmountRequest
	GetSource() *string
}

type ChangeResellerConsumeAmountRequest struct {
	// The type of the consumption amount adjustment. Valid values: increase: The consumption amount increases because new consumption occurs. decrease: The consumption amount decreases because funds are added to the account. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// increase/decrease
	AdjustType *string `json:"AdjustType,omitempty" xml:"AdjustType,omitempty"`
	// The amount to be adjusted. Unit: CNY
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.00
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The type of the business.
	//
	// This parameter is required.
	//
	// example:
	//
	// quota_amount_adjust
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The type of the currency.
	//
	// This parameter is required.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The extended field of a message.
	//
	// example:
	//
	// {}
	ExtendMap *string `json:"ExtendMap,omitempty" xml:"ExtendMap,omitempty"`
	// The ID of the primary key for external business. The ID is used for idempotence verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1647396865
	OutBizId *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	// This parameter is required.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The source of the request. Specify the system name for the parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// system
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ChangeResellerConsumeAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResellerConsumeAmountRequest) GoString() string {
	return s.String()
}

func (s *ChangeResellerConsumeAmountRequest) GetAdjustType() *string {
	return s.AdjustType
}

func (s *ChangeResellerConsumeAmountRequest) GetAmount() *string {
	return s.Amount
}

func (s *ChangeResellerConsumeAmountRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ChangeResellerConsumeAmountRequest) GetCurrency() *string {
	return s.Currency
}

func (s *ChangeResellerConsumeAmountRequest) GetExtendMap() *string {
	return s.ExtendMap
}

func (s *ChangeResellerConsumeAmountRequest) GetOutBizId() *string {
	return s.OutBizId
}

func (s *ChangeResellerConsumeAmountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChangeResellerConsumeAmountRequest) GetSource() *string {
	return s.Source
}

func (s *ChangeResellerConsumeAmountRequest) SetAdjustType(v string) *ChangeResellerConsumeAmountRequest {
	s.AdjustType = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetAmount(v string) *ChangeResellerConsumeAmountRequest {
	s.Amount = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetBusinessType(v string) *ChangeResellerConsumeAmountRequest {
	s.BusinessType = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetCurrency(v string) *ChangeResellerConsumeAmountRequest {
	s.Currency = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetExtendMap(v string) *ChangeResellerConsumeAmountRequest {
	s.ExtendMap = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetOutBizId(v string) *ChangeResellerConsumeAmountRequest {
	s.OutBizId = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetOwnerId(v int64) *ChangeResellerConsumeAmountRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetSource(v string) *ChangeResellerConsumeAmountRequest {
	s.Source = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) Validate() error {
	return dara.Validate(s)
}
