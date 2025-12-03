// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAliasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *SendByAliasShrinkRequest
	GetAlias() *string
	SetAliasType(v string) *SendByAliasShrinkRequest
	GetAliasType() *string
	SetAndroidPayloadShrink(v string) *SendByAliasShrinkRequest
	GetAndroidPayloadShrink() *string
	SetAndroidShortPayloadShrink(v string) *SendByAliasShrinkRequest
	GetAndroidShortPayloadShrink() *string
	SetChannelPropertiesShrink(v string) *SendByAliasShrinkRequest
	GetChannelPropertiesShrink() *string
	SetDescription(v string) *SendByAliasShrinkRequest
	GetDescription() *string
	SetHarmonyPayloadShrink(v string) *SendByAliasShrinkRequest
	GetHarmonyPayloadShrink() *string
	SetIosPayloadShrink(v string) *SendByAliasShrinkRequest
	GetIosPayloadShrink() *string
	SetPolicyShrink(v string) *SendByAliasShrinkRequest
	GetPolicyShrink() *string
	SetProductionMode(v bool) *SendByAliasShrinkRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByAliasShrinkRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByAliasShrinkRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByAliasShrinkRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByAliasShrinkRequest
	GetCallbackParams() *string
}

type SendByAliasShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Alias                     *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AliasType                 *string `json:"AliasType,omitempty" xml:"AliasType,omitempty"`
	AndroidPayloadShrink      *string `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayloadShrink *string `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelPropertiesShrink   *string `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HarmonyPayloadShrink      *string `json:"HarmonyPayload,omitempty" xml:"HarmonyPayload,omitempty"`
	IosPayloadShrink          *string `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	PolicyShrink              *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByAliasShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByAliasShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByAliasShrinkRequest) GetAlias() *string {
	return s.Alias
}

func (s *SendByAliasShrinkRequest) GetAliasType() *string {
	return s.AliasType
}

func (s *SendByAliasShrinkRequest) GetAndroidPayloadShrink() *string {
	return s.AndroidPayloadShrink
}

func (s *SendByAliasShrinkRequest) GetAndroidShortPayloadShrink() *string {
	return s.AndroidShortPayloadShrink
}

func (s *SendByAliasShrinkRequest) GetChannelPropertiesShrink() *string {
	return s.ChannelPropertiesShrink
}

func (s *SendByAliasShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByAliasShrinkRequest) GetHarmonyPayloadShrink() *string {
	return s.HarmonyPayloadShrink
}

func (s *SendByAliasShrinkRequest) GetIosPayloadShrink() *string {
	return s.IosPayloadShrink
}

func (s *SendByAliasShrinkRequest) GetPolicyShrink() *string {
	return s.PolicyShrink
}

func (s *SendByAliasShrinkRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByAliasShrinkRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByAliasShrinkRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByAliasShrinkRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByAliasShrinkRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByAliasShrinkRequest) SetAlias(v string) *SendByAliasShrinkRequest {
	s.Alias = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetAliasType(v string) *SendByAliasShrinkRequest {
	s.AliasType = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetAndroidPayloadShrink(v string) *SendByAliasShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetAndroidShortPayloadShrink(v string) *SendByAliasShrinkRequest {
	s.AndroidShortPayloadShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetChannelPropertiesShrink(v string) *SendByAliasShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetDescription(v string) *SendByAliasShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetHarmonyPayloadShrink(v string) *SendByAliasShrinkRequest {
	s.HarmonyPayloadShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetIosPayloadShrink(v string) *SendByAliasShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetPolicyShrink(v string) *SendByAliasShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetProductionMode(v bool) *SendByAliasShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetReceiptType(v int32) *SendByAliasShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetReceiptUrl(v string) *SendByAliasShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetThirdPartyId(v string) *SendByAliasShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetCallbackParams(v string) *SendByAliasShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByAliasShrinkRequest) Validate() error {
	return dara.Validate(s)
}
