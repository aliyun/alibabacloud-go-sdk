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
