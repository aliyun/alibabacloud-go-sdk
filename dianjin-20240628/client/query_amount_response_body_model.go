// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAmountResponseBody
	GetCode() *string
	SetData(v *QueryAmountResponseBodyData) *QueryAmountResponseBody
	GetData() *QueryAmountResponseBodyData
	SetMessage(v string) *QueryAmountResponseBody
	GetMessage() *string
	SetRetryAble(v bool) *QueryAmountResponseBody
	GetRetryAble() *bool
	SetSuccess(v bool) *QueryAmountResponseBody
	GetSuccess() *bool
}

type QueryAmountResponseBody struct {
	Code      *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data      *QueryAmountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                      `json:"message,omitempty" xml:"message,omitempty"`
	RetryAble *bool                        `json:"retryAble,omitempty" xml:"retryAble,omitempty"`
	Success   *bool                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAmountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAmountResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAmountResponseBody) GetData() *QueryAmountResponseBodyData {
	return s.Data
}

func (s *QueryAmountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAmountResponseBody) GetRetryAble() *bool {
	return s.RetryAble
}

func (s *QueryAmountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAmountResponseBody) SetCode(v string) *QueryAmountResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAmountResponseBody) SetData(v *QueryAmountResponseBodyData) *QueryAmountResponseBody {
	s.Data = v
	return s
}

func (s *QueryAmountResponseBody) SetMessage(v string) *QueryAmountResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAmountResponseBody) SetRetryAble(v bool) *QueryAmountResponseBody {
	s.RetryAble = &v
	return s
}

func (s *QueryAmountResponseBody) SetSuccess(v bool) *QueryAmountResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAmountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAmountResponseBodyData struct {
	EndDate   *string                             `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Items     []*QueryAmountResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	ScopeNote *string                             `json:"scopeNote,omitempty" xml:"scopeNote,omitempty"`
	StartDate *string                             `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Total     *QueryAmountResponseBodyDataTotal   `json:"total,omitempty" xml:"total,omitempty" type:"Struct"`
}

func (s QueryAmountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAmountResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAmountResponseBodyData) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryAmountResponseBodyData) GetItems() []*QueryAmountResponseBodyDataItems {
	return s.Items
}

func (s *QueryAmountResponseBodyData) GetScopeNote() *string {
	return s.ScopeNote
}

func (s *QueryAmountResponseBodyData) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryAmountResponseBodyData) GetTotal() *QueryAmountResponseBodyDataTotal {
	return s.Total
}

func (s *QueryAmountResponseBodyData) SetEndDate(v string) *QueryAmountResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *QueryAmountResponseBodyData) SetItems(v []*QueryAmountResponseBodyDataItems) *QueryAmountResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryAmountResponseBodyData) SetScopeNote(v string) *QueryAmountResponseBodyData {
	s.ScopeNote = &v
	return s
}

func (s *QueryAmountResponseBodyData) SetStartDate(v string) *QueryAmountResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *QueryAmountResponseBodyData) SetTotal(v *QueryAmountResponseBodyDataTotal) *QueryAmountResponseBodyData {
	s.Total = v
	return s
}

func (s *QueryAmountResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Total != nil {
		if err := s.Total.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAmountResponseBodyDataItems struct {
	AliyunUid   *string `json:"aliyunUid,omitempty" xml:"aliyunUid,omitempty"`
	Amount      *string `json:"amount,omitempty" xml:"amount,omitempty"`
	AmountRatio *string `json:"amountRatio,omitempty" xml:"amountRatio,omitempty"`
	ListFee     *string `json:"listFee,omitempty" xml:"listFee,omitempty"`
	Price       *string `json:"price,omitempty" xml:"price,omitempty"`
	Tier        *string `json:"tier,omitempty" xml:"tier,omitempty"`
	TotalAmount *string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s QueryAmountResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QueryAmountResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryAmountResponseBodyDataItems) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *QueryAmountResponseBodyDataItems) GetAmount() *string {
	return s.Amount
}

func (s *QueryAmountResponseBodyDataItems) GetAmountRatio() *string {
	return s.AmountRatio
}

func (s *QueryAmountResponseBodyDataItems) GetListFee() *string {
	return s.ListFee
}

func (s *QueryAmountResponseBodyDataItems) GetPrice() *string {
	return s.Price
}

func (s *QueryAmountResponseBodyDataItems) GetTier() *string {
	return s.Tier
}

func (s *QueryAmountResponseBodyDataItems) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *QueryAmountResponseBodyDataItems) SetAliyunUid(v string) *QueryAmountResponseBodyDataItems {
	s.AliyunUid = &v
	return s
}

func (s *QueryAmountResponseBodyDataItems) SetAmount(v string) *QueryAmountResponseBodyDataItems {
	s.Amount = &v
	return s
}

func (s *QueryAmountResponseBodyDataItems) SetAmountRatio(v string) *QueryAmountResponseBodyDataItems {
	s.AmountRatio = &v
	return s
}

func (s *QueryAmountResponseBodyDataItems) SetListFee(v string) *QueryAmountResponseBodyDataItems {
	s.ListFee = &v
	return s
}

func (s *QueryAmountResponseBodyDataItems) SetPrice(v string) *QueryAmountResponseBodyDataItems {
	s.Price = &v
	return s
}

func (s *QueryAmountResponseBodyDataItems) SetTier(v string) *QueryAmountResponseBodyDataItems {
	s.Tier = &v
	return s
}

func (s *QueryAmountResponseBodyDataItems) SetTotalAmount(v string) *QueryAmountResponseBodyDataItems {
	s.TotalAmount = &v
	return s
}

func (s *QueryAmountResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}

type QueryAmountResponseBodyDataTotal struct {
	Amount      *string `json:"amount,omitempty" xml:"amount,omitempty"`
	ListFee     *string `json:"listFee,omitempty" xml:"listFee,omitempty"`
	TotalAmount *string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s QueryAmountResponseBodyDataTotal) String() string {
	return dara.Prettify(s)
}

func (s QueryAmountResponseBodyDataTotal) GoString() string {
	return s.String()
}

func (s *QueryAmountResponseBodyDataTotal) GetAmount() *string {
	return s.Amount
}

func (s *QueryAmountResponseBodyDataTotal) GetListFee() *string {
	return s.ListFee
}

func (s *QueryAmountResponseBodyDataTotal) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *QueryAmountResponseBodyDataTotal) SetAmount(v string) *QueryAmountResponseBodyDataTotal {
	s.Amount = &v
	return s
}

func (s *QueryAmountResponseBodyDataTotal) SetListFee(v string) *QueryAmountResponseBodyDataTotal {
	s.ListFee = &v
	return s
}

func (s *QueryAmountResponseBodyDataTotal) SetTotalAmount(v string) *QueryAmountResponseBodyDataTotal {
	s.TotalAmount = &v
	return s
}

func (s *QueryAmountResponseBodyDataTotal) Validate() error {
	return dara.Validate(s)
}
