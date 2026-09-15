// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpMarketItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMcpMarketItemsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMcpMarketItemsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListMcpMarketItemsResponseBodyItems) *ListMcpMarketItemsResponseBody
	GetItems() []*ListMcpMarketItemsResponseBodyItems
	SetMaxResults(v int32) *ListMcpMarketItemsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMcpMarketItemsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMcpMarketItemsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMcpMarketItemsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMcpMarketItemsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListMcpMarketItemsResponseBody
	GetTotalCount() *int64
}

type ListMcpMarketItemsResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of query results.
	Items []*ListMcpMarketItemsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records to return in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token used to retrieve the next page of results.
	//
	// example:
	//
	// 20
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that match the filter conditions.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMcpMarketItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpMarketItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpMarketItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMcpMarketItemsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMcpMarketItemsResponseBody) GetItems() []*ListMcpMarketItemsResponseBodyItems {
	return s.Items
}

func (s *ListMcpMarketItemsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpMarketItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMcpMarketItemsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpMarketItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpMarketItemsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcpMarketItemsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcpMarketItemsResponseBody) SetCode(v string) *ListMcpMarketItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpMarketItemsResponseBody) SetHttpStatusCode(v int32) *ListMcpMarketItemsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMcpMarketItemsResponseBody) SetItems(v []*ListMcpMarketItemsResponseBodyItems) *ListMcpMarketItemsResponseBody {
	s.Items = v
	return s
}

func (s *ListMcpMarketItemsResponseBody) SetMaxResults(v int32) *ListMcpMarketItemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMcpMarketItemsResponseBody) SetMessage(v string) *ListMcpMarketItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpMarketItemsResponseBody) SetNextToken(v string) *ListMcpMarketItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMcpMarketItemsResponseBody) SetRequestId(v string) *ListMcpMarketItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpMarketItemsResponseBody) SetSuccess(v bool) *ListMcpMarketItemsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMcpMarketItemsResponseBody) SetTotalCount(v int64) *ListMcpMarketItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMcpMarketItemsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcpMarketItemsResponseBodyItems struct {
	// The MCP marketplace template category.
	//
	// example:
	//
	// knowledge
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The MCP service description.
	//
	// example:
	//
	// An MCP service for querying knowledge bases
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display metadata of the template.
	DisplayMetadata map[string]interface{} `json:"displayMetadata,omitempty" xml:"displayMetadata,omitempty"`
	// The icon URL of the MCP marketplace template.
	//
	// example:
	//
	// https://example.com/mcp-icon.png
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// The number of times the template has been installed.
	//
	// example:
	//
	// 12
	InstallCount *int64 `json:"installCount,omitempty" xml:"installCount,omitempty"`
	// The MCP marketplace template ID.
	//
	// example:
	//
	// market-1
	MarketItemId *string `json:"marketItemId,omitempty" xml:"marketItemId,omitempty"`
	// The MCP type.
	//
	// example:
	//
	// CODE_PACKAGE
	McpType *string `json:"mcpType,omitempty" xml:"mcpType,omitempty"`
	// The MCP marketplace template name.
	//
	// example:
	//
	// Knowledge
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The official usage tag.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	OfficialTag *string `json:"officialTag,omitempty" xml:"officialTag,omitempty"`
	// The MCP protocol.
	//
	// example:
	//
	// StreamableHTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The usage instructions for the MCP marketplace template.
	//
	// example:
	//
	// # Knowledge\\nKnowledge base query service
	Readme *string `json:"readme,omitempty" xml:"readme,omitempty"`
	// The template schema version.
	//
	// example:
	//
	// 1.0
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The template input schema, represented as a JSON Schema string.
	//
	// example:
	//
	// {"type":"object","properties":{"addresses":{"type":"array","items":{"type":"string"}}}}
	TemplateInputSchema *string `json:"templateInputSchema,omitempty" xml:"templateInputSchema,omitempty"`
	// The MCP marketplace template version.
	//
	// example:
	//
	// 1.0.0
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s ListMcpMarketItemsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMcpMarketItemsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMcpMarketItemsResponseBodyItems) GetCategory() *string {
	return s.Category
}

func (s *ListMcpMarketItemsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListMcpMarketItemsResponseBodyItems) GetDisplayMetadata() map[string]interface{} {
	return s.DisplayMetadata
}

func (s *ListMcpMarketItemsResponseBodyItems) GetIconUrl() *string {
	return s.IconUrl
}

func (s *ListMcpMarketItemsResponseBodyItems) GetInstallCount() *int64 {
	return s.InstallCount
}

func (s *ListMcpMarketItemsResponseBodyItems) GetMarketItemId() *string {
	return s.MarketItemId
}

func (s *ListMcpMarketItemsResponseBodyItems) GetMcpType() *string {
	return s.McpType
}

func (s *ListMcpMarketItemsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListMcpMarketItemsResponseBodyItems) GetOfficialTag() *string {
	return s.OfficialTag
}

func (s *ListMcpMarketItemsResponseBodyItems) GetProtocol() *string {
	return s.Protocol
}

func (s *ListMcpMarketItemsResponseBodyItems) GetReadme() *string {
	return s.Readme
}

func (s *ListMcpMarketItemsResponseBodyItems) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *ListMcpMarketItemsResponseBodyItems) GetTemplateInputSchema() *string {
	return s.TemplateInputSchema
}

func (s *ListMcpMarketItemsResponseBodyItems) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListMcpMarketItemsResponseBodyItems) SetCategory(v string) *ListMcpMarketItemsResponseBodyItems {
	s.Category = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetDescription(v string) *ListMcpMarketItemsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetDisplayMetadata(v map[string]interface{}) *ListMcpMarketItemsResponseBodyItems {
	s.DisplayMetadata = v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetIconUrl(v string) *ListMcpMarketItemsResponseBodyItems {
	s.IconUrl = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetInstallCount(v int64) *ListMcpMarketItemsResponseBodyItems {
	s.InstallCount = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetMarketItemId(v string) *ListMcpMarketItemsResponseBodyItems {
	s.MarketItemId = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetMcpType(v string) *ListMcpMarketItemsResponseBodyItems {
	s.McpType = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetName(v string) *ListMcpMarketItemsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetOfficialTag(v string) *ListMcpMarketItemsResponseBodyItems {
	s.OfficialTag = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetProtocol(v string) *ListMcpMarketItemsResponseBodyItems {
	s.Protocol = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetReadme(v string) *ListMcpMarketItemsResponseBodyItems {
	s.Readme = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetSchemaVersion(v string) *ListMcpMarketItemsResponseBodyItems {
	s.SchemaVersion = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetTemplateInputSchema(v string) *ListMcpMarketItemsResponseBodyItems {
	s.TemplateInputSchema = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) SetTemplateVersion(v string) *ListMcpMarketItemsResponseBodyItems {
	s.TemplateVersion = &v
	return s
}

func (s *ListMcpMarketItemsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
