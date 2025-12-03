// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *SendByAliasRequest
	GetAlias() *string
	SetAliasType(v string) *SendByAliasRequest
	GetAliasType() *string
	SetAndroidPayload(v *AndroidPayload) *SendByAliasRequest
	GetAndroidPayload() *AndroidPayload
	SetAndroidShortPayload(v *AndroidShortPayload) *SendByAliasRequest
	GetAndroidShortPayload() *AndroidShortPayload
	SetChannelProperties(v *ChannelProperties) *SendByAliasRequest
	GetChannelProperties() *ChannelProperties
	SetDescription(v string) *SendByAliasRequest
	GetDescription() *string
	SetHarmonyPayload(v *HarmonyPayload) *SendByAliasRequest
	GetHarmonyPayload() *HarmonyPayload
	SetIosPayload(v *IosPayload) *SendByAliasRequest
	GetIosPayload() *IosPayload
	SetPolicy(v *Policy) *SendByAliasRequest
	GetPolicy() *Policy
	SetProductionMode(v bool) *SendByAliasRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByAliasRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByAliasRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByAliasRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByAliasRequest
	GetCallbackParams() *string
}

type SendByAliasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Alias               *string              `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AliasType           *string              `json:"AliasType,omitempty" xml:"AliasType,omitempty"`
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

func (s SendByAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByAliasRequest) GoString() string {
	return s.String()
}

func (s *SendByAliasRequest) GetAlias() *string {
	return s.Alias
}

func (s *SendByAliasRequest) GetAliasType() *string {
	return s.AliasType
}

func (s *SendByAliasRequest) GetAndroidPayload() *AndroidPayload {
	return s.AndroidPayload
}

func (s *SendByAliasRequest) GetAndroidShortPayload() *AndroidShortPayload {
	return s.AndroidShortPayload
}

func (s *SendByAliasRequest) GetChannelProperties() *ChannelProperties {
	return s.ChannelProperties
}

func (s *SendByAliasRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByAliasRequest) GetHarmonyPayload() *HarmonyPayload {
	return s.HarmonyPayload
}

func (s *SendByAliasRequest) GetIosPayload() *IosPayload {
	return s.IosPayload
}

func (s *SendByAliasRequest) GetPolicy() *Policy {
	return s.Policy
}

func (s *SendByAliasRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByAliasRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByAliasRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByAliasRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByAliasRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByAliasRequest) SetAlias(v string) *SendByAliasRequest {
	s.Alias = &v
	return s
}

func (s *SendByAliasRequest) SetAliasType(v string) *SendByAliasRequest {
	s.AliasType = &v
	return s
}

func (s *SendByAliasRequest) SetAndroidPayload(v *AndroidPayload) *SendByAliasRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByAliasRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByAliasRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByAliasRequest) SetChannelProperties(v *ChannelProperties) *SendByAliasRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByAliasRequest) SetDescription(v string) *SendByAliasRequest {
	s.Description = &v
	return s
}

func (s *SendByAliasRequest) SetHarmonyPayload(v *HarmonyPayload) *SendByAliasRequest {
	s.HarmonyPayload = v
	return s
}

func (s *SendByAliasRequest) SetIosPayload(v *IosPayload) *SendByAliasRequest {
	s.IosPayload = v
	return s
}

func (s *SendByAliasRequest) SetPolicy(v *Policy) *SendByAliasRequest {
	s.Policy = v
	return s
}

func (s *SendByAliasRequest) SetProductionMode(v bool) *SendByAliasRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAliasRequest) SetReceiptType(v int32) *SendByAliasRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAliasRequest) SetReceiptUrl(v string) *SendByAliasRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAliasRequest) SetThirdPartyId(v string) *SendByAliasRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAliasRequest) SetCallbackParams(v string) *SendByAliasRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByAliasRequest) Validate() error {
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
