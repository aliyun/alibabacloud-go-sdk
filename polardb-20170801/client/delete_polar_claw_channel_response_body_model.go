// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeletePolarClawChannelResponseBody
	GetApplicationId() *string
	SetChannelId(v string) *DeletePolarClawChannelResponseBody
	GetChannelId() *string
	SetCode(v int32) *DeletePolarClawChannelResponseBody
	GetCode() *int32
	SetMessage(v string) *DeletePolarClawChannelResponseBody
	GetMessage() *string
	SetOk(v bool) *DeletePolarClawChannelResponseBody
	GetOk() *bool
	SetPluginId(v string) *DeletePolarClawChannelResponseBody
	GetPluginId() *string
	SetPluginUninstalled(v bool) *DeletePolarClawChannelResponseBody
	GetPluginUninstalled() *bool
	SetRequestId(v string) *DeletePolarClawChannelResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *DeletePolarClawChannelResponseBody
	GetRestarted() *bool
}

type DeletePolarClawChannelResponseBody struct {
	// **The application ID.**
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the deleted channel.
	//
	// example:
	//
	// feishu
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// The ID of the uninstalled plugin.
	//
	// example:
	//
	// openclaw-lark
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// Indicates whether the channel plugin was uninstalled.
	//
	// example:
	//
	// true
	PluginUninstalled *bool `json:"PluginUninstalled,omitempty" xml:"PluginUninstalled,omitempty"`
	// **The request ID.**
	//
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the gateway was restarted.
	//
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s DeletePolarClawChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolarClawChannelResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeletePolarClawChannelResponseBody) GetChannelId() *string {
	return s.ChannelId
}

func (s *DeletePolarClawChannelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeletePolarClawChannelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePolarClawChannelResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *DeletePolarClawChannelResponseBody) GetPluginId() *string {
	return s.PluginId
}

func (s *DeletePolarClawChannelResponseBody) GetPluginUninstalled() *bool {
	return s.PluginUninstalled
}

func (s *DeletePolarClawChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolarClawChannelResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *DeletePolarClawChannelResponseBody) SetApplicationId(v string) *DeletePolarClawChannelResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DeletePolarClawChannelResponseBody) SetChannelId(v string) *DeletePolarClawChannelResponseBody {
	s.ChannelId = &v
	return s
}

func (s *DeletePolarClawChannelResponseBody) SetCode(v int32) *DeletePolarClawChannelResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolarClawChannelResponseBody) SetMessage(v string) *DeletePolarClawChannelResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePolarClawChannelResponseBody) SetOk(v bool) *DeletePolarClawChannelResponseBody {
	s.Ok = &v
	return s
}

func (s *DeletePolarClawChannelResponseBody) SetPluginId(v string) *DeletePolarClawChannelResponseBody {
	s.PluginId = &v
	return s
}

func (s *DeletePolarClawChannelResponseBody) SetPluginUninstalled(v bool) *DeletePolarClawChannelResponseBody {
	s.PluginUninstalled = &v
	return s
}

func (s *DeletePolarClawChannelResponseBody) SetRequestId(v string) *DeletePolarClawChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolarClawChannelResponseBody) SetRestarted(v bool) *DeletePolarClawChannelResponseBody {
	s.Restarted = &v
	return s
}

func (s *DeletePolarClawChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
