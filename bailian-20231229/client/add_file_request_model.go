// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *AddFileRequest
	GetCategoryId() *string
	SetCategoryType(v string) *AddFileRequest
	GetCategoryType() *string
	SetLeaseId(v string) *AddFileRequest
	GetLeaseId() *string
	SetOriginalFileUrl(v string) *AddFileRequest
	GetOriginalFileUrl() *string
	SetParser(v string) *AddFileRequest
	GetParser() *string
	SetParserConfig(v *AddFileRequestParserConfig) *AddFileRequest
	GetParserConfig() *AddFileRequestParserConfig
	SetTags(v []*string) *AddFileRequest
	GetTags() []*string
}

type AddFileRequest struct {
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
	ParserConfig *AddFileRequestParserConfig `json:"ParserConfig,omitempty" xml:"ParserConfig,omitempty" type:"Struct"`
	// - A list of tags for the file. You can specify up to 100 tags. The total length of all tags cannot exceed 700 characters.
	//
	// - If this parameter is not specified, no tags are added.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AddFileRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFileRequest) GoString() string {
	return s.String()
}

func (s *AddFileRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *AddFileRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *AddFileRequest) GetLeaseId() *string {
	return s.LeaseId
}

func (s *AddFileRequest) GetOriginalFileUrl() *string {
	return s.OriginalFileUrl
}

func (s *AddFileRequest) GetParser() *string {
	return s.Parser
}

func (s *AddFileRequest) GetParserConfig() *AddFileRequestParserConfig {
	return s.ParserConfig
}

func (s *AddFileRequest) GetTags() []*string {
	return s.Tags
}

func (s *AddFileRequest) SetCategoryId(v string) *AddFileRequest {
	s.CategoryId = &v
	return s
}

func (s *AddFileRequest) SetCategoryType(v string) *AddFileRequest {
	s.CategoryType = &v
	return s
}

func (s *AddFileRequest) SetLeaseId(v string) *AddFileRequest {
	s.LeaseId = &v
	return s
}

func (s *AddFileRequest) SetOriginalFileUrl(v string) *AddFileRequest {
	s.OriginalFileUrl = &v
	return s
}

func (s *AddFileRequest) SetParser(v string) *AddFileRequest {
	s.Parser = &v
	return s
}

func (s *AddFileRequest) SetParserConfig(v *AddFileRequestParserConfig) *AddFileRequest {
	s.ParserConfig = v
	return s
}

func (s *AddFileRequest) SetTags(v []*string) *AddFileRequest {
	s.Tags = v
	return s
}

func (s *AddFileRequest) Validate() error {
	if s.ParserConfig != nil {
		if err := s.ParserConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFileRequestParserConfig struct {
	// The model name.
	//
	// example:
	//
	// qwen-vl-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The prompt to use when calling the Qwen-VL parser.
	//
	// example:
	//
	// #角色
	//
	// 你是一个专业的图片内容标注人员，擅长识别并描述出图片中的内容。
	//
	// # 任务目标
	//
	// 请结合输入图片，详细描述图片中的内容。
	ModelPrompt *string `json:"ModelPrompt,omitempty" xml:"ModelPrompt,omitempty"`
}

func (s AddFileRequestParserConfig) String() string {
	return dara.Prettify(s)
}

func (s AddFileRequestParserConfig) GoString() string {
	return s.String()
}

func (s *AddFileRequestParserConfig) GetModelName() *string {
	return s.ModelName
}

func (s *AddFileRequestParserConfig) GetModelPrompt() *string {
	return s.ModelPrompt
}

func (s *AddFileRequestParserConfig) SetModelName(v string) *AddFileRequestParserConfig {
	s.ModelName = &v
	return s
}

func (s *AddFileRequestParserConfig) SetModelPrompt(v string) *AddFileRequestParserConfig {
	s.ModelPrompt = &v
	return s
}

func (s *AddFileRequestParserConfig) Validate() error {
	return dara.Validate(s)
}
