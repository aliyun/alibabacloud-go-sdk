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
	// The category ID. This is the `CategoryId` returned by the **AddCategory*	- operation. You can also obtain the ID from the <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) - File tab<props="intl">[Application Data](https://bailian.console.alibabacloud.com/?tab=app#/data-center) - Unstructured Data tab by clicking the ID icon next to the category name.
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The file type, specified by its extension. Valid values:
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
	// The identifier for the parser. Different parsers are suitable for different scenarios. For more information, refer to the knowledge base. Valid values:
	//
	// - DOCMIND (intelligent document parsing)
	//
	// - DOCMIND_DIGITAL (digital document parsing)
	//
	// - DOCMIND_LLM_VERSION (LLM-based document parsing)
	//
	// - DASH_QWEN_VL_PARSER (Qwen VL Parser)
	//
	// This parameter is required.
	//
	// example:
	//
	// DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The parser configuration. This parameter is required only when the `Parser` parameter is set to `DASH_QWEN_VL_PARSER`.
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
