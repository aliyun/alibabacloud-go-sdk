// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProductResponseBody
	GetCode() *string
	SetData(v *GetProductResponseBodyData) *GetProductResponseBody
	GetData() *GetProductResponseBodyData
	SetMessage(v string) *GetProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetProductResponseBody
	GetRequestId() *string
}

type GetProductResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProductResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProductResponseBody) GetData() *GetProductResponseBodyData {
	return s.Data
}

func (s *GetProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProductResponseBody) SetCode(v string) *GetProductResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductResponseBody) SetData(v *GetProductResponseBodyData) *GetProductResponseBody {
	s.Data = v
	return s
}

func (s *GetProductResponseBody) SetMessage(v string) *GetProductResponseBody {
	s.Message = &v
	return s
}

func (s *GetProductResponseBody) SetRequestId(v string) *GetProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProductResponseBodyData struct {
	// The Agent-related configurations of the API product. The agentAPIConfig member includes agentProtocols, which is a list of protocols supported by the agent, and routes, which is a list of routes.
	AgentConfig *HiMarketAgentConfig `json:"agentConfig,omitempty" xml:"agentConfig,omitempty"`
	// Indicates whether subscription requests for the product are automatically approved.
	AutoApprove *bool `json:"autoApprove,omitempty" xml:"autoApprove,omitempty"`
	// The description of the API product.
	//
	// example:
	//
	// An API product that provides weather query capabilities
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The documentation content of the API product.
	//
	// example:
	//
	// The invocation methods, parameters, and billing instructions for this product
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// The capability configurations of the API product. Currently, this includes only the modelFeature member for model capabilities, which contains model, maxTokens, temperature, streaming, webSearch, enableThinking, and enableMultiModal.
	Feature *HiMarketProductFeature `json:"feature,omitempty" xml:"feature,omitempty"`
	// The icon configurations of the API product, which consist of two fields: type for the icon type and value for the icon information.
	Icon *HiMarketIcon `json:"icon,omitempty" xml:"icon,omitempty"`
	// The MCP-related configurations of the API product. The members include mcpServerName for the MCP server name, tools, mcpServerConfig for the path and domains, and meta for the protocol.
	McpConfig *HiMarketMcpConfig `json:"mcpConfig,omitempty" xml:"mcpConfig,omitempty"`
	// The model API-related configurations of the API product. The modelAPIConfig member includes modelCategory for the AI model category, aiProtocols for the list of AI API protocols, and routes for the list of routes.
	ModelConfig *HiMarketModelConfig `json:"modelConfig,omitempty" xml:"modelConfig,omitempty"`
	// The name of the API product.
	//
	// example:
	//
	// weather-query-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the API product, which is the primary ID of the API product resource. You can use this ID to locate the product in operations such as GetProduct, UpdateProduct, DeleteProduct, PublishProduct, SetProductCategories, and GetProductRef.
	//
	// example:
	//
	// p-xxxxxxxx
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The ID of the region to which the API product belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The status of the API product. Valid values: Pending, Ready, and Published.
	//
	// example:
	//
	// Published
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the API product. The set of valid values is not declared in the API definition. Refer to the types displayed in the console.
	//
	// example:
	//
	// API
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetProductResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductResponseBodyData) GetAgentConfig() *HiMarketAgentConfig {
	return s.AgentConfig
}

func (s *GetProductResponseBodyData) GetAutoApprove() *bool {
	return s.AutoApprove
}

func (s *GetProductResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetProductResponseBodyData) GetDocument() *string {
	return s.Document
}

func (s *GetProductResponseBodyData) GetFeature() *HiMarketProductFeature {
	return s.Feature
}

func (s *GetProductResponseBodyData) GetIcon() *HiMarketIcon {
	return s.Icon
}

func (s *GetProductResponseBodyData) GetMcpConfig() *HiMarketMcpConfig {
	return s.McpConfig
}

func (s *GetProductResponseBodyData) GetModelConfig() *HiMarketModelConfig {
	return s.ModelConfig
}

func (s *GetProductResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetProductResponseBodyData) GetProductId() *string {
	return s.ProductId
}

func (s *GetProductResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetProductResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetProductResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetProductResponseBodyData) SetAgentConfig(v *HiMarketAgentConfig) *GetProductResponseBodyData {
	s.AgentConfig = v
	return s
}

func (s *GetProductResponseBodyData) SetAutoApprove(v bool) *GetProductResponseBodyData {
	s.AutoApprove = &v
	return s
}

func (s *GetProductResponseBodyData) SetDescription(v string) *GetProductResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetProductResponseBodyData) SetDocument(v string) *GetProductResponseBodyData {
	s.Document = &v
	return s
}

func (s *GetProductResponseBodyData) SetFeature(v *HiMarketProductFeature) *GetProductResponseBodyData {
	s.Feature = v
	return s
}

func (s *GetProductResponseBodyData) SetIcon(v *HiMarketIcon) *GetProductResponseBodyData {
	s.Icon = v
	return s
}

func (s *GetProductResponseBodyData) SetMcpConfig(v *HiMarketMcpConfig) *GetProductResponseBodyData {
	s.McpConfig = v
	return s
}

func (s *GetProductResponseBodyData) SetModelConfig(v *HiMarketModelConfig) *GetProductResponseBodyData {
	s.ModelConfig = v
	return s
}

func (s *GetProductResponseBodyData) SetName(v string) *GetProductResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProductResponseBodyData) SetProductId(v string) *GetProductResponseBodyData {
	s.ProductId = &v
	return s
}

func (s *GetProductResponseBodyData) SetRegionId(v string) *GetProductResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetProductResponseBodyData) SetStatus(v string) *GetProductResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetProductResponseBodyData) SetType(v string) *GetProductResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetProductResponseBodyData) Validate() error {
	if s.AgentConfig != nil {
		if err := s.AgentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Feature != nil {
		if err := s.Feature.Validate(); err != nil {
			return err
		}
	}
	if s.Icon != nil {
		if err := s.Icon.Validate(); err != nil {
			return err
		}
	}
	if s.McpConfig != nil {
		if err := s.McpConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ModelConfig != nil {
		if err := s.ModelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
