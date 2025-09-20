// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailableParserTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAvailableParserTypesResponseBody
	GetCode() *string
	SetData(v *GetAvailableParserTypesResponseBodyData) *GetAvailableParserTypesResponseBody
	GetData() *GetAvailableParserTypesResponseBodyData
	SetMessage(v string) *GetAvailableParserTypesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAvailableParserTypesResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetAvailableParserTypesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetAvailableParserTypesResponseBody
	GetSuccess() *bool
}

type GetAvailableParserTypesResponseBody struct {
	// example:
	//
	// DataCenter.Throttling
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAvailableParserTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 17204B98-7734-4F9A-8464-2446XXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAvailableParserTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableParserTypesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAvailableParserTypesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAvailableParserTypesResponseBody) GetData() *GetAvailableParserTypesResponseBodyData {
	return s.Data
}

func (s *GetAvailableParserTypesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAvailableParserTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAvailableParserTypesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAvailableParserTypesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAvailableParserTypesResponseBody) SetCode(v string) *GetAvailableParserTypesResponseBody {
	s.Code = &v
	return s
}

func (s *GetAvailableParserTypesResponseBody) SetData(v *GetAvailableParserTypesResponseBodyData) *GetAvailableParserTypesResponseBody {
	s.Data = v
	return s
}

func (s *GetAvailableParserTypesResponseBody) SetMessage(v string) *GetAvailableParserTypesResponseBody {
	s.Message = &v
	return s
}

func (s *GetAvailableParserTypesResponseBody) SetRequestId(v string) *GetAvailableParserTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAvailableParserTypesResponseBody) SetStatus(v string) *GetAvailableParserTypesResponseBody {
	s.Status = &v
	return s
}

func (s *GetAvailableParserTypesResponseBody) SetSuccess(v bool) *GetAvailableParserTypesResponseBody {
	s.Success = &v
	return s
}

func (s *GetAvailableParserTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAvailableParserTypesResponseBodyData struct {
	// example:
	//
	// pdf
	FileType   *string                                              `json:"FileType,omitempty" xml:"FileType,omitempty"`
	ParserList []*GetAvailableParserTypesResponseBodyDataParserList `json:"ParserList,omitempty" xml:"ParserList,omitempty" type:"Repeated"`
}

func (s GetAvailableParserTypesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableParserTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAvailableParserTypesResponseBodyData) GetFileType() *string {
	return s.FileType
}

func (s *GetAvailableParserTypesResponseBodyData) GetParserList() []*GetAvailableParserTypesResponseBodyDataParserList {
	return s.ParserList
}

func (s *GetAvailableParserTypesResponseBodyData) SetFileType(v string) *GetAvailableParserTypesResponseBodyData {
	s.FileType = &v
	return s
}

func (s *GetAvailableParserTypesResponseBodyData) SetParserList(v []*GetAvailableParserTypesResponseBodyDataParserList) *GetAvailableParserTypesResponseBodyData {
	s.ParserList = v
	return s
}

func (s *GetAvailableParserTypesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetAvailableParserTypesResponseBodyDataParserList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
}

func (s GetAvailableParserTypesResponseBodyDataParserList) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableParserTypesResponseBodyDataParserList) GoString() string {
	return s.String()
}

func (s *GetAvailableParserTypesResponseBodyDataParserList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetAvailableParserTypesResponseBodyDataParserList) GetParser() *string {
	return s.Parser
}

func (s *GetAvailableParserTypesResponseBodyDataParserList) SetDisplayName(v string) *GetAvailableParserTypesResponseBodyDataParserList {
	s.DisplayName = &v
	return s
}

func (s *GetAvailableParserTypesResponseBodyDataParserList) SetParser(v string) *GetAvailableParserTypesResponseBodyDataParserList {
	s.Parser = &v
	return s
}

func (s *GetAvailableParserTypesResponseBodyDataParserList) Validate() error {
	return dara.Validate(s)
}
