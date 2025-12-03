// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAliasFileIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasType(v string) *SendByAliasFileIdRequest
	GetAliasType() *string
	SetAndroidPayload(v *AndroidPayload) *SendByAliasFileIdRequest
	GetAndroidPayload() *AndroidPayload
	SetAndroidShortPayload(v *AndroidShortPayload) *SendByAliasFileIdRequest
	GetAndroidShortPayload() *AndroidShortPayload
	SetChannelProperties(v *ChannelProperties) *SendByAliasFileIdRequest
	GetChannelProperties() *ChannelProperties
	SetDescription(v string) *SendByAliasFileIdRequest
	GetDescription() *string
	SetFileId(v string) *SendByAliasFileIdRequest
	GetFileId() *string
	SetHarmonyPayload(v *HarmonyPayload) *SendByAliasFileIdRequest
	GetHarmonyPayload() *HarmonyPayload
	SetIosPayload(v *IosPayload) *SendByAliasFileIdRequest
	GetIosPayload() *IosPayload
	SetPolicy(v *Policy) *SendByAliasFileIdRequest
	GetPolicy() *Policy
	SetProductionMode(v bool) *SendByAliasFileIdRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByAliasFileIdRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByAliasFileIdRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByAliasFileIdRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByAliasFileIdRequest
	GetCallbackParams() *string
}

type SendByAliasFileIdRequest struct {
	AliasType           *string              `json:"AliasType,omitempty" xml:"AliasType,omitempty"`
	AndroidPayload      *AndroidPayload      `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelProperties   *ChannelProperties   `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description         *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PF835431668603208261
	FileId         *string         `json:"FileId,omitempty" xml:"FileId,omitempty"`
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

func (s SendByAliasFileIdRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByAliasFileIdRequest) GoString() string {
	return s.String()
}

func (s *SendByAliasFileIdRequest) GetAliasType() *string {
	return s.AliasType
}

func (s *SendByAliasFileIdRequest) GetAndroidPayload() *AndroidPayload {
	return s.AndroidPayload
}

func (s *SendByAliasFileIdRequest) GetAndroidShortPayload() *AndroidShortPayload {
	return s.AndroidShortPayload
}

func (s *SendByAliasFileIdRequest) GetChannelProperties() *ChannelProperties {
	return s.ChannelProperties
}

func (s *SendByAliasFileIdRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByAliasFileIdRequest) GetFileId() *string {
	return s.FileId
}

func (s *SendByAliasFileIdRequest) GetHarmonyPayload() *HarmonyPayload {
	return s.HarmonyPayload
}

func (s *SendByAliasFileIdRequest) GetIosPayload() *IosPayload {
	return s.IosPayload
}

func (s *SendByAliasFileIdRequest) GetPolicy() *Policy {
	return s.Policy
}

func (s *SendByAliasFileIdRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByAliasFileIdRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByAliasFileIdRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByAliasFileIdRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByAliasFileIdRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByAliasFileIdRequest) SetAliasType(v string) *SendByAliasFileIdRequest {
	s.AliasType = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetAndroidPayload(v *AndroidPayload) *SendByAliasFileIdRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByAliasFileIdRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByAliasFileIdRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByAliasFileIdRequest) SetChannelProperties(v *ChannelProperties) *SendByAliasFileIdRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByAliasFileIdRequest) SetDescription(v string) *SendByAliasFileIdRequest {
	s.Description = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetFileId(v string) *SendByAliasFileIdRequest {
	s.FileId = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetHarmonyPayload(v *HarmonyPayload) *SendByAliasFileIdRequest {
	s.HarmonyPayload = v
	return s
}

func (s *SendByAliasFileIdRequest) SetIosPayload(v *IosPayload) *SendByAliasFileIdRequest {
	s.IosPayload = v
	return s
}

func (s *SendByAliasFileIdRequest) SetPolicy(v *Policy) *SendByAliasFileIdRequest {
	s.Policy = v
	return s
}

func (s *SendByAliasFileIdRequest) SetProductionMode(v bool) *SendByAliasFileIdRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetReceiptType(v int32) *SendByAliasFileIdRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetReceiptUrl(v string) *SendByAliasFileIdRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetThirdPartyId(v string) *SendByAliasFileIdRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetCallbackParams(v string) *SendByAliasFileIdRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByAliasFileIdRequest) Validate() error {
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
