// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParseSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetParseSettingsResponseBody
	GetCode() *string
	SetData(v []*GetParseSettingsResponseBodyData) *GetParseSettingsResponseBody
	GetData() []*GetParseSettingsResponseBodyData
	SetMessage(v string) *GetParseSettingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetParseSettingsResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetParseSettingsResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetParseSettingsResponseBody
	GetSuccess() *bool
}

type GetParseSettingsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data fields.
	Data []*GetParseSettingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// workspace id is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35A267BF-xxxx-54DB-8394-AA3B0742D833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetParseSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetParseSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetParseSettingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetParseSettingsResponseBody) GetData() []*GetParseSettingsResponseBodyData {
	return s.Data
}

func (s *GetParseSettingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetParseSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetParseSettingsResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetParseSettingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetParseSettingsResponseBody) SetCode(v string) *GetParseSettingsResponseBody {
	s.Code = &v
	return s
}

func (s *GetParseSettingsResponseBody) SetData(v []*GetParseSettingsResponseBodyData) *GetParseSettingsResponseBody {
	s.Data = v
	return s
}

func (s *GetParseSettingsResponseBody) SetMessage(v string) *GetParseSettingsResponseBody {
	s.Message = &v
	return s
}

func (s *GetParseSettingsResponseBody) SetRequestId(v string) *GetParseSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParseSettingsResponseBody) SetStatus(v string) *GetParseSettingsResponseBody {
	s.Status = &v
	return s
}

func (s *GetParseSettingsResponseBody) SetSuccess(v bool) *GetParseSettingsResponseBody {
	s.Success = &v
	return s
}

func (s *GetParseSettingsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetParseSettingsResponseBodyData struct {
	// The file type. Valid values are: pdf, docx, doc, etc. All supported file types in the category are listed here.
	//
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The parser used for files of the current type. Valid values:
	//
	// 	- DOCMIND (Intelligent parsing)
	//
	// 	- DOCMIND_DIGITAL (Digital parsing)
	//
	// 	- DOCMIND_LLM_VERSION (LLM parsing)
	//
	// 	- DASH_QWEN_VL_PARSER (Qwen VL parsing)
	//
	// example:
	//
	// DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The parser configuration. Currently, this is available only for Qwen VL parsing.
	ParserConfig *GetParseSettingsResponseBodyDataParserConfig `json:"ParserConfig,omitempty" xml:"ParserConfig,omitempty" type:"Struct"`
	// The display name of the parsing method.
	//
	// example:
	//
	// Digital parsing
	ParserDisplayName *string `json:"ParserDisplayName,omitempty" xml:"ParserDisplayName,omitempty"`
}

func (s GetParseSettingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetParseSettingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetParseSettingsResponseBodyData) GetFileType() *string {
	return s.FileType
}

func (s *GetParseSettingsResponseBodyData) GetParser() *string {
	return s.Parser
}

func (s *GetParseSettingsResponseBodyData) GetParserConfig() *GetParseSettingsResponseBodyDataParserConfig {
	return s.ParserConfig
}

func (s *GetParseSettingsResponseBodyData) GetParserDisplayName() *string {
	return s.ParserDisplayName
}

func (s *GetParseSettingsResponseBodyData) SetFileType(v string) *GetParseSettingsResponseBodyData {
	s.FileType = &v
	return s
}

func (s *GetParseSettingsResponseBodyData) SetParser(v string) *GetParseSettingsResponseBodyData {
	s.Parser = &v
	return s
}

func (s *GetParseSettingsResponseBodyData) SetParserConfig(v *GetParseSettingsResponseBodyDataParserConfig) *GetParseSettingsResponseBodyData {
	s.ParserConfig = v
	return s
}

func (s *GetParseSettingsResponseBodyData) SetParserDisplayName(v string) *GetParseSettingsResponseBodyData {
	s.ParserDisplayName = &v
	return s
}

func (s *GetParseSettingsResponseBodyData) Validate() error {
	if s.ParserConfig != nil {
		if err := s.ParserConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetParseSettingsResponseBodyDataParserConfig struct {
	// The model name.
	//
	// example:
	//
	// - qwen-vl-max
	//
	// - qwen-vl-plus
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The prompt used for parsing.
	//
	// example:
	//
	// # Role
	//
	// You are a professional image content annotator, skilled in identifying and describing the contents of images.
	//
	// # Task Objective
	//
	// Based on the input image, provide a detailed description of its contents.
	ModelPrompt *string `json:"ModelPrompt,omitempty" xml:"ModelPrompt,omitempty"`
}

func (s GetParseSettingsResponseBodyDataParserConfig) String() string {
	return dara.Prettify(s)
}

func (s GetParseSettingsResponseBodyDataParserConfig) GoString() string {
	return s.String()
}

func (s *GetParseSettingsResponseBodyDataParserConfig) GetModelName() *string {
	return s.ModelName
}

func (s *GetParseSettingsResponseBodyDataParserConfig) GetModelPrompt() *string {
	return s.ModelPrompt
}

func (s *GetParseSettingsResponseBodyDataParserConfig) SetModelName(v string) *GetParseSettingsResponseBodyDataParserConfig {
	s.ModelName = &v
	return s
}

func (s *GetParseSettingsResponseBodyDataParserConfig) SetModelPrompt(v string) *GetParseSettingsResponseBodyDataParserConfig {
	s.ModelPrompt = &v
	return s
}

func (s *GetParseSettingsResponseBodyDataParserConfig) Validate() error {
	return dara.Validate(s)
}
