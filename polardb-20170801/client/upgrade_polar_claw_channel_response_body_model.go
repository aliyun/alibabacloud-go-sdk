// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePolarClawChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpgradePolarClawChannelResponseBody
	GetApplicationId() *string
	SetChannelId(v string) *UpgradePolarClawChannelResponseBody
	GetChannelId() *string
	SetCode(v int32) *UpgradePolarClawChannelResponseBody
	GetCode() *int32
	SetMessage(v string) *UpgradePolarClawChannelResponseBody
	GetMessage() *string
	SetNpmPackage(v string) *UpgradePolarClawChannelResponseBody
	GetNpmPackage() *string
	SetOk(v bool) *UpgradePolarClawChannelResponseBody
	GetOk() *bool
	SetPluginId(v string) *UpgradePolarClawChannelResponseBody
	GetPluginId() *string
	SetPluginUpgraded(v bool) *UpgradePolarClawChannelResponseBody
	GetPluginUpgraded() *bool
	SetRequestId(v string) *UpgradePolarClawChannelResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *UpgradePolarClawChannelResponseBody
	GetRestarted() *bool
}

type UpgradePolarClawChannelResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// feishu
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// @larksuite/openclaw-lark@2026.4.7
	NpmPackage *string `json:"NpmPackage,omitempty" xml:"NpmPackage,omitempty"`
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// example:
	//
	// openclaw-lark
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// example:
	//
	// true
	PluginUpgraded *bool `json:"PluginUpgraded,omitempty" xml:"PluginUpgraded,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s UpgradePolarClawChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradePolarClawChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradePolarClawChannelResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpgradePolarClawChannelResponseBody) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpgradePolarClawChannelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpgradePolarClawChannelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradePolarClawChannelResponseBody) GetNpmPackage() *string {
	return s.NpmPackage
}

func (s *UpgradePolarClawChannelResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *UpgradePolarClawChannelResponseBody) GetPluginId() *string {
	return s.PluginId
}

func (s *UpgradePolarClawChannelResponseBody) GetPluginUpgraded() *bool {
	return s.PluginUpgraded
}

func (s *UpgradePolarClawChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradePolarClawChannelResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *UpgradePolarClawChannelResponseBody) SetApplicationId(v string) *UpgradePolarClawChannelResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *UpgradePolarClawChannelResponseBody) SetChannelId(v string) *UpgradePolarClawChannelResponseBody {
	s.ChannelId = &v
	return s
}

func (s *UpgradePolarClawChannelResponseBody) SetCode(v int32) *UpgradePolarClawChannelResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradePolarClawChannelResponseBody) SetMessage(v string) *UpgradePolarClawChannelResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradePolarClawChannelResponseBody) SetNpmPackage(v string) *UpgradePolarClawChannelResponseBody {
	s.NpmPackage = &v
	return s
}

func (s *UpgradePolarClawChannelResponseBody) SetOk(v bool) *UpgradePolarClawChannelResponseBody {
	s.Ok = &v
	return s
}

func (s *UpgradePolarClawChannelResponseBody) SetPluginId(v string) *UpgradePolarClawChannelResponseBody {
	s.PluginId = &v
	return s
}

func (s *UpgradePolarClawChannelResponseBody) SetPluginUpgraded(v bool) *UpgradePolarClawChannelResponseBody {
	s.PluginUpgraded = &v
	return s
}

func (s *UpgradePolarClawChannelResponseBody) SetRequestId(v string) *UpgradePolarClawChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradePolarClawChannelResponseBody) SetRestarted(v bool) *UpgradePolarClawChannelResponseBody {
	s.Restarted = &v
	return s
}

func (s *UpgradePolarClawChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
