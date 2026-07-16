// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeParseSettingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *ChangeParseSettingShrinkRequest
	GetCategoryId() *string
	SetFileType(v string) *ChangeParseSettingShrinkRequest
	GetFileType() *string
	SetParser(v string) *ChangeParseSettingShrinkRequest
	GetParser() *string
	SetParserConfigShrink(v string) *ChangeParseSettingShrinkRequest
	GetParserConfigShrink() *string
}

type ChangeParseSettingShrinkRequest struct {
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
	// - DOCMIND (Intelligent Document Parsing)
	//
	// - DOCMIND_DIGITAL (Electronic Document Parsing)
	//
	// - DOCMIND_LLM_VERSION (Large Model Document Parsing)
	//
	// - DASH_QWEN_VL_PARSER (Qwen VL Parsing)
	//
	// - DOCMIND_LLM_VERSION_MEDIA (Audio/Video Parsing)
	//
	// This parameter is required.
	//
	// example:
	//
	// DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The parser configuration. This parameter is required only when the parser is set to Qwen VL Parsing.
	ParserConfigShrink *string `json:"ParserConfig,omitempty" xml:"ParserConfig,omitempty"`
}

func (s ChangeParseSettingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeParseSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChangeParseSettingShrinkRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ChangeParseSettingShrinkRequest) GetFileType() *string {
	return s.FileType
}

func (s *ChangeParseSettingShrinkRequest) GetParser() *string {
	return s.Parser
}

func (s *ChangeParseSettingShrinkRequest) GetParserConfigShrink() *string {
	return s.ParserConfigShrink
}

func (s *ChangeParseSettingShrinkRequest) SetCategoryId(v string) *ChangeParseSettingShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *ChangeParseSettingShrinkRequest) SetFileType(v string) *ChangeParseSettingShrinkRequest {
	s.FileType = &v
	return s
}

func (s *ChangeParseSettingShrinkRequest) SetParser(v string) *ChangeParseSettingShrinkRequest {
	s.Parser = &v
	return s
}

func (s *ChangeParseSettingShrinkRequest) SetParserConfigShrink(v string) *ChangeParseSettingShrinkRequest {
	s.ParserConfigShrink = &v
	return s
}

func (s *ChangeParseSettingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
