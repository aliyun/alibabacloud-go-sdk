// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvail(v string) *CheckDomainResponseBody
	GetAvail() *string
	SetDomainName(v string) *CheckDomainResponseBody
	GetDomainName() *string
	SetDynamicCheck(v bool) *CheckDomainResponseBody
	GetDynamicCheck() *bool
	SetPremium(v string) *CheckDomainResponseBody
	GetPremium() *string
	SetPrice(v int64) *CheckDomainResponseBody
	GetPrice() *int64
	SetReason(v string) *CheckDomainResponseBody
	GetReason() *string
	SetRequestId(v string) *CheckDomainResponseBody
	GetRequestId() *string
	SetStaticPriceInfo(v *CheckDomainResponseBodyStaticPriceInfo) *CheckDomainResponseBody
	GetStaticPriceInfo() *CheckDomainResponseBodyStaticPriceInfo
}

type CheckDomainResponseBody struct {
	// example:
	//
	// 1
	Avail *string `json:"Avail,omitempty" xml:"Avail,omitempty"`
	// example:
	//
	// test**.xin
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// true
	DynamicCheck *bool `json:"DynamicCheck,omitempty" xml:"DynamicCheck,omitempty"`
	// example:
	//
	// true
	Premium *string `json:"Premium,omitempty" xml:"Premium,omitempty"`
	// example:
	//
	// 1286
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// In use
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// BA7A4FD4-EB9A-4A20-BB0C-9AEB15634DC1
	RequestId       *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StaticPriceInfo *CheckDomainResponseBodyStaticPriceInfo `json:"StaticPriceInfo,omitempty" xml:"StaticPriceInfo,omitempty" type:"Struct"`
}

func (s CheckDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBody) GetAvail() *string {
	return s.Avail
}

func (s *CheckDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *CheckDomainResponseBody) GetDynamicCheck() *bool {
	return s.DynamicCheck
}

func (s *CheckDomainResponseBody) GetPremium() *string {
	return s.Premium
}

func (s *CheckDomainResponseBody) GetPrice() *int64 {
	return s.Price
}

func (s *CheckDomainResponseBody) GetReason() *string {
	return s.Reason
}

func (s *CheckDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDomainResponseBody) GetStaticPriceInfo() *CheckDomainResponseBodyStaticPriceInfo {
	return s.StaticPriceInfo
}

func (s *CheckDomainResponseBody) SetAvail(v string) *CheckDomainResponseBody {
	s.Avail = &v
	return s
}

func (s *CheckDomainResponseBody) SetDomainName(v string) *CheckDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *CheckDomainResponseBody) SetDynamicCheck(v bool) *CheckDomainResponseBody {
	s.DynamicCheck = &v
	return s
}

func (s *CheckDomainResponseBody) SetPremium(v string) *CheckDomainResponseBody {
	s.Premium = &v
	return s
}

func (s *CheckDomainResponseBody) SetPrice(v int64) *CheckDomainResponseBody {
	s.Price = &v
	return s
}

func (s *CheckDomainResponseBody) SetReason(v string) *CheckDomainResponseBody {
	s.Reason = &v
	return s
}

func (s *CheckDomainResponseBody) SetRequestId(v string) *CheckDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDomainResponseBody) SetStaticPriceInfo(v *CheckDomainResponseBodyStaticPriceInfo) *CheckDomainResponseBody {
	s.StaticPriceInfo = v
	return s
}

func (s *CheckDomainResponseBody) Validate() error {
	if s.StaticPriceInfo != nil {
		if err := s.StaticPriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckDomainResponseBodyStaticPriceInfo struct {
	PriceInfo []*CheckDomainResponseBodyStaticPriceInfoPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Repeated"`
}

func (s CheckDomainResponseBodyStaticPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainResponseBodyStaticPriceInfo) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBodyStaticPriceInfo) GetPriceInfo() []*CheckDomainResponseBodyStaticPriceInfoPriceInfo {
	return s.PriceInfo
}

func (s *CheckDomainResponseBodyStaticPriceInfo) SetPriceInfo(v []*CheckDomainResponseBodyStaticPriceInfoPriceInfo) *CheckDomainResponseBodyStaticPriceInfo {
	s.PriceInfo = v
	return s
}

func (s *CheckDomainResponseBodyStaticPriceInfo) Validate() error {
	if s.PriceInfo != nil {
		for _, item := range s.PriceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckDomainResponseBodyStaticPriceInfoPriceInfo struct {
	Action *string  `json:"action,omitempty" xml:"action,omitempty"`
	Money  *float64 `json:"money,omitempty" xml:"money,omitempty"`
	Period *int64   `json:"period,omitempty" xml:"period,omitempty"`
}

func (s CheckDomainResponseBodyStaticPriceInfoPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainResponseBodyStaticPriceInfoPriceInfo) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBodyStaticPriceInfoPriceInfo) GetAction() *string {
	return s.Action
}

func (s *CheckDomainResponseBodyStaticPriceInfoPriceInfo) GetMoney() *float64 {
	return s.Money
}

func (s *CheckDomainResponseBodyStaticPriceInfoPriceInfo) GetPeriod() *int64 {
	return s.Period
}

func (s *CheckDomainResponseBodyStaticPriceInfoPriceInfo) SetAction(v string) *CheckDomainResponseBodyStaticPriceInfoPriceInfo {
	s.Action = &v
	return s
}

func (s *CheckDomainResponseBodyStaticPriceInfoPriceInfo) SetMoney(v float64) *CheckDomainResponseBodyStaticPriceInfoPriceInfo {
	s.Money = &v
	return s
}

func (s *CheckDomainResponseBodyStaticPriceInfoPriceInfo) SetPeriod(v int64) *CheckDomainResponseBodyStaticPriceInfoPriceInfo {
	s.Period = &v
	return s
}

func (s *CheckDomainResponseBodyStaticPriceInfoPriceInfo) Validate() error {
	return dara.Validate(s)
}
