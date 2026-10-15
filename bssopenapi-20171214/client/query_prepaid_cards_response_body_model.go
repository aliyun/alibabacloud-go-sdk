// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPrepaidCardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPrepaidCardsResponseBody
	GetCode() *string
	SetData(v *QueryPrepaidCardsResponseBodyData) *QueryPrepaidCardsResponseBody
	GetData() *QueryPrepaidCardsResponseBodyData
	SetMessage(v string) *QueryPrepaidCardsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPrepaidCardsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPrepaidCardsResponseBody
	GetSuccess() *bool
}

type QueryPrepaidCardsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryPrepaidCardsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7EA6C02D-06D0-4213-9C3B-E67910F7D1EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPrepaidCardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPrepaidCardsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPrepaidCardsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPrepaidCardsResponseBody) GetData() *QueryPrepaidCardsResponseBodyData {
	return s.Data
}

func (s *QueryPrepaidCardsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPrepaidCardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPrepaidCardsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPrepaidCardsResponseBody) SetCode(v string) *QueryPrepaidCardsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPrepaidCardsResponseBody) SetData(v *QueryPrepaidCardsResponseBodyData) *QueryPrepaidCardsResponseBody {
	s.Data = v
	return s
}

func (s *QueryPrepaidCardsResponseBody) SetMessage(v string) *QueryPrepaidCardsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPrepaidCardsResponseBody) SetRequestId(v string) *QueryPrepaidCardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPrepaidCardsResponseBody) SetSuccess(v bool) *QueryPrepaidCardsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPrepaidCardsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPrepaidCardsResponseBodyData struct {
	PrepaidCard []*QueryPrepaidCardsResponseBodyDataPrepaidCard `json:"PrepaidCard,omitempty" xml:"PrepaidCard,omitempty" type:"Repeated"`
}

func (s QueryPrepaidCardsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPrepaidCardsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPrepaidCardsResponseBodyData) GetPrepaidCard() []*QueryPrepaidCardsResponseBodyDataPrepaidCard {
	return s.PrepaidCard
}

func (s *QueryPrepaidCardsResponseBodyData) SetPrepaidCard(v []*QueryPrepaidCardsResponseBodyDataPrepaidCard) *QueryPrepaidCardsResponseBodyData {
	s.PrepaidCard = v
	return s
}

func (s *QueryPrepaidCardsResponseBodyData) Validate() error {
	if s.PrepaidCard != nil {
		for _, item := range s.PrepaidCard {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPrepaidCardsResponseBodyDataPrepaidCard struct {
	ApplicableProducts  *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	ApplicableScenarios *string `json:"ApplicableScenarios,omitempty" xml:"ApplicableScenarios,omitempty"`
	Balance             *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	EffectiveTime       *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpiryTime          *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	GrantedTime         *string `json:"GrantedTime,omitempty" xml:"GrantedTime,omitempty"`
	NominalValue        *string `json:"NominalValue,omitempty" xml:"NominalValue,omitempty"`
	PrepaidCardId       *int64  `json:"PrepaidCardId,omitempty" xml:"PrepaidCardId,omitempty"`
	PrepaidCardNo       *string `json:"PrepaidCardNo,omitempty" xml:"PrepaidCardNo,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryPrepaidCardsResponseBodyDataPrepaidCard) String() string {
	return dara.Prettify(s)
}

func (s QueryPrepaidCardsResponseBodyDataPrepaidCard) GoString() string {
	return s.String()
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) GetApplicableProducts() *string {
	return s.ApplicableProducts
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) GetApplicableScenarios() *string {
	return s.ApplicableScenarios
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) GetBalance() *string {
	return s.Balance
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) GetExpiryTime() *string {
	return s.ExpiryTime
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) GetGrantedTime() *string {
	return s.GrantedTime
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) GetNominalValue() *string {
	return s.NominalValue
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) GetPrepaidCardId() *int64 {
	return s.PrepaidCardId
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) GetPrepaidCardNo() *string {
	return s.PrepaidCardNo
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) GetStatus() *string {
	return s.Status
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetApplicableProducts(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.ApplicableProducts = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetApplicableScenarios(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.ApplicableScenarios = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetBalance(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.Balance = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetEffectiveTime(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.EffectiveTime = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetExpiryTime(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.ExpiryTime = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetGrantedTime(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.GrantedTime = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetNominalValue(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.NominalValue = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetPrepaidCardId(v int64) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.PrepaidCardId = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetPrepaidCardNo(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.PrepaidCardNo = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetStatus(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.Status = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) Validate() error {
	return dara.Validate(s)
}
