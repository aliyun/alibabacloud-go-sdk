// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCssProduce interface {
	dara.Model
	String() string
	GoString() string
	SetBid(v string) *CssProduce
	GetBid() *string
	SetBuyerId(v int64) *CssProduce
	GetBuyerId() *int64
	SetChildId(v int64) *CssProduce
	GetChildId() *int64
	SetFromApp(v string) *CssProduce
	GetFromApp() *string
	SetOrderId(v int64) *CssProduce
	GetOrderId() *int64
	SetPayerId(v int64) *CssProduce
	GetPayerId() *int64
	SetPurchases(v []*CssPurchase) *CssProduce
	GetPurchases() []*CssPurchase
	SetRequestId(v string) *CssProduce
	GetRequestId() *string
	SetSkipChannel(v bool) *CssProduce
	GetSkipChannel() *bool
	SetToken(v string) *CssProduce
	GetToken() *string
	SetUserId(v int64) *CssProduce
	GetUserId() *int64
}

type CssProduce struct {
	Bid         *string        `json:"bid,omitempty" xml:"bid,omitempty"`
	BuyerId     *int64         `json:"buyerId,omitempty" xml:"buyerId,omitempty"`
	ChildId     *int64         `json:"childId,omitempty" xml:"childId,omitempty"`
	FromApp     *string        `json:"fromApp,omitempty" xml:"fromApp,omitempty"`
	OrderId     *int64         `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PayerId     *int64         `json:"payerId,omitempty" xml:"payerId,omitempty"`
	Purchases   []*CssPurchase `json:"purchases,omitempty" xml:"purchases,omitempty" type:"Repeated"`
	RequestId   *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SkipChannel *bool          `json:"skipChannel,omitempty" xml:"skipChannel,omitempty"`
	Token       *string        `json:"token,omitempty" xml:"token,omitempty"`
	UserId      *int64         `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CssProduce) String() string {
	return dara.Prettify(s)
}

func (s CssProduce) GoString() string {
	return s.String()
}

func (s *CssProduce) GetBid() *string {
	return s.Bid
}

func (s *CssProduce) GetBuyerId() *int64 {
	return s.BuyerId
}

func (s *CssProduce) GetChildId() *int64 {
	return s.ChildId
}

func (s *CssProduce) GetFromApp() *string {
	return s.FromApp
}

func (s *CssProduce) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CssProduce) GetPayerId() *int64 {
	return s.PayerId
}

func (s *CssProduce) GetPurchases() []*CssPurchase {
	return s.Purchases
}

func (s *CssProduce) GetRequestId() *string {
	return s.RequestId
}

func (s *CssProduce) GetSkipChannel() *bool {
	return s.SkipChannel
}

func (s *CssProduce) GetToken() *string {
	return s.Token
}

func (s *CssProduce) GetUserId() *int64 {
	return s.UserId
}

func (s *CssProduce) SetBid(v string) *CssProduce {
	s.Bid = &v
	return s
}

func (s *CssProduce) SetBuyerId(v int64) *CssProduce {
	s.BuyerId = &v
	return s
}

func (s *CssProduce) SetChildId(v int64) *CssProduce {
	s.ChildId = &v
	return s
}

func (s *CssProduce) SetFromApp(v string) *CssProduce {
	s.FromApp = &v
	return s
}

func (s *CssProduce) SetOrderId(v int64) *CssProduce {
	s.OrderId = &v
	return s
}

func (s *CssProduce) SetPayerId(v int64) *CssProduce {
	s.PayerId = &v
	return s
}

func (s *CssProduce) SetPurchases(v []*CssPurchase) *CssProduce {
	s.Purchases = v
	return s
}

func (s *CssProduce) SetRequestId(v string) *CssProduce {
	s.RequestId = &v
	return s
}

func (s *CssProduce) SetSkipChannel(v bool) *CssProduce {
	s.SkipChannel = &v
	return s
}

func (s *CssProduce) SetToken(v string) *CssProduce {
	s.Token = &v
	return s
}

func (s *CssProduce) SetUserId(v int64) *CssProduce {
	s.UserId = &v
	return s
}

func (s *CssProduce) Validate() error {
	if s.Purchases != nil {
		for _, item := range s.Purchases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
