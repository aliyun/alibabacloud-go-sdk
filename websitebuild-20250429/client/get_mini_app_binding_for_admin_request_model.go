// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiniAppBindingForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetMiniAppBindingForAdminRequest
	GetBizId() *string
	SetChannel(v string) *GetMiniAppBindingForAdminRequest
	GetChannel() *string
	SetPlatformAppid(v string) *GetMiniAppBindingForAdminRequest
	GetPlatformAppid() *string
}

type GetMiniAppBindingForAdminRequest struct {
	// example:
	//
	// WS20250801152639000005
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// WECHAT
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// xxxx
	PlatformAppid *string `json:"PlatformAppid,omitempty" xml:"PlatformAppid,omitempty"`
}

func (s GetMiniAppBindingForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppBindingForAdminRequest) GoString() string {
	return s.String()
}

func (s *GetMiniAppBindingForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetMiniAppBindingForAdminRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetMiniAppBindingForAdminRequest) GetPlatformAppid() *string {
	return s.PlatformAppid
}

func (s *GetMiniAppBindingForAdminRequest) SetBizId(v string) *GetMiniAppBindingForAdminRequest {
	s.BizId = &v
	return s
}

func (s *GetMiniAppBindingForAdminRequest) SetChannel(v string) *GetMiniAppBindingForAdminRequest {
	s.Channel = &v
	return s
}

func (s *GetMiniAppBindingForAdminRequest) SetPlatformAppid(v string) *GetMiniAppBindingForAdminRequest {
	s.PlatformAppid = &v
	return s
}

func (s *GetMiniAppBindingForAdminRequest) Validate() error {
	return dara.Validate(s)
}
