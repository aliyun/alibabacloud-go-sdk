// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidPayload(v *AndroidPayload) *SendByDeviceRequest
	GetAndroidPayload() *AndroidPayload
	SetAndroidShortPayload(v *AndroidShortPayload) *SendByDeviceRequest
	GetAndroidShortPayload() *AndroidShortPayload
	SetChannelProperties(v *ChannelProperties) *SendByDeviceRequest
	GetChannelProperties() *ChannelProperties
	SetDescription(v string) *SendByDeviceRequest
	GetDescription() *string
	SetDeviceTokens(v string) *SendByDeviceRequest
	GetDeviceTokens() *string
	SetHarmonyPayload(v *HarmonyPayload) *SendByDeviceRequest
	GetHarmonyPayload() *HarmonyPayload
	SetIosPayload(v *IosPayload) *SendByDeviceRequest
	GetIosPayload() *IosPayload
	SetPolicy(v *Policy) *SendByDeviceRequest
	GetPolicy() *Policy
	SetProductionMode(v bool) *SendByDeviceRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByDeviceRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByDeviceRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByDeviceRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByDeviceRequest
	GetCallbackParams() *string
}

type SendByDeviceRequest struct {
	AndroidPayload      *AndroidPayload      `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelProperties   *ChannelProperties   `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description         *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ArdNyIzFCH2K3szXA8arpu0Y7ywOdA67mCSumtpnMnmf
	DeviceTokens   *string         `json:"DeviceTokens,omitempty" xml:"DeviceTokens,omitempty"`
	HarmonyPayload *HarmonyPayload `json:"HarmonyPayload,omitempty" xml:"HarmonyPayload,omitempty"`
	IosPayload     *IosPayload     `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	Policy         *Policy         `json:"Policy,omitempty" xml:"Policy,omitempty"`
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

func (s SendByDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByDeviceRequest) GoString() string {
	return s.String()
}

func (s *SendByDeviceRequest) GetAndroidPayload() *AndroidPayload {
	return s.AndroidPayload
}

func (s *SendByDeviceRequest) GetAndroidShortPayload() *AndroidShortPayload {
	return s.AndroidShortPayload
}

func (s *SendByDeviceRequest) GetChannelProperties() *ChannelProperties {
	return s.ChannelProperties
}

func (s *SendByDeviceRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByDeviceRequest) GetDeviceTokens() *string {
	return s.DeviceTokens
}

func (s *SendByDeviceRequest) GetHarmonyPayload() *HarmonyPayload {
	return s.HarmonyPayload
}

func (s *SendByDeviceRequest) GetIosPayload() *IosPayload {
	return s.IosPayload
}

func (s *SendByDeviceRequest) GetPolicy() *Policy {
	return s.Policy
}

func (s *SendByDeviceRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByDeviceRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByDeviceRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByDeviceRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByDeviceRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByDeviceRequest) SetAndroidPayload(v *AndroidPayload) *SendByDeviceRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByDeviceRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByDeviceRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByDeviceRequest) SetChannelProperties(v *ChannelProperties) *SendByDeviceRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByDeviceRequest) SetDescription(v string) *SendByDeviceRequest {
	s.Description = &v
	return s
}

func (s *SendByDeviceRequest) SetDeviceTokens(v string) *SendByDeviceRequest {
	s.DeviceTokens = &v
	return s
}

func (s *SendByDeviceRequest) SetHarmonyPayload(v *HarmonyPayload) *SendByDeviceRequest {
	s.HarmonyPayload = v
	return s
}

func (s *SendByDeviceRequest) SetIosPayload(v *IosPayload) *SendByDeviceRequest {
	s.IosPayload = v
	return s
}

func (s *SendByDeviceRequest) SetPolicy(v *Policy) *SendByDeviceRequest {
	s.Policy = v
	return s
}

func (s *SendByDeviceRequest) SetProductionMode(v bool) *SendByDeviceRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByDeviceRequest) SetReceiptType(v int32) *SendByDeviceRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByDeviceRequest) SetReceiptUrl(v string) *SendByDeviceRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByDeviceRequest) SetThirdPartyId(v string) *SendByDeviceRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByDeviceRequest) SetCallbackParams(v string) *SendByDeviceRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByDeviceRequest) Validate() error {
	if s.AndroidPayload != nil {
		if err := s.AndroidPayload.Validate(); err != nil {
			return err
		}
	}
	if s.AndroidShortPayload != nil {
		if err := s.AndroidShortPayload.Validate(); err != nil {
			return err
		}
	}
	if s.ChannelProperties != nil {
		if err := s.ChannelProperties.Validate(); err != nil {
			return err
		}
	}
	if s.HarmonyPayload != nil {
		if err := s.HarmonyPayload.Validate(); err != nil {
			return err
		}
	}
	if s.IosPayload != nil {
		if err := s.IosPayload.Validate(); err != nil {
			return err
		}
	}
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
