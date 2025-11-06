// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainRealTimePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *QueryDomainRealTimePriceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryDomainRealTimePriceResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *QueryDomainRealTimePriceResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryDomainRealTimePriceResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *QueryDomainRealTimePriceResponseBody
	GetHttpStatusCode() *int32
	SetModule(v *QueryDomainRealTimePriceResponseBodyModule) *QueryDomainRealTimePriceResponseBody
	GetModule() *QueryDomainRealTimePriceResponseBodyModule
	SetRequestId(v string) *QueryDomainRealTimePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDomainRealTimePriceResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *QueryDomainRealTimePriceResponseBody
	GetSynchro() *bool
}

type QueryDomainRealTimePriceResponseBody struct {
	// example:
	//
	// 4000
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// 非法参数
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 3002
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 非法参数
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         *QueryDomainRealTimePriceResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// A4A4F72C-8E8E-19D2-BCC1-21E41D334A75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *bool   `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryDomainRealTimePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainRealTimePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainRealTimePriceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryDomainRealTimePriceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryDomainRealTimePriceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryDomainRealTimePriceResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryDomainRealTimePriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryDomainRealTimePriceResponseBody) GetModule() *QueryDomainRealTimePriceResponseBodyModule {
	return s.Module
}

func (s *QueryDomainRealTimePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainRealTimePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDomainRealTimePriceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryDomainRealTimePriceResponseBody) SetDynamicCode(v string) *QueryDomainRealTimePriceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBody) SetDynamicMessage(v string) *QueryDomainRealTimePriceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBody) SetErrorCode(v string) *QueryDomainRealTimePriceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBody) SetErrorMsg(v string) *QueryDomainRealTimePriceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBody) SetHttpStatusCode(v int32) *QueryDomainRealTimePriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBody) SetModule(v *QueryDomainRealTimePriceResponseBodyModule) *QueryDomainRealTimePriceResponseBody {
	s.Module = v
	return s
}

func (s *QueryDomainRealTimePriceResponseBody) SetRequestId(v string) *QueryDomainRealTimePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBody) SetSuccess(v bool) *QueryDomainRealTimePriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBody) SetSynchro(v bool) *QueryDomainRealTimePriceResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainRealTimePriceResponseBodyModule struct {
	DomainPrices []*QueryDomainRealTimePriceResponseBodyModuleDomainPrices `json:"DomainPrices,omitempty" xml:"DomainPrices,omitempty" type:"Repeated"`
}

func (s QueryDomainRealTimePriceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainRealTimePriceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryDomainRealTimePriceResponseBodyModule) GetDomainPrices() []*QueryDomainRealTimePriceResponseBodyModuleDomainPrices {
	return s.DomainPrices
}

func (s *QueryDomainRealTimePriceResponseBodyModule) SetDomainPrices(v []*QueryDomainRealTimePriceResponseBodyModuleDomainPrices) *QueryDomainRealTimePriceResponseBodyModule {
	s.DomainPrices = v
	return s
}

func (s *QueryDomainRealTimePriceResponseBodyModule) Validate() error {
	if s.DomainPrices != nil {
		for _, item := range s.DomainPrices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDomainRealTimePriceResponseBodyModuleDomainPrices struct {
	// example:
	//
	// renew
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// 1
	Avail *int32 `json:"Avail,omitempty" xml:"Avail,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// xxx.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1
	Period  *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	Premium *bool  `json:"Premium,omitempty" xml:"Premium,omitempty"`
	// example:
	//
	// 16.22
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// 不可续费
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s QueryDomainRealTimePriceResponseBodyModuleDomainPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainRealTimePriceResponseBodyModuleDomainPrices) GoString() string {
	return s.String()
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) GetAction() *string {
	return s.Action
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) GetAvail() *int32 {
	return s.Avail
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) GetCurrency() *string {
	return s.Currency
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) GetPeriod() *int32 {
	return s.Period
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) GetPremium() *bool {
	return s.Premium
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) GetPrice() *float64 {
	return s.Price
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) GetReason() *string {
	return s.Reason
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) SetAction(v string) *QueryDomainRealTimePriceResponseBodyModuleDomainPrices {
	s.Action = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) SetAvail(v int32) *QueryDomainRealTimePriceResponseBodyModuleDomainPrices {
	s.Avail = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) SetCurrency(v string) *QueryDomainRealTimePriceResponseBodyModuleDomainPrices {
	s.Currency = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) SetDomainName(v string) *QueryDomainRealTimePriceResponseBodyModuleDomainPrices {
	s.DomainName = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) SetPeriod(v int32) *QueryDomainRealTimePriceResponseBodyModuleDomainPrices {
	s.Period = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) SetPremium(v bool) *QueryDomainRealTimePriceResponseBodyModuleDomainPrices {
	s.Premium = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) SetPrice(v float64) *QueryDomainRealTimePriceResponseBodyModuleDomainPrices {
	s.Price = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) SetReason(v string) *QueryDomainRealTimePriceResponseBodyModuleDomainPrices {
	s.Reason = &v
	return s
}

func (s *QueryDomainRealTimePriceResponseBodyModuleDomainPrices) Validate() error {
	return dara.Validate(s)
}
