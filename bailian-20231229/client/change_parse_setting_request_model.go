// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeParseSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *ChangeParseSettingRequest
	GetCategoryId() *string
	SetFileType(v string) *ChangeParseSettingRequest
	GetFileType() *string
	SetParser(v string) *ChangeParseSettingRequest
	GetParser() *string
	SetParserConfig(v *ChangeParseSettingRequestParserConfig) *ChangeParseSettingRequest
	GetParserConfig() *ChangeParseSettingRequestParserConfig
}

type ChangeParseSettingRequest struct {
	// The category ID, which is the `CategoryId` returned by the **AddCategory*	- operation. You can also obtain it by clicking the ID icon next to the category name on the <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) - Files tab<props="intl">[Application Data](https://bailian.console.alibabacloud.com/?tab=app#/data-center) - Unstructured Data tab.
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The file type (extension). Valid values:
	//
	// - doc
	//
	// - docx
	//
	// - ppt
	//
	// - pptx
	//
	// - xls
	//
	// - xlsx
	//
	// - md
	//
	// - txt
	//
	// - pdf
	//
	// - png
	//
	// - jpg
	//
	// - jpeg
	//
	// - bmp
	//
	// - gif
	//
	// - html
	//
	// This parameter is required.
	//
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The parser identifier code. Different parsers are applicable to different scenarios. For more information, see "Knowledge Base". Valid values:
	//
	// - DOCMIND (intelligent document parsing)
	//
	// - DOCMIND_DIGITAL (electronic document parsing)
	//
	// - DOCMIND_LLM_VERSION (large model document parsing)
	//
	// - DASH_QWEN_VL_PARSER (Qwen VL parsing)
	//
	// - DOCMIND_LLM_VERSION_MEDIA (audio and video parsing)
	//
	// This parameter is required.
	//
	// example:
	//
	// DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The parser configuration. This parameter is required only when the parser is set to Qwen VL parsing.
	ParserConfig *ChangeParseSettingRequestParserConfig `json:"ParserConfig,omitempty" xml:"ParserConfig,omitempty" type:"Struct"`
}

func (s ChangeParseSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeParseSettingRequest) GoString() string {
	return s.String()
}

func (s *ChangeParseSettingRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ChangeParseSettingRequest) GetFileType() *string {
	return s.FileType
}

func (s *ChangeParseSettingRequest) GetParser() *string {
	return s.Parser
}

func (s *ChangeParseSettingRequest) GetParserConfig() *ChangeParseSettingRequestParserConfig {
	return s.ParserConfig
}

func (s *ChangeParseSettingRequest) SetCategoryId(v string) *ChangeParseSettingRequest {
	s.CategoryId = &v
	return s
}

func (s *ChangeParseSettingRequest) SetFileType(v string) *ChangeParseSettingRequest {
	s.FileType = &v
	return s
}

func (s *ChangeParseSettingRequest) SetParser(v string) *ChangeParseSettingRequest {
	s.Parser = &v
	return s
}

func (s *ChangeParseSettingRequest) SetParserConfig(v *ChangeParseSettingRequestParserConfig) *ChangeParseSettingRequest {
	s.ParserConfig = v
	return s
}

func (s *ChangeParseSettingRequest) Validate() error {
	if s.ParserConfig != nil {
		if err := s.ParserConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeParseSettingRequestParserConfig struct {
	// The model name.
	//
	// example:
	//
	// qwen-vl-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The prompt used when calling Qwen VL parsing.
	//
	// example:
	//
	// # Role
	//
	// You are a professional image content annotator who excels at identifying and describing the content in images.
	//
	// # Task Objective
	//
	// Based on the input image, describe the content in the image in detail
	ModelPrompt *string `json:"modelPrompt,omitempty" xml:"modelPrompt,omitempty"`
}

func (s ChangeParseSettingRequestParserConfig) String() string {
	return dara.Prettify(s)
}

func (s ChangeParseSettingRequestParserConfig) GoString() string {
	return s.String()
}

func (s *ChangeParseSettingRequestParserConfig) GetModelName() *string {
	return s.ModelName
}

func (s *ChangeParseSettingRequestParserConfig) GetModelPrompt() *string {
	return s.ModelPrompt
}

func (s *ChangeParseSettingRequestParserConfig) SetModelName(v string) *ChangeParseSettingRequestParserConfig {
	s.ModelName = &v
	return s
}

func (s *ChangeParseSettingRequestParserConfig) SetModelPrompt(v string) *ChangeParseSettingRequestParserConfig {
	s.ModelPrompt = &v
	return s
}

func (s *ChangeParseSettingRequestParserConfig) Validate() error {
	return dara.Validate(s)
}
