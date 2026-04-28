// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCssCreateOrderParam interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CssCreateOrderParam
	GetAgentId() *string
	SetAutoPay(v bool) *CssCreateOrderParam
	GetAutoPay() *bool
	SetAutoUseCoupon(v bool) *CssCreateOrderParam
	GetAutoUseCoupon() *bool
	SetBid(v string) *CssCreateOrderParam
	GetBid() *string
	SetBuyerId(v int64) *CssCreateOrderParam
	GetBuyerId() *int64
	SetCertificate(v string) *CssCreateOrderParam
	GetCertificate() *string
	SetChildId(v int64) *CssCreateOrderParam
	GetChildId() *int64
	SetCilentIp(v string) *CssCreateOrderParam
	GetCilentIp() *string
	SetCommodities(v []*CssInstanceCommodity) *CssCreateOrderParam
	GetCommodities() []*CssInstanceCommodity
	SetCreaterNick(v string) *CssCreateOrderParam
	GetCreaterNick() *string
	SetCssAuthRequestParam(v interface{}) *CssCreateOrderParam
	GetCssAuthRequestParam() interface{}
	SetFromApp(v string) *CssCreateOrderParam
	GetFromApp() *string
	SetLanguage(v string) *CssCreateOrderParam
	GetLanguage() *string
	SetMarketType(v int64) *CssCreateOrderParam
	GetMarketType() *int64
	SetMemo(v string) *CssCreateOrderParam
	GetMemo() *string
	SetOrderOrigin(v string) *CssCreateOrderParam
	GetOrderOrigin() *string
	SetOrderParams(v map[string]*string) *CssCreateOrderParam
	GetOrderParams() map[string]*string
	SetPayerId(v int64) *CssCreateOrderParam
	GetPayerId() *int64
	SetPlanGroupId(v int64) *CssCreateOrderParam
	GetPlanGroupId() *int64
	SetPlanId(v int64) *CssCreateOrderParam
	GetPlanId() *int64
	SetPlanInstanceId(v string) *CssCreateOrderParam
	GetPlanInstanceId() *string
	SetPromotionCode(v string) *CssCreateOrderParam
	GetPromotionCode() *string
	SetPromotionInputParam(v interface{}) *CssCreateOrderParam
	GetPromotionInputParam() interface{}
	SetRequestId(v string) *CssCreateOrderParam
	GetRequestId() *string
	SetSkipChannel(v bool) *CssCreateOrderParam
	GetSkipChannel() *bool
	SetToken(v string) *CssCreateOrderParam
	GetToken() *string
	SetTransientAccess(v interface{}) *CssCreateOrderParam
	GetTransientAccess() interface{}
	SetUmidToken(v string) *CssCreateOrderParam
	GetUmidToken() *string
	SetUserId(v int64) *CssCreateOrderParam
	GetUserId() *int64
}

type CssCreateOrderParam struct {
	AgentId             *string                 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AutoPay             *bool                   `json:"autoPay,omitempty" xml:"autoPay,omitempty"`
	AutoUseCoupon       *bool                   `json:"autoUseCoupon,omitempty" xml:"autoUseCoupon,omitempty"`
	Bid                 *string                 `json:"bid,omitempty" xml:"bid,omitempty"`
	BuyerId             *int64                  `json:"buyerId,omitempty" xml:"buyerId,omitempty"`
	Certificate         *string                 `json:"certificate,omitempty" xml:"certificate,omitempty"`
	ChildId             *int64                  `json:"childId,omitempty" xml:"childId,omitempty"`
	CilentIp            *string                 `json:"cilentIp,omitempty" xml:"cilentIp,omitempty"`
	Commodities         []*CssInstanceCommodity `json:"commodities,omitempty" xml:"commodities,omitempty" type:"Repeated"`
	CreaterNick         *string                 `json:"createrNick,omitempty" xml:"createrNick,omitempty"`
	CssAuthRequestParam interface{}             `json:"cssAuthRequestParam,omitempty" xml:"cssAuthRequestParam,omitempty"`
	FromApp             *string                 `json:"fromApp,omitempty" xml:"fromApp,omitempty"`
	Language            *string                 `json:"language,omitempty" xml:"language,omitempty"`
	MarketType          *int64                  `json:"marketType,omitempty" xml:"marketType,omitempty"`
	Memo                *string                 `json:"memo,omitempty" xml:"memo,omitempty"`
	OrderOrigin         *string                 `json:"orderOrigin,omitempty" xml:"orderOrigin,omitempty"`
	OrderParams         map[string]*string      `json:"orderParams,omitempty" xml:"orderParams,omitempty"`
	PayerId             *int64                  `json:"payerId,omitempty" xml:"payerId,omitempty"`
	PlanGroupId         *int64                  `json:"planGroupId,omitempty" xml:"planGroupId,omitempty"`
	PlanId              *int64                  `json:"planId,omitempty" xml:"planId,omitempty"`
	PlanInstanceId      *string                 `json:"planInstanceId,omitempty" xml:"planInstanceId,omitempty"`
	PromotionCode       *string                 `json:"promotionCode,omitempty" xml:"promotionCode,omitempty"`
	PromotionInputParam interface{}             `json:"promotionInputParam,omitempty" xml:"promotionInputParam,omitempty"`
	RequestId           *string                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SkipChannel         *bool                   `json:"skipChannel,omitempty" xml:"skipChannel,omitempty"`
	Token               *string                 `json:"token,omitempty" xml:"token,omitempty"`
	TransientAccess     interface{}             `json:"transientAccess,omitempty" xml:"transientAccess,omitempty"`
	UmidToken           *string                 `json:"umidToken,omitempty" xml:"umidToken,omitempty"`
	UserId              *int64                  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CssCreateOrderParam) String() string {
	return dara.Prettify(s)
}

func (s CssCreateOrderParam) GoString() string {
	return s.String()
}

func (s *CssCreateOrderParam) GetAgentId() *string {
	return s.AgentId
}

func (s *CssCreateOrderParam) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CssCreateOrderParam) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CssCreateOrderParam) GetBid() *string {
	return s.Bid
}

func (s *CssCreateOrderParam) GetBuyerId() *int64 {
	return s.BuyerId
}

func (s *CssCreateOrderParam) GetCertificate() *string {
	return s.Certificate
}

func (s *CssCreateOrderParam) GetChildId() *int64 {
	return s.ChildId
}

func (s *CssCreateOrderParam) GetCilentIp() *string {
	return s.CilentIp
}

func (s *CssCreateOrderParam) GetCommodities() []*CssInstanceCommodity {
	return s.Commodities
}

func (s *CssCreateOrderParam) GetCreaterNick() *string {
	return s.CreaterNick
}

func (s *CssCreateOrderParam) GetCssAuthRequestParam() interface{} {
	return s.CssAuthRequestParam
}

func (s *CssCreateOrderParam) GetFromApp() *string {
	return s.FromApp
}

func (s *CssCreateOrderParam) GetLanguage() *string {
	return s.Language
}

func (s *CssCreateOrderParam) GetMarketType() *int64 {
	return s.MarketType
}

func (s *CssCreateOrderParam) GetMemo() *string {
	return s.Memo
}

func (s *CssCreateOrderParam) GetOrderOrigin() *string {
	return s.OrderOrigin
}

func (s *CssCreateOrderParam) GetOrderParams() map[string]*string {
	return s.OrderParams
}

func (s *CssCreateOrderParam) GetPayerId() *int64 {
	return s.PayerId
}

func (s *CssCreateOrderParam) GetPlanGroupId() *int64 {
	return s.PlanGroupId
}

func (s *CssCreateOrderParam) GetPlanId() *int64 {
	return s.PlanId
}

func (s *CssCreateOrderParam) GetPlanInstanceId() *string {
	return s.PlanInstanceId
}

func (s *CssCreateOrderParam) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CssCreateOrderParam) GetPromotionInputParam() interface{} {
	return s.PromotionInputParam
}

func (s *CssCreateOrderParam) GetRequestId() *string {
	return s.RequestId
}

func (s *CssCreateOrderParam) GetSkipChannel() *bool {
	return s.SkipChannel
}

func (s *CssCreateOrderParam) GetToken() *string {
	return s.Token
}

func (s *CssCreateOrderParam) GetTransientAccess() interface{} {
	return s.TransientAccess
}

func (s *CssCreateOrderParam) GetUmidToken() *string {
	return s.UmidToken
}

func (s *CssCreateOrderParam) GetUserId() *int64 {
	return s.UserId
}

func (s *CssCreateOrderParam) SetAgentId(v string) *CssCreateOrderParam {
	s.AgentId = &v
	return s
}

func (s *CssCreateOrderParam) SetAutoPay(v bool) *CssCreateOrderParam {
	s.AutoPay = &v
	return s
}

func (s *CssCreateOrderParam) SetAutoUseCoupon(v bool) *CssCreateOrderParam {
	s.AutoUseCoupon = &v
	return s
}

func (s *CssCreateOrderParam) SetBid(v string) *CssCreateOrderParam {
	s.Bid = &v
	return s
}

func (s *CssCreateOrderParam) SetBuyerId(v int64) *CssCreateOrderParam {
	s.BuyerId = &v
	return s
}

func (s *CssCreateOrderParam) SetCertificate(v string) *CssCreateOrderParam {
	s.Certificate = &v
	return s
}

func (s *CssCreateOrderParam) SetChildId(v int64) *CssCreateOrderParam {
	s.ChildId = &v
	return s
}

func (s *CssCreateOrderParam) SetCilentIp(v string) *CssCreateOrderParam {
	s.CilentIp = &v
	return s
}

func (s *CssCreateOrderParam) SetCommodities(v []*CssInstanceCommodity) *CssCreateOrderParam {
	s.Commodities = v
	return s
}

func (s *CssCreateOrderParam) SetCreaterNick(v string) *CssCreateOrderParam {
	s.CreaterNick = &v
	return s
}

func (s *CssCreateOrderParam) SetCssAuthRequestParam(v interface{}) *CssCreateOrderParam {
	s.CssAuthRequestParam = v
	return s
}

func (s *CssCreateOrderParam) SetFromApp(v string) *CssCreateOrderParam {
	s.FromApp = &v
	return s
}

func (s *CssCreateOrderParam) SetLanguage(v string) *CssCreateOrderParam {
	s.Language = &v
	return s
}

func (s *CssCreateOrderParam) SetMarketType(v int64) *CssCreateOrderParam {
	s.MarketType = &v
	return s
}

func (s *CssCreateOrderParam) SetMemo(v string) *CssCreateOrderParam {
	s.Memo = &v
	return s
}

func (s *CssCreateOrderParam) SetOrderOrigin(v string) *CssCreateOrderParam {
	s.OrderOrigin = &v
	return s
}

func (s *CssCreateOrderParam) SetOrderParams(v map[string]*string) *CssCreateOrderParam {
	s.OrderParams = v
	return s
}

func (s *CssCreateOrderParam) SetPayerId(v int64) *CssCreateOrderParam {
	s.PayerId = &v
	return s
}

func (s *CssCreateOrderParam) SetPlanGroupId(v int64) *CssCreateOrderParam {
	s.PlanGroupId = &v
	return s
}

func (s *CssCreateOrderParam) SetPlanId(v int64) *CssCreateOrderParam {
	s.PlanId = &v
	return s
}

func (s *CssCreateOrderParam) SetPlanInstanceId(v string) *CssCreateOrderParam {
	s.PlanInstanceId = &v
	return s
}

func (s *CssCreateOrderParam) SetPromotionCode(v string) *CssCreateOrderParam {
	s.PromotionCode = &v
	return s
}

func (s *CssCreateOrderParam) SetPromotionInputParam(v interface{}) *CssCreateOrderParam {
	s.PromotionInputParam = v
	return s
}

func (s *CssCreateOrderParam) SetRequestId(v string) *CssCreateOrderParam {
	s.RequestId = &v
	return s
}

func (s *CssCreateOrderParam) SetSkipChannel(v bool) *CssCreateOrderParam {
	s.SkipChannel = &v
	return s
}

func (s *CssCreateOrderParam) SetToken(v string) *CssCreateOrderParam {
	s.Token = &v
	return s
}

func (s *CssCreateOrderParam) SetTransientAccess(v interface{}) *CssCreateOrderParam {
	s.TransientAccess = v
	return s
}

func (s *CssCreateOrderParam) SetUmidToken(v string) *CssCreateOrderParam {
	s.UmidToken = &v
	return s
}

func (s *CssCreateOrderParam) SetUserId(v int64) *CssCreateOrderParam {
	s.UserId = &v
	return s
}

func (s *CssCreateOrderParam) Validate() error {
	if s.Commodities != nil {
		for _, item := range s.Commodities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
