// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByDeviceFileIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidPayload(v *AndroidPayload) *SendByDeviceFileIdRequest
	GetAndroidPayload() *AndroidPayload
	SetAndroidShortPayload(v *AndroidShortPayload) *SendByDeviceFileIdRequest
	GetAndroidShortPayload() *AndroidShortPayload
	SetChannelProperties(v *ChannelProperties) *SendByDeviceFileIdRequest
	GetChannelProperties() *ChannelProperties
	SetDescription(v string) *SendByDeviceFileIdRequest
	GetDescription() *string
	SetFileId(v string) *SendByDeviceFileIdRequest
	GetFileId() *string
	SetHarmonyPayload(v *HarmonyPayload) *SendByDeviceFileIdRequest
	GetHarmonyPayload() *HarmonyPayload
	SetIosPayload(v *IosPayload) *SendByDeviceFileIdRequest
	GetIosPayload() *IosPayload
	SetPolicy(v *Policy) *SendByDeviceFileIdRequest
	GetPolicy() *Policy
	SetProductionMode(v bool) *SendByDeviceFileIdRequest
	GetProductionMode() *bool
	SetReceiptType(v int32) *SendByDeviceFileIdRequest
	GetReceiptType() *int32
	SetReceiptUrl(v string) *SendByDeviceFileIdRequest
	GetReceiptUrl() *string
	SetThirdPartyId(v string) *SendByDeviceFileIdRequest
	GetThirdPartyId() *string
	SetCallbackParams(v string) *SendByDeviceFileIdRequest
	GetCallbackParams() *string
}

type SendByDeviceFileIdRequest struct {
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

func (s SendByDeviceFileIdRequest) String() string {
	return dara.Prettify(s)
}

func (s SendByDeviceFileIdRequest) GoString() string {
	return s.String()
}

func (s *SendByDeviceFileIdRequest) GetAndroidPayload() *AndroidPayload {
	return s.AndroidPayload
}

func (s *SendByDeviceFileIdRequest) GetAndroidShortPayload() *AndroidShortPayload {
	return s.AndroidShortPayload
}

func (s *SendByDeviceFileIdRequest) GetChannelProperties() *ChannelProperties {
	return s.ChannelProperties
}

func (s *SendByDeviceFileIdRequest) GetDescription() *string {
	return s.Description
}

func (s *SendByDeviceFileIdRequest) GetFileId() *string {
	return s.FileId
}

func (s *SendByDeviceFileIdRequest) GetHarmonyPayload() *HarmonyPayload {
	return s.HarmonyPayload
}

func (s *SendByDeviceFileIdRequest) GetIosPayload() *IosPayload {
	return s.IosPayload
}

func (s *SendByDeviceFileIdRequest) GetPolicy() *Policy {
	return s.Policy
}

func (s *SendByDeviceFileIdRequest) GetProductionMode() *bool {
	return s.ProductionMode
}

func (s *SendByDeviceFileIdRequest) GetReceiptType() *int32 {
	return s.ReceiptType
}

func (s *SendByDeviceFileIdRequest) GetReceiptUrl() *string {
	return s.ReceiptUrl
}

func (s *SendByDeviceFileIdRequest) GetThirdPartyId() *string {
	return s.ThirdPartyId
}

func (s *SendByDeviceFileIdRequest) GetCallbackParams() *string {
	return s.CallbackParams
}

func (s *SendByDeviceFileIdRequest) SetAndroidPayload(v *AndroidPayload) *SendByDeviceFileIdRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByDeviceFileIdRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetChannelProperties(v *ChannelProperties) *SendByDeviceFileIdRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetDescription(v string) *SendByDeviceFileIdRequest {
	s.Description = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetFileId(v string) *SendByDeviceFileIdRequest {
	s.FileId = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetHarmonyPayload(v *HarmonyPayload) *SendByDeviceFileIdRequest {
	s.HarmonyPayload = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetIosPayload(v *IosPayload) *SendByDeviceFileIdRequest {
	s.IosPayload = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetPolicy(v *Policy) *SendByDeviceFileIdRequest {
	s.Policy = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetProductionMode(v bool) *SendByDeviceFileIdRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetReceiptType(v int32) *SendByDeviceFileIdRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetReceiptUrl(v string) *SendByDeviceFileIdRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetThirdPartyId(v string) *SendByDeviceFileIdRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetCallbackParams(v string) *SendByDeviceFileIdRequest {
	s.CallbackParams = &v
	return s
}

func (s *SendByDeviceFileIdRequest) Validate() error {
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
