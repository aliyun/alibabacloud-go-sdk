// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByDeviceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidPayloadShrink(v string) *SendByDeviceShrinkRequest
	GetAndroidPayloadShrink() *string
	SetAndroidShortPayloadShrink(v string) *SendByDeviceShrinkRequest
	GetAndroidShortPayloadShrink() *string
	SetChannelPropertiesShrink(v string) *SendByDeviceShrinkRequest
	GetChannelPropertiesShrink() *string
	SetDescription(v string) *SendByDeviceShrinkRequest
	GetDescription() *string
	SetDeviceTokens(v string) *SendByDeviceShrinkRequest
	GetDeviceTokens() *string
	SetHarmonyPayloadShrink(v string) *SendByDeviceShrinkRequest
	GetHarmonyPayloadShrink() *string
	SetIosPayloadShrink(v string) *SendByDeviceShrinkRequest
	GetIosPayloadShrink() *string
	SetPolicyShrink(v string) *SendByDeviceShrinkRequest
	GetPolicyShrink() *string
	SetProductionMode(v bool) *SendByDeviceShrinkRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByDeviceShrinkRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByDeviceShrinkRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByDeviceShrinkRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByDeviceShrinkRequest
	GetCallbackParams() *string
}

type SendByDeviceShrinkRequest struct {
	AndroidPayloadShrink      *string `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayloadShrink *string `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelPropertiesShrink   *string `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ArdNyIzFCH2K3szXA8arpu0Y7ywOdA67mCSumtpnMnmf
	DeviceTokens         *string `json:"DeviceTokens,omitempty" xml:"DeviceTokens,omitempty"`
	HarmonyPayloadShrink *string `json:"HarmonyPayload,omitempty" xml:"HarmonyPayload,omitempty"`
	IosPayloadShrink     *string `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	PolicyShrink         *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
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

func (s SendByDeviceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByDeviceShrinkRequest) GetAndroidPayloadShrink() *string {
	return s.AndroidPayloadShrink
}

func (s *SendByDeviceShrinkRequest) GetAndroidShortPayloadShrink() *string {
	return s.AndroidShortPayloadShrink
}

func (s *SendByDeviceShrinkRequest) GetChannelPropertiesShrink() *string {
	return s.ChannelPropertiesShrink
}

func (s *SendByDeviceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByDeviceShrinkRequest) GetDeviceTokens() *string {
	return s.DeviceTokens
}

func (s *SendByDeviceShrinkRequest) GetHarmonyPayloadShrink() *string {
	return s.HarmonyPayloadShrink
}

func (s *SendByDeviceShrinkRequest) GetIosPayloadShrink() *string {
	return s.IosPayloadShrink
}

func (s *SendByDeviceShrinkRequest) GetPolicyShrink() *string {
	return s.PolicyShrink
}

func (s *SendByDeviceShrinkRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByDeviceShrinkRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByDeviceShrinkRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByDeviceShrinkRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByDeviceShrinkRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByDeviceShrinkRequest) SetAndroidPayloadShrink(v string) *SendByDeviceShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetAndroidShortPayloadShrink(v string) *SendByDeviceShrinkRequest {
	s.AndroidShortPayloadShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetChannelPropertiesShrink(v string) *SendByDeviceShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetDescription(v string) *SendByDeviceShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetDeviceTokens(v string) *SendByDeviceShrinkRequest {
	s.DeviceTokens = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetHarmonyPayloadShrink(v string) *SendByDeviceShrinkRequest {
	s.HarmonyPayloadShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetIosPayloadShrink(v string) *SendByDeviceShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetPolicyShrink(v string) *SendByDeviceShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetProductionMode(v bool) *SendByDeviceShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetReceiptType(v int32) *SendByDeviceShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetReceiptUrl(v string) *SendByDeviceShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetThirdPartyId(v string) *SendByDeviceShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetCallbackParams(v string) *SendByDeviceShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByDeviceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
