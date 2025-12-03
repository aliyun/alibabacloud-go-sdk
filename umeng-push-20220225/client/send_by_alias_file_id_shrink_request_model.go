// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAliasFileIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasType(v string) *SendByAliasFileIdShrinkRequest
	GetAliasType() *string
	SetAndroidPayloadShrink(v string) *SendByAliasFileIdShrinkRequest
	GetAndroidPayloadShrink() *string
	SetAndroidShortPayloadShrink(v string) *SendByAliasFileIdShrinkRequest
	GetAndroidShortPayloadShrink() *string
	SetChannelPropertiesShrink(v string) *SendByAliasFileIdShrinkRequest
	GetChannelPropertiesShrink() *string
	SetDescription(v string) *SendByAliasFileIdShrinkRequest
	GetDescription() *string
	SetFileId(v string) *SendByAliasFileIdShrinkRequest
	GetFileId() *string
	SetHarmonyPayloadShrink(v string) *SendByAliasFileIdShrinkRequest
	GetHarmonyPayloadShrink() *string
	SetIosPayloadShrink(v string) *SendByAliasFileIdShrinkRequest
	GetIosPayloadShrink() *string
	SetPolicyShrink(v string) *SendByAliasFileIdShrinkRequest
	GetPolicyShrink() *string
	SetProductionMode(v bool) *SendByAliasFileIdShrinkRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByAliasFileIdShrinkRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByAliasFileIdShrinkRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByAliasFileIdShrinkRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByAliasFileIdShrinkRequest
	GetCallbackParams() *string
}

type SendByAliasFileIdShrinkRequest struct {
	AliasType                 *string `json:"AliasType,omitempty" xml:"AliasType,omitempty"`
	AndroidPayloadShrink      *string `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayloadShrink *string `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelPropertiesShrink   *string `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PF835431668603208261
	FileId               *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
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

func (s SendByAliasFileIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByAliasFileIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByAliasFileIdShrinkRequest) GetAliasType() *string {
	return s.AliasType
}

func (s *SendByAliasFileIdShrinkRequest) GetAndroidPayloadShrink() *string {
	return s.AndroidPayloadShrink
}

func (s *SendByAliasFileIdShrinkRequest) GetAndroidShortPayloadShrink() *string {
	return s.AndroidShortPayloadShrink
}

func (s *SendByAliasFileIdShrinkRequest) GetChannelPropertiesShrink() *string {
	return s.ChannelPropertiesShrink
}

func (s *SendByAliasFileIdShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByAliasFileIdShrinkRequest) GetFileId() *string {
	return s.FileId
}

func (s *SendByAliasFileIdShrinkRequest) GetHarmonyPayloadShrink() *string {
	return s.HarmonyPayloadShrink
}

func (s *SendByAliasFileIdShrinkRequest) GetIosPayloadShrink() *string {
	return s.IosPayloadShrink
}

func (s *SendByAliasFileIdShrinkRequest) GetPolicyShrink() *string {
	return s.PolicyShrink
}

func (s *SendByAliasFileIdShrinkRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByAliasFileIdShrinkRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByAliasFileIdShrinkRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByAliasFileIdShrinkRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByAliasFileIdShrinkRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByAliasFileIdShrinkRequest) SetAliasType(v string) *SendByAliasFileIdShrinkRequest {
	s.AliasType = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetAndroidPayloadShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetAndroidShortPayloadShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.AndroidShortPayloadShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetChannelPropertiesShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetDescription(v string) *SendByAliasFileIdShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetFileId(v string) *SendByAliasFileIdShrinkRequest {
	s.FileId = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetHarmonyPayloadShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.HarmonyPayloadShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetIosPayloadShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetPolicyShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetProductionMode(v bool) *SendByAliasFileIdShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetReceiptType(v int32) *SendByAliasFileIdShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetReceiptUrl(v string) *SendByAliasFileIdShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetThirdPartyId(v string) *SendByAliasFileIdShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetCallbackParams(v string) *SendByAliasFileIdShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
