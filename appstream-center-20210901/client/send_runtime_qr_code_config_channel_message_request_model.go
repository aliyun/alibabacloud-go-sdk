// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRuntimeQrCodeConfigChannelMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *SendRuntimeQrCodeConfigChannelMessageRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *SendRuntimeQrCodeConfigChannelMessageRequest
	GetAgentProvider() *string
	SetCode(v string) *SendRuntimeQrCodeConfigChannelMessageRequest
	GetCode() *string
	SetRuntimeIds(v []*string) *SendRuntimeQrCodeConfigChannelMessageRequest
	GetRuntimeIds() []*string
	SetRuntimeType(v string) *SendRuntimeQrCodeConfigChannelMessageRequest
	GetRuntimeType() *string
}

type SendRuntimeQrCodeConfigChannelMessageRequest struct {
	// The agent platform type.
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
	// The third-party channel code, such as dingtalk-connector (DingTalk), feishu (Lark), or wecom (WeCom). If the current agent platform and provider do not support the channel, the error code ChannelConfig.NotSupported is returned, and no messages are sent.
	//
	// This parameter is required.
	//
	// example:
	//
	// feishu
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of the runtime resources, such as cloud desktop IDs. You can specify 1 to 200 IDs.
	//
	// This parameter is required.
	RuntimeIds []*string `json:"RuntimeIds,omitempty" xml:"RuntimeIds,omitempty" type:"Repeated"`
	// The runtime environment type of the runtime resource. Currently, only CloudDesktop is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s SendRuntimeQrCodeConfigChannelMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendRuntimeQrCodeConfigChannelMessageRequest) GoString() string {
	return s.String()
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) GetCode() *string {
	return s.Code
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) GetRuntimeIds() []*string {
	return s.RuntimeIds
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) SetAgentPlatform(v string) *SendRuntimeQrCodeConfigChannelMessageRequest {
	s.AgentPlatform = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) SetAgentProvider(v string) *SendRuntimeQrCodeConfigChannelMessageRequest {
	s.AgentProvider = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) SetCode(v string) *SendRuntimeQrCodeConfigChannelMessageRequest {
	s.Code = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) SetRuntimeIds(v []*string) *SendRuntimeQrCodeConfigChannelMessageRequest {
	s.RuntimeIds = v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) SetRuntimeType(v string) *SendRuntimeQrCodeConfigChannelMessageRequest {
	s.RuntimeType = &v
	return s
}

func (s *SendRuntimeQrCodeConfigChannelMessageRequest) Validate() error {
	return dara.Validate(s)
}
