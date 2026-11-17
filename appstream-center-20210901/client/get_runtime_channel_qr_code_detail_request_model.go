// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuntimeChannelQrCodeDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *GetRuntimeChannelQrCodeDetailRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *GetRuntimeChannelQrCodeDetailRequest
	GetAgentProvider() *string
	SetCode(v string) *GetRuntimeChannelQrCodeDetailRequest
	GetCode() *string
	SetRuntimeId(v string) *GetRuntimeChannelQrCodeDetailRequest
	GetRuntimeId() *string
	SetRuntimeType(v string) *GetRuntimeChannelQrCodeDetailRequest
	GetRuntimeType() *string
}

type GetRuntimeChannelQrCodeDetailRequest struct {
	// The agent platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The channel code. It must be the same as the channel code used when initiating the QR code configuration channel, such as dingtalk-connector (DingTalk), feishu (Lark), or wecom (WeCom).
	//
	// This parameter is required.
	//
	// example:
	//
	// dingtalk-connector
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the runtime, such as the ID of a cloud desktop.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-sdfdsx
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// The runtime environment type of the runtime resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s GetRuntimeChannelQrCodeDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeChannelQrCodeDetailRequest) GoString() string {
	return s.String()
}

func (s *GetRuntimeChannelQrCodeDetailRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *GetRuntimeChannelQrCodeDetailRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *GetRuntimeChannelQrCodeDetailRequest) GetCode() *string {
	return s.Code
}

func (s *GetRuntimeChannelQrCodeDetailRequest) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *GetRuntimeChannelQrCodeDetailRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *GetRuntimeChannelQrCodeDetailRequest) SetAgentPlatform(v string) *GetRuntimeChannelQrCodeDetailRequest {
	s.AgentPlatform = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailRequest) SetAgentProvider(v string) *GetRuntimeChannelQrCodeDetailRequest {
	s.AgentProvider = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailRequest) SetCode(v string) *GetRuntimeChannelQrCodeDetailRequest {
	s.Code = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailRequest) SetRuntimeId(v string) *GetRuntimeChannelQrCodeDetailRequest {
	s.RuntimeId = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailRequest) SetRuntimeType(v string) *GetRuntimeChannelQrCodeDetailRequest {
	s.RuntimeType = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailRequest) Validate() error {
	return dara.Validate(s)
}
