// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreatePolarClawChannelResponseBody
	GetApplicationId() *string
	SetChannelId(v string) *CreatePolarClawChannelResponseBody
	GetChannelId() *string
	SetCode(v int32) *CreatePolarClawChannelResponseBody
	GetCode() *int32
	SetMessage(v string) *CreatePolarClawChannelResponseBody
	GetMessage() *string
	SetNpmPackage(v string) *CreatePolarClawChannelResponseBody
	GetNpmPackage() *string
	SetOk(v bool) *CreatePolarClawChannelResponseBody
	GetOk() *bool
	SetPluginId(v string) *CreatePolarClawChannelResponseBody
	GetPluginId() *string
	SetPluginInstalled(v bool) *CreatePolarClawChannelResponseBody
	GetPluginInstalled() *bool
	SetRequestId(v string) *CreatePolarClawChannelResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *CreatePolarClawChannelResponseBody
	GetRestarted() *bool
}

type CreatePolarClawChannelResponseBody struct {
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
	PluginInstalled *bool `json:"PluginInstalled,omitempty" xml:"PluginInstalled,omitempty"`
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s CreatePolarClawChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolarClawChannelResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreatePolarClawChannelResponseBody) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreatePolarClawChannelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePolarClawChannelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePolarClawChannelResponseBody) GetNpmPackage() *string {
	return s.NpmPackage
}

func (s *CreatePolarClawChannelResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *CreatePolarClawChannelResponseBody) GetPluginId() *string {
	return s.PluginId
}

func (s *CreatePolarClawChannelResponseBody) GetPluginInstalled() *bool {
	return s.PluginInstalled
}

func (s *CreatePolarClawChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolarClawChannelResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *CreatePolarClawChannelResponseBody) SetApplicationId(v string) *CreatePolarClawChannelResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreatePolarClawChannelResponseBody) SetChannelId(v string) *CreatePolarClawChannelResponseBody {
	s.ChannelId = &v
	return s
}

func (s *CreatePolarClawChannelResponseBody) SetCode(v int32) *CreatePolarClawChannelResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolarClawChannelResponseBody) SetMessage(v string) *CreatePolarClawChannelResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePolarClawChannelResponseBody) SetNpmPackage(v string) *CreatePolarClawChannelResponseBody {
	s.NpmPackage = &v
	return s
}

func (s *CreatePolarClawChannelResponseBody) SetOk(v bool) *CreatePolarClawChannelResponseBody {
	s.Ok = &v
	return s
}

func (s *CreatePolarClawChannelResponseBody) SetPluginId(v string) *CreatePolarClawChannelResponseBody {
	s.PluginId = &v
	return s
}

func (s *CreatePolarClawChannelResponseBody) SetPluginInstalled(v bool) *CreatePolarClawChannelResponseBody {
	s.PluginInstalled = &v
	return s
}

func (s *CreatePolarClawChannelResponseBody) SetRequestId(v string) *CreatePolarClawChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolarClawChannelResponseBody) SetRestarted(v bool) *CreatePolarClawChannelResponseBody {
	s.Restarted = &v
	return s
}

func (s *CreatePolarClawChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
