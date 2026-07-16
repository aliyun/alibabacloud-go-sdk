// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *AddFileShrinkRequest
	GetCategoryId() *string
	SetCategoryType(v string) *AddFileShrinkRequest
	GetCategoryType() *string
	SetLeaseId(v string) *AddFileShrinkRequest
	GetLeaseId() *string
	SetOriginalFileUrl(v string) *AddFileShrinkRequest
	GetOriginalFileUrl() *string
	SetParser(v string) *AddFileShrinkRequest
	GetParser() *string
	SetParserConfigShrink(v string) *AddFileShrinkRequest
	GetParserConfigShrink() *string
	SetTagsShrink(v string) *AddFileShrinkRequest
	GetTagsShrink() *string
}

type AddFileShrinkRequest struct {
	// <props="china">
	//
	// - When CategoryType is set to UNSTRUCTURED, set this parameter to the category ID of the uploaded file, which is the `CategoryId` returned by the **AddCategory*	- operation. You can also go to [Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center), click the File tab, and then click the ID icon next to the category name to obtain the category ID. You can set this parameter to default to use the system-created default category.
	//
	// - When CategoryType is set to SESSION_FILE, set this parameter to "default".
	//
	//
	// <props="intl">
	//
	// Set this parameter to the category ID of the uploaded file, which is the `CategoryId` returned by the **AddCategory*	- operation. You can also go to [Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center), click the File tab, and then click the ID icon next to the category name to obtain the category ID. You can set this parameter to default to use the system-created default category.
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The category type. This parameter is optional. Default value: UNSTRUCTURED. Valid values:
	//
	// - UNSTRUCTURED: category used for building knowledge base scenarios.
	//
	// <props="china">
	//
	// - SESSION_FILE: file used for [session interaction](https://www.alibabacloud.com/help/en/model-studio/user-guide/file-interaction) in agent applications.
	//
	// <note>When using `SESSION_FILE`, set the CategoryType parameter to `SESSION_FILE` when calling the ApplyFileUploadLease operation as well.</note>
	//
	// <note>The file is valid only for the current user session. After the user closes the session, the file expires. The maximum validity period is 7 days. Long-term storage is not supported.</note>
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The upload lease ID, which corresponds to the `FileUploadLeaseId` returned by the **ApplyFileUploadLease*	- operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68abd1dea7b6404d8f7d7b9f7fbd332d.17166xxxxxxxx
	LeaseId *string `json:"LeaseId,omitempty" xml:"LeaseId,omitempty"`
	// <props="china">
	//
	// Specifies a URL for the file. The system records this URL when building a [document search knowledge base](https://help.aliyun.com/document_detail/2807740.html). When you use the Alibaba Cloud Model Studio console to interact with an [agent application](https://help.aliyun.com/document_detail/2842749.html), this URL is returned with the retrieval results of the file through the `docUrl` field.
	//
	// > The agent application must have **Knowledge Base*	- enabled and the **Show answer sources*	- feature turned on. Otherwise, this parameter does not take effect.
	//
	//
	//
	// <props="intl">
	//
	// Specifies a URL for the file. The system records this URL when building a [document search knowledge base](https://help.aliyun.com/document_detail/2807740.html). When you use the Alibaba Cloud Model Studio console to interact with an [agent application](https://help.aliyun.com/document_detail/2842749.html), this URL is returned with the retrieval results of the file through the `docUrl` field.
	//
	// > The agent application must have **Knowledge Base*	- enabled and the **Show answer sources*	- feature turned on. Otherwise, this parameter does not take effect.
	//
	// example:
	//
	// www.test.com/111.docx
	OriginalFileUrl *string `json:"OriginalFileUrl,omitempty" xml:"OriginalFileUrl,omitempty"`
	// The parser type. Valid values:
	//
	// - DOCMIND: intelligent document parsing
	//
	// - DOCMIND_DIGITAL: electronic document parsing
	//
	// - DOCMIND_LLM_VERSION: large language model document parsing
	//
	// - DASH_QWEN_VL_PARSER: Qwen VL parsing
	//
	// - DOCMIND_LLM_VERSION_MEDIA: audio and video parsing
	//
	// - AUTO_SELECT: automatic parser selection
	//
	// <props="intl">
	//
	// <note>The uploaded file is parsed by using the currently specified parser. If you set this parameter to AUTO_SELECT, the parser configured for the category is used.</note>
	//
	//
	// <props="china">
	//
	// <note>When CategoryType is set to UNSTRUCTURED, the parser parses the uploaded file based on the data parsing settings of the current category.</note>
	//
	// <note>When CategoryType is set to SESSION_FILE, the system parses the file content by using the default method, which cannot be changed.</note>
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTO_SELECT
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The parser configuration. This parameter is required only when the parser type is set to Qwen VL parsing.
	ParserConfigShrink *string `json:"ParserConfig,omitempty" xml:"ParserConfig,omitempty"`
	// - The list of tags associated with the file. You can specify up to 100 tags, and the total character length of all tags cannot exceed 700.
	//
	// - Default value: empty, which means no tags are set.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddFileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddFileShrinkRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *AddFileShrinkRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *AddFileShrinkRequest) GetLeaseId() *string {
	return s.LeaseId
}

func (s *AddFileShrinkRequest) GetOriginalFileUrl() *string {
	return s.OriginalFileUrl
}

func (s *AddFileShrinkRequest) GetParser() *string {
	return s.Parser
}

func (s *AddFileShrinkRequest) GetParserConfigShrink() *string {
	return s.ParserConfigShrink
}

func (s *AddFileShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *AddFileShrinkRequest) SetCategoryId(v string) *AddFileShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *AddFileShrinkRequest) SetCategoryType(v string) *AddFileShrinkRequest {
	s.CategoryType = &v
	return s
}

func (s *AddFileShrinkRequest) SetLeaseId(v string) *AddFileShrinkRequest {
	s.LeaseId = &v
	return s
}

func (s *AddFileShrinkRequest) SetOriginalFileUrl(v string) *AddFileShrinkRequest {
	s.OriginalFileUrl = &v
	return s
}

func (s *AddFileShrinkRequest) SetParser(v string) *AddFileShrinkRequest {
	s.Parser = &v
	return s
}

func (s *AddFileShrinkRequest) SetParserConfigShrink(v string) *AddFileShrinkRequest {
	s.ParserConfigShrink = &v
	return s
}

func (s *AddFileShrinkRequest) SetTagsShrink(v string) *AddFileShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *AddFileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
