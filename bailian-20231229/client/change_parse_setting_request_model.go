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
	// The category ID, which is the `CategoryId` returned by **AddCategory**. To view the category ID, click the ID icon next to the category name on the Unstructured Data tab of the [Application Data](https://bailian.console.alibabacloud.com/?tab=app#/data-center) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The file type. Valid values: pdf, docx, and doc.
	//
	// This parameter is required.
	//
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The parser code. Valid values:
	//
	// 	- DOCMIND (Intelligent parsing)
	//
	// 	- DOCMIND_DIGITAL (Digital parsing)
	//
	// 	- DOCMIND_LLM_VERSION (LLM parsing)
	//
	// 	- DASH_QWEN_VL_PARSER (Qwen VL parsing)
	//
	// This parameter is required.
	//
	// example:
	//
	// DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The parser configuration. Currently, this is available only for Qwen VL parsing.
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
	// Valid values:
	//
	// 	- qwen-vl-max
	//
	// 	- qwen-vl-plus
	//
	// example:
	//
	// qwen-vl-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The prompt used for parsing.
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
