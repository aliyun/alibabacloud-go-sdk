// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPayOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuyerId(v int64) *PayOrderRequest
	GetBuyerId() *int64
	SetEcIdAccountIds(v []*PayOrderRequestEcIdAccountIds) *PayOrderRequest
	GetEcIdAccountIds() []*PayOrderRequestEcIdAccountIds
	SetNbid(v string) *PayOrderRequest
	GetNbid() *string
	SetOrderId(v int64) *PayOrderRequest
	GetOrderId() *int64
	SetPaySubmitUid(v int64) *PayOrderRequest
	GetPaySubmitUid() *int64
	SetPayerId(v int64) *PayOrderRequest
	GetPayerId() *int64
	SetToken(v string) *PayOrderRequest
	GetToken() *string
}

type PayOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1178418164369806
	BuyerId        *int64                           `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	EcIdAccountIds []*PayOrderRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 26888345
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 260591304500425
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 201437655683478597
	PaySubmitUid *int64 `json:"PaySubmitUid,omitempty" xml:"PaySubmitUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 201437655683478597
	PayerId *int64 `json:"PayerId,omitempty" xml:"PayerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s PayOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s PayOrderRequest) GoString() string {
	return s.String()
}

func (s *PayOrderRequest) GetBuyerId() *int64 {
	return s.BuyerId
}

func (s *PayOrderRequest) GetEcIdAccountIds() []*PayOrderRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *PayOrderRequest) GetNbid() *string {
	return s.Nbid
}

func (s *PayOrderRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *PayOrderRequest) GetPaySubmitUid() *int64 {
	return s.PaySubmitUid
}

func (s *PayOrderRequest) GetPayerId() *int64 {
	return s.PayerId
}

func (s *PayOrderRequest) GetToken() *string {
	return s.Token
}

func (s *PayOrderRequest) SetBuyerId(v int64) *PayOrderRequest {
	s.BuyerId = &v
	return s
}

func (s *PayOrderRequest) SetEcIdAccountIds(v []*PayOrderRequestEcIdAccountIds) *PayOrderRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *PayOrderRequest) SetNbid(v string) *PayOrderRequest {
	s.Nbid = &v
	return s
}

func (s *PayOrderRequest) SetOrderId(v int64) *PayOrderRequest {
	s.OrderId = &v
	return s
}

func (s *PayOrderRequest) SetPaySubmitUid(v int64) *PayOrderRequest {
	s.PaySubmitUid = &v
	return s
}

func (s *PayOrderRequest) SetPayerId(v int64) *PayOrderRequest {
	s.PayerId = &v
	return s
}

func (s *PayOrderRequest) SetToken(v string) *PayOrderRequest {
	s.Token = &v
	return s
}

func (s *PayOrderRequest) Validate() error {
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PayOrderRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s PayOrderRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s PayOrderRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *PayOrderRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *PayOrderRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *PayOrderRequestEcIdAccountIds) SetAccountIds(v []*int64) *PayOrderRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *PayOrderRequestEcIdAccountIds) SetEcId(v string) *PayOrderRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *PayOrderRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
