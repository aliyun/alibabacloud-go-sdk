// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAtiAgentRegisterInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentDescription(v string) *CreateAtiAgentRegisterInfoRequest
	GetAgentDescription() *string
	SetAgentDisplayName(v string) *CreateAtiAgentRegisterInfoRequest
	GetAgentDisplayName() *string
	SetAgentHost(v string) *CreateAtiAgentRegisterInfoRequest
	GetAgentHost() *string
	SetAgentVersion(v string) *CreateAtiAgentRegisterInfoRequest
	GetAgentVersion() *string
	SetClientToken(v string) *CreateAtiAgentRegisterInfoRequest
	GetClientToken() *string
	SetEndpoints(v []*CreateAtiAgentRegisterInfoRequestEndpoints) *CreateAtiAgentRegisterInfoRequest
	GetEndpoints() []*CreateAtiAgentRegisterInfoRequestEndpoints
	SetRegistrantId(v string) *CreateAtiAgentRegisterInfoRequest
	GetRegistrantId() *string
}

type CreateAtiAgentRegisterInfoRequest struct {
	// The description of the agent capabilities.
	//
	// example:
	//
	// 支付服务
	AgentDescription *string `json:"AgentDescription,omitempty" xml:"AgentDescription,omitempty"`
	// The display name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试Agent
	AgentDisplayName *string `json:"AgentDisplayName,omitempty" xml:"AgentDisplayName,omitempty"`
	// The endpoint domain name through which the agent provides services.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	AgentHost *string `json:"AgentHost,omitempty" xml:"AgentHost,omitempty"`
	// The version of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.1
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// Provides idempotency. Within 3 minutes, the same value takes effect only once.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The endpoint information of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"EndpointValue\\":\\"http://www.baidu.com\\",\\"EndpointType\\":\\"http\\"}]
	Endpoints []*CreateAtiAgentRegisterInfoRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The ID of the verified registrant. Obtain this ID by invoking the identity verification API operation or from the ATS console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2072277378616354816
	RegistrantId *string `json:"RegistrantId,omitempty" xml:"RegistrantId,omitempty"`
}

func (s CreateAtiAgentRegisterInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAtiAgentRegisterInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateAtiAgentRegisterInfoRequest) GetAgentDescription() *string {
	return s.AgentDescription
}

func (s *CreateAtiAgentRegisterInfoRequest) GetAgentDisplayName() *string {
	return s.AgentDisplayName
}

func (s *CreateAtiAgentRegisterInfoRequest) GetAgentHost() *string {
	return s.AgentHost
}

func (s *CreateAtiAgentRegisterInfoRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *CreateAtiAgentRegisterInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAtiAgentRegisterInfoRequest) GetEndpoints() []*CreateAtiAgentRegisterInfoRequestEndpoints {
	return s.Endpoints
}

func (s *CreateAtiAgentRegisterInfoRequest) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *CreateAtiAgentRegisterInfoRequest) SetAgentDescription(v string) *CreateAtiAgentRegisterInfoRequest {
	s.AgentDescription = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequest) SetAgentDisplayName(v string) *CreateAtiAgentRegisterInfoRequest {
	s.AgentDisplayName = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequest) SetAgentHost(v string) *CreateAtiAgentRegisterInfoRequest {
	s.AgentHost = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequest) SetAgentVersion(v string) *CreateAtiAgentRegisterInfoRequest {
	s.AgentVersion = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequest) SetClientToken(v string) *CreateAtiAgentRegisterInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequest) SetEndpoints(v []*CreateAtiAgentRegisterInfoRequestEndpoints) *CreateAtiAgentRegisterInfoRequest {
	s.Endpoints = v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequest) SetRegistrantId(v string) *CreateAtiAgentRegisterInfoRequest {
	s.RegistrantId = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequest) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAtiAgentRegisterInfoRequestEndpoints struct {
	// The actual service address of the agent endpoint, which is the HTTPS entry point where the agent runs online. This is a required field with a maximum of 500 characters.
	//
	// Example: https://my-agent.example.com/mcp
	//
	// After a caller discovers this agent through DNS, it can use this URL to initiate a connection directly. This is the address where the agent is actually online.
	//
	// example:
	//
	// https://www.example.com/mcp
	AgentUrl *string `json:"AgentUrl,omitempty" xml:"AgentUrl,omitempty"`
	// An optional URL that points to the metadata description file of the agent (typically in JSON format). This allows other agents or clients to automatically discover the agent capabilities before connecting, including:
	//
	// - Functions supported by the agent
	//
	// - Input/output formats
	//
	// - Version information
	//
	// - Other capability declarations
	//
	// example:
	//
	// // https://my-agent.example.com/.well-known/agent.json
	//
	// {
	//
	//   "name": "智能客服助手",
	//
	//   "version": "1.2.0",
	//
	//   "protocol": "MCP",
	//
	//   "description": "提供产品咨询、订单查询、售后服务的智能客服Agent",
	//
	//   "capabilities": {
	//
	//     "tools": [
	//
	//       { "name": "order_query", "description": "查询订单状态" },
	//
	//       { "name": "product_search", "description": "搜索产品信息" },
	//
	//       { "name": "refund_request", "description": "发起退款申请" }
	//
	//     ],
	//
	//     "resources": [
	//
	//       { "uri": "faq://knowledge-base", "description": "FAQ知识库" }
	//
	//     ]
	//
	//   },
	//
	//   "endpoint": {
	//
	//     "url": "https://my-agent.example.com/mcp",
	//
	//     "transport": ["STREAMABLE-HTTP", "SSE"]
	//
	//   },
	//
	//   "policy": {
	//
	//     "authentication": "mTLS",
	//
	//     "rateLimit": "100/min"
	//
	//   }
	//
	// }
	MetadataUrl *string `json:"MetadataUrl,omitempty" xml:"MetadataUrl,omitempty"`
	// The communication protocol standard that the agent endpoint follows. This determines how the invoker interacts with the agent.
	//
	// Valid values:
	//
	// - MCP: Model Context Protocol, an agent tool invocation protocol developed by Anthropic.
	//
	// - A2A: Agent-to-Agent Protocol, a cross-agent communication protocol developed by Google.
	//
	// - OpenAPI: Standard RESTful API specification (Swagger/OpenAPI).
	//
	// When other agents or clients see this protocol identity, they know which method to use to communicate with the agent. For example, MCP uses the MCP SDK, A2A uses the A2A SDK, and OpenAPI uses standard HTTP requests.
	//
	// example:
	//
	// A2A
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The transport methods.
	//
	// example:
	//
	// STREAMABLE-HTTP
	Transports []*string `json:"Transports,omitempty" xml:"Transports,omitempty" type:"Repeated"`
}

func (s CreateAtiAgentRegisterInfoRequestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s CreateAtiAgentRegisterInfoRequestEndpoints) GoString() string {
	return s.String()
}

func (s *CreateAtiAgentRegisterInfoRequestEndpoints) GetAgentUrl() *string {
	return s.AgentUrl
}

func (s *CreateAtiAgentRegisterInfoRequestEndpoints) GetMetadataUrl() *string {
	return s.MetadataUrl
}

func (s *CreateAtiAgentRegisterInfoRequestEndpoints) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateAtiAgentRegisterInfoRequestEndpoints) GetTransports() []*string {
	return s.Transports
}

func (s *CreateAtiAgentRegisterInfoRequestEndpoints) SetAgentUrl(v string) *CreateAtiAgentRegisterInfoRequestEndpoints {
	s.AgentUrl = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequestEndpoints) SetMetadataUrl(v string) *CreateAtiAgentRegisterInfoRequestEndpoints {
	s.MetadataUrl = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequestEndpoints) SetProtocol(v string) *CreateAtiAgentRegisterInfoRequestEndpoints {
	s.Protocol = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequestEndpoints) SetTransports(v []*string) *CreateAtiAgentRegisterInfoRequestEndpoints {
	s.Transports = v
	return s
}

func (s *CreateAtiAgentRegisterInfoRequestEndpoints) Validate() error {
	return dara.Validate(s)
}
