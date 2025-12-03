// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidPayload(v *AndroidPayload) *SendByAppRequest
	GetAndroidPayload() *AndroidPayload
	SetAndroidShortPayload(v *AndroidShortPayload) *SendByAppRequest
	GetAndroidShortPayload() *AndroidShortPayload
	SetChannelProperties(v *ChannelProperties) *SendByAppRequest
	GetChannelProperties() *ChannelProperties
	SetDescription(v string) *SendByAppRequest
	GetDescription() *string
	SetHarmonyPayload(v *HarmonyPayload) *SendByAppRequest
	GetHarmonyPayload() *HarmonyPayload
	SetIosPayload(v *IosPayload) *SendByAppRequest
	GetIosPayload() *IosPayload
	SetPolicy(v *Policy) *SendByAppRequest
	GetPolicy() *Policy
	SetProductionMode(v bool) *SendByAppRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByAppRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByAppRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByAppRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByAppRequest
	GetCallbackParams() *string
}

type SendByAppRequest struct {
	AndroidPayload      *AndroidPayload      `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelProperties   *ChannelProperties   `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description         *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	HarmonyPayload      *HarmonyPayload      `json:"HarmonyPayload,omitempty" xml:"HarmonyPayload,omitempty"`
	IosPayload          *IosPayload          `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	Policy              *Policy              `json:"Policy,omitempty" xml:"Policy,omitempty"`
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

func (s SendByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByAppRequest) GoString() string {
	return s.String()
}

func (s *SendByAppRequest) GetAndroidPayload() *AndroidPayload {
	return s.AndroidPayload
}

func (s *SendByAppRequest) GetAndroidShortPayload() *AndroidShortPayload {
	return s.AndroidShortPayload
}

func (s *SendByAppRequest) GetChannelProperties() *ChannelProperties {
	return s.ChannelProperties
}

func (s *SendByAppRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByAppRequest) GetHarmonyPayload() *HarmonyPayload {
	return s.HarmonyPayload
}

func (s *SendByAppRequest) GetIosPayload() *IosPayload {
	return s.IosPayload
}

func (s *SendByAppRequest) GetPolicy() *Policy {
	return s.Policy
}

func (s *SendByAppRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByAppRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByAppRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByAppRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByAppRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByAppRequest) SetAndroidPayload(v *AndroidPayload) *SendByAppRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByAppRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByAppRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByAppRequest) SetChannelProperties(v *ChannelProperties) *SendByAppRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByAppRequest) SetDescription(v string) *SendByAppRequest {
	s.Description = &v
	return s
}

func (s *SendByAppRequest) SetHarmonyPayload(v *HarmonyPayload) *SendByAppRequest {
	s.HarmonyPayload = v
	return s
}

func (s *SendByAppRequest) SetIosPayload(v *IosPayload) *SendByAppRequest {
	s.IosPayload = v
	return s
}

func (s *SendByAppRequest) SetPolicy(v *Policy) *SendByAppRequest {
	s.Policy = v
	return s
}

func (s *SendByAppRequest) SetProductionMode(v bool) *SendByAppRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAppRequest) SetReceiptType(v int32) *SendByAppRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAppRequest) SetReceiptUrl(v string) *SendByAppRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAppRequest) SetThirdPartyId(v string) *SendByAppRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAppRequest) SetCallbackParams(v string) *SendByAppRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByAppRequest) Validate() error {
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
