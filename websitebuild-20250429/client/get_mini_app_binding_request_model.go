// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiniAppBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetMiniAppBindingRequest
	GetBizId() *string
	SetChannel(v string) *GetMiniAppBindingRequest
	GetChannel() *string
	SetSettingKeys(v []*string) *GetMiniAppBindingRequest
	GetSettingKeys() []*string
}

type GetMiniAppBindingRequest struct {
	// example:
	//
	// WS20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// WECHAT
	Channel     *string   `json:"Channel,omitempty" xml:"Channel,omitempty"`
	SettingKeys []*string `json:"SettingKeys,omitempty" xml:"SettingKeys,omitempty" type:"Repeated"`
}

func (s GetMiniAppBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppBindingRequest) GoString() string {
	return s.String()
}

func (s *GetMiniAppBindingRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetMiniAppBindingRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetMiniAppBindingRequest) GetSettingKeys() []*string {
	return s.SettingKeys
}

func (s *GetMiniAppBindingRequest) SetBizId(v string) *GetMiniAppBindingRequest {
	s.BizId = &v
	return s
}

func (s *GetMiniAppBindingRequest) SetChannel(v string) *GetMiniAppBindingRequest {
	s.Channel = &v
	return s
}

func (s *GetMiniAppBindingRequest) SetSettingKeys(v []*string) *GetMiniAppBindingRequest {
	s.SettingKeys = v
	return s
}

func (s *GetMiniAppBindingRequest) Validate() error {
	return dara.Validate(s)
}
