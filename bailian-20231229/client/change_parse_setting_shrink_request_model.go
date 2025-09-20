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
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DOCMIND
	Parser             *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
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
