// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchConfigRuntimeChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *BatchConfigRuntimeChannelRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *BatchConfigRuntimeChannelRequest
	GetAgentProvider() *string
	SetChannelConfigs(v []*BatchConfigRuntimeChannelRequestChannelConfigs) *BatchConfigRuntimeChannelRequest
	GetChannelConfigs() []*BatchConfigRuntimeChannelRequestChannelConfigs
	SetRuntimeType(v string) *BatchConfigRuntimeChannelRequest
	GetRuntimeType() *string
}

type BatchConfigRuntimeChannelRequest struct {
	// The agent platform, such as ENTERPRISE. If the value is invalid, the error code InvalidParameter.AgentPlatform is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider.
	//
	// - JVS Computer: Enter OpenClaw.
	//
	// - OpenClaw: Enter OpenClaw.
	//
	// - Hermes Agent: Enter HermesAgent.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The list of channel configurations. A maximum of 200 configurations are supported. Each configuration item specifies an agent runtime and the third-party channel to be configured for it. Different channel configurations can be delivered to different runtimes. Multiple configuration items for the same runtime are processed sequentially.
	//
	// This parameter is required.
	ChannelConfigs []*BatchConfigRuntimeChannelRequestChannelConfigs `json:"ChannelConfigs,omitempty" xml:"ChannelConfigs,omitempty" type:"Repeated"`
	// The agent runtime type, such as CloudDesktop (cloud desktop).
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s BatchConfigRuntimeChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchConfigRuntimeChannelRequest) GoString() string {
	return s.String()
}

func (s *BatchConfigRuntimeChannelRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *BatchConfigRuntimeChannelRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *BatchConfigRuntimeChannelRequest) GetChannelConfigs() []*BatchConfigRuntimeChannelRequestChannelConfigs {
	return s.ChannelConfigs
}

func (s *BatchConfigRuntimeChannelRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *BatchConfigRuntimeChannelRequest) SetAgentPlatform(v string) *BatchConfigRuntimeChannelRequest {
	s.AgentPlatform = &v
	return s
}

func (s *BatchConfigRuntimeChannelRequest) SetAgentProvider(v string) *BatchConfigRuntimeChannelRequest {
	s.AgentProvider = &v
	return s
}

func (s *BatchConfigRuntimeChannelRequest) SetChannelConfigs(v []*BatchConfigRuntimeChannelRequestChannelConfigs) *BatchConfigRuntimeChannelRequest {
	s.ChannelConfigs = v
	return s
}

func (s *BatchConfigRuntimeChannelRequest) SetRuntimeType(v string) *BatchConfigRuntimeChannelRequest {
	s.RuntimeType = &v
	return s
}

func (s *BatchConfigRuntimeChannelRequest) Validate() error {
	if s.ChannelConfigs != nil {
		for _, item := range s.ChannelConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchConfigRuntimeChannelRequestChannelConfigs struct {
	// The third-party channel code, such as dingtalk-connector (DingTalk), feishu (Lark), and wecom (WeCom). System channel codes are not case-sensitive. Configuring wechat and jvsclaw is not supported. If these values are passed, the entire request fails and the error code ChannelConfig.NotSupported is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// dingtalk-connector
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The JSON string for the channel configuration.
	//
	// - Simple mode: only appKey and appSecret are allowed.
	//
	// - Custom mode: all configuration items required for the channel under the current agent provider must be included.
	//
	// If the string is not valid JSON, the error code InvalidParameter.ValueInvalid is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"appKey":"testKey","appSecret":"testSecret"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The configuration mode. Valid values:
	//
	// - Simple: the simple configuration mode.
	//
	// - Custom: the custom configuration mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// Simple
	ConfigMode *string `json:"ConfigMode,omitempty" xml:"ConfigMode,omitempty"`
	// The channel name.
	//
	// If this parameter is not specified, the system automatically uses the code as the name.
	//
	// example:
	//
	// DingTalk
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The agent runtime ID. The ID mappings are as follows:
	//
	// - JVS Computer: the JVS Computer ID, in the format of jvs-xxxx.
	//
	// - OpenClaw: the cloud desktop ID, in the format of ecd-xxxx.
	//
	// - Hermes Agent: the Hermes Agent ID, in the format of jvs-xxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-xgads
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
}

func (s BatchConfigRuntimeChannelRequestChannelConfigs) String() string {
	return dara.Prettify(s)
}

func (s BatchConfigRuntimeChannelRequestChannelConfigs) GoString() string {
	return s.String()
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) GetCode() *string {
	return s.Code
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) GetConfig() *string {
	return s.Config
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) GetConfigMode() *string {
	return s.ConfigMode
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) GetName() *string {
	return s.Name
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) SetCode(v string) *BatchConfigRuntimeChannelRequestChannelConfigs {
	s.Code = &v
	return s
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) SetConfig(v string) *BatchConfigRuntimeChannelRequestChannelConfigs {
	s.Config = &v
	return s
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) SetConfigMode(v string) *BatchConfigRuntimeChannelRequestChannelConfigs {
	s.ConfigMode = &v
	return s
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) SetName(v string) *BatchConfigRuntimeChannelRequestChannelConfigs {
	s.Name = &v
	return s
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) SetRuntimeId(v string) *BatchConfigRuntimeChannelRequestChannelConfigs {
	s.RuntimeId = &v
	return s
}

func (s *BatchConfigRuntimeChannelRequestChannelConfigs) Validate() error {
	return dara.Validate(s)
}
