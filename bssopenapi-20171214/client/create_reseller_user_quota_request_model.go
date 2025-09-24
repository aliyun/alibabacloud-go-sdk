// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResellerUserQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v string) *CreateResellerUserQuotaRequest
	GetAmount() *string
	SetCurrency(v string) *CreateResellerUserQuotaRequest
	GetCurrency() *string
	SetOutBizId(v string) *CreateResellerUserQuotaRequest
	GetOutBizId() *string
	SetOwnerId(v int64) *CreateResellerUserQuotaRequest
	GetOwnerId() *int64
}

type CreateResellerUserQuotaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 7ed0bedc-056e-4a71-9249-4581615c028f
	OutBizId *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	// This parameter is required.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateResellerUserQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResellerUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *CreateResellerUserQuotaRequest) GetAmount() *string {
	return s.Amount
}

func (s *CreateResellerUserQuotaRequest) GetCurrency() *string {
	return s.Currency
}

func (s *CreateResellerUserQuotaRequest) GetOutBizId() *string {
	return s.OutBizId
}

func (s *CreateResellerUserQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateResellerUserQuotaRequest) SetAmount(v string) *CreateResellerUserQuotaRequest {
	s.Amount = &v
	return s
}

func (s *CreateResellerUserQuotaRequest) SetCurrency(v string) *CreateResellerUserQuotaRequest {
	s.Currency = &v
	return s
}

func (s *CreateResellerUserQuotaRequest) SetOutBizId(v string) *CreateResellerUserQuotaRequest {
	s.OutBizId = &v
	return s
}

func (s *CreateResellerUserQuotaRequest) SetOwnerId(v int64) *CreateResellerUserQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateResellerUserQuotaRequest) Validate() error {
	return dara.Validate(s)
}
