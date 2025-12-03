// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidPayloadShrink(v string) *SendByAppShrinkRequest
	GetAndroidPayloadShrink() *string
	SetAndroidShortPayloadShrink(v string) *SendByAppShrinkRequest
	GetAndroidShortPayloadShrink() *string
	SetChannelPropertiesShrink(v string) *SendByAppShrinkRequest
	GetChannelPropertiesShrink() *string
	SetDescription(v string) *SendByAppShrinkRequest
	GetDescription() *string
	SetHarmonyPayloadShrink(v string) *SendByAppShrinkRequest
	GetHarmonyPayloadShrink() *string
	SetIosPayloadShrink(v string) *SendByAppShrinkRequest
	GetIosPayloadShrink() *string
	SetPolicyShrink(v string) *SendByAppShrinkRequest
	GetPolicyShrink() *string
	SetProductionMode(v bool) *SendByAppShrinkRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByAppShrinkRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByAppShrinkRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByAppShrinkRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByAppShrinkRequest
	GetCallbackParams() *string
}

type SendByAppShrinkRequest struct {
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

func (s SendByAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByAppShrinkRequest) GetAndroidPayloadShrink() *string {
	return s.AndroidPayloadShrink
}

func (s *SendByAppShrinkRequest) GetAndroidShortPayloadShrink() *string {
	return s.AndroidShortPayloadShrink
}

func (s *SendByAppShrinkRequest) GetChannelPropertiesShrink() *string {
	return s.ChannelPropertiesShrink
}

func (s *SendByAppShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByAppShrinkRequest) GetHarmonyPayloadShrink() *string {
	return s.HarmonyPayloadShrink
}

func (s *SendByAppShrinkRequest) GetIosPayloadShrink() *string {
	return s.IosPayloadShrink
}

func (s *SendByAppShrinkRequest) GetPolicyShrink() *string {
	return s.PolicyShrink
}

func (s *SendByAppShrinkRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByAppShrinkRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByAppShrinkRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByAppShrinkRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByAppShrinkRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByAppShrinkRequest) SetAndroidPayloadShrink(v string) *SendByAppShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetAndroidShortPayloadShrink(v string) *SendByAppShrinkRequest {
	s.AndroidShortPayloadShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetChannelPropertiesShrink(v string) *SendByAppShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetDescription(v string) *SendByAppShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByAppShrinkRequest) SetHarmonyPayloadShrink(v string) *SendByAppShrinkRequest {
	s.HarmonyPayloadShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetIosPayloadShrink(v string) *SendByAppShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetPolicyShrink(v string) *SendByAppShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetProductionMode(v bool) *SendByAppShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAppShrinkRequest) SetReceiptType(v int32) *SendByAppShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAppShrinkRequest) SetReceiptUrl(v string) *SendByAppShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAppShrinkRequest) SetThirdPartyId(v string) *SendByAppShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAppShrinkRequest) SetCallbackParams(v string) *SendByAppShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
