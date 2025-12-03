// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByDeviceFileIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest
	GetAndroidPayloadShrink() *string
	SetAndroidShortPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest
	GetAndroidShortPayloadShrink() *string
	SetChannelPropertiesShrink(v string) *SendByDeviceFileIdShrinkRequest
	GetChannelPropertiesShrink() *string
	SetDescription(v string) *SendByDeviceFileIdShrinkRequest
	GetDescription() *string
	SetFileId(v string) *SendByDeviceFileIdShrinkRequest
	GetFileId() *string
	SetHarmonyPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest
	GetHarmonyPayloadShrink() *string
	SetIosPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest
	GetIosPayloadShrink() *string
	SetPolicyShrink(v string) *SendByDeviceFileIdShrinkRequest
	GetPolicyShrink() *string
	SetProductionMode(v bool) *SendByDeviceFileIdShrinkRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByDeviceFileIdShrinkRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByDeviceFileIdShrinkRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByDeviceFileIdShrinkRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByDeviceFileIdShrinkRequest
	GetCallbackParams() *string
}

type SendByDeviceFileIdShrinkRequest struct {
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

func (s SendByDeviceFileIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByDeviceFileIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByDeviceFileIdShrinkRequest) GetAndroidPayloadShrink() *string {
	return s.AndroidPayloadShrink
}

func (s *SendByDeviceFileIdShrinkRequest) GetAndroidShortPayloadShrink() *string {
	return s.AndroidShortPayloadShrink
}

func (s *SendByDeviceFileIdShrinkRequest) GetChannelPropertiesShrink() *string {
	return s.ChannelPropertiesShrink
}

func (s *SendByDeviceFileIdShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByDeviceFileIdShrinkRequest) GetFileId() *string {
	return s.FileId
}

func (s *SendByDeviceFileIdShrinkRequest) GetHarmonyPayloadShrink() *string {
	return s.HarmonyPayloadShrink
}

func (s *SendByDeviceFileIdShrinkRequest) GetIosPayloadShrink() *string {
	return s.IosPayloadShrink
}

func (s *SendByDeviceFileIdShrinkRequest) GetPolicyShrink() *string {
	return s.PolicyShrink
}

func (s *SendByDeviceFileIdShrinkRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByDeviceFileIdShrinkRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByDeviceFileIdShrinkRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByDeviceFileIdShrinkRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByDeviceFileIdShrinkRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByDeviceFileIdShrinkRequest) SetAndroidPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetAndroidShortPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.AndroidShortPayloadShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetChannelPropertiesShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetDescription(v string) *SendByDeviceFileIdShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetFileId(v string) *SendByDeviceFileIdShrinkRequest {
	s.FileId = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetHarmonyPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.HarmonyPayloadShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetIosPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetPolicyShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetProductionMode(v bool) *SendByDeviceFileIdShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetReceiptType(v int32) *SendByDeviceFileIdShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetReceiptUrl(v string) *SendByDeviceFileIdShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetThirdPartyId(v string) *SendByDeviceFileIdShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetCallbackParams(v string) *SendByDeviceFileIdShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
