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
	// example:
	//
	// success
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetParseSettingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// workspace id is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 35A267BF-xxxx-54DB-8394-AA3B0742D833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	return dara.Validate(s)
}

type GetParseSettingsResponseBodyData struct {
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// DOCMIND
	Parser            *string                                       `json:"Parser,omitempty" xml:"Parser,omitempty"`
	ParserConfig      *GetParseSettingsResponseBodyDataParserConfig `json:"ParserConfig,omitempty" xml:"ParserConfig,omitempty" type:"Struct"`
	ParserDisplayName *string                                       `json:"ParserDisplayName,omitempty" xml:"ParserDisplayName,omitempty"`
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
	return dara.Validate(s)
}

type GetParseSettingsResponseBodyDataParserConfig struct {
	// example:
	//
	// - qwen-vl-max
	//
	// - qwen-vl-plus
	ModelName   *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
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
