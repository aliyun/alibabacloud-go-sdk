// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartAuditResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSmartAuditResultResponseBody
	GetCode() *string
	SetData(v *GetSmartAuditResultResponseBodyData) *GetSmartAuditResultResponseBody
	GetData() *GetSmartAuditResultResponseBodyData
	SetHttpStatusCode(v int32) *GetSmartAuditResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSmartAuditResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSmartAuditResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSmartAuditResultResponseBody
	GetSuccess() *bool
}

type GetSmartAuditResultResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSmartAuditResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSmartAuditResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAuditResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmartAuditResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSmartAuditResultResponseBody) GetData() *GetSmartAuditResultResponseBodyData {
	return s.Data
}

func (s *GetSmartAuditResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSmartAuditResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSmartAuditResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmartAuditResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSmartAuditResultResponseBody) SetCode(v string) *GetSmartAuditResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmartAuditResultResponseBody) SetData(v *GetSmartAuditResultResponseBodyData) *GetSmartAuditResultResponseBody {
	s.Data = v
	return s
}

func (s *GetSmartAuditResultResponseBody) SetHttpStatusCode(v int32) *GetSmartAuditResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSmartAuditResultResponseBody) SetMessage(v string) *GetSmartAuditResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmartAuditResultResponseBody) SetRequestId(v string) *GetSmartAuditResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmartAuditResultResponseBody) SetSuccess(v bool) *GetSmartAuditResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetSmartAuditResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSmartAuditResultResponseBodyData struct {
	ErrorItemDetails []*GetSmartAuditResultResponseBodyDataErrorItemDetails `json:"ErrorItemDetails,omitempty" xml:"ErrorItemDetails,omitempty" type:"Repeated"`
	ErrorMessage     *string                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// SUCCESSED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSmartAuditResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAuditResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSmartAuditResultResponseBodyData) GetErrorItemDetails() []*GetSmartAuditResultResponseBodyDataErrorItemDetails {
	return s.ErrorItemDetails
}

func (s *GetSmartAuditResultResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSmartAuditResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSmartAuditResultResponseBodyData) SetErrorItemDetails(v []*GetSmartAuditResultResponseBodyDataErrorItemDetails) *GetSmartAuditResultResponseBodyData {
	s.ErrorItemDetails = v
	return s
}

func (s *GetSmartAuditResultResponseBodyData) SetErrorMessage(v string) *GetSmartAuditResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyData) SetStatus(v string) *GetSmartAuditResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetSmartAuditResultResponseBodyDataErrorItemDetails struct {
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// example:
	//
	// 0
	ContextOffset *int32 `json:"ContextOffset,omitempty" xml:"ContextOffset,omitempty"`
	// example:
	//
	// 2
	ErrorLevel *int32 `json:"ErrorLevel,omitempty" xml:"ErrorLevel,omitempty"`
	// example:
	//
	// ”xxx“
	ErrorWord *string `json:"ErrorWord,omitempty" xml:"ErrorWord,omitempty"`
	// example:
	//
	// ContentAccuracy
	MajorCode     *string `json:"MajorCode,omitempty" xml:"MajorCode,omitempty"`
	MajorCodeDesc *string `json:"MajorCodeDesc,omitempty" xml:"MajorCodeDesc,omitempty"`
	// example:
	//
	// 0
	Offset *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// “xxx”
	RightWord *string `json:"RightWord,omitempty" xml:"RightWord,omitempty"`
	// example:
	//
	// PunctuationError
	SubClassCode *string `json:"SubClassCode,omitempty" xml:"SubClassCode,omitempty"`
	SubClassDesc *string `json:"SubClassDesc,omitempty" xml:"SubClassDesc,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetSmartAuditResultResponseBodyDataErrorItemDetails) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAuditResultResponseBodyDataErrorItemDetails) GoString() string {
	return s.String()
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetCheckId() *string {
	return s.CheckId
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetContext() *string {
	return s.Context
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetContextOffset() *int32 {
	return s.ContextOffset
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetErrorLevel() *int32 {
	return s.ErrorLevel
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetErrorWord() *string {
	return s.ErrorWord
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetMajorCode() *string {
	return s.MajorCode
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetMajorCodeDesc() *string {
	return s.MajorCodeDesc
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetOffset() *int32 {
	return s.Offset
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetReason() *string {
	return s.Reason
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetRightWord() *string {
	return s.RightWord
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetSubClassCode() *string {
	return s.SubClassCode
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetSubClassDesc() *string {
	return s.SubClassDesc
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) GetUrl() *string {
	return s.Url
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetCheckId(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.CheckId = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetContext(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.Context = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetContextOffset(v int32) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.ContextOffset = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetErrorLevel(v int32) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.ErrorLevel = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetErrorWord(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.ErrorWord = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetMajorCode(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.MajorCode = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetMajorCodeDesc(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.MajorCodeDesc = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetOffset(v int32) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.Offset = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetReason(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.Reason = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetRightWord(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.RightWord = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetSubClassCode(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.SubClassCode = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetSubClassDesc(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.SubClassDesc = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) SetUrl(v string) *GetSmartAuditResultResponseBodyDataErrorItemDetails {
	s.Url = &v
	return s
}

func (s *GetSmartAuditResultResponseBodyDataErrorItemDetails) Validate() error {
	return dara.Validate(s)
}
