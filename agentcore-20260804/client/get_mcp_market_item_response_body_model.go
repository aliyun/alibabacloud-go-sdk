// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpMarketItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpMarketItemResponseBody
	GetCode() *string
	SetData(v *GetMcpMarketItemResponseBodyData) *GetMcpMarketItemResponseBody
	GetData() *GetMcpMarketItemResponseBodyData
	SetHttpStatusCode(v int32) *GetMcpMarketItemResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMcpMarketItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpMarketItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpMarketItemResponseBody
	GetSuccess() *bool
}

type GetMcpMarketItemResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetMcpMarketItemResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMcpMarketItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpMarketItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpMarketItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpMarketItemResponseBody) GetData() *GetMcpMarketItemResponseBodyData {
	return s.Data
}

func (s *GetMcpMarketItemResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMcpMarketItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpMarketItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpMarketItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpMarketItemResponseBody) SetCode(v string) *GetMcpMarketItemResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpMarketItemResponseBody) SetData(v *GetMcpMarketItemResponseBodyData) *GetMcpMarketItemResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpMarketItemResponseBody) SetHttpStatusCode(v int32) *GetMcpMarketItemResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpMarketItemResponseBody) SetMessage(v string) *GetMcpMarketItemResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpMarketItemResponseBody) SetRequestId(v string) *GetMcpMarketItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpMarketItemResponseBody) SetSuccess(v bool) *GetMcpMarketItemResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpMarketItemResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpMarketItemResponseBodyData struct {
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
	// The template display metadata.
	DisplayMetadata map[string]interface{} `json:"displayMetadata,omitempty" xml:"displayMetadata,omitempty"`
	// The multilingual display content, organized by BCP-47 language tags. Falls back to default fields when the specified language is not matched.
	I18n map[string]*DataI18nValue `json:"i18n,omitempty" xml:"i18n,omitempty"`
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

func (s GetMcpMarketItemResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpMarketItemResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpMarketItemResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *GetMcpMarketItemResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetMcpMarketItemResponseBodyData) GetDisplayMetadata() map[string]interface{} {
	return s.DisplayMetadata
}

func (s *GetMcpMarketItemResponseBodyData) GetI18n() map[string]*DataI18nValue {
	return s.I18n
}

func (s *GetMcpMarketItemResponseBodyData) GetIconUrl() *string {
	return s.IconUrl
}

func (s *GetMcpMarketItemResponseBodyData) GetInstallCount() *int64 {
	return s.InstallCount
}

func (s *GetMcpMarketItemResponseBodyData) GetMarketItemId() *string {
	return s.MarketItemId
}

func (s *GetMcpMarketItemResponseBodyData) GetMcpType() *string {
	return s.McpType
}

func (s *GetMcpMarketItemResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMcpMarketItemResponseBodyData) GetOfficialTag() *string {
	return s.OfficialTag
}

func (s *GetMcpMarketItemResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetMcpMarketItemResponseBodyData) GetReadme() *string {
	return s.Readme
}

func (s *GetMcpMarketItemResponseBodyData) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *GetMcpMarketItemResponseBodyData) GetTemplateInputSchema() *string {
	return s.TemplateInputSchema
}

func (s *GetMcpMarketItemResponseBodyData) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetMcpMarketItemResponseBodyData) SetCategory(v string) *GetMcpMarketItemResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetDescription(v string) *GetMcpMarketItemResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetDisplayMetadata(v map[string]interface{}) *GetMcpMarketItemResponseBodyData {
	s.DisplayMetadata = v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetI18n(v map[string]*DataI18nValue) *GetMcpMarketItemResponseBodyData {
	s.I18n = v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetIconUrl(v string) *GetMcpMarketItemResponseBodyData {
	s.IconUrl = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetInstallCount(v int64) *GetMcpMarketItemResponseBodyData {
	s.InstallCount = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetMarketItemId(v string) *GetMcpMarketItemResponseBodyData {
	s.MarketItemId = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetMcpType(v string) *GetMcpMarketItemResponseBodyData {
	s.McpType = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetName(v string) *GetMcpMarketItemResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetOfficialTag(v string) *GetMcpMarketItemResponseBodyData {
	s.OfficialTag = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetProtocol(v string) *GetMcpMarketItemResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetReadme(v string) *GetMcpMarketItemResponseBodyData {
	s.Readme = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetSchemaVersion(v string) *GetMcpMarketItemResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetTemplateInputSchema(v string) *GetMcpMarketItemResponseBodyData {
	s.TemplateInputSchema = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) SetTemplateVersion(v string) *GetMcpMarketItemResponseBodyData {
	s.TemplateVersion = &v
	return s
}

func (s *GetMcpMarketItemResponseBodyData) Validate() error {
	return dara.Validate(s)
}
