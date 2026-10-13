// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProductsResponseBody
	GetCode() *string
	SetData(v *ListProductsResponseBodyData) *ListProductsResponseBody
	GetData() *ListProductsResponseBodyData
	SetMaxResults(v int32) *ListProductsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListProductsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListProductsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListProductsResponseBody
	GetRequestId() *string
}

type ListProductsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListProductsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The maximum number of entries returned in this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The cursor token for the next page.
	//
	// example:
	//
	// next-token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProductsResponseBody) GetData() *ListProductsResponseBodyData {
	return s.Data
}

func (s *ListProductsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProductsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsResponseBody) SetCode(v string) *ListProductsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductsResponseBody) SetData(v *ListProductsResponseBodyData) *ListProductsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductsResponseBody) SetMaxResults(v int32) *ListProductsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductsResponseBody) SetMessage(v string) *ListProductsResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductsResponseBody) SetNextToken(v string) *ListProductsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductsResponseBodyData struct {
	// The list of API products. The list element is named productInfo, and each element contains the information of an API product.
	Content []*ListProductsResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProductsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyData) GetContent() []*ListProductsResponseBodyDataContent {
	return s.Content
}

func (s *ListProductsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductsResponseBodyData) SetContent(v []*ListProductsResponseBodyDataContent) *ListProductsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListProductsResponseBodyData) SetPageNumber(v int32) *ListProductsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListProductsResponseBodyData) SetPageSize(v int32) *ListProductsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProductsResponseBodyData) SetTotalCount(v int32) *ListProductsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListProductsResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsResponseBodyDataContent struct {
	// The agent-related configurations of the API product. The agentAPIConfig member includes agentProtocols (the list of protocols supported by the agent) and routes (the list of routes).
	AgentConfig *HiMarketAgentConfig `json:"agentConfig,omitempty" xml:"agentConfig,omitempty"`
	// Specifies whether subscription requests for the API product are automatically approved. This is the AutoApprove property of the API product resource.
	AutoApprove *bool `json:"autoApprove,omitempty" xml:"autoApprove,omitempty"`
	// The list of product categories to which the API product belongs. The list element is named category.
	Categories []*HiMarketProductCategory `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	// The description of the API product, which is the Description property of the API product resource.
	//
	// example:
	//
	// An API product that provides weather query capabilities
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The documentation of the API product, which is the Document property of the API product resource.
	//
	// example:
	//
	// The invocation methods, parameters, and billing instructions for this product
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// The capability configurations of the API product. Currently, it includes only the modelFeature (model capabilities) member, which contains model, maxTokens, temperature, streaming, webSearch, enableThinking, and enableMultiModal.
	Feature *HiMarketProductFeature `json:"feature,omitempty" xml:"feature,omitempty"`
	// The icon of the API product, which is the Icon property of the API product resource. It consists of two fields: type (the icon type) and value (the icon information).
	Icon *HiMarketIcon `json:"icon,omitempty" xml:"icon,omitempty"`
	// The MCP-related configurations of the API product. The members include mcpServerName (the MCP server name), tools, mcpServerConfig (path and domains), and meta (protocol).
	McpConfig *HiMarketMcpConfig `json:"mcpConfig,omitempty" xml:"mcpConfig,omitempty"`
	// The model API-related configurations of the API product. The modelAPIConfig member includes modelCategory (the AI model category), aiProtocols (the list of AI API protocols), and routes (the list of routes).
	ModelConfig *HiMarketModelConfig `json:"modelConfig,omitempty" xml:"modelConfig,omitempty"`
	// The name of the API product, which is the Name property of the API product resource.
	//
	// example:
	//
	// weather-query-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the API product, which is the primary ID of the API product resource. You can use it to locate the product in operations such as GetProduct, UpdateProduct, DeleteProduct, PublishProduct, SetProductCategories, and ListProductPublications.
	//
	// example:
	//
	// product-xxxxxxxx
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The status of the API product, which is the Status property of the API product resource. Valid values: Pending, Ready, and Published.
	//
	// example:
	//
	// Published
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the API product, which is the Type property of the API product resource. The valid values are not declared in the API definition. Refer to the types displayed in the console.
	//
	// example:
	//
	// API
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProductsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataContent) GetAgentConfig() *HiMarketAgentConfig {
	return s.AgentConfig
}

func (s *ListProductsResponseBodyDataContent) GetAutoApprove() *bool {
	return s.AutoApprove
}

func (s *ListProductsResponseBodyDataContent) GetCategories() []*HiMarketProductCategory {
	return s.Categories
}

func (s *ListProductsResponseBodyDataContent) GetDescription() *string {
	return s.Description
}

func (s *ListProductsResponseBodyDataContent) GetDocument() *string {
	return s.Document
}

func (s *ListProductsResponseBodyDataContent) GetFeature() *HiMarketProductFeature {
	return s.Feature
}

func (s *ListProductsResponseBodyDataContent) GetIcon() *HiMarketIcon {
	return s.Icon
}

func (s *ListProductsResponseBodyDataContent) GetMcpConfig() *HiMarketMcpConfig {
	return s.McpConfig
}

func (s *ListProductsResponseBodyDataContent) GetModelConfig() *HiMarketModelConfig {
	return s.ModelConfig
}

func (s *ListProductsResponseBodyDataContent) GetName() *string {
	return s.Name
}

func (s *ListProductsResponseBodyDataContent) GetProductId() *string {
	return s.ProductId
}

func (s *ListProductsResponseBodyDataContent) GetStatus() *string {
	return s.Status
}

func (s *ListProductsResponseBodyDataContent) GetType() *string {
	return s.Type
}

func (s *ListProductsResponseBodyDataContent) SetAgentConfig(v *HiMarketAgentConfig) *ListProductsResponseBodyDataContent {
	s.AgentConfig = v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetAutoApprove(v bool) *ListProductsResponseBodyDataContent {
	s.AutoApprove = &v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetCategories(v []*HiMarketProductCategory) *ListProductsResponseBodyDataContent {
	s.Categories = v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetDescription(v string) *ListProductsResponseBodyDataContent {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetDocument(v string) *ListProductsResponseBodyDataContent {
	s.Document = &v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetFeature(v *HiMarketProductFeature) *ListProductsResponseBodyDataContent {
	s.Feature = v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetIcon(v *HiMarketIcon) *ListProductsResponseBodyDataContent {
	s.Icon = v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetMcpConfig(v *HiMarketMcpConfig) *ListProductsResponseBodyDataContent {
	s.McpConfig = v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetModelConfig(v *HiMarketModelConfig) *ListProductsResponseBodyDataContent {
	s.ModelConfig = v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetName(v string) *ListProductsResponseBodyDataContent {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetProductId(v string) *ListProductsResponseBodyDataContent {
	s.ProductId = &v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetStatus(v string) *ListProductsResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *ListProductsResponseBodyDataContent) SetType(v string) *ListProductsResponseBodyDataContent {
	s.Type = &v
	return s
}

func (s *ListProductsResponseBodyDataContent) Validate() error {
	if s.AgentConfig != nil {
		if err := s.AgentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Categories != nil {
		for _, item := range s.Categories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
