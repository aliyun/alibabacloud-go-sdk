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
	// - If `CategoryType` is set to `UNSTRUCTURED`, you must specify the ID of the category to which the file belongs. This is the `CategoryId` returned by the **AddCategory*	- API. You can also obtain the category ID by navigating to the \\*\\*Application data\\*\\	- > \\*\\*Files\\*\\	- tab and clicking the ID icon next to the category name. You can specify `default` to use the default category.
	//
	// - If `CategoryType` is set to `SESSION_FILE`, specify `default`.
	//
	//
	//
	// <props="intl">
	//
	// The ID of the category to which the file belongs. This is the `CategoryId` returned by the **AddCategory*	- API. You can also obtain the category ID by navigating to the \\*\\*Application data\\*\\	- > \\*\\*Files\\*\\	- tab and clicking the ID icon next to the category name. You can specify `default` to use the default category.
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The type of category. This parameter is optional. Default value: `UNSTRUCTURED`. Valid values:
	//
	// - `UNSTRUCTURED`: A category used for building a knowledge base.
	//
	// <props="china">
	//
	// - `SESSION_FILE`: A file used for interactions within an agent [session](https://help.aliyun.com/zh/model-studio/user-guide/file-interaction).
	//
	//   > If you set this parameter to `SESSION_FILE`, you must also set the `CategoryType` parameter to `SESSION_FILE` when you call the ApplyFileUploadLease API.
	//
	//   > Files of this type are valid only for the current session and expire after the session is closed, with a maximum validity of 7 days. These files are not intended for long-term storage.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The upload lease ID. This value maps to the `FileUploadLeaseId` returned by the **ApplyFileUploadLease*	- API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68abd1dea7b6404d8f7d7b9f7fbd332d.17166xxxxxxxx
	LeaseId *string `json:"LeaseId,omitempty" xml:"LeaseId,omitempty"`
	// <props="china">
	//
	// The URL of the file. The system records this link when building a [document retrieval-based knowledge base](https://help.aliyun.com/document_detail/2807740.html). When you interact with an [agent](https://help.aliyun.com/document_detail/2842749.html) in the Alibaba Cloud Model Studio console, this URL is returned with the retrieval results for the file in the `docUrl` field.
	//
	// > For this parameter to take effect, the **knowledge base*	- feature must be enabled for the agent, and the **display the source of the answer*	- option must be enabled.
	//
	//
	//
	// <props="intl">
	//
	// The URL of the file. The system records this link when building a [document retrieval-based knowledge base](https://help.aliyun.com/document_detail/2807740.html). When you interact with an [agent](https://help.aliyun.com/document_detail/2842749.html) in the Alibaba Cloud Model Studio console, this URL is returned with the retrieval results for the file in the `docUrl` field.
	//
	// > For this parameter to take effect, the **knowledge base*	- feature must be enabled for the agent, and the **display the source of the answer*	- option must be enabled.
	//
	// example:
	//
	// www.test.com/111.docx
	OriginalFileUrl *string `json:"OriginalFileUrl,omitempty" xml:"OriginalFileUrl,omitempty"`
	// The type of parser. Valid values:
	//
	// - DOCMIND: Intelligent Document Parsing
	//
	// - DOCMIND_DIGITAL: Digital Document Parsing
	//
	// - DOCMIND_LLM_VERSION: Large Language Model-based Document Parsing
	//
	// - DASH_QWEN_VL_PARSER: Qwen-VL Parsing
	//
	// - DOCMIND_LLM_VERSION_MEDIA: Audio and Video Parsing
	//
	// - AUTO_SELECT: Automatic Parser Selection
	//
	// <props="intl">
	//
	// > The system uses the specified parser to parse the uploaded file. If you set this parameter to `AUTO_SELECT`, the parser configured for the category is used.
	//
	//
	//
	// <props="china">
	//
	// > If `CategoryType` is set to `UNSTRUCTURED`, the parser parses your uploaded file based on the category’s data parsing settings.
	//
	// > If `CategoryType` is set to `SESSION_FILE`, the system uses a default parsing method that cannot be changed.
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTO_SELECT
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The parser configuration. This parameter is required only if you set `Parser` to `DASH_QWEN_VL_PARSER`.
	ParserConfigShrink *string `json:"ParserConfig,omitempty" xml:"ParserConfig,omitempty"`
	// - A list of tags for the file. You can specify up to 100 tags. The total length of all tags cannot exceed 700 characters.
	//
	// - If this parameter is not specified, no tags are added.
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
