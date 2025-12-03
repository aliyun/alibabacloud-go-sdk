// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByFilterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidPayloadShrink(v string) *SendByFilterShrinkRequest
	GetAndroidPayloadShrink() *string
	SetAndroidShortPayload(v *AndroidShortPayload) *SendByFilterShrinkRequest
	GetAndroidShortPayload() *AndroidShortPayload
	SetChannelPropertiesShrink(v string) *SendByFilterShrinkRequest
	GetChannelPropertiesShrink() *string
	SetDescription(v string) *SendByFilterShrinkRequest
	GetDescription() *string
	SetFilter(v string) *SendByFilterShrinkRequest
	GetFilter() *string
	SetHarmonyPayloadShrink(v string) *SendByFilterShrinkRequest
	GetHarmonyPayloadShrink() *string
	SetIosPayloadShrink(v string) *SendByFilterShrinkRequest
	GetIosPayloadShrink() *string
	SetPolicyShrink(v string) *SendByFilterShrinkRequest
	GetPolicyShrink() *string
	SetProductionMode(v bool) *SendByFilterShrinkRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByFilterShrinkRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByFilterShrinkRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByFilterShrinkRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByFilterShrinkRequest
	GetCallbackParams() *string
}

type SendByFilterShrinkRequest struct {
	AndroidPayloadShrink    *string              `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload     *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelPropertiesShrink *string              `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description             *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// "where":{"and":[{"or":[{"app_version":">=1.0"}]}]}
	Filter               *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
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

func (s SendByFilterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByFilterShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByFilterShrinkRequest) GetAndroidPayloadShrink() *string {
	return s.AndroidPayloadShrink
}

func (s *SendByFilterShrinkRequest) GetAndroidShortPayload() *AndroidShortPayload {
	return s.AndroidShortPayload
}

func (s *SendByFilterShrinkRequest) GetChannelPropertiesShrink() *string {
	return s.ChannelPropertiesShrink
}

func (s *SendByFilterShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByFilterShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *SendByFilterShrinkRequest) GetHarmonyPayloadShrink() *string {
	return s.HarmonyPayloadShrink
}

func (s *SendByFilterShrinkRequest) GetIosPayloadShrink() *string {
	return s.IosPayloadShrink
}

func (s *SendByFilterShrinkRequest) GetPolicyShrink() *string {
	return s.PolicyShrink
}

func (s *SendByFilterShrinkRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByFilterShrinkRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByFilterShrinkRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByFilterShrinkRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByFilterShrinkRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByFilterShrinkRequest) SetAndroidPayloadShrink(v string) *SendByFilterShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByFilterShrinkRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByFilterShrinkRequest) SetChannelPropertiesShrink(v string) *SendByFilterShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetDescription(v string) *SendByFilterShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetFilter(v string) *SendByFilterShrinkRequest {
	s.Filter = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetHarmonyPayloadShrink(v string) *SendByFilterShrinkRequest {
	s.HarmonyPayloadShrink = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetIosPayloadShrink(v string) *SendByFilterShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetPolicyShrink(v string) *SendByFilterShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetProductionMode(v bool) *SendByFilterShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetReceiptType(v int32) *SendByFilterShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetReceiptUrl(v string) *SendByFilterShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetThirdPartyId(v string) *SendByFilterShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetCallbackParams(v string) *SendByFilterShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByFilterShrinkRequest) Validate() error {
	if s.AndroidShortPayload != nil {
		if err := s.AndroidShortPayload.Validate(); err != nil {
			return err
		}
	}
	return nil
}
