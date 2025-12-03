// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidPayload(v *AndroidPayload) *SendByFilterRequest
	GetAndroidPayload() *AndroidPayload
	SetAndroidShortPayload(v *AndroidShortPayload) *SendByFilterRequest
	GetAndroidShortPayload() *AndroidShortPayload
	SetChannelProperties(v *ChannelProperties) *SendByFilterRequest
	GetChannelProperties() *ChannelProperties
	SetDescription(v string) *SendByFilterRequest
	GetDescription() *string
	SetFilter(v string) *SendByFilterRequest
	GetFilter() *string
	SetHarmonyPayload(v *HarmonyPayload) *SendByFilterRequest
	GetHarmonyPayload() *HarmonyPayload
	SetIosPayload(v *IosPayload) *SendByFilterRequest
	GetIosPayload() *IosPayload
	SetPolicy(v *Policy) *SendByFilterRequest
	GetPolicy() *Policy
	SetProductionMode(v bool) *SendByFilterRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByFilterRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByFilterRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByFilterRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByFilterRequest
	GetCallbackParams() *string
}

type SendByFilterRequest struct {
	AndroidPayload      *AndroidPayload      `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelProperties   *ChannelProperties   `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description         *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// "where":{"and":[{"or":[{"app_version":">=1.0"}]}]}
	Filter         *string         `json:"Filter,omitempty" xml:"Filter,omitempty"`
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

func (s SendByFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByFilterRequest) GoString() string {
	return s.String()
}

func (s *SendByFilterRequest) GetAndroidPayload() *AndroidPayload {
	return s.AndroidPayload
}

func (s *SendByFilterRequest) GetAndroidShortPayload() *AndroidShortPayload {
	return s.AndroidShortPayload
}

func (s *SendByFilterRequest) GetChannelProperties() *ChannelProperties {
	return s.ChannelProperties
}

func (s *SendByFilterRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByFilterRequest) GetFilter() *string {
	return s.Filter
}

func (s *SendByFilterRequest) GetHarmonyPayload() *HarmonyPayload {
	return s.HarmonyPayload
}

func (s *SendByFilterRequest) GetIosPayload() *IosPayload {
	return s.IosPayload
}

func (s *SendByFilterRequest) GetPolicy() *Policy {
	return s.Policy
}

func (s *SendByFilterRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByFilterRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByFilterRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByFilterRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByFilterRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByFilterRequest) SetAndroidPayload(v *AndroidPayload) *SendByFilterRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByFilterRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByFilterRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByFilterRequest) SetChannelProperties(v *ChannelProperties) *SendByFilterRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByFilterRequest) SetDescription(v string) *SendByFilterRequest {
	s.Description = &v
	return s
}

func (s *SendByFilterRequest) SetFilter(v string) *SendByFilterRequest {
	s.Filter = &v
	return s
}

func (s *SendByFilterRequest) SetHarmonyPayload(v *HarmonyPayload) *SendByFilterRequest {
	s.HarmonyPayload = v
	return s
}

func (s *SendByFilterRequest) SetIosPayload(v *IosPayload) *SendByFilterRequest {
	s.IosPayload = v
	return s
}

func (s *SendByFilterRequest) SetPolicy(v *Policy) *SendByFilterRequest {
	s.Policy = v
	return s
}

func (s *SendByFilterRequest) SetProductionMode(v bool) *SendByFilterRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByFilterRequest) SetReceiptType(v int32) *SendByFilterRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByFilterRequest) SetReceiptUrl(v string) *SendByFilterRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByFilterRequest) SetThirdPartyId(v string) *SendByFilterRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByFilterRequest) SetCallbackParams(v string) *SendByFilterRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByFilterRequest) Validate() error {
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
