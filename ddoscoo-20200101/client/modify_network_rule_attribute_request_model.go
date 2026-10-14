// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkRuleAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyNetworkRuleAttributeRequest
	GetConfig() *string
	SetForwardProtocol(v string) *ModifyNetworkRuleAttributeRequest
	GetForwardProtocol() *string
	SetFrontendPort(v int32) *ModifyNetworkRuleAttributeRequest
	GetFrontendPort() *int32
	SetInstanceId(v string) *ModifyNetworkRuleAttributeRequest
	GetInstanceId() *string
	SetModule(v string) *ModifyNetworkRuleAttributeRequest
	GetModule() *string
}

type ModifyNetworkRuleAttributeRequest struct {
	// The advanced settings for the port forwarding rule. Specify the settings as a JSON string. The structure is described below.
	//
	// The structure of \\`TcpConfig\\` is as follows:
	//
	// - **PersistenceTimeout**: An integer. This parameter is required. The timeout period for session persistence. Unit: seconds. Valid values: **30*	- to **3600**. The default value is **0**, which disables session persistence.
	//
	// - **Synproxy**: A string. This parameter is required. The status of the False Source setting for the DDoS mitigation policy. Valid values: \\`off\\` and \\`on\\`.
	//
	// - **NodataConn**: A string. This parameter is required. The status of the Empty Connection setting for the DDoS mitigation policy. Valid values: \\`off\\` and \\`on\\`.
	//
	// - **Sla**: A struct. This parameter is required. The configuration for destination rate limiting. For more information, see the \\`Sla\\` section below.
	//
	// - **Slimit**: A struct. This parameter is required. The configuration for source rate limiting. For more information, see the \\`Slimit\\` section below.
	//
	// - **PayloadLen**: A struct. This parameter is required. The packet length filter configuration. For more information, see the \\`PayloadLen\\` section below.
	//
	// The structure of \\`Sla\\` is as follows:
	//
	// - **Cps**: An integer. This parameter is required. The maximum number of new connections per second to a destination IP address. Valid values: 100 to 100,000.
	//
	// - **Maxconn**: An integer. This parameter is required. The maximum number of concurrent connections to a destination IP address. Valid values: 1,000 to 1,000,000.
	//
	// - **CpsEnable**: An integer. This parameter is required. Specifies whether to enable the limit on new connections. Valid values: 0 (disabled) and 1 (enabled). The default value is 1.
	//
	// - **MaxconnEnable**: An integer. This parameter is required. Specifies whether to enable the limit on concurrent connections. Valid values: 0 (disabled) and 1 (enabled). The default value is 1.
	//
	// The structure of \\`Slimit\\` is as follows:
	//
	// - **Cps**: An integer. This parameter is required. The maximum number of new connections per second from a source IP address. Valid values: 1 to 50,000.
	//
	// - **Maxconn**: An integer. This parameter is required. The maximum number of concurrent connections from a source IP address. Valid values: 1 to 50,000.
	//
	// - **CpsEnable**: An integer. This parameter is required. Specifies whether to enable the limit on new connections. Valid values: 0 (disabled) and 1 (enabled). The default value is 1.
	//
	// - **MaxconnEnable**: An integer. This parameter is required. Specifies whether to enable the limit on concurrent connections. Valid values: 0 (disabled) and 1 (enabled). The default value is 1.
	//
	// - **CpsMode**: An integer. This parameter is required. The mode for the new connection rate limit from a source IP address. Valid values: 1 (off) and 2 (auto).
	//
	// The structure of \\`PayloadLen\\` is as follows:
	//
	// - **Min**: An integer. This parameter is required. The minimum packet length. Valid values: 0 to 1,500.
	//
	// - **Max**: An integer. This parameter is required. The maximum packet length. Valid values: 0 to 1,500.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Slimit":{"CpsEnable":1,"MaxconnEnable":1,"Cps":1,"Maxconn":1},"Sla":{"CpsEnable":1,"MaxconnEnable":1,"Cps":100,"Maxconn":1000},"PayloadLen":{"Min":0,"Max":1500}}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The forwarding protocol. Valid values:
	//
	// - **tcp**
	//
	// - **udp**
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	// The forwarding port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// > Call [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the IDs of all Anti-DDoS Pro and Anti-DDoS Premium instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The module for Layer 7 port support. The default value is sla. Valid value:
	//
	// **sla**: Speed Limit for Destination
	//
	// example:
	//
	// sla
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
}

func (s ModifyNetworkRuleAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkRuleAttributeRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyNetworkRuleAttributeRequest) GetForwardProtocol() *string {
	return s.ForwardProtocol
}

func (s *ModifyNetworkRuleAttributeRequest) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *ModifyNetworkRuleAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyNetworkRuleAttributeRequest) GetModule() *string {
	return s.Module
}

func (s *ModifyNetworkRuleAttributeRequest) SetConfig(v string) *ModifyNetworkRuleAttributeRequest {
	s.Config = &v
	return s
}

func (s *ModifyNetworkRuleAttributeRequest) SetForwardProtocol(v string) *ModifyNetworkRuleAttributeRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *ModifyNetworkRuleAttributeRequest) SetFrontendPort(v int32) *ModifyNetworkRuleAttributeRequest {
	s.FrontendPort = &v
	return s
}

func (s *ModifyNetworkRuleAttributeRequest) SetInstanceId(v string) *ModifyNetworkRuleAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyNetworkRuleAttributeRequest) SetModule(v string) *ModifyNetworkRuleAttributeRequest {
	s.Module = &v
	return s
}

func (s *ModifyNetworkRuleAttributeRequest) Validate() error {
	return dara.Validate(s)
}
