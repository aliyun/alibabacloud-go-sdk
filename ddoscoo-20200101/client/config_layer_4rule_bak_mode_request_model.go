// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RuleBakModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBakMode(v string) *ConfigLayer4RuleBakModeRequest
	GetBakMode() *string
	SetListeners(v string) *ConfigLayer4RuleBakModeRequest
	GetListeners() *string
}

type ConfigLayer4RuleBakModeRequest struct {
	// The back-to-origin mode. Valid values:
	//
	// - **0**: The default back-to-origin mode. Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the origin server IP address that you specified when you called [CreateNetworkRules](https://help.aliyun.com/document_detail/157482.html) to create the port forwarding rule.
	//
	// - **1**: The origin redundancy mode. Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the primary origin server IP address or secondary origin server IP address that you configured by calling [ConfigLayer4RulePolicy](https://help.aliyun.com/document_detail/312684.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BakMode *string `json:"BakMode,omitempty" xml:"BakMode,omitempty"`
	// The port forwarding rule to configure.
	//
	// This parameter is a string converted from a JSON array. Each element in the JSON array represents a created port forwarding rule. You can configure up to one port forwarding rule at a time.
	//
	// > You can call [DescribeNetworkRules](https://help.aliyun.com/document_detail/157484.html) to query all created port forwarding rules.
	//
	// Each port forwarding rule contains the following fields:
	//
	// - **InstanceId**: String type. Required. The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// - **Protocol**: String type. Required. The forwarding protocol type. Valid values: **tcp**, **udp**.
	//
	// - **FrontendPort**: Integer type. Required. The Redirection Port.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"InstanceId\\":\\"ddosDip-sg-4hr2b3l****\\",\\"FrontendPort\\":2020,\\"Protocol\\":\\"udp\\"}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s ConfigLayer4RuleBakModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleBakModeRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleBakModeRequest) GetBakMode() *string {
	return s.BakMode
}

func (s *ConfigLayer4RuleBakModeRequest) GetListeners() *string {
	return s.Listeners
}

func (s *ConfigLayer4RuleBakModeRequest) SetBakMode(v string) *ConfigLayer4RuleBakModeRequest {
	s.BakMode = &v
	return s
}

func (s *ConfigLayer4RuleBakModeRequest) SetListeners(v string) *ConfigLayer4RuleBakModeRequest {
	s.Listeners = &v
	return s
}

func (s *ConfigLayer4RuleBakModeRequest) Validate() error {
	return dara.Validate(s)
}
