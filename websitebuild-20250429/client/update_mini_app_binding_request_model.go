// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMiniAppBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *UpdateMiniAppBindingRequest
	GetBizId() *string
	SetChannel(v string) *UpdateMiniAppBindingRequest
	GetChannel() *string
	SetSettingKey(v string) *UpdateMiniAppBindingRequest
	GetSettingKey() *string
	SetSettingValue(v string) *UpdateMiniAppBindingRequest
	GetSettingValue() *string
}

type UpdateMiniAppBindingRequest struct {
	// Business ID
	//
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Channel
	//
	// example:
	//
	// WECHAT
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Configuration key
	//
	// example:
	//
	// xxxx
	SettingKey *string `json:"SettingKey,omitempty" xml:"SettingKey,omitempty"`
	// Configuration value
	//
	// example:
	//
	// xxxx
	SettingValue *string `json:"SettingValue,omitempty" xml:"SettingValue,omitempty"`
}

func (s UpdateMiniAppBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMiniAppBindingRequest) GoString() string {
	return s.String()
}

func (s *UpdateMiniAppBindingRequest) GetBizId() *string {
	return s.BizId
}

func (s *UpdateMiniAppBindingRequest) GetChannel() *string {
	return s.Channel
}

func (s *UpdateMiniAppBindingRequest) GetSettingKey() *string {
	return s.SettingKey
}

func (s *UpdateMiniAppBindingRequest) GetSettingValue() *string {
	return s.SettingValue
}

func (s *UpdateMiniAppBindingRequest) SetBizId(v string) *UpdateMiniAppBindingRequest {
	s.BizId = &v
	return s
}

func (s *UpdateMiniAppBindingRequest) SetChannel(v string) *UpdateMiniAppBindingRequest {
	s.Channel = &v
	return s
}

func (s *UpdateMiniAppBindingRequest) SetSettingKey(v string) *UpdateMiniAppBindingRequest {
	s.SettingKey = &v
	return s
}

func (s *UpdateMiniAppBindingRequest) SetSettingValue(v string) *UpdateMiniAppBindingRequest {
	s.SettingValue = &v
	return s
}

func (s *UpdateMiniAppBindingRequest) Validate() error {
	return dara.Validate(s)
}
